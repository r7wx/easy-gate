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

package service

import (
	"net"

	"github.com/r7wx/easy-gate/internal/models"
)

func isAllowed(groups []models.Group, allowedGroups []string, addr string) bool {
	if len(allowedGroups) == 0 {
		return true
	}

	for _, allowedGroup := range allowedGroups {
		for _, group := range groups {
			if group.Name != allowedGroup {
				continue
			}

			_, groupNet, err := net.ParseCIDR(group.Subnet)
			if err != nil {
				continue
			}

			ipAddr := net.ParseIP(addr)
			if groupNet.Contains(ipAddr) {
				return true
			}
		}
	}

	return false
}
