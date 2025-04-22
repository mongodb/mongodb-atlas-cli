// Copyright 2022 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"fmt"
	"io"
	"os"
)

type Level int

const (
	NoneLevel Level = iota
	WarningLevel
	DebugLevel
)

type Logger interface {
	SetLevel(level Level)
	IsDebugLevel() bool
	IsWarningLevel() bool
	Debug(a ...any)
	Debugln(a ...any)
	Debugf(format string, a ...any)
	Warning(a ...any)
	Warningln(a ...any)
	Warningf(format string, a ...any)
}

type IOLogger struct {
	w     io.Writer
	level Level
}

func New(w io.Writer, l Level) *IOLogger {
	return &IOLogger{
		level: l,
		w:     w,
	}
}

func (l *IOLogger) Writer() io.Writer {
	return l.w
}

func (l *IOLogger) SetWriter(w io.Writer) {
	l.w = w
}

func (l *IOLogger) SetLevel(level Level) {
	l.level = level
}

func (l *IOLogger) Level() Level {
	return l.level
}

func (l *IOLogger) IsDebugLevel() bool {
	return l.level >= DebugLevel
}

func (l *IOLogger) IsWarningLevel() bool {
	return l.level >= WarningLevel
}

func (l *IOLogger) Debug(a ...any) {
	if !l.IsDebugLevel() {
		return
	}
	fmt.Fprint(l.w, a...)
}

func (l *IOLogger) Debugln(a ...any) {
	if !l.IsDebugLevel() {
		return
	}
	fmt.Fprintln(l.w, a...)
}

func (l *IOLogger) Debugf(format string, a ...any) {
	if !l.IsDebugLevel() {
		return
	}
	fmt.Fprintf(l.w, format, a...)
}

func (l *IOLogger) Warning(a ...any) {
	if !l.IsWarningLevel() {
		return
	}
	fmt.Fprint(l.w, a...)
}

func (l *IOLogger) Warningln(a ...any) {
	if !l.IsWarningLevel() {
		return
	}
	fmt.Fprintln(l.w, a...)
}

func (l *IOLogger) Warningf(format string, a ...any) {
	if !l.IsWarningLevel() {
		return
	}
	fmt.Fprintf(l.w, format, a...)
}

var std = New(os.Stderr, WarningLevel)

func Writer() io.Writer {
	return std.Writer()
}

func SetWriter(w io.Writer) {
	std.SetWriter(w)
}

func SetLevel(level Level) {
	std.SetLevel(level)
}

func IsDebugLevel() bool {
	return std.IsDebugLevel()
}

func IsWarningLevel() bool {
	return std.IsWarningLevel()
}

func Default() *IOLogger {
	return std
}

func Debug(a ...any) {
	std.Debug(a...)
}

func Debugln(a ...any) {
	std.Debugln(a...)
}

func Debugf(format string, a ...any) {
	std.Debugf(format, a...)
}

func Warning(a ...any) {
	std.Warning(a...)
}

func Warningln(a ...any) {
	std.Warningln(a...)
}

func Warningf(format string, a ...any) {
	std.Warningf(format, a...)
}
