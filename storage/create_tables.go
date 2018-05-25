package storage

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createDeposit(svc *dynamodb.DynamoDB) {
	// Create table Deposit
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("deposit"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		TableName: aws.String("Deposit"),
	}

	_, err := svc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling Deposit:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created Deposit")
}

func createTransaction(svc *dynamodb.DynamoDB) {
	// Create table Transaction
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("deposit"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("status"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("deadline"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("transaction_hash"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("offer_id"),
				KeyType:       aws.String("S"),
			},
			{
				AttributeName: aws.String("provider"),
				KeyType:       aws.String("S"),
			},
			{
				AttributeName: aws.String("requester"),
				KeyType:       aws.String("S"),
			},
		},
		TableName: aws.String("Deposit"),
	}

	_, err := svc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling Transaction:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created Transaction")
}

// UploadedFileTable for uploaded data
const UploadedFileTable = "UploadedFile"

func createUploadedFile(svc *dynamodb.DynamoDB) {
	// Create table UploadedFile
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("json"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("offer_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		TableName: aws.String(UploadedFileTable),
	}

	_, err := svc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling " + UploadedFileTable + ":")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created " + UploadedFileTable)
}

func main() {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	createDeposit(svc)
	createTransaction(svc)
	createUploadedFile(svc)
}
