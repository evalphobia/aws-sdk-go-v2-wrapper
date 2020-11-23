package pinpointemail

import (
	"context"
)

// XSendEmailText sends text type email.
func (svc *PinpointEmail) XSendEmailText(ctx context.Context, subject, body, from string, to ...string) error {
	return svc.xSendEmail(ctx, false, subject, body, from, to...)
}

// XSendEmailHTML sends HTML type email.
func (svc *PinpointEmail) XSendEmailHTML(ctx context.Context, subject, body, from string, to ...string) error {
	return svc.xSendEmail(ctx, true, subject, body, from, to...)
}

func (svc *PinpointEmail) xSendEmail(ctx context.Context, useHTML bool, subject, body, from string, to ...string) error {
	_, err := svc.SendEmail(ctx, SendEmailRequest{
		Subject: subject,
		Body:    body,
		From:    from,
		To:      to,
		UseHTML: useHTML,
	})
	return err
}
