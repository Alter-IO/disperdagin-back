package service

import (
	"alter-io-go/domain"
	"alter-io-go/helpers/derrors"
	"alter-io-go/helpers/logger"
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

/*
	UploadFile processes an uploaded file by generating a structured file path.

	It performs the following:
	- Extracts the original filename and extension.
	- Slugifies the filename (lowercase, spaces and special characters replaced).
	- Generates a unique timestamp down to milliseconds.
	- Constructs a file path in the format: `/uploads/{type}/{timestamp}_{slug}{ext}`

	Example:
	Original: "The Term Vocabulary and Postings Lists.pdf"
	Type: "legal-document"
	Output: `/uploads/legal-document/20250329_203119.923_the_term_vocabulary_and_postings_lists.pdf`

*
*/
func (s *Service) UploadFile(ctx context.Context, reqBody domain.UploadReq) (domain.Upload, error) {
	// validate type
	allowedTypes := map[string]struct{}{
		"employee":           {},
		"ikm-type":           {},
		"legal-document":     {},
		"photo":              {},
		"public-information": {},
	}

	if _, exists := allowedTypes[reqBody.Type]; !exists {
		return domain.Upload{}, derrors.NewErrorf(derrors.ErrorCodeBadRequest, "Invalid Type")
	}

	// Extract base filename and extension
	filename := strings.TrimSpace(filepath.Base(reqBody.File.Filename))
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)

	// Slugify filename: lowercase + replace spaces with underscores
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "_"))

	// Use timestamp down to milliseconds for make that unique
	timestamp := time.Now().Format("20060102_150405.000")

	// Construct the file URL with structured format
	fileURL := fmt.Sprintf("/uploads/%s/%s_%s%s", reqBody.Type, timestamp, slug, ext)

	// Loggin that file
	logger.Get().With().Info(fmt.Sprintf("Uploading file: original='%s', path='%s'", filename, fileURL))

	return domain.Upload{
		FileName: strings.TrimSpace(name),
		FileURL:  fileURL,
	}, nil
}
