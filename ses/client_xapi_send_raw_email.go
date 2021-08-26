package ses

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"strings"
)

// XSendRawEmail sends raw email.
func (svc *SES) XSendRawEmail(ctx context.Context, r XSendRawEmailRequest) error {
	req, err := r.ToRequest()
	if err != nil {
		return err
	}
	if req.ConfigurationSetName == "" {
		req.ConfigurationSetName = svc.defaultConfigurationSet
	}

	_, err = svc.SendRawEmail(ctx, req)
	return err
}

// XSendRawEmailRequest has parameters for `XSendRawEmail` function.
type XSendRawEmailRequest struct {
	From            string
	To              []string
	ReturnPath      string
	Subject         string
	TextBody        string
	TextCharset     string
	HTMLBody        string
	HTMLCharset     string
	ContentLanguage string
	MIMEVersion     string
	Attachments     []XAttachment

	// optional
	ConfigurationSetName string
	FromARN              string
	ReturnPathARN        string
	SourceARN            string
	Tags                 []MessageTag
}

// ToRequest converts to SendRawEmailRequest.
// most header codes are based on,
//   - https://gist.github.com/carelvwyk/60100f2421c6284391d08374bc887dca
//   - https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-raw.html
func (r XSendRawEmailRequest) ToRequest() (SendRawEmailRequest, error) {
	req := SendRawEmailRequest{
		ConfigurationSetName: r.ConfigurationSetName,
		Destinations:         r.To,
		FromARN:              r.FromARN,
		ReturnPathARN:        r.ReturnPathARN,
		Source:               r.From,
		SourceARN:            r.SourceARN,
		Tags:                 r.Tags,
	}

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// email header part
	h := make(textproto.MIMEHeader)
	if r.From != "" {
		h.Set("From", r.From)
	}
	if len(r.To) != 0 {
		h.Set("To", strings.Join(r.To, ", "))
	}
	switch {
	case r.ReturnPath != "":
		h.Set("Return-Path", r.ReturnPath)
	case r.From != "":
		h.Set("Return-Path", r.From)
	}

	if r.Subject != "" {
		h.Set("Subject", r.Subject)
	}

	if r.ContentLanguage != "" {
		h.Set("Content-Language", r.ContentLanguage)
	}

	if r.MIMEVersion != "" {
		h.Set("MIME-Version", r.MIMEVersion)
	}

	h.Set("Content-Type", "multipart/mixed; boundary=\""+writer.Boundary()+"\"")
	_, err := writer.CreatePart(h)
	if err != nil {
		return req, err
	}

	// text body part
	if r.TextBody != "" {
		h := make(textproto.MIMEHeader)
		charset := r.TextCharset
		if charset == "" {
			charset = "utf-8"
		}
		h.Set("Content-Type", fmt.Sprintf("text/plain; charset=%s", charset))
		part, err := writer.CreatePart(h)
		if err != nil {
			return req, err
		}
		_, err = part.Write([]byte(r.TextBody))
		if err != nil {
			return req, err
		}
	}

	// html body part
	if r.HTMLBody != "" {
		h := make(textproto.MIMEHeader)
		charset := r.HTMLCharset
		if charset == "" {
			charset = "utf-8"
		}
		h.Set("Content-Type", fmt.Sprintf("text/html; charset=%s", charset))
		part, err := writer.CreatePart(h)
		if err != nil {
			return req, err
		}
		_, err = part.Write([]byte(r.HTMLBody))
		if err != nil {
			return req, err
		}
	}

	// attachment part
	for _, v := range r.Attachments {
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, v.Filename))
		mimeType := v.MIMEType
		if mimeType == "" {
			mimeType = "text/plain"
		}
		h.Set("Content-Type", fmt.Sprintf(`%s; name="%s"`, mimeType, v.Filename))
		h.Set("Content-Description", v.Filename)
		if v.ContentTransferEncoding != "" {
			h.Set("Content-Transfer-Encoding", v.ContentTransferEncoding)
		}
		part, err := writer.CreatePart(h)
		if err != nil {
			return req, err
		}
		_, err = part.Write(v.Data)
		if err != nil {
			return req, err
		}
		err = writer.Close()
		if err != nil {
			return req, err
		}
	}

	s := buf.String()
	ss := strings.SplitN(s, "\n", 2)
	if len(ss) < 2 {
		return req, fmt.Errorf("invalid e-mail content: [%+v]", ss)
	}
	req.RawMessageData = []byte(ss[1])
	return req, nil
}

type XAttachment struct {
	Data     []byte
	Filename string
	MIMEType string
	// if 'ContentTransferEncoding' is "base64", then 'Data' must be base64 encoded.
	ContentTransferEncoding string
}
