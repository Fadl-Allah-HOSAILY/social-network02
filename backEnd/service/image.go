
package service

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofrs/uuid"
)

func CheckImageExtension(fileName string) bool {
	extension := strings.ToLower(filepath.Ext(fileName))

	switch extension {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}

func CheckImageContent(file io.ReadSeeker) (bool, error) {
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return false, err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return false, err
	}

	contentType := http.DetectContentType(buffer)
	return strings.HasPrefix(contentType, "image/"), nil
}

func ImageTreatment(file multipart.File, header *multipart.FileHeader) (string, error) {
	if file == nil || header == nil {
		return "", nil
	}

	defer file.Close()

	if !CheckImageExtension(header.Filename) {
		return "", errors.New("extension not allowed")
	}

	const maxSize = 5 << 20

	if header.Size > maxSize {
		return "", errors.New("image too big")
	}

	isImage, err := CheckImageContent(file)
	if err != nil {
		return "", err
	}

	if !isImage {
		return "", errors.New("not an image")
	}

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return "", err
	}

	uniqueName, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))
	fileName := fmt.Sprintf("%s%s", uniqueName.String(), ext)
	imagePath := filepath.Join("uploads", fileName)

	destination, err := os.Create(imagePath)
	if err != nil {
		return "", err
	}
	defer destination.Close()

	if _, err := io.Copy(destination, file); err != nil {
		return "", err
	}

	return imagePath, nil
}
