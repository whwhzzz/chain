package storage

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dynamo_con *dynamodb.DynamoDB

type Item struct {
	OfferId  string `json:"offer_id"`
	Content  string `json:"content"`
	DeadLine int64  `json:"dead_line"`
}

func AddPayload(svc *dynamodb.DynamoDB, offer_id string, content string) {
	item := Item{
		OfferId: offer_id,
		Content: content,
		// Save for 90 days.
		DeadLine: time.Now().Add(time.Hour * 24 * 90).Unix(),
	}
	av, err := dynamodbattribute.MarshalMap(item)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(UploadedFileTable),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully added")
}

func UpdatePayload(svc *dynamodb.DynamoDB, offer_id string, content string) {
	// Create item in table UploadedFileTable
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String(content),
			},
		},
		TableName: aws.String(UploadedFileTable),
		Key: map[string]*dynamodb.AttributeValue{
			"json": {
				S: aws.String(offer_id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set json = :r"),
	}

	_, err := svc.UpdateItem(input)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated")
}

func FetchPayload(svc *dynamodb.DynamoDB, offer_id string) (content string) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(UploadedFileTable),
		Key: map[string]*dynamodb.AttributeValue{
			"offer_id": {
				S: aws.String(offer_id),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	content = item.Content
	fmt.Println("Successfully fetch offer " + offer_id)
	return
}

func ConnectDB() (svc *dynamodb.DynamoDB) {
	if dynamo_con != nil {
		svc = dynamo_con
		return
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	if err != nil {
		fmt.Println("Error creating session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create DynamoDB client
	svc = dynamodb.New(sess)
	return
}
