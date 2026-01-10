package storage

import (
	"context"
	"io"

	"github.com/okoye-dev/marchive/internal/files"
)

// S3Client implements files.ObjectStorage for AWS S3
type S3Client struct {
	bucket string
	// Add other S3-specific fields: region, credentials, etc.
}

// NewS3Client creates a new S3 client
func NewS3Client(bucket string) *S3Client {
	return &S3Client{bucket: bucket}
}

// Ensure S3Client implements files.ObjectStorage interface
// This is a compile-time check - if it doesn't implement all methods, code won't compile
var _ files.ObjectStorage = (*S3Client)(nil)

// Upload implements files.ObjectStorage interface
func (s *S3Client) Upload(ctx context.Context, bucket, key string, body io.Reader) error {
	// TODO: Implement AWS S3 upload logic
	// Using AWS SDK: s3.PutObject, etc.
	return nil
}

// Download implements files.ObjectStorage interface
func (s *S3Client) Download(ctx context.Context, bucket, key string) (io.ReadCloser, error) {
	// TODO: Implement AWS S3 download logic
	// Using AWS SDK: s3.GetObject, etc.
	return nil, nil
}

// Delete implements files.ObjectStorage interface
func (s *S3Client) Delete(ctx context.Context, bucket, key string) error {
	// TODO: Implement AWS S3 delete logic
	// Using AWS SDK: s3.DeleteObject, etc.
	return nil
}
