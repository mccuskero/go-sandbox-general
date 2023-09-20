package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// use awscli to setup session
// ➜  ~ aws configure sso
// SSO session name (Recommended): my-sso
// SSO start URL [None]: https://d-9067a8ef02.awsapps.com/start
// SSO region [None]: us-east-1
// SSO registration scopes [sso:account:access]:

// ➜  go-sandbox-general git:(main) ✗ export AWS_DEFAULT_PROFILE=AdministratorAccess-773577149418
// ➜  go-sandbox-general git:(main) ✗ go run ./cmd/aws/main-s3.go                                
// Let's list up to 10 buckets for your account.
//        mccuskerowen-test-bucket
func main() {
	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration")
		fmt.Println(err)
		return
	}
	s3Client := s3.NewFromConfig(awsConfig)
	count := 10
	fmt.Printf("Listing up to %v buckets.\n", count)
	result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Printf("Couldn't list buckets, error: %v\n", err)
		return
	}
	if len(result.Buckets) == 0 {
		fmt.Println("No buckets found")
		return
	} 
	
	if count > len(result.Buckets) {
		count = len(result.Buckets)
	}
	
	for _, bucket := range result.Buckets[:count] {
		fmt.Printf("\t%v\n", *bucket.Name)
	}
}