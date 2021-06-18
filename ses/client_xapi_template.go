package ses

import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
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

// XExistsTemplate checks if the template already exists or not.
func (svc *SES) XExistsTemplate(ctx context.Context, templateName string) (bool, error) {
	_, err := svc.GetTemplate(ctx, GetTemplateRequest{
		TemplateName: templateName,
	})

	if e, ok := err.(errors.ErrorData); ok {
		if e.GetAWSErrCode() == "TemplateDoesNotExist" {
			return false, nil
		}
	}
	if err != nil {
		return false, err
	}
	return true, nil
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
