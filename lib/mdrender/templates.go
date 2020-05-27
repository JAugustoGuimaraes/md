package mdrender

import (
	"fmt"
	"io/ioutil"

	"github.com/markbates/pkger"
)

func getIndexHtml() (string, error) {
	indexHtmlBuf, err := pkger.Open("/lib/mdrender/templates/index.html")
	if err != nil {
		return "", fmt.Errorf("Failed to get index.html: %w", err)
	}

	indexHtml, err := ioutil.ReadAll(indexHtmlBuf)
	if err != nil {
		return "", fmt.Errorf("Failed to read index.html: %w", err)
	}

	return string(indexHtml), nil
}

func getWebSocketJs() (string, error) {
	webSocketJsBuf, err := pkger.Open("/lib/mdrender/templates/webSocket.js")
	if err != nil {
		return "", fmt.Errorf("Failed to get websocket script: %w", err)
	}

	webSocketJs, err := ioutil.ReadAll(webSocketJsBuf)
	if err != nil {
		return "", fmt.Errorf("Failed to read websocket script: %w", err)
	}

	return string(webSocketJs), nil
}

func getCssSwitchJs() (string, error) {
	cssSwitchJsBuf, err := pkger.Open("/lib/mdrender/templates/cssSwitch.js")
	if err != nil {
		return "", fmt.Errorf("Failed to get css switch script: %w", err)
	}

	cssSwitchJs, err := ioutil.ReadAll(cssSwitchJsBuf)
	if err != nil {
		return "", fmt.Errorf("Failed to get read switch script: %w", err)
	}

	return string(cssSwitchJs), nil
}
