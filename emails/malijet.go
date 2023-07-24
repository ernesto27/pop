package emails

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/mailjet/mailjet-apiv3-go"
	"github.com/resendlabs/resend-go"
)

type Malijet struct {
	apiKeyPublic  string
	apiKeyPrivate string
}

func NewMalijet() Malijet {
	return Malijet{
		apiKeyPublic:  "",
		apiKeyPrivate: "",
	}
}

func (m Malijet) SendEmail(to []string, from, subject, body string, paths []string) error {
	mailjetClient := NewMailjetClient(m.apiKeyPublic, m.apiKeyPrivate)

	repients := make([]Recipient, len(to))
	for i, r := range to {
		repients[i] = Recipient{
			Email: r,
		}
	}

	email := &InfoSendMail{
		FromEmail:  from,
		FromName:   "malijet",
		Subject:    subject,
		TextPart:   body,
		HTMLPart:   "",
		Recipients: repients,
	}

	res, err := mailjetClient.SendMail(email)
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}

func (m Malijet) MakeAttachments(paths []string) []resend.Attachment {
	if len(paths) == 0 {
		return nil
	}

	attachments := make([]resend.Attachment, len(paths))
	for i, a := range paths {
		f, err := os.ReadFile(a)
		if err != nil {
			continue
		}
		attachments[i] = resend.Attachment{
			Content:  string(f),
			Filename: filepath.Base(a),
		}
	}

	return attachments
}
