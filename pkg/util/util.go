package util

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
	parsedURL, err := url.Parse(u)

	if err != nil {
		return "", err
	}

	return parsedURL.Hostname(), nil
}

// IsURLValid will exclude those url containing `.gif` and `logo`.
func IsURLValid(url string) bool {
	invalidValues := []string{".gif", "logo", "mobilebanner", "wp-content"}
	check := true

	for _, v := range invalidValues {
		if strings.Contains(url, v) {
			check = false
			break
		}
	}

	if check {
		return strings.HasPrefix(url, "http") || strings.HasPrefix(url, "https")
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
	case "image/png", "png":
		tp = "png"
	case "image/jpg", "jpg":
		tp = "jpg"
	case "image/jpeg", "jpeg":
		tp = "jpg"
	case "image/gif", "gif":
		tp = "gif"
	case "img":
		tp = "img"
	default:
		tp = "unknown"
	}
	return
}

// PathSetup will create the folders where the comic will be saved
func PathSetup(outputFolder, source, name string) (string, error) {
	var dir string
	var err error

	dir, err = filepath.Abs(fmt.Sprintf("%s/comics/%s/%s/", outputFolder, source, name))

	if err != nil {
		return dir, err
	}

	// create folders
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return dir, err
	}

	return dir, err
}

// ImagesPathSetup setup the folders for the images to be saved.
func ImagesPathSetup(outputFolder, source, name, issueNumber string) (string, error) {
	var dir string
	var err error

	dir, err = filepath.Abs(fmt.Sprintf("%s/comics/%s/%s/images-%s/", outputFolder, source, name, issueNumber))

	if err != nil {
		return dir, err
	}

	// create folders
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return dir, err
	}

	return dir, err
}

// CurrentDir return the path where the executable is
func CurrentDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// DirectoryOrFileDoesNotExist check if a directory/file exist or not
func DirectoryOrFileDoesNotExist(filePath string) bool {
	_, err := os.Stat(filePath)

	return os.IsNotExist(err)
}

// GenerateFileName will return the path where the file should be saved
func GenerateFileName(dir, name, issueNumber, format string) string {
	return fmt.Sprintf("%s/%s-%s.%s", dir, name, issueNumber, format)
}

// Parse is used to escape characters
func Parse(s string) string {
	replacer := strings.NewReplacer(
		".", " ",
		"/", "_",
		"[", "",
		"]", "",
		":", "",
		";", "",
		"!", "",
		"?", "",
	)

	return strings.Trim(replacer.Replace(s), " ")
}

// SaveImage saves an image from a given format
func SaveImage(w io.Writer, content io.Reader, format string) error {
	img, _, err := image.Decode(content)

	if err != nil {
		return err
	}

	switch strings.ToLower(format) {
	case "img":
		_, err = io.Copy(w, content)
		return err
	case "gif":
		return gif.Encode(w, img, nil)
	case "jpg", "jpeg":
		return jpeg.Encode(w, img, nil)
	case "png":
		return png.Encode(w, img)
	default:
		return errors.New("format not found")
	}
}
