package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	lambda.Start(handleS3Event)
}

// handleS3Event is the main Lambda handler function
func handleS3Event(ctx context.Context, event events.S3Event) error {
	// Create an EC2 and SSM client using the default session
	ec2Client := ec2.New(session.Must(session.NewSession()))
	ssmClient := ssm.New(session.Must(session.NewSession()))

	// Loop over each S3 event record
	for _, record := range event.Records {
		// Get the S3 bucket and key from the event record
		s3Bucket := record.S3.Bucket.Name
		s3Key := record.S3.Object.Key

		// Find the instances with the tag "intuit:app" equals to "spl"
		instancesWithSplTag, err := getInstancesWithSplTag(ctx, ec2Client)
		if err != nil {
			return err
		}

		// Execute the SSM command on the instances
		for _, instance := range instancesWithSplTag {
			// Define the SSM command and parameters
			documentName := "<ssm-document-name>" // Replace with your SSM document name
			commandInput := map[string][]*string{
				"Bucket": []*string{aws.String(s3Bucket)},
				"Key":    []*string{aws.String(s3Key)},
			}

			// Execute the SSM command using the instance ID and command input
			_, err := ssmClient.SendCommandWithContext(ctx, &ssm.SendCommandInput{
				InstanceIds:  []*string{aws.String(*instance.InstanceId)},
				DocumentName: aws.String(documentName),
				Parameters:   commandInput,
			})
			if err != nil {
				return err
			}
		}
	}

	// Return nil to indicate success
	return nil
}

// getInstancesWithSplTag returns a list of instances that have the tag "intuit:app" equals to "spl"
func getInstancesWithSplTag(ctx context.Context, ec2Client *ec2.EC2) ([]*ec2.Instance, error) {
	// Define the filters to find instances with the "intuit:app" tag equals to "spl"
	filters := []*ec2.Filter{
		{
			Name:   aws.String("tag:intuit:app"),
			Values: []*string{aws.String("spl")},
		},
	}

	// Get the instance IDs that match the filters
	instancesOutput, err := ec2Client.DescribeInstancesWithContext(ctx, &ec2.DescribeInstancesInput{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}

	// Create a slice to hold the instances
	var instances []*ec2.Instance

	// Loop over each reservation in the output to get the instances
	for _, reservation := range instancesOutput.Reservations {
		for _, instance := range reservation.Instances {
			if instance.State.Name != nil && strings.ToLower(*instance.State.Name) == "running" {
				instances = append(instances, instance)
			}
		}
	}

	return instances, nil
}
