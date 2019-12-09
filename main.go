package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	}))
	s3 := s3.New(sess)
	fmt.Println(s3.ListBuckets())
}
