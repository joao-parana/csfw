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

// Generates code for all EAV types
package main

/*
   Data will be "carved in stone" because it only changes during development.
   - DONE: entity_type with translation of some columns to the Go type
   - attribute_set related tables: eav_attribute_set, eav_entity_attribute, eav_attribute_group, etc
   - label and option tables will not be hard coded
   - eav_attribute full config and from that one the flat table structure
*/

const tplEav = `// Copyright 2015 CoreStore Authors
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

// Package {{ .Package }} file is auto generated via eavToStruct
package {{ .Package }}
import (
    "github.com/corestoreio/csfw/eav"
    {{ range .ETypeData }}{{if ne .ImportPath "" }}"{{.ImportPath}}"
{{end}}{{end}}
)

var (
    // CSEntityTypeCollection contains all entity types mapped to their Go types/interfaces
    CSEntityTypeCollection = eav.CSEntityTypeSlice{
        {{ range .ETypeData }} &eav.CSEntityType {
            EntityTypeID: {{ .EntityTypeID }},
            EntityTypeCode: "{{ .EntityTypeCode }}",
            EntityModel: {{ .EntityModel }},
            AttributeModel: {{ .AttributeModel }},
            EntityTable: {{ .EntityTable }},
            ValueTablePrefix: "{{ .ValueTablePrefix }}",
            IsDataSharing: {{ .IsDataSharing }},
            DataSharingKey: "{{ .DataSharingKey }}",
            DefaultAttributeSetID: {{ .DefaultAttributeSetID }},
            {{ if ne "" .IncrementModel }}IncrementModel: {{ .IncrementModel }},{{ end }}
            IncrementPerStore: {{ .IncrementPerStore }},
            IncrementPadLength: {{ .IncrementPadLength }},
            IncrementPadChar: "{{ .IncrementPadChar }}",
            AdditionalAttributeTable: {{ .AdditionalAttributeTable }},
            EntityAttributeCollection: {{ .EntityAttributeCollection }},
        },
        {{ end }}
    }
)
`

var defaultMapping = []byte(`{
  "customer": {
    "import_path": "github.com/corestoreio/csfw/customer",
    "entity_model": "customer.Customer()",
    "attribute_model": "customer.Attribute()",
    "entity_table": "customer.Customer()",
    "increment_model": "customer.Customer()",
    "additional_attribute_table": "customer.Customer()",
    "entity_attribute_collection": "customer.Customer()"
  },
  "customer_address": {
    "import_path": "",
    "entity_model": "customer.Address()",
    "attribute_model": "customer.AddressAttribute()",
    "entity_table": "customer.Address()",
    "additional_attribute_table": "customer.Address()",
    "entity_attribute_collection": "customer.Address()"
  },
  "catalog_category": {
    "import_path": "github.com/corestoreio/csfw/catalog",
    "entity_model": "catalog.Category()",
    "attribute_model": "catalog.Attribute()",
    "entity_table": "catalog.Category()",
    "additional_attribute_table": "catalog.Category()",
    "entity_attribute_collection": "catalog.Category()"
  },
  "catalog_product": {
    "import_path": "",
    "entity_model": "catalog.Product()",
    "attribute_model": "catalog.Attribute()",
    "entity_table": "catalog.Product()",
    "additional_attribute_table": "catalog.Product()",
    "entity_attribute_collection": "catalog.Product()"
  }
}`)
