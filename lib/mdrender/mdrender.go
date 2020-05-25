package mdrender

import (
  "fmt"
  "net/http"
  "text/template"
)

type Renderer struct {
  template *template.Template
  data Data
}

type Data struct {
  Title string
  Content string
  WebSocketScript string
  CssSwitchScript string
}

func New(title, content string) (*Renderer, error) {
  indexHtml, err := getIndexHtml()
  if err != nil {
    return &Renderer{}, fmt.Errorf("Failed to create renderer: %w", err)
  }

  websocketJs, err := getWebSocketJs()
  if err != nil {
    return &Renderer{}, fmt.Errorf("Failed to create renderer: %w", err)
  }

  cssSwitchJs, err := getCssSwitchJs()
  if err != nil {
    return &Renderer{}, fmt.Errorf("Failed to create renderer: %w", err)
  }

  tmpl, err := template.New("index").Parse(indexHtml)
  if err != nil {
    return &Renderer{}, fmt.Errorf("Failed to create renderer: %w", err)
  }

  return &Renderer{
    template: tmpl,
    data: Data{
      Title: title,
      Content: content,
      WebSocketScript: websocketJs,
      CssSwitchScript: cssSwitchJs,
    },
  }, nil
}

func (r *Renderer) Render(w http.ResponseWriter) error {
  return r.template.Execute(w, r.data)
}
