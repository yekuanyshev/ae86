package connections

import (
	"context"

	"github.com/supernova0730/ae86/pkg/minio"
)

var MinioConn *minio.Client

func MinioConnect(ctx context.Context, config minio.Config) error {
	minioConn, err := minio.Connect(ctx, config)
	if err != nil {
		return err
	}
	MinioConn = minioConn
	return nil
}
