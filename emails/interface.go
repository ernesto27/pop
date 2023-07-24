package emails

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/pop/types"
)

type ServiceEmail interface {
	SendEmail(to []string, from, subject, body string, paths []string) error
}

func MakeAttachments(paths []string) []types.Attachment {
	if len(paths) == 0 {
		return nil
	}

	attachments := make([]types.Attachment, len(paths))
	for i, a := range paths {
		f, err := os.ReadFile(a)
		if err != nil {
			continue
		}
		attachments[i] = types.Attachment{
			Content:  string(f),
			Filename: filepath.Base(a),
		}
	}

	return attachments
}
