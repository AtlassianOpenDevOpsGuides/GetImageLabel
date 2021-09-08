package opendevopslambda

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"net/url"
)

type ImageLabelTuple struct {
	Id string
	Label string
}

type Dependency struct {
	DepDynamoDB dynamodbiface.DynamoDBAPI
}

func (d *Dependency) processRequest(imageId string) (string, error) {
	dynamoResult, dynamoErr := d.DepDynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("ImageLabels"),
		Key: map[string]*dynamodb.AttributeValue{
		"Id": {
				S: aws.String(imageId),
			},
		},
	})

	if dynamoErr != nil {
		fmt.Printf("failed to get item from dynamodb %s\n", dynamoErr.Error())
	}

	imageLabelTuple := ImageLabelTuple{}

	dynamoErr = dynamodbattribute.UnmarshalMap(dynamoResult.Item, &imageLabelTuple)
	if dynamoErr != nil {
		fmt.Printf("failed to unmarshal record %s\n", dynamoErr.Error())
	}

	return imageLabelTuple.Label, nil
}

func (d* Dependency) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	imageId, found := request.QueryStringParameters["imageId"]
	if found {
		imageId, err := url.QueryUnescape(imageId)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 500,
				Body: `{"ImageLabel":"null"}`,
			}, err
		}

		processString, processErr := d.processRequest(imageId)
		return events.APIGatewayProxyResponse{StatusCode: 200,
			Body: fmt.Sprintf(`"ImageLabel":"%s"`, processString),
		}, processErr
	}

	return events.APIGatewayProxyResponse{StatusCode: 500,
		Body: `{"ImageLabel":"null"}`,
	}, errors.New("url parameter not found")
}
