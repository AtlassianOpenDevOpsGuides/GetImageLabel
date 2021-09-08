package main

import (
	"log"
	"os"
	"get-image-label/opendevopslambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	sess := session.Must(session.NewSession())

	d := opendevopslambda.Dependency{
		DepDynamoDB: dynamodb.New(sess),
	}

	lambda.Start(d.Handler)
}
