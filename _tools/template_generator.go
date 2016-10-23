package main

import (
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
	Name    string
	Type    string
	Example interface{}
}

type MapSchemaProperty struct {
	Self      string
	JsonName  string
	ModelName string
}

var APP_PATH = fmt.Sprintf("%s/src/github.com/fujimisakari/turntable-build", os.Getenv("GOPATH"))

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Usage: ", path.Base(args[0]), " model_name")
		return
	}
	target := args[1]

	// Load yaml data
	yamlPath := fmt.Sprintf("%s/yaml/%s.yaml", APP_PATH, target)
	if _, err := os.Stat(yamlPath); err != nil {
		fmt.Println("Yaml File not found: ", yamlPath)
		return
	}

	context := createContext(yamlPath)

	// Create model template
	modelTplPath := fmt.Sprintf("%s/_templates/model.go.tpl", APP_PATH)
	outputModelPath := fmt.Sprintf("%s/model/%s_master.go", APP_PATH, target)
	templateWriter(context, modelTplPath, outputModelPath)

	// Create service template
	serviceTplPath := fmt.Sprintf("%s/_templates/service.go.tpl", APP_PATH)
	outputServicePath := fmt.Sprintf("%s/domain/%s/service_master.go", APP_PATH, target)
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
		if s, ok := c["json-example"].(string); ok {
			jsonExample = s
		} else if i, ok := c["json-example"].(int); ok {
			jsonExample = i
		} else {
			jsonExample = c["json-example"]
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
