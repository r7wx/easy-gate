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

package engine

import (
	"github.com/r7wx/easy-gate/internal/group"
	"github.com/r7wx/easy-gate/internal/note"
	"github.com/r7wx/easy-gate/internal/routine"
)

func getNotes(status *routine.Status, addr string) []note.Note {
	notes := []note.Note{}
	for _, statusNote := range status.Notes {
		if group.IsAllowed(status.Groups, statusNote.Groups, addr) {
			note := note.Note{
				Name:     statusNote.Name,
				Text:     statusNote.Text,
				Category: statusNote.Category,
			}
			notes = append(notes, note)
		}
	}

	return notes
}
