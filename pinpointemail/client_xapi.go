package pinpointemail

import (
	"context"
)

// XSendEmailText sends text type email.
func (svc *PinpointEmail) XSendEmailText(ctx context.Context, subject, body, from string, to ...string) error {
	return svc.xSendEmail(ctx, subject, "", body, from, to...)
}

// XSendEmailHTML sends HTML type email.
func (svc *PinpointEmail) XSendEmailHTML(ctx context.Context, subject, body, from string, to ...string) error {
	return svc.xSendEmail(ctx, subject, body, "", from, to...)
}

func (svc *PinpointEmail) xSendEmail(ctx context.Context, subject, htmlBody, textBody, from string, to ...string) error {
	_, err := svc.SendEmail(ctx, SendEmailRequest{
		Subject:  subject,
		HTMLBody: htmlBody,
		From:     from,
		To:       to,
		TextBody: textBody,
	})
	return err
}
