package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

const typeTemplateString = `package types
// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
{{ range $alias, $import := .Imports }}
	{{ $alias }} "{{ $import }}"
{{ end }}
)

{{ godoc .Type.StructName .TypeDoc "" }}type {{ .Type.StructName }} struct {
{{ range $i, $property := .DefaultProperties }}
{{ godoc $property.Name $property.Doc "\t" }}	{{ upper $property.Name }} {{ if $property.TypeArray }}[]{{ end }}{{ if $property.Type.GetGoType.Pointer }}*{{ end }}{{ typePath $property.Type.GetGoType }}{{ $property.Tag }}
{{ end }}

{{ range $i, $property := .ExtendedProperties }}
{{ godoc $property.Name $property.Doc "\t" }}	{{ upper $property.Name }} {{ if $property.TypeArray }}[]{{ end }}{{ if $property.Type.GetGoType.Pointer }}*{{ end }}{{ typePath $property.Type.GetGoType }}{{ $property.Tag }}
{{ end }}

}

func ({{ $.Type.Lower }} *{{ $.Type.StructName }}) String() string {
	return "{{ $.Type.StructName }}"
}
`

func getBases(entity MetadataEntity, entities map[string]MetadataEntity) []string {
	if entity.Base == "" {
		return []string{string(entity.Type)}
	}
	return append(getBases(entities[entity.Base], entities), string(entity.Type))
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.softlayer.com/metadata/v3.1", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var entities map[string]MetadataEntity
	err = json.Unmarshal(body, &entities)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, service := range entities {
		for _, base := range getBases(service, entities) {
			if service.Properties == nil {
				service.Properties = make(map[string]Property)
			}
			for prop_name, prop := range entities[base].Properties {
				service.Properties[prop_name] = prop
			}
		}
	}

	typeTemplate := template.Must(
		template.New("typeTemplateString").Funcs(template.FuncMap{
			"upper":     upper,
			"godoc":     godoc,
			"safeParam": safeParam,
			"typePath":  typePath,
		}).Parse(typeTemplateString))

	err = os.RemoveAll("types")
	if err != nil {
		log.Fatal(err)
	}
	for _, service := range entities {
		err = os.MkdirAll(service.Type.DirectoryPath(), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(service.Type.FileName())
		if err != nil {
			log.Fatal(err)
		}

		// Gather imports

		err = typeTemplate.Execute(f, service)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
