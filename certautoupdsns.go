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

type SNSMessage struct {
	Records []struct {
		S3 struct {
			Object struct {
				Key string `json:"key"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

func handler(ctx context.Context, snsEvent events.SNSEvent) error {
	sess := session.Must(session.NewSession())
	ssmSvc := ssm.New(sess)
	ec2Svc := ec2.New(sess)

	for _, record := range snsEvent.Records {
		log.Printf("Processing SNS message record: %v\n", record)

		// Extract the S3 key from the message payload
		var snsMsg SNSMessage
		if err := json.Unmarshal([]byte(record.SNS.Message), &snsMsg); err != nil {
			log.Printf("Failed to unmarshal SNS message: %v\n", err)
			continue
		}
		key := snsMsg.Records[0].S3.Object.Key
		log.Printf("S3 key=%s\n", key)

		input := &ssm.SendCommandInput{
			DocumentName: aws.String("AWS-RunRemoteScript"),
			InstanceIds:  []*string{},
			Parameters: map[string][]*string{
				"sourceType":       {aws.String("S3")},
				"sourceInfo":       {aws.String(fmt.Sprintf("{\"path\":\"s3://bucket/%s\"}", key))}, // replace 'bucket' with your own S3 bucket name
				"commandLine":      {aws.String("sh /path/to/your/script")},
				"executionTimeout": {aws.String("1800")},
			},
			Comment:            aws.String("Execute script from S3"),
			OutputS3BucketName: aws.String("bucket"), // replace 'bucket' with your own S3 bucket name
			OutputS3KeyPrefix:  aws.String("ssm-output/"),
		}

		// Get instances matching the tag "intuit:app","spl"
		tagName := "intuit:app"
		tagValue := "spl"
		filters := []*ec2.Filter{
			{
				Name:   aws.String("tag:" + tagName),
				Values: []*string{aws.String(tagValue)},
			},
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running")},
			},
		}

		// Get the list of matching instances
		result, err := ec2Svc.DescribeInstances(&ec2.DescribeInstancesInput{
			Filters: filters,
		})
		if err != nil {
			log.Printf("Failed to get instances with tag=%s, value=%s: %v\n", tagName, tagValue, err)
			continue
		}

		if len(result.Reservations) == 0 {
			log.Printf("No instances found with tag=%s, value=%s\n", tagName, tagValue)
		} else {
			// Extract instance IDs
			for _, reservation := range result.Reservations {
				for _, instance := range reservation.Instances {
					input.InstanceIds = append(input.InstanceIds, instance.InstanceId)
				}
			}

			// Run SSM command
			log.Printf("Running SSM command on instances with tag=%s, value=%s\n", tagName, tagValue)
			_, err := ssmSvc.SendCommand(input)
			if err != nil {
				log.Printf("Failed to run SSM command: %v\n", err)
				continue
			}
		}
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
