package opendevopslambda

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"testing"
)

type mockedGetItem struct {
	dynamodbiface.DynamoDBAPI
	Response dynamodb.GetItemOutput
}

func (d mockedGetItem) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return &d.Response, nil
}

func TestHandler(t *testing.T) {
	t.Run("Successful Request", func(t *testing.T) {
		mgi := mockedGetItem {
			Response: dynamodb.GetItemOutput{},
		}

		d := Dependency{
			DepDynamoDB: mgi,
		}

		qsp := map[string]string{}
		qsp["imageId"] = "33ba5573-4a2f-4367-a430-0fed005719b1"

		request := events.APIGatewayProxyRequest{
			QueryStringParameters: qsp,
		}

		_, err := d.Handler(request)
		if err != nil {
			t.Fatal(fmt.Sprintf("TestHandler failed with %s", err.Error()))
		}
	})
}