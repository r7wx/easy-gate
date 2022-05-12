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

package errors

import "testing"

func TestErrors(t *testing.T) {
	err := NewEasyGateError(InvalidIcon, Root, "")
	if err.Error() != "Invalid icon for root" {
		t.Fatalf("Expected 'Invalid icon for root', got '%s'",
			err.Error())
	}

	err = NewEasyGateError(InvalidIcon, Service, "service1")
	if err.Error() != "Invalid icon for service: service1" {
		t.Fatalf("Expected 'Invalid icon for service: service1', got '%s'",
			err.Error())
	}

	err = NewEasyGateError(InvalidURL, Service, "service1")
	if err.Error() != "Invalid url for service: service1" {
		t.Fatalf("Expected 'Invalid url for service: service1', got '%s'",
			err.Error())
	}

	err = NewEasyGateError(InvalidColor, Root, "background")
	if err.Error() != "Invalid color for root: background" {
		t.Fatalf("Expected 'Invalid color for root: background', got '%s'",
			err.Error())
	}

	err = NewEasyGateError(InvalidFormat, ConfigurationFile, "")
	if err.Error() != "Invalid format for configuration file" {
		t.Fatalf("Expected 'Invalid format for configuration file', got '%s'",
			err.Error())
	}
}
