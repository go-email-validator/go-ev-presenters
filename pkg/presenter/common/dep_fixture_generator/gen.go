package dep_fixture_generator

import (
	"strings"
	"text/template"
	"time"
)

type Template struct {
	Timestamp     time.Time
	PackageName   string
	PresenterName string
	Presenters    []interface{}
	Import        string
}

func replace(old, new, src string) string {
	return strings.Replace(src, old, new, -1)
}

var funcMap = template.FuncMap{
	"replace": replace,
}

/*
TODO do representation more readable, may be acceptable with spew.NewFormatter
*/
var PackageTemplate = template.Must(template.New("").Funcs(funcMap).Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package {{ .PackageName }}

{{ .Import }}

func depPresenters() []{{ .PresenterName }} {
	return []{{ .PresenterName }}{
		{{- range .Presenters }}
		{{ printf "%#v" . | replace (printf "%T" .) "" | replace (printf "%v." $.PackageName) ""|}},
		{{- end }}
	}
}
`))
