package util

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

// IMAGEREGEX to extract the image html tag
const IMAGEREGEX = `<img[^>]+src="([^">]+)"`

// TrimAndSplitURL will trim tailing "/" and split url
func TrimAndSplitURL(u string) []string {
	u = strings.TrimSuffix(u, "/")
	return strings.Split(u, "/")
}

// URLSource will retrieve the url hostname.
func URLSource(u string) (string, error) {
	parsedUrl, err := url.Parse(u)

	if err != nil {
		return "", err
	}

	return parsedUrl.Hostname(), nil
}

// IsURLValid will exclude those url containing `.gif` and `logo`.
func IsURLValid(value string) bool {
	check := value != "" && !strings.Contains(value, ".gif") && !strings.Contains(value, "logo") && !strings.Contains(value, "mobilebanner")

	if check {
		return strings.HasPrefix(value, "http") || strings.HasPrefix(value, "https")
	}

	return check
}

// IsValueInSlice will check if a value is already in a slice.
func IsValueInSlice(valueToCheck string, values []string) bool {
	for _, v := range values {
		if v == valueToCheck {
			return true
		}
	}
	return false
}

// ConvertToJPG converts an image to jpeg
func ConvertToJPG(img image.Image, imgData *bytes.Buffer) error {
	err := jpeg.Encode(imgData, img, nil)
	if err != nil {
		return err
	}

	return nil
}

// ImageType return the image type
func ImageType(mimeStr string) (tp string) {
	switch mimeStr {
	case "image/png":
		tp = "png"
	case "image/jpg":
		tp = "jpg"
	case "image/jpeg":
		tp = "jpg"
	case "image/gif":
		tp = "gif"
	default:
		tp = "unknown"
	}
	return
}

// PathSetup will create the folders where the comic will be saved
func PathSetup(source, name string) (string, error) {
	dir, err := filepath.Abs(fmt.Sprintf("%s/%s/%s/%s/", filepath.Dir(os.Args[0]), "comics", source, name))

	if err != nil {
		log.Error(err)
	}

	// create folders
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Error(err)
	}

	return dir, err
}

// CurrentDir
func CurrentDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// FileDoesNotExist
func FileDoesNotExist(filePath string) bool {
	_, err := os.Stat(filePath)

	return os.IsNotExist(err)
}

// generateFileName will return the path where the file should be saved
func GenerateFileName(dir, name, issueNumber, format string) string {
	return fmt.Sprintf("%s/%s-%s.%s", dir, name, issueNumber, format)
}

func Parse(s string) string {
	return strings.Replace(s, "/", "_", -1)
}
