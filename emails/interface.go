package emails

import "github.com/charmbracelet/pop/types"

type ServiceEmail interface {
	SendEmail(to []string, from, subject, body string, paths []string) error
	MakeAttachments(paths []string) []types.Attachment
}
