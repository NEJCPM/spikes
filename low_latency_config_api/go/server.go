package main

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

    "fmt"
    "net/http"
	"log"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":3001", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	type Item struct {
		name   string
		value  string
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("spike_low_latency"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String("test"),
			},
		},
	})
	
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		msg := "Could not find"
		fmt.Println(msg)
		// return nil
	}
		
	item := Item{}
	
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	
	fmt.Println("result: ", result)

	fmt.Fprintf(w, "Result: %s!", result)
}