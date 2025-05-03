package ratelimiter

import "context"

type BucketIface interface {
	GetTokens(userIP string) (int, error)
	AddToken(userIP string)
	RemoveToken(userIP string) error
	GetMaxTokens(userIP string) (int, error)
	GetRate(userIP string) (int, error)
	SetMaxTokens(userIP string, max int) error
	SetRate(userIP string, rate int) error
	AddUser(userIP string) error
	StopAllTickers(ctx context.Context)
}

type BucketDB interface {
	FindOne(userIP string) (result Bucket, err error)
	InsertOne(userIP string, bucket Bucket) error
	UpdateOne(userIP string, updatedBucket Bucket) error
}
