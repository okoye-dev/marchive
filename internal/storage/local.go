package storage

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// LocalClient implements files.ObjectStorage for local filesystem.
type LocalClient struct {
	root string
}

func NewLocalClient(root string) *LocalClient {
	return &LocalClient{root: root}
}

func (l *LocalClient) Upload(ctx context.Context, bucket, key string, body io.Reader) error {
	path, err := l.safePath(bucket, key)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, body)
	return err
}

func (l *LocalClient) Download(ctx context.Context, bucket, key string) (io.ReadCloser, error) {
	path, err := l.safePath(bucket, key)
	if err != nil {
		return nil, err
	}
	return os.Open(path)
}

func (l *LocalClient) Delete(ctx context.Context, bucket, key string) error {
	path, err := l.safePath(bucket, key)
	if err != nil {
		return err
	}
	return os.Remove(path)
}

func (l *LocalClient) List(ctx context.Context, bucket, prefix string) ([]string, error) {
	base, err := l.safePath(bucket, "")
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(base); err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	var out []string
	err = filepath.WalkDir(base, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(base, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if prefix == "" || strings.HasPrefix(rel, prefix) {
			out = append(out, rel)
		}
		return nil
	})
	return out, err
}

func (l *LocalClient) safePath(bucket, key string) (string, error) {
	if bucket == "" {
		return "", errors.New("bucket is required")
	}
	if strings.Contains(bucket, "..") || strings.Contains(key, "..") {
		return "", errors.New("invalid path")
	}
	root := filepath.Clean(l.root)
	bucket = filepath.Clean(bucket)
	key = filepath.Clean(key)
	if key == "." {
		key = ""
	}
	path := filepath.Join(root, bucket, key)
	if root != "." && !strings.HasPrefix(path, root+string(os.PathSeparator)) && path != root {
		return "", errors.New("invalid path")
	}
	return path, nil
}
