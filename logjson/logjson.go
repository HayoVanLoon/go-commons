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
// Logs according to specified log level.
//
// DEFAULT   (0) The log entry has no assigned Severity level.
// DEBUG     (100) Debug or trace information.
// INFO      (200) Routine information, such as ongoing status or performance.
// NOTICE    (300) Normal but significant events, such as start up, shut down, or a configuration change.
// WARNING   (400) Warning events might cause problems.
// ERROR     (500) Error events are likely to cause problems.
// CRITICAL  (600) Critical events cause more severe problems or outages.
// ALERT     (700) A person must take an action immediately.
// EMERGENCY (800) One or more systems are unusable.

package logjson

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Logger interface {
	// Gets the minimum severity level that will be reported
	GetLevel() Severity
	// Sets the minimum severity level reported
	SetLevel(sev Severity)
	// Logs a message on DEBUG level.
	Debug(v interface{})
	// Logs a message on INFO level.
	Info(v interface{})
	// Logs a message on NOTICE level.
	Notice(v interface{})
	// Logs a message on WARN level.
	Warn(v interface{})
	// Logs a message on ERROR level.
	Error(v interface{})
	// Logs a message on CRITICAL level and exits application (like log.Fatal).
	Critical(v interface{})
	// Logs a message on ALERT level and exits application (like log.Fatal).
	Alert(v interface{})
	// Logs a message on EMERGENCY level and exits application (like log.Fatal).
	Emergency(v interface{})
}

type Severity int

const (
	LevelDebug Severity = iota
	LevelInfo           = iota + 100
	LevelNotice
	LevelWarning
	LevelError
	LevelCritical
	LevelAlert
	LevelEmergency
)

var toSeverity = map[string]Severity{
	"DEBUG":     LevelDebug,
	"INFO":      LevelInfo,
	"NOTICE":    LevelNotice,
	"WARNING":   LevelWarning,
	"ERROR":     LevelError,
	"CRITICAL":  LevelCritical,
	"ALERT":     LevelAlert,
	"EMERGENCY": LevelEmergency,
}

var toName = func() map[Severity]string {
	m := map[Severity]string{}
	for k := range toSeverity {
		m[toSeverity[k]] = k
	}
	return m
}()

type logger struct {
	projectId string
	component string
	level     Severity
}

func getLogLevel() Severity {
	l := os.Getenv("GCP_LOG_LEVEL")
	sev, ok := toSeverity[l]
	if ok {
		return sev
	}
	return LevelDebug
}

func NewDefaultLogger(project, component string) Logger {
	return &logger{
		projectId: project,
		component: component,
		level:     getLogLevel(),
	}
}

func NewLogger(project, component string, sev Severity) Logger {
	return &logger{
		projectId: project,
		component: component,
		level:     sev,
	}
}

var instance Logger = NewDefaultLogger("", "")

// Based on https://github.com/GoogleCloudPlatform/golang-samples/blob/master/run/logging-manual/main.go
type entry struct {
	Message  interface{} `json:"message"`
	Severity string      `json:"Severity"`
	Trace    string      `json:"logging.googleapis.com/trace,omitempty"` // TODO decide whether or not to implement it
	// Stackdriver Log Viewer allows filtering and display of this as `jsonPayload.component`.
	Component string `json:"component,omitempty"`
}

// String renders an entry structure to the JSON format expected by Stackdriver.
func (e entry) String() string {
	// JSON serialisability already verified on creation
	out, _ := json.Marshal(e)
	return string(out)
}

func (l logger) log(v interface{}, sev Severity, trace string) {
	// TODO: check if message can be nested or if we need to add extra fields
	e := entry{Severity: toName[sev]}

	_, err := json.Marshal(v)
	if err != nil {
		e.Message = fmt.Sprintf("%v", v)
	} else {
		e.Message = v
	}

	if trace != "" {
		e.Trace = "/projects/" + l.projectId + "/traces/" + trace
	}
	if l.component != "" {
		e.Component = l.component
	}

	if sev < LevelCritical {
		log.Println(e)
	} else {
		log.Fatalln(e)
	}
}

func (l logger) GetLevel() Severity {
	return l.level
}

func (l *logger) SetLevel(sev Severity) {
	l.level = sev
}

func (l logger) Debug(v interface{}) {
	l.log(v, LevelDebug, "")
}

func (l logger) Info(v interface{}) {
	l.log(v, LevelInfo, "")
}

func (l logger) Notice(v interface{}) {
	l.log(v, LevelNotice, "")
}

func (l logger) Warn(v interface{}) {
	l.log(v, LevelWarning, "")
}

func (l logger) Error(v interface{}) {
	l.log(v, LevelError, "")
}

func (l logger) Critical(v interface{}) {
	l.log(v, LevelCritical, "")
}

func (l logger) Alert(v interface{}) {
	l.log(v, LevelAlert, "")
}

func (l logger) Emergency(v interface{}) {
	l.log(v, LevelEmergency, "")
}

func Debug(v interface{}) {
	instance.Debug(v)
}

func Info(v interface{}) {
	instance.Info(v)
}

func Notice(v interface{}) {
	instance.Notice(v)
}

func Warn(v interface{}) {
	instance.Warn(v)
}

func Error(v interface{}) {
	instance.Error(v)
}

func Critical(v interface{}) {
	instance.Critical(v)
}

func Alert(v interface{}) {
	instance.Alert(v)
}

func Emergency(v interface{}) {
	instance.Emergency(v)
}
