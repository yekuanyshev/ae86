package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	minioClient *minio.Client
	bucket      string
}

func Connect(ctx context.Context, config Config) (*Client, error) {
	minioClient, err := minio.New(config.Endpoint(), &minio.Options{
		Creds:  credentials.NewStaticV4(config.User, config.Password, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	found, err := minioClient.BucketExists(ctx, config.Bucket)
	if err != nil {
		return nil, err
	}

	if !found {
		err = minioClient.MakeBucket(ctx, config.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		minioClient: minioClient,
		bucket:      config.Bucket,
	}, nil
}

func (c *Client) Upload(ctx context.Context, object *Object) error {
	_, err := c.minioClient.PutObject(
		ctx,
		c.bucket,
		object.Name,
		object.Content,
		object.Size,
		minio.PutObjectOptions{
			ContentType: object.ContentType,
		},
	)
	return err
}

func (c *Client) Download(ctx context.Context, filename string) (*Object, error) {
	object, err := c.minioClient.GetObject(
		ctx,
		c.bucket,
		filename,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, err
	}

	stat, err := object.Stat()
	if err != nil {
		return nil, err
	}

	return &Object{
		Content:     object,
		Name:        filename,
		Size:        stat.Size,
		ContentType: stat.ContentType,
	}, nil
}
