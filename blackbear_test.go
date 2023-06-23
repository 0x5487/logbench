// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"io"
	"strings"

	"github.com/nite-coder/blackbear/pkg/log"
)

const _1b = "a"

var (
	_200b = strings.Repeat(_1b, 200)
	_10kb = strings.Repeat(_1b, 10*1024)

	shortText = _200b
	longText  = _10kb
)

func fakeBlackBearLogContext(c log.Context) log.Context {
	return c.
		Int("int", _tenInts[0]).
		Ints("ints", _tenInts).
		Str("string", _tenStrings[0]).
		Strs("strings", _tenStrings).
		Time("time", _tenTimes[0]).
		Times("times", _tenTimes).
		Any("user1", _oneUser).
		Any("user2", _oneUser).
		Any("users", _tenUsers).
		Err(errExample)
}

func fakeBlackBearFields(e *log.Entry) *log.Entry {
	return e.
		Int("int", _tenInts[0]).
		Ints("ints", _tenInts).
		Str("string", _tenStrings[0]).
		Strs("strings", _tenStrings).
		Time("time", _tenTimes[0]).
		Times("times", _tenTimes).
		Any("user1", _oneUser).
		Any("user2", _oneUser).
		Any("users", _tenUsers).
		Err(errExample)
}

func newBlackbearLog() *log.Logger {
	opts := log.HandlerOptions{
		Level:       log.DebugLevel,
		DisableTime: true,
	}
	return log.New(log.NewJSONHandler(io.Discard, &opts))
}

func newDisabledBlackbearLog() *log.Logger {
	opts := log.HandlerOptions{
		Level:       log.FatalLevel,
		DisableTime: true,
	}
	return log.New(log.NewJSONHandler(io.Discard, &opts))
}
