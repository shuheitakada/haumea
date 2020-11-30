package elbv2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

type Client struct {
	svc *elbv2.ELBV2
}

func NewClient(role string) *Client {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	var serviceClientValue *elbv2.ELBV2
	if role == "" {
		serviceClientValue = elbv2.New(sess)
	} else {
		creds := stscreds.NewCredentials(sess, role)
		serviceClientValue = elbv2.New(sess, &aws.Config{Credentials: creds})
	}
	return &Client{svc: serviceClientValue}
}

func (client *Client) DescribeTargetHealth(targetGroupArn string) {
	input := &elbv2.DescribeTargetHealthInput{TargetGroupArn: aws.String(targetGroupArn)}
	result, err := client.svc.DescribeTargetHealth(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}

func (client *Client) DescribeAllTargetHealth(targetGroupArns []string) {
	for _, targetGroupArn := range targetGroupArns {
		client.DescribeTargetHealth(targetGroupArn)
	}
}
