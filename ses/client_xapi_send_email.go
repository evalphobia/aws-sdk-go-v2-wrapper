package ses

import (
	"context"
)

// XSendEmailText sends text type email.
func (svc *SES) XSendEmailText(ctx context.Context, subject, body, from string, to ...string) error {
	return svc.xSendEmail(ctx, subject, "", body, from, to...)
}

// XSendEmailHTML sends HTML type email.
func (svc *SES) XSendEmailHTML(ctx context.Context, subject, body, from string, to ...string) error {
	return svc.xSendEmail(ctx, subject, body, "", from, to...)
}

func (svc *SES) xSendEmail(ctx context.Context, subject, htmlBody, textBody, from string, to ...string) error {
	_, err := svc.SendEmail(ctx, SendEmailRequest{
		Destination: Destination{
			ToAddresses: to,
		},
		Source:               from,
		Subject:              subject,
		HTMLBody:             htmlBody,
		TextBody:             textBody,
		ConfigurationSetName: svc.defaultConfigurationSet,
	})
	return err
}
