package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"udhar/constant"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type AWSClient struct {
	config *aws.Config
}

func createTopic(ctx context.Context, awsConfig *aws.Config, topicName string, snsClient *sns.Client) (string, error) {
	if topicName == "" {
		log.Fatal("Topic Name is Empty")
		return "", errors.New("Topic Name Not Defined")
	}
	c := snsClient
	input := &sns.CreateTopicInput{
		Name: &topicName,
	}

	results, err := c.CreateTopic(ctx, input)

	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return *results.TopicArn, nil

}

func createQueue(ctx context.Context, awsConfig *aws.Config, qName string, sqsClient *sqs.Client, delayInSeconds string) (string, error) {
	if qName == "" {
		log.Fatal("Topic Name is Empty")
		return "", errors.New("Topic Name is not defined")
	}
	c := sqsClient
	input := &sqs.CreateQueueInput{
		QueueName: &qName,
		Attributes: map[string]string{
			"DelaySeconds":           delayInSeconds,
			"MessageRetentionPeriod": "86400",
		},
	}

	result, err := c.CreateQueue(ctx, input)

	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return *result.QueueUrl, nil

}

func SetupInfra(c *aws.Config) (*sqs.Client, *sns.Client, map[string]string) {
	ctx := context.Background()
	sqsClient := sqs.NewFromConfig(*c)
	snsClient := sns.NewFromConfig(*c)

	queueData := make(map[string]string)

	for _, v := range constant.SETUP {
		if v.PublisherType == "SNS" && v.RecieverType == "SQS" {
			topicArn, err := createTopic(ctx, c, v.Name, snsClient)
			if err != nil {
				panic(err)
			}
			queueUrl, err := createQueue(ctx, c, v.Name, sqsClient, v.SNSDelayInSeconds)
			if err != nil {
				panic(err)
			}
			queueData[queueUrl] = v.Name
			// AWS go-client doesn't have any method to get Queue ARN
			queueArn := fmt.Sprintf("arn:aws:sqs:%s:%s:%s", "us-east-1", "000000000000", v.Name)
			// Subscribing to SNS
			subscribeInput := &sns.SubscribeInput{
				Protocol: aws.String("sqs"),
				TopicArn: &topicArn,
				Endpoint: &queueArn,
			}

			_, err = snsClient.Subscribe(ctx, subscribeInput)
			if err != nil {
				log.Fatalf("Failed to subscribe SQS queue to SNS topic: %v", err)
			}

		}
	}
	return sqsClient, snsClient, queueData
}
