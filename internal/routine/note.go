package routine

import (
	"github.com/r7wx/easy-gate/internal/config"
	"github.com/r7wx/easy-gate/internal/note"
)

func (r *Routine) getNotes(cfg *config.Config) []note.Note {
	notes := []note.Note{}
	for _, cfgNote := range cfg.Notes {
		notes = append(notes, note.Note{
			Name:   cfgNote.Name,
			Text:   cfgNote.Text,
			Groups: cfgNote.Groups,
		})
	}

	return notes
}
