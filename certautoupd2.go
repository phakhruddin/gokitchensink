package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type S3EventRecord struct {
	S3 struct {
		Bucket struct {
			Name string `json:"name"`
		} `json:"bucket"`
		Object struct {
			Key string `json:"key"`
		} `json:"object"`
	} `json:"s3"`
}

func handleRequest(ctx context.Context, s3Event events.S3Event) error {
	for _, record := range s3Event.Records {
		s3Record := S3EventRecord{}
		if err := json.Unmarshal([]byte(record.S3), &s3Record); err != nil {
			return fmt.Errorf("failed to unmarshal S3 event record: %v", err)
		}
		bucket := s3Record.S3.Bucket.Name
		key := s3Record.S3.Object.Key
		log.Printf("S3 bucket: %s, key: %s", bucket, key)
		if err := executeSSMCommand(); err != nil {
			return fmt.Errorf("failed to execute SSM command: %v", err)
		}
	}
	return nil
}
func executeSSMCommand() error {
	sess := session.Must(session.NewSession())
	ec2Svc := ec2.New(sess)
	ssmSvc := ssm.New(sess)
	// Describe instances with the specific tag
	describeInstancesInput := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:intuit:app"),
				Values: []*string{aws.String("spl")},
			},
		},
	}
	result, err := ec2Svc.DescribeInstances(describeInstancesInput)
	if err != nil {
		return fmt.Errorf("failed to describe instances: %v", err)
	}
	var instanceIDs []*string
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instanceIDs = append(instanceIDs, instance.InstanceId)
		}
	}
	if len(instanceIDs) == 0 {
		log.Println("No instances found with the specified tag")
		return nil
	}
	// Execute SSM command
	sendCommandInput := &ssm.SendCommandInput{
		DocumentName: aws.String("AWS-RunShellScript"),
		Parameters: map[string][]*string{
			"commands": {
				aws.String("echo 'Hello from SSM'"), // Replace with your desired command
			},
		},
		InstanceIds: instanceIDs,
	}
	_, err = ssmSvc.SendCommand(sendCommandInput)
	if err != nil {
		return fmt.Errorf("failed to send SSM command: %v", err)
	}
	log.Println("SSM command executed successfully")
	return nil
}
func main() {
	lambda.Start(handleRequest)
}
