package note

// Note - Note model
type Note struct {
	Name   string   `json:"name"`
	Text   string   `json:"text"`
	Groups []string `json:"-"`
}
