package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type SES struct {
	svc *ses.SES
}

func NewSES(AccessKeyId, SecretAccessKey, Region string) *SES {
	awsSession := session.New(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKeyId, SecretAccessKey, ""),
	})
	if awsSession == nil {
		return nil
	}
	svc := ses.New(awsSession)
	if svc == nil {
		return nil
	}
	return &SES{svc: svc}
}

func (s *SES) Send(to, from, sub, body string) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(sub),
			},
		},
		Source: aws.String(from),
	}
	_, err := s.svc.SendEmail(input)
	return err
}
