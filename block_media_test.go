package notion_test

import (
	"testing"
	"time"

	notion "github.com/amberpixels/notion-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestPdfBlockImplementsDownloadableFileBlock(t *testing.T) {
	// Test setup
	now := time.Now()
	pdfBlock := &notion.PdfBlock{
		Pdf: notion.File{
			File: &notion.FileData{
				URL:        "https://example.com/file.pdf",
				ExpiryTime: &now,
			},
		},
	}

	// Test GetURL
	assert.Equal(t, "https://example.com/file.pdf", pdfBlock.GetURL(), "Unexpected URL")

	// Test GetExpiryTime
	assert.Equal(t, &now, pdfBlock.GetExpiryTime(), "Unexpected expiry time")
}

func TestFileBlockImplementsDownloadableFileBlock(t *testing.T) {
	// Test setup
	now := time.Now()
	fileBlock := &notion.FileBlock{
		File: notion.File{
			File: &notion.FileData{
				URL:        "https://example.com/file.txt",
				ExpiryTime: &now,
			},
		},
	}

	// Test GetURL
	assert.Equal(t, "https://example.com/file.txt", fileBlock.GetURL(), "Unexpected URL")

	// Test GetExpiryTime
	assert.Equal(t, &now, fileBlock.GetExpiryTime(), "Unexpected expiry time")
}

func TestImageBlockImplementsDownloadableFileBlock(t *testing.T) {
	// Test setup
	now := time.Now()
	imageBlock := &notion.ImageBlock{
		Image: notion.File{
			File: &notion.FileData{
				URL:        "https://example.com/image.jpg",
				ExpiryTime: &now,
			},
		},
	}

	// Test GetURL
	assert.Equal(t, "https://example.com/image.jpg", imageBlock.GetURL(), "Unexpected URL")

	// Test GetExpiryTime
	assert.Equal(t, &now, imageBlock.GetExpiryTime(), "Unexpected expiry time")
}

func TestExternalURLCases(t *testing.T) {
	// Test External URLs for each block type
	testCases := []struct {
		name     string
		block    notion.Media
		expected string
	}{
		{
			name: "PDF with external URL",
			block: &notion.PdfBlock{
				Pdf: notion.File{
					External: &notion.FileData{
						URL: "https://external.com/file.pdf",
					},
				},
			},
			expected: "https://external.com/file.pdf",
		},
		{
			name: "File with external URL",
			block: &notion.FileBlock{
				File: notion.File{
					External: &notion.FileData{
						URL: "https://external.com/file.txt",
					},
				},
			},
			expected: "https://external.com/file.txt",
		},
		{
			name: "Image with external URL",
			block: &notion.ImageBlock{
				Image: notion.File{
					External: &notion.FileData{
						URL: "https://external.com/image.jpg",
					},
				},
			},
			expected: "https://external.com/image.jpg",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.block.GetURL(), "Unexpected URL")
		})
	}
}
