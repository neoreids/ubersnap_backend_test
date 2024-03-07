package utilities

import (
	"bytes"
	"github.com/gabriel-vasile/mimetype"
	"gocv.io/x/gocv"
)

// DetectMimeType detects the MIME type of an image represented by gocv.Mat
func DetectMimeTypeOfMat(img gocv.Mat) (string, error) {
	// Convert the image to JPEG format
	buf, err := gocv.IMEncode(".jpg", img)
	if err != nil {
		return "", err
	}
	defer buf.Close()

	// Read the first few bytes to determine the format
	reader := bytes.NewReader(buf.GetBytes())
	// Detect MIME type
	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		return "", err
	}
	return mime.String(), nil
}
