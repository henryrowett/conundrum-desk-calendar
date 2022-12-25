package sms

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
	"strings"
)

type AWS struct{}

// uses aws sns service to send sms messages
// the following env variables are required
// to be set
// - AWS_ACCESS_KEY_ID
// - AWS_REGION
// - AWS_SDK_LOAD_CONFIG
// - AWS_SECRET_ACCESS_KEY
func (a *AWS) Send(body string) error {
	sess := session.Must(session.NewSession())

	svc := sns.New(sess, &aws.Config{
		Region: aws.String("eu-west-2"),
	})

	numbers := strings.Split(os.Getenv("NUMBERS"), ",")

	for _, n := range numbers {
		req := &sns.PublishInput{
			Message:     aws.String(body),
			PhoneNumber: aws.String(n),
		}
		rsp, err := svc.Publish(req)
		if err != nil {
			return err
		}

		fmt.Println(rsp.String())
	}

	return nil
}
