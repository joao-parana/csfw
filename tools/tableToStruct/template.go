// Copyright 2015 CoreStore Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "github.com/corestoreio/csfw/tools"

const tplCode = tools.Copyright + `
package {{ .Package }}

// Package {{ .Package }} is auto generated via tableToStruct

import (
	"time"
    {{ if not .TypeCodeValueTables.Empty }}
	"github.com/corestoreio/csfw/eav"{{end}}
	"github.com/corestoreio/csfw/storage/csdb"
	"github.com/corestoreio/csfw/storage/dbr"
)

const (
    {{ range $k,$v := .Tables }} // TableIndex{{.name | prepareVar}} is the index to {{.table}}
    TableIndex{{.name | prepareVar}} {{ if eq $k 0 }}csdb.Index = iota // must start with 0{{ end }}
{{ end }} // TableIndexZZZ represents the maximum index, which is not available.
TableIndexZZZ
)

var (
    // Always reference these packages, just in case the auto-generated code
    // below doesn't.
    _ = time.Time{}

    tableMap = csdb.TableStructureSlice{
{{ range .Tables }}TableIndex{{.name | prepareVar}} : csdb.NewTableStructure(
        "{{.table}}",
        []string{
        {{ range .columns }}{{ if eq .Key.String "PRI" }} "{{.Field.String}}",{{end}}
        {{ end }} },
        []string {
        {{ range .columns }}{{ if ne .Key.String "PRI" }} "{{.Field.String}}",{{end}}
        {{ end }} },
    ),
    {{ end }}
    }
)

// GetTableStructure returns for a given index i the table structure or an error it not found.
func GetTableStructure(i csdb.Index) (*csdb.TableStructure, error) {
    if i < TableIndexZZZ { return tableMap.Structure(i) }
	return nil, csdb.ErrTableNotFound
}

// GetTableName returns for a given index the table name. If not found an empty string.
func GetTableName(i csdb.Index) string {
    if i < TableIndexZZZ { return tableMap.Name(i) }
	return ""
}

{{ if not .TypeCodeValueTables.Empty }}
{{range $typeCode,$valueTables := .TypeCodeValueTables}}
// Get{{ $typeCode | prepareVar }}ValueStructure returns for an eav value index the table structure.
// Important also if you have custom value tables
func Get{{ $typeCode | prepareVar }}ValueStructure(i eav.ValueIndex) (*csdb.TableStructure, error) {
	switch i {
	{{range $vt,$v := $valueTables }}case eav.EntityType{{ $v | prepareVar }}:
		return GetTableStructure(TableIndex{{ $vt | prepareVar }})
    {{end}}	}
	return nil, eav.ErrEntityTypeValueNotFound
}
{{end}}{{end}}

type (

{{ range .Tables }}
    // Table{{.name | prepareVar}}Slice contains pointers to Table{{.name | prepareVar}} types
    Table{{.name | prepareVar}}Slice []*Table{{.name | prepareVar}}
    // Table{{.name | prepareVar}} a type for the MySQL table {{ .table }}
    Table{{.name | prepareVar}} struct {
        {{ range .columns }}{{.GoName}} {{.GoType}} {{ $.Tick }}db:"{{.Field.String}}"{{ $.Tick }} {{.Comment}}
        {{ end }} }
{{ end }}
)
`
