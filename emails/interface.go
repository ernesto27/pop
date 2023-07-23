package emails

type ServiceEmail interface {
	SendEmail(to []string, from, subject, body string, paths []string) error
}
