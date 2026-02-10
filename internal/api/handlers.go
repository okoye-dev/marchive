package api

import "github.com/okoye-dev/marchive/internal/files"

type Handlers struct {
	Files         *files.FileService
	DefaultBucket string
}
