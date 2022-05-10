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

import "fmt"

// ErrorType - Easy Gate error type
type ErrorType string

// Easy Gate errors enum
const (
	InvalidFormat ErrorType = "format"
	InvalidIcon   ErrorType = "icon"
	InvalidURL    ErrorType = "url"
	InvalidColor  ErrorType = "color"
)

// ErrorElement - Easy Gate error element
type ErrorElement string

// Easy Gate error context enum
const (
	Root              ErrorElement = "root"
	Service           ErrorElement = "service"
	ConfigurationFile ErrorElement = "configuration file"
)

// EasyGateError - Easy Gate Error struct
type EasyGateError struct {
	ErrorType ErrorType
	Element   ErrorElement
	Name      string
}

// NewEasyGateError - Create a new Easy Gate error
func NewEasyGateError(errorType ErrorType, element ErrorElement, name string) error {
	return EasyGateError{errorType, element, name}
}

func (e EasyGateError) Error() string {
	message := fmt.Sprintf("Invalid %s for %s element",
		e.ErrorType, e.Element)
	if e.Name != "" {
		message = fmt.Sprintf("%s: %s", message, e.Name)
	}
	return message
}
