package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const TO_SEPARATOR = ","

// sendEmailSuccessMsg is the tea.Msg handled by Bubble Tea when the email has
// been sent successfully.
type sendEmailSuccessMsg struct{}

// sendEmailFailureMsg is the tea.Msg handled by Bubble Tea when the email has
// failed to send.
type sendEmailFailureMsg error

// sendEmailCmd returns a tea.Cmd that sends the email.
func (m Model) sendEmailCmd() tea.Cmd {
	return func() tea.Msg {
		attachments := make([]string, len(m.Attachments.Items()))
		for i, a := range m.Attachments.Items() {
			attachments[i] = a.FilterValue()
		}
		err := m.serviceEmail.SendEmail(strings.Split(m.To.Value(), TO_SEPARATOR), m.From.Value(), m.Subject.Value(), m.Body.Value(), attachments)
		if err != nil {
			return sendEmailFailureMsg(err)
		}
		return sendEmailSuccessMsg{}
	}
}
