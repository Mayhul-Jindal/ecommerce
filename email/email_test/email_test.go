package email_test

import (
	"testing"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/email"
	"github.com/stretchr/testify/require"
)

func TestSendEmail(t *testing.T) {
	gmailSender := email.NewGmailSender(config.EMAIL_SENDER_NAME, config.EMAIL_SENDER_ADDRESS, config.EMAIL_SENDER_PASSWORD)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="https://www.balkan.id/">BalkanID</a></p>
	`
	to := []string{"mayhuljindal@gmail.com"}
	attachFiles := []string{"../../README.md"}

	err := gmailSender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
