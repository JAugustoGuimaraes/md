package mdrender

import (
  "path"
  "path/filepath"
  "runtime"
  "io/ioutil"
  "errors"
  "fmt"
)

func getBasePath() (string, bool) {
  _, b, _, ok := runtime.Caller(0)
  if !ok {
    return "", false
  }

  return filepath.Dir(b), true
}

func getIndexHtml() (string, error) {
  basePath, ok := getBasePath()
  if !ok {
    return "", errors.New("Failed to get base path")
  }

  indexHtmlPath := path.Join(basePath, "/templates/index.html")
  indexHtml, err := ioutil.ReadFile(indexHtmlPath)
  if err != nil {
    return "", fmt.Errorf("Failed to get index.html: %w", err)
  }

  return string(indexHtml), nil
}

func getWebSocketJs() (string, error) {
  basePath, ok := getBasePath()
  if !ok {
    return "", errors.New("Failed to get base path")
  }

  webSocketJsPath := path.Join(basePath, "/templates/webSocket.js")
  webSocketJs, err := ioutil.ReadFile(webSocketJsPath)
  if err != nil {
    return "", fmt.Errorf("Failed to get websocket script: %w", err)
  }

  return string(webSocketJs), nil
}

func getCssSwitchJs() (string, error) {
  basePath, ok := getBasePath()
  if !ok {
    return "", errors.New("Failed to get base path")
  }

  cssSwitchJsPath := path.Join(basePath, "/templates/cssSwitch.js")
  cssSwitchJs, err := ioutil.ReadFile(cssSwitchJsPath)
  if err != nil {
    return "", fmt.Errorf("Failed to get css switch script: %w", err)
  }

  return string(cssSwitchJs), nil
}
