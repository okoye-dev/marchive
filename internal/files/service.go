package files

import (
	"context"
	"io"
)

// ObjectStorage defines what FileService needs from storage
// This interface is defined HERE (where it's used), not in storage package
type ObjectStorage interface {
	Upload(ctx context.Context, bucket, key string, body io.Reader) error
	Download(ctx context.Context, bucket, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, bucket, key string) error
}

// FileService is the high-level service that uses ObjectStorage
// It depends on the interface, not a concrete implementation
type FileService struct {
	storage ObjectStorage
}

// NewFileService creates a new FileService with the given storage
// Notice: accepts ObjectStorage interface, not *storage.S3Client
func NewFileService(storage ObjectStorage) *FileService {
	return &FileService{storage: storage}
}

// UploadFile uploads a file using the storage interface
func (fs *FileService) UploadFile(ctx context.Context, bucket, key string, body io.Reader) error {
	return fs.storage.Upload(ctx, bucket, key, body)
}

// DownloadFile downloads a file using the storage interface
func (fs *FileService) DownloadFile(ctx context.Context, bucket, key string) (io.ReadCloser, error) {
	return fs.storage.Download(ctx, bucket, key)
}

// DeleteFile deletes a file using the storage interface
func (fs *FileService) DeleteFile(ctx context.Context, bucket, key string) error {
	return fs.storage.Delete(ctx, bucket, key)
}