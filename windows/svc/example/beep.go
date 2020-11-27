// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package main

import (
	"syscall"
)

const (
  simple      = 0xFFFFFFFF
  asterisk    = 0x00000040
  exclamation = 0x00000030
  err         = 0x00000010
  information = 0x00000040
  question    = 0x00000020
  warning     = 0x00000030
  ok          = 0x00000000
)

// BUG(brainman): MessageBeep Windows api is broken on Windows 7,
// so this example does not beep when runs as service on Windows 7.

var (
	beepFunc = syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
)

func beep() {
	beepFunc.Call(simple)
}
