package emails

import (
	"encoding/base64"
	"errors"

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

	attachment := []mailjet.AttachmentV31{}
	if len(paths) > 0 {
		ap := MakeAttachments(paths)
		for _, p := range ap {
			attachment = append(attachment, mailjet.AttachmentV31{
				ContentType:   "text/plain",
				Filename:      p.Filename,
				Base64Content: base64.StdEncoding.EncodeToString([]byte(p.Content)),
			})
		}

		messagesInfo[0].Attachments = (*AttachmentsV31)(&attachment)
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		return err
	}

	return nil
}
