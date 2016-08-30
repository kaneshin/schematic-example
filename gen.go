package main

import (
	"bytes"
	"go/format"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/kaneshin/schematic"
	bundle "github.com/kaneshin/schematic-example/templates"
)

var (
	templates *template.Template
	newlines  = regexp.MustCompile(`(?m:\s*$)`)
)

func init() {
	templates = template.Must(bundle.Parse(schematic.Templates()))
}

// Generate generates code according to the schema.
func Generate(s *schematic.Schema) ([]byte, error) {
	var buf bytes.Buffer

	for i := 0; i < 2; i++ {
		s.Resolve(nil)
	}

	name := strings.ToLower(strings.Split(s.Title, " ")[0])
	templates.ExecuteTemplate(&buf, "package.tmpl", name)
	templates.ExecuteTemplate(&buf, "imports.tmpl", []string{
		"fmt",
		"bytes",
	})

	for _, name := range sortedKeys(s.Properties) {
		schema := s.Properties[name]
		// Skipping definitions because there is no links, nor properties.
		if schema.Links == nil && schema.Properties == nil {
			continue
		}

		context := struct {
			Name       string
			Definition *schematic.Schema
		}{
			Name:       name,
			Definition: schema,
		}

		templates.ExecuteTemplate(&buf, "struct.tmpl", context)
	}

	// Remove blank lines added by text/template
	bytes := newlines.ReplaceAll(buf.Bytes(), []byte(""))

	// Format sources
	clean, err := format.Source(bytes)
	if err != nil {
		return buf.Bytes(), err
	}
	return clean, nil
}

func sortedKeys(m map[string]*schematic.Schema) (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return
}
