package repository

import "context"

type Repository interface {
	PublishToTopic(ctx context.Context, message []byte, arn string)
}
