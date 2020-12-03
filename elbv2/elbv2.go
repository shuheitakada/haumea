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

func (client *Client) RegisterTargets(targetGroupArn string, targets []interface{}) {
	input := &elbv2.RegisterTargetsInput{
		TargetGroupArn: aws.String(targetGroupArn),
		Targets: targetEC2IDs(targets),
	}
	result, err := client.svc.RegisterTargets(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}

func (client *Client) DeregisterTargets(targetGroupArn string, targets []interface{}) {
	input := &elbv2.DeregisterTargetsInput{
		TargetGroupArn: aws.String(targetGroupArn),
		Targets: targetEC2IDs(targets),
	}
	result, err := client.svc.DeregisterTargets(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}

func targetEC2IDs(targets []interface{}) (result []*elbv2.TargetDescription) {
	for _, target := range targets {
		targetEC2ID := &elbv2.TargetDescription{Id: aws.String(target.(string))}
		result = append(result, targetEC2ID)
	}
	return
}
