package theme

// Theme - Easy Gate theme model
type Theme struct {
	Background string `json:"background" yaml:"background"`
	Foreground string `json:"foreground" yaml:"foreground"`
	CustomCSS  string `json:"custom_css" yaml:"custom_css"`
}
