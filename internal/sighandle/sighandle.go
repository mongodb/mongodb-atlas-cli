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

package sighandle

import (
	"os"
	"os/signal"
)

type handler struct {
	sig      []os.Signal
	f        func(os.Signal)
	c        chan os.Signal
	notified bool
}

func (h *handler) routine() {
	h.f(<-h.c)
}

func (h *handler) notify() {
	if h.notified {
		h.reset()
	}
	h.notified = true
	h.c = make(chan os.Signal, 1)
	go h.routine()
	signal.Notify(h.c, h.sig...)
}

func (h *handler) reset() {
	signal.Reset(h.sig...)
	h.notified = false
}

var std = &handler{}

func Notify(f func(os.Signal), sig ...os.Signal) {
	std.f = f
	std.sig = sig
	std.notify()
}

func Reset() {
	std.reset()
}
