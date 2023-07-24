package emails

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"

	"github.com/charmbracelet/pop/types"
	"github.com/mailjet/mailjet-apiv3-go"
	. "github.com/mailjet/mailjet-apiv3-go"
)

type Mailjet struct {
	apiKeyPublic  string
	apiKeyPrivate string
}

func NewMailjet(apiKeyPublic string, apiKeyPrivate string) Mailjet {
	return Mailjet{
		apiKeyPublic:  apiKeyPublic,
		apiKeyPrivate: apiKeyPrivate,
	}
}

func (m Mailjet) SendEmail(to []string, from, subject, body string, paths []string) error {
	if len(to) == 0 {
		return errors.New("to is empty")
	}

	mailjetClient := NewMailjetClient(m.apiKeyPublic, m.apiKeyPrivate)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: from,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to[0],
				},
			},
			Subject:  subject,
			TextPart: body,
		},
	}

	attachment := &mailjet.AttachmentV31{}
	if len(paths) > 0 {
		ap := m.MakeAttachments(paths)
		for _, p := range ap {
			// TODO - get content type from file
			attachment.ContentType = "text/plain"
			attachment.Filename = p.Filename
			attachment.Base64Content = base64.StdEncoding.EncodeToString([]byte(p.Content))
		}

		// TODO - support multiple attachments
		messagesInfo[0].Attachments = &mailjet.AttachmentsV31{*attachment}
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		return err
	}
	return nil
}

func (m Mailjet) MakeAttachments(paths []string) []types.Attachment {
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
