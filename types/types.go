package types

type Attachment struct {
	Content string `json:"content,omitempty"`

	Filename string `json:"filename"`

	Path string `json:"path,omitempty"`
}

type EmailParams struct {
	From        string
	To          []string
	Subject     string
	Body        string
	Text        string
	Attachments []Attachment
}
