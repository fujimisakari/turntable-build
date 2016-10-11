package main

import (
	"strconv"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

type Context struct {
	Self                 string
	UpperName            string
	LowerName            string
	Description          string
	ModelProperties      []ModelProperty
	JsonSchemaProperties []JsonSchemaProperty
	MapSchemaProperties  []MapSchemaProperty
}

type ModelProperty struct {
	DBColumn string
	Name     string
	Type     string
}

type JsonSchemaProperty struct {
	Name string
	Type string
	Example interface{}
}

type MapSchemaProperty struct {
	Self      string
	JsonName  string
	ModelName string
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Usage: ", path.Base(args[0]), " model_name")
		return
	}
	target := args[1]

	// Load yaml data
	yamlPath := fmt.Sprintf("../yaml/%s.yaml", target)
	if _, err := os.Stat(yamlPath); err != nil {
		fmt.Println("Yaml File not found: ", yamlPath)
		return
	}

	context := createContext(yamlPath)

	// Create model template
	modelTplPath := "../_templates/model.go.tpl"
	outputModelPath := fmt.Sprintf("../model/%s_master.go", target)
	templateWriter(context, modelTplPath, outputModelPath)

	// Create service template
	serviceTplPath := "../_templates/service.go.tpl"
	outputServicePath := fmt.Sprintf("../domain/%s/service_master.go", target)
	templateWriter(context, serviceTplPath, outputServicePath)
}

func getYamlData(yamlPath string) map[interface{}]interface{} {
	buf, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		panic(err)
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}

	return m
}

func createContext(yamlPath string) Context {
	yamlMap := getYamlData(yamlPath)
	schema := yamlMap["schema"].(map[interface{}]interface{})
	upperName, _ := schema["name"].(string)
	lowerName := strings.ToLower(upperName)
	columns := schema["columns"].([]interface{})
	colCount := len(columns)
	jsonColCount := getJsonColumnCount(columns)

	// Create ModelProperty list
	modelProperties := make([]ModelProperty, colCount)
	for i, cData := range columns {
		c, _ := cData.(map[interface{}]interface{})
		mp := ModelProperty{
			c["db-column-name"].(string),
			c["model-name"].(string),
			c["model-type"].(string),
		}
		modelProperties[i] = mp
	}

	// Create JsonSchemaProperty list
	jsProperties := make([]JsonSchemaProperty, jsonColCount)
	jIdx := 0
	for _, cData := range columns {
		c, _ := cData.(map[interface{}]interface{})
		var jsonExample interface{}
		jsonExample, ok := strconv.ParseInt(c["json-example"].(string), 0, 64)
		if ok == nil {
			jsonExample = c["json-example"].(string)
		}
		if c["json-name"] != nil {
			js := JsonSchemaProperty{
				c["json-name"].(string),
				c["json-type"].(string),
				jsonExample,
			}
			jsProperties[jIdx] = js
			jIdx += 1
		}
	}

	// Create MapSchemaProperty list
	mapProperties := make([]MapSchemaProperty, jsonColCount)
	mIdx := 0
	for _, cData := range columns {
		c, _ := cData.(map[interface{}]interface{})
		if c["json-name"] != nil {
			ms := MapSchemaProperty{
				lowerName[0:1],
				c["json-name"].(string),
				c["model-name"].(string),
			}
			mapProperties[mIdx] = ms
			mIdx += 1
		}
	}

	context := Context{
		lowerName[0:1],
		upperName,
		lowerName,
		schema["description"].(string),
		modelProperties,
		jsProperties,
		mapProperties,
	}
	return context
}

func getJsonColumnCount(columns []interface{}) int {
	count := 0
	for _, cData := range columns {
		c, _ := cData.(map[interface{}]interface{})
		if c["json-name"] != nil {
			count += 1
		}
	}
	return count
}

func templateWriter(context Context, tplPath string, outputPath string) {
	tpl := template.Must(template.ParseFiles(tplPath))
	output, _ := os.Create(outputPath)
	if err := tpl.Execute(output, context); err != nil {
		fmt.Println(err)
	}
	output.Close()
}
