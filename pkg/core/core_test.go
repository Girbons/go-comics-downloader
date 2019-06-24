package core

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Girbons/mangarock"
	"github.com/stretchr/testify/assert"
)

func exists(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func TestNewComic(t *testing.T) {
	comic := new(Comic)
	// links
	links := []string{"foo.example.com"}

	comic.Name = "foo"
	comic.IssueNumber = "2"
	comic.Links = links
	comic.Source = "bar"

	assert.Equal(t, "foo", comic.Name)
	assert.Equal(t, "2", comic.IssueNumber)
	assert.Equal(t, "bar", comic.Source)

	assert.Equal(t, 1, len(comic.Links))
}

func TestMakeComicPDF(t *testing.T) {
	comic := new(Comic)

	comic.Name = "foo"
	comic.Format = "pdf"
	comic.IssueNumber = "example-chapter-1"
	comic.Links = []string{"http://via.placeholder.com/150", "http://via.placeholder.com/150", "http://via.placeholder.com/150"}

	err := comic.MakeComic()
	assert.Nil(t, err)

	dir, _ := filepath.Abs(fmt.Sprintf("%s/%s/%s/%s/", filepath.Dir(os.Args[0]), "comics", "foo", "foo-example-chapter-1.pdf"))
	assert.True(t, exists(dir))
}

func TestMakeComicEPUB(t *testing.T) {
	comic := new(Comic)

	comic.Name = "foo"
	comic.Format = "epub"
	comic.IssueNumber = "example-chapter-1"
	comic.Author = "author"

	comic.Links = []string{"http://via.placeholder.com/150", "http://via.placeholder.com/150", "http://via.placeholder.com/150"}
	err := comic.MakeComic()
	assert.Nil(t, err)

	dir, _ := filepath.Abs(fmt.Sprintf("%s/%s/%s/%s/", filepath.Dir(os.Args[0]), "comics", "foo", "foo-example-chapter-1.epub"))
	assert.True(t, exists(dir))
}

func TestMakeComicEPUBMangarock(t *testing.T) {
	options := map[string]string{"country": "italy"}
	client := mangarock.NewClient(options)
	result, _ := client.Pages("mrs-chapter-100051049")

	comic := new(Comic)

	comic.Name = "Boruto"
	comic.Format = "epub"
	comic.IssueNumber = "chapter-13"
	comic.Source = "mangarock.com"
	comic.Links = result.Data

	err := comic.MakeComic()
	assert.Nil(t, err)

	dir, _ := filepath.Abs(fmt.Sprintf("%s/%s/%s/%s/%s/", filepath.Dir(os.Args[0]), "comics", "mangarock.com", "Boruto", "Boruto-chapter-13.epub"))
	assert.True(t, exists(dir))
}

func TestMakeComicCBZMangarock(t *testing.T) {
	options := map[string]string{"country": "italy"}
	client := mangarock.NewClient(options)
	result, _ := client.Pages("mrs-chapter-100051049")

	comic := new(Comic)
	comic.Name = "Boruto"
	comic.Format = "cbz"
	comic.IssueNumber = "chapter-13"
	comic.Source = "mangarock.com"
	comic.Links = result.Data

	err := comic.MakeComic()
	assert.Nil(t, err)

	dir, _ := filepath.Abs(fmt.Sprintf("%s/%s/%s/%s/%s/", filepath.Dir(os.Args[0]), "comics", "mangarock.com", "Boruto", "Boruto-chapter-13.cbz"))
	assert.True(t, exists(dir))
}

func TestMakeComicCBRMangarock(t *testing.T) {
	options := map[string]string{"country": "italy"}
	client := mangarock.NewClient(options)
	result, _ := client.Pages("mrs-chapter-100051049")

	comic := new(Comic)
	comic.Name = "Boruto"
	comic.Format = "cbr"
	comic.IssueNumber = "chapter-13"
	comic.Source = "mangarock.com"
	comic.Links = result.Data

	err := comic.MakeComic()
	assert.Nil(t, err)

	dir, _ := filepath.Abs(fmt.Sprintf("%s/%s/%s/%s/%s/", filepath.Dir(os.Args[0]), "comics", "mangarock.com", "Boruto", "Boruto-chapter-13.cbr"))
	assert.True(t, exists(dir))
}

func TestMakeComicPDFMangarock(t *testing.T) {
	options := map[string]string{"country": "italy"}
	client := mangarock.NewClient(options)
	result, _ := client.Pages("mrs-chapter-100051049")

	comic := new(Comic)

	comic.Name = "Boruto"
	comic.Format = "pdf"
	comic.IssueNumber = "chapter-13"
	comic.Source = "mangarock.com"
	comic.Links = result.Data

	err := comic.MakeComic()
	assert.Nil(t, err)

	dir, _ := filepath.Abs(fmt.Sprintf("%s/%s/%s/%s/%s/", filepath.Dir(os.Args[0]), "comics", "mangarock.com", "Boruto", "Boruto-chapter-13.pdf"))
	assert.True(t, exists(dir))
}
