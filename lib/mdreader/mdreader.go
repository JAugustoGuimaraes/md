package mdreader

import (
  "io/ioutil"
  "bytes"
  "fmt"

  "github.com/yuin/goldmark"
)

func ReadMDFile(filePath string) (string, error) {
  mdBuf, err := ioutil.ReadFile(filePath)
  if err != nil {
    return "", fmt.Errorf("Failed to read file: %w", err)
  }

  var htmlBuf bytes.Buffer
  if err := goldmark.Convert(mdBuf, &htmlBuf); err != nil {
    return "", fmt.Errorf("Failed to convert Markdown to HTML: %w", err)
  }

  return htmlBuf.String(), nil
}
