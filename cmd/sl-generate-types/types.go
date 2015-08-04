package main

import (
	"bytes"
	"go/doc"
	"strings"

	"github.com/aarzilli/sandblast"
	"golang.org/x/net/html"
)

type GoType struct {
	Literal string
	BuiltIn bool
	Pointer bool
	Import  string
	Alias   string
}

type Type string

func (t Type) GetGoType() *GoType {
	switch string(t) {
	case "string":
		return &GoType{Literal: "string", BuiltIn: true}
	case "int":
		return &GoType{Literal: "int", BuiltIn: true}
	case "long":
		return &GoType{Literal: "int64", BuiltIn: true}
	case "nonNegativeInteger":
		return &GoType{Literal: "uint", BuiltIn: true}
	case "unsignedInt":
		return &GoType{Literal: "uint", BuiltIn: true}
	case "unsignedLong":
		return &GoType{Literal: "uint64", BuiltIn: true}
	case "boolean":
		return &GoType{Literal: "bool", BuiltIn: true}
	case "decimal":
		return &GoType{Literal: "Float64", BuiltIn: false, Import: "github.com/sudorandom/softlayer-go/slapi", Alias: "slapi"}
	case "float":
		return &GoType{Literal: "Float64", BuiltIn: false, Import: "github.com/sudorandom/softlayer-go/slapi", Alias: "slapi"}
	case "dateTime":
		return &GoType{Literal: "Time", BuiltIn: false, Import: "time", Alias: "time", Pointer: true}
	case "base64Binary":
		return &GoType{Literal: "string", BuiltIn: true}
	case "enum":
		return &GoType{Literal: "string", BuiltIn: true}
	case "json":
		return &GoType{Literal: "string", BuiltIn: true}
	case "void":
		return nil
	default:
		return &GoType{Literal: t.StructName(), BuiltIn: true, Import: "github.com/sudorandom/softlayer-go/slapi/types" + t.DirectoryPath(), Alias: t.Alias(), Pointer: true}
	}
}

func (t Type) PackageName() string {
	lowerName := strings.ToLower(string(t))
	nameParts := strings.Split(lowerName, "_")
	return nameParts[len(nameParts)-2]
}

func (t Type) DirectoryPath() string {
	return "types/"
}

func (t Type) Alias() string {
	lowerName := strings.ToLower(string(t))
	nameParts := strings.Split(lowerName, "_")
	return "types_" + strings.Join(nameParts[:len(nameParts)-1], "_")
}

func (t Type) StructName() string {
	return string(t)
}

func (t Type) FileName() string {
	lowerName := strings.ToLower(string(t))
	nameParts := strings.Split(lowerName, "_")
	return "types/" + strings.Join(nameParts[:len(nameParts)], "_") + ".go"
}

func (t Type) Lower() string {
	return strings.ToLower(string(t))
}

type Parameter struct {
	Name         string      `json:"name"`
	Type         Type        `json:"type"`
	TypeArray    bool        `json:"typeArray"`
	Doc          string      `json:"doc"`
	DefaultValue interface{} `json:"defaultValue"`
}

type Method struct {
	Name       string      `json:"name"`
	Type       Type        `json:"type"`
	TypeArray  bool        `json:"typeArray"`
	Doc        string      `json:"doc"`
	Static     bool        `json:"static"`
	Maskable   bool        `json:"maskable"`
	Parameters []Parameter `json:"parameters"`
}

func (m Method) Arguments() string {
	argumentString := ""
	for i, param := range m.Parameters {
		argumentString += safeParam(param.Name)
		argumentString += " "
		if param.TypeArray == true {
			argumentString += "[]"
		}
		gotype := param.Type.GetGoType()
		argumentString += typePath(*gotype)
		if i != len(m.Parameters)-1 {
			argumentString += ", "
		}
	}

	return argumentString
}

func (m Method) ReturnArguments() string {
	returnArgumentString := ""
	gotype := m.Type.GetGoType()
	if gotype != nil {
		if m.TypeArray == true {
			returnArgumentString += "[]"
		}
		if gotype.Pointer == true {
			returnArgumentString += "*"
		}
		returnArgumentString += typePath(*gotype) + ", "
	}
	returnArgumentString += "error"
	return returnArgumentString
}

func (m Method) Body() string {
	body := ""
	gotype := m.Type.GetGoType()
	if gotype != nil {
		body += "var returnValue "
		if m.TypeArray == true {
			body += "[]"
		}
		if gotype.Pointer == true {
			body += "*"
		}
		body += typePath(*gotype) + "\n"
	}
	// TODO ACTUAL API CALL
	body += "	return "
	if m.Type != "void" {
		body += "returnValue, "
	}

	// TODO: Return err when there is one
	body += "nil"
	return body
}

type Property struct {
	Name      string `json:"name"`
	Form      string `json:"form"`
	Type      Type   `json:"type"`
	TypeArray bool   `json:"typeArray"`
	Doc       string `json:"doc"`
	Maskable  bool   `json:"maskable"`
}

func (p Property) Tag() string {
	return " `json:\"" + p.Name + ",omitempty\"`"
}

type MetadataEntity struct {
	Type       Type                `json:"name"`
	Base       string              `json:"base"`
	ServiceDoc string              `json:"serviceDoc"`
	TypeDoc    string              `json:"typeDoc"`
	Methods    map[string]Method   `json:"methods"`
	Properties map[string]Property `json:"properties"`
}

func (e MetadataEntity) DefaultProperties() []Property {
	properties := make([]Property, 0)
	for _, prop := range e.Properties {
		if prop.Form == "local" {
			properties = append(properties, prop)
		}
	}
	return properties
}

func (e MetadataEntity) ExtendedProperties() []Property {
	properties := make([]Property, 0)
	for _, prop := range e.Properties {
		if prop.Form != "local" {
			properties = append(properties, prop)
		}
	}
	return properties
}

func (e MetadataEntity) Imports() map[string]string {
	imports := map[string]string{}
	for _, property := range e.Properties {

		propertyGoType := property.Type.GetGoType()
		if propertyGoType == nil {
			continue
		}
		if propertyGoType.BuiltIn == true {
			continue
		}
		imports[propertyGoType.Alias] = propertyGoType.Import
	}

	return imports
}

func upper(name string) string {
	return strings.ToUpper(name[0:1]) + name[1:]
}

func godoc(name, helpText string, indent string) string {
	node, err := html.Parse(strings.NewReader(helpText))
	if err != nil {
		return "no documentation"
	}

	_, helpText, err = sandblast.Extract(node)
	if err != nil {
		return "no documentation"
	}

	helpText = strings.TrimSpace(helpText)
	if helpText == "" {
		helpText = "no documentation"
	}

	text := upper(name) + " - " + helpText
	out := bytes.NewBuffer(nil)
	doc.ToText(out, text, indent+"// ", "", 100)
	return out.String()
}

func safeParam(name string) string {
	if name == "type" {
		return "type_"
	}
	return name
}

func typePath(t GoType) string {
	path := ""
	if t.BuiltIn == false {
		path += t.Alias + "."
	}
	path += t.Literal
	return path
}
