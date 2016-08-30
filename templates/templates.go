package templates

import "text/template"

var templates = map[string]string{"imports.tmpl": `{{if .}}
  {{if len . | eq 1}}
    import {{range .}}"{{.}}"{{end}}
  {{else}}
    import (
      {{range .}}
      	"{{.}}"
      {{end}}
    )
  {{end}}
{{end}}`,
	"package.tmpl": `package {{.}}
`,
	"struct.tmpl": `{{asComment .Definition.Description}}
type {{initialCap .Name}} {{goType .Definition}}
`,
}

// Parse parses declared templates.
func Parse(t *template.Template) (*template.Template, error) {
	for name, s := range templates {
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		if _, err := tmpl.Parse(s); err != nil {
			return nil, err
		}
	}
	return t, nil
}

