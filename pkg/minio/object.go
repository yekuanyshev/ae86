package minio

import "io"

type Object struct {
	Content     io.Reader
	Name        string
	Size        int64
	ContentType string
}
