package ses

import (
	"context"
)

// XCreateTemplate creates an email template.
func (svc *SES) XCreateTemplate(ctx context.Context, templateName, subject, htmlBody, textBody string) error {
	_, err := svc.CreateTemplate(ctx, CreateTemplateRequest{
		Template: Template{
			TemplateName: templateName,
			SubjectPart:  subject,
			HTMLPart:     htmlBody,
			TextPart:     textBody,
		},
	})
	return err
}
