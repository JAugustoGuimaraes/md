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
}

func New() (*Renderer, error) {
  indexHtml := fmt.Sprintf(baseIndexHtml, websocketJs, cssSwitchJs)

  tmpl, err := template.New("index").Parse(indexHtml)
  if err != nil {
    return &Renderer{}, fmt.Errorf("Failed to create renderer: %w", err)
  }

  return &Renderer{
    template: tmpl,
    data: Data{},
  }, nil
}

func (r *Renderer) SetTitle(title string) {
  r.data.Title = title
}

func (r *Renderer) SetContent(content string) {
  r.data.Content = content
}

func (r *Renderer) Render(w http.ResponseWriter) error {
  return r.template.Execute(w, r.data)
}
