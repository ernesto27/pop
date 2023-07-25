package emails

import (
	"strconv"

	"gopkg.in/mail.v2"
)

type Smtp struct {
	host     string
	port     int
	username string
	password string
}

func NewSmtp(host string, port string, username string, password string) (Smtp, error) {
	p, err := strconv.Atoi(port)
	if err != nil {
		return Smtp{}, err
	}

	return Smtp{
		host:     host,
		port:     p,
		username: username,
		password: password,
	}, nil
}

func (s Smtp) SendEmail(to []string, from, subject, body string, paths []string) error {
	m := mail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	for _, p := range paths {
		m.Attach(p)
	}

	d := mail.NewDialer(s.host, s.port, s.username, s.password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
