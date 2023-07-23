package emails

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/resendlabs/resend-go"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	renderer "github.com/yuin/goldmark/renderer/html"
)

type Resend struct {
	apiKey string
	unsafe bool
}

func NewResend(apiKey string, unsafe bool) Resend {
	return Resend{
		apiKey: apiKey,
		unsafe: unsafe,
	}
}

func (r Resend) SendEmail(to []string, from, subject, body string, paths []string) error {
	client := resend.NewClient(r.apiKey)

	html := bytes.NewBufferString("")
	// If the conversion fails, we'll simply send the plain-text body.
	if r.unsafe {
		markdown := goldmark.New(
			goldmark.WithRendererOptions(
				renderer.WithUnsafe(),
			),
			goldmark.WithExtensions(
				extension.Strikethrough,
				extension.Table,
				extension.Linkify,
			),
		)
		_ = markdown.Convert([]byte(body), html)
	} else {
		_ = goldmark.Convert([]byte(body), html)
	}

	request := &resend.SendEmailRequest{
		From:        from,
		To:          to,
		Subject:     subject,
		Html:        html.String() + "dfsfsfdsfs",
		Text:        body,
		Attachments: r.MakeAttachments(paths),
	}

	_, err := client.Emails.Send(request)
	if err != nil {
		return err
	}

	return nil
}

func (r Resend) MakeAttachments(paths []string) []resend.Attachment {
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
