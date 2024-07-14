package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	s3buckets := flag.Args()

	fmt.Printf("String: %+s\n", s3buckets[len(s3buckets)-1])
	fmt.Printf("String Slice: %+q\n", s3buckets)

	for _, element := range s3buckets {

		//add present element to sum
		fmt.Printf("Checking content of: %+s\n", element)

		// Create an Amazon S3 service client
		client := s3.NewFromConfig(cfg)

		// Get the first page of results for s3buckets last value
		output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			//Bucket: aws.String(s3buckets[len(s3buckets)-1]),
			Bucket: aws.String(element),
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Println("element results:")
		for _, object := range output.Contents {
			log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
		}
	}

}
