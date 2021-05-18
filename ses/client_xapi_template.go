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

// XGetTemplate gets an email template.
func (svc *SES) XGetTemplate(ctx context.Context, templateName string) (*GetTemplateResult, error) {
	return svc.GetTemplate(ctx, GetTemplateRequest{
		TemplateName: templateName,
	})
}

// XUpdateTemplate updates an email template.
func (svc *SES) XUpdateTemplate(ctx context.Context, templateName, subject, htmlBody, textBody string) error {
	_, err := svc.UpdateTemplate(ctx, UpdateTemplateRequest{
		Template: Template{
			TemplateName: templateName,
			SubjectPart:  subject,
			HTMLPart:     htmlBody,
			TextPart:     textBody,
		},
	})
	return err
}
