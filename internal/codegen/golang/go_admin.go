package golang

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/kyleconroy/sqlc/internal/codegen"
	"github.com/kyleconroy/sqlc/internal/config"
	"go/format"
	"strings"
	"text/template"
)


type AdminTableGenerator struct {
	TableName string
	Method 	  string
}

func generateGoAdmin(settings config.CombinedSettings, enums []Enum, structs []Struct, queries []Query) (map[string]string, error) {
	i := &importer{
		Settings: settings,
		Queries:  queries,
		Enums:    enums,
		Structs:  structs,
	}

	funcMap := template.FuncMap{
		"lowerTitle": codegen.LowerTitle,
		"comment":    codegen.DoubleSlashComment,
		"escape":     codegen.EscapeBacktick,
		"imports":    i.Imports,
	}

	tmpl := template.Must(
		template.New("table").
			Funcs(funcMap).
			ParseFS(
				templates,
				"templates/*.tmpl",
				"templates/*/*.tmpl",
			),
	)

	golang := settings.Go
	tctx := tmplCtx{
		Settings:                  settings.Global,
		EmitInterface:             golang.EmitInterface,
		EmitJSONTags:              golang.EmitJSONTags,
		EmitDBTags:                golang.EmitDBTags,
		EmitPreparedQueries:       golang.EmitPreparedQueries,
		EmitEmptySlices:           golang.EmitEmptySlices,
		EmitMethodsWithDBArgument: golang.EmitMethodsWithDBArgument,
		SQLPackage:                SQLPackageFromString(golang.SQLPackage),
		Q:                         "`",
		Package:                   golang.Package,
		GoQueries:                 queries,
		Enums:                     enums,
		Structs:                   structs,
	}

	output := map[string]string{}

	execute := func(name, templateName string) error {
		var b bytes.Buffer
		w := bufio.NewWriter(&b)
		tctx.SourceName = name
		err := tmpl.ExecuteTemplate(w, templateName, &tctx)
		w.Flush()
		if err != nil {
			return err
		}
		code, err := format.Source(b.Bytes())
		if err != nil {
			fmt.Println(b.String())
			return fmt.Errorf("source error: %w", err)
		}

		if templateName == "queryFile" && golang.OutputFilesSuffix != "" {
			name += golang.OutputFilesSuffix
		}

		if !strings.HasSuffix(name, ".go") {
			name += ".go"
		}
		output[name] = string(code)
		return nil
	}

	return nil, nil
}
