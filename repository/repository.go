package repository

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type awsConfig struct {
	awscfg *aws.Config
}

func NewAWSClient(cfg *aws.Config) *awsConfig {
	return &awsConfig{
		awscfg: cfg,
	}
}

func (cfg awsConfig) PublishToTopic(ctx context.Context, message []byte, arn string) {
	c := sns.NewFromConfig(*cfg.awscfg)
	messageData := string(message)
	input := &sns.PublishInput{
		Message:  &messageData,
		TopicArn: &arn,
	}

	_, err := c.Publish(ctx, input)
	if err != nil {
		log.Fatalf("Failed to publish to Topic %v", err)
		return
	}

}
