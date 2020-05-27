package mdreader

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/yuin/goldmark"
)

func ReadMDFile(filePath string) ([]byte, error) {
	mdBuf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to read file: %w", err)
	}

	var htmlBuf bytes.Buffer
	if err := goldmark.Convert(mdBuf, &htmlBuf); err != nil {
		return []byte{}, fmt.Errorf("Failed to convert Markdown to HTML: %w", err)
	}

	return htmlBuf.Bytes(), nil
}
