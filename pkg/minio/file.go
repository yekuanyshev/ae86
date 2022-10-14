package minio

import "io"

type File struct {
	Content     io.Reader
	Name        string
	Size        int64
	ContentType string
}
