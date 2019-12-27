/*
 * Copyright 2019 Hayo van Loon
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package logjson contains methods that use the standard Go logger to output
// JSON rather than plain text.
//
// The compatibility target is StackDriver Logging.
package logjson

import (
	"encoding/json"
	"log"
)

type Logger interface {
	// Logs a message on DEBUG level.
	Debug(msg string)
	// Logs a message on INFO level.
	Info(msg string)
	// Logs a message on WARN level.
	Warn(msg string)
	// Logs a message on ERROR level.
	Error(msg string)
	// Logs a message on CRITICAL level and exits application (like log.Fatal).
	Critical(msg string)
}

type logger struct {
	projectId string
	component string
}

type severity string

const (
	debug    severity = "DEBUG"
	info     severity = "INFO"
	warning  severity = "WARNING"
	_error   severity = "ERROR"
	critical severity = "CRITICAL"
)

var instance Logger = logger{}

// Based on https://github.com/GoogleCloudPlatform/golang-samples/blob/master/run/logging-manual/main.go
type entry struct {
	Message  string `json:"message"`
	Severity string `json:"severity"`
	Trace    string `json:"logging.googleapis.com/trace,omitempty"` // TODO decide whether or not to implement it
	// Stackdriver Log Viewer allows filtering and display of this as `jsonPayload.component`.
	Component string `json:"component,omitempty"`
}

// String renders an entry structure to the JSON format expected by Stackdriver.
func (e entry) String() string {
	out, err := json.Marshal(e)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
	}
	return string(out)
}

func (l logger) log(msg string, sev severity, trace string) {
	e := entry{
		Message:  msg,
		Severity: string(sev),
	}
	if trace != "" {
		e.Trace = "/projects/" + l.projectId + "/traces/" + trace
	}
	if l.component != "" {
		e.Component = l.component
	}
	if sev == critical {
		log.Fatalln(e)
	} else {
		log.Println(e)
	}
}

func (l logger) Info(msg string) {
	l.log(msg, info, "")
}

func (l logger) Debug(msg string) {
	l.log(msg, debug, "")
}

func (l logger) Warn(msg string) {
	l.log(msg, warning, "")
}

func (l logger) Error(msg string) {
	l.log(msg, _error, "")
}

func (l logger) Critical(msg string) {
	l.log(msg, critical, "")
}

func Info(msg string) {
	instance.Info(msg)
}

func Debug(msg string) {
	instance.Debug(msg)
}

func Warn(msg string) {
	instance.Warn(msg)
}

func Error(msg string) {
	instance.Error(msg)
}

func Critical(msg string) {
	instance.Critical(msg)
}
