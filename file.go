package aquagram

import (
	"io"
)

type Files = map[string]*InputFile

type InputFile struct {
	MediaType  string
	FileName   string
	FromReader io.Reader
	FromPath   string
	FromFileID string
	FromURL    string
}

func InputFileFromReader(r io.Reader) *InputFile {
	file := new(InputFile)
	file.FromReader = r

	return file
}

func InputFileFromPath(p string) *InputFile {
	file := new(InputFile)
	file.FromPath = p

	return file
}

func InputFileFromFileID(id string) *InputFile {
	file := new(InputFile)
	file.FromFileID = id

	return file
}

func InputFileFromURL(u string) *InputFile {
	file := new(InputFile)
	file.FromURL = u

	return file
}
