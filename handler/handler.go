package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func PollMessages(sqsClient *sqs.Client, queueURL string, channel chan<- map[string]interface{}) {

	receiveParams := &sqs.ReceiveMessageInput{
		QueueUrl:            &queueURL,
		MaxNumberOfMessages: 1,  // Maximum number of messages to retrieve
		WaitTimeSeconds:     10, // Long polling (wait for up to 10 seconds)
	}

	for {
		resp, err := sqsClient.ReceiveMessage(context.TODO(), receiveParams)
		if err != nil {
			log.Fatalf("Error occurred while recieving message from queue %v", err)
			panic(err)
		}
		if len(resp.Messages) == 0 {
			log.Println("No messages pulled")
			continue
		}

		for _, message := range resp.Messages {
			messageBody := *message.Body
			byteSlice := []byte(messageBody)

			var data map[string]interface{}
			err := json.Unmarshal(byteSlice, &data)
			if err != nil {
				fmt.Printf("Error unmarshaling JSON: %v\n", err)
				return
			}

			data = map[string]interface{}{
				"queueUrl": queueURL,
				"message":  data["Message"],
			}
			channel <- data
			deleteParams := &sqs.DeleteMessageInput{
				QueueUrl:      &queueURL,
				ReceiptHandle: message.ReceiptHandle,
			}
			_, err = sqsClient.DeleteMessage(context.TODO(), deleteParams)
			if err != nil {
				log.Printf("Error deleting message: %v", err)
			}
		}
	}
}
