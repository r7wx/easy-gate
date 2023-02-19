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
				Name: statusNote.Name,
				Text: statusNote.Text,
			}
			notes = append(notes, note)
		}
	}

	return notes
}
