/*
Copyright 2019 Doug Edgar.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mailer

import (
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	asession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/rhdedgar/emailer/datastores"
	"github.com/rhdedgar/emailer/localfile"
)

var (
	subjectString = fmt.Sprintf("Activate your new %v account", datastores.AppSecrets.AppName)
)

func ConfirmEmail(stringChan <-chan string) error {
	for {
		select {
		case url := <-stringChan:
			err := sendConfirmation(url)
			if err != nil {
				fmt.Printf("Error sending confirmation email: %v\n", err)
			}
		}
	}
}

func sendConfirmation(msg string) error {
	u, err := url.Parse(msg)
	if err != nil {
		return fmt.Errorf("URL received was not able to be parsed:%q %v\n", msg, err)
	}

	q := u.Query()
	email := q.Get("email")
	fmt.Println("sending to email: ", email)

	templateOptions := map[string]string{
		"ActivationLink": msg,
		"CompanyInfo":    datastores.AppSecrets.CompanyInfo,
		"AppName":        datastores.AppSecrets.AppName,
	}

	TextBody, err := localfile.GetTemplate(datastores.AppSecrets.EmailTemplate, templateOptions)
	if err != nil {
		return fmt.Errorf("Error getting confirmation email text template: %v\n", err)
	}

	HtmlBody, err := localfile.GetTemplate(datastores.AppSecrets.HTMLEmailTemplate, templateOptions)
	if err != nil {
		return fmt.Errorf("Error getting confirmation email HTML template: %v\n", err)
	}

	sess, err := asession.NewSession(&aws.Config{
		Region:      aws.String(datastores.AppSecrets.SESRegion),
		Credentials: credentials.NewStaticCredentials(datastores.AppSecrets.SESKeyID, datastores.AppSecrets.SESKey, ""),
	})

	if err != nil {
		return fmt.Errorf("Error getting confirmation email session: %v\n", err)
	}

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(datastores.CharSet),
					Data:    aws.String(HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(datastores.CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(datastores.CharSet),
				Data:    aws.String(subjectString),
			},
		},
		Source: aws.String(datastores.Sender),
	}

	result, err := svc.SendEmail(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("ses result: ", result)
	return nil
}
