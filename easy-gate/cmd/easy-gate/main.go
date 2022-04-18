/*
MIT License

Copyright (c) 2022 r7wx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"os"
	"time"

	"github.com/r7wx/easy-gate/internal/config"
	"github.com/r7wx/easy-gate/internal/service"
	"github.com/r7wx/easy-gate/internal/share"
)

func main() {
	share.PrintBanner()

	cfgFilePath := "easy-gate.json"
	args := os.Args[1:]
	if len(args) >= 1 {
		cfgFilePath = args[0]
	}

	cfgRoutine := config.NewRoutine(cfgFilePath,
		1*time.Minute)
	go cfgRoutine.Start()

	service := service.Service{
		ConfigRoutine: cfgRoutine,
	}
	service.Serve()
}
