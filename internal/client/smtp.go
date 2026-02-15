package client

import (
	"fmt"
	"net/smtp"
)

// SMTPClient sends emails via SMTP.
type SMTPClient struct {
	Host     string
	Port     string
	From     string
	Password string
	To       string
}

// NewSMTPClient creates a new SMTPClient.
func NewSMTPClient(host, port, from, password, to string) *SMTPClient {
	return &SMTPClient{
		Host:     host,
		Port:     port,
		From:     from,
		Password: password,
		To:       to,
	}
}

// Send sends an email with the given subject and plain-text body.
func (c *SMTPClient) Send(subject, body string) error {
	auth := smtp.PlainAuth("", c.From, c.Password, c.Host)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"UTF-8\"\r\n\r\n%s",
		c.From, c.To, subject, body)

	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)
	return smtp.SendMail(addr, auth, c.From, []string{c.To}, []byte(msg))
}
