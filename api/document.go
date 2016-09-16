package api

import (
	_ "strconv"

	_ "github.com/Sirupsen/logrus"
	"github.com/fujimisakari/turntable-build/jsonschema"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type schemaMap map[string]interface{}

func GetAllURL() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		schemaMapper := jsonschema.GetSchemaMapper()

		idx := 0
		urlList := make([]string, len(schemaMapper))
		for url, _ := range schemaMapper {
			urlList[idx] = url
			idx += 1
		}

		context := map[string]interface{}{
			"urlList": urlList,
		}
		return c.JSON(fasthttp.StatusOK, context)
	}
}

func GetSchema() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		schemaMapper := jsonschema.GetSchemaMapper()

		context := make(map[string]schemaMap, len(schemaMapper))
		for url, schema := range schemaMapper {
			apiSchema, _ := schema.(jsonschema.APISchema)
			context[url] = map[string]interface{}{
				"request":  createRequestDocSchema(apiSchema.GetRequestSchema()),
				"response": createResponseDocSchema(apiSchema.GetResponseSchema()),
			}
		}

		return c.JSON(fasthttp.StatusOK, context)
	}
}

func createRequestDocSchema(s schemaMap) map[string]interface{} {
	properties, _ := s["properties"].(map[string]interface{})
	arguments := make([]map[string]string, len(properties))

	idx := 0
	for name, detail := range properties {
		argument := make(map[string]string)
		detail := detail.(map[string]string)
		argument["argument"] = name
		argument["example"] = detail["example"]
		argument["description"] = detail["description"]

		required := "Optional"
		for _, arg := range s["required"].([]string) {
			if arg == name {
				required = "Required"
			}
		}
		argument["required"] = required
		arguments[idx] = argument
		idx++
	}

	docSchema := map[string]interface{}{
		"title":       s["title"],
		"description": s["description"],
		"arguments":   arguments,
	}
	return docSchema
}

func createResponseDocSchema(s schemaMap) map[string]interface{} {
	docSchema := make(schemaMap)
	properties := s["properties"].(map[string]interface{})
	docSchema = getPropertiesRecursion(docSchema, properties)
	return docSchema
}

func getPropertiesRecursion(docSchema schemaMap, properties schemaMap) schemaMap {

	for name, detail := range properties {
		detail := detail.(map[string]interface{})

		switch detail["type"] {
		case "array":
			arraySchema := make(schemaMap)
			itemSchema := detail["items"].(map[string]interface{})
			itemProperties := itemSchema["properties"].(map[string]interface{})
			item := getPropertiesRecursion(arraySchema, itemProperties)
			schemaList := []schemaMap{item}
			docSchema[name] := schemaList

		case "object":
			objSchema := make(schemaMap)
			docSchema[name] = getPropertiesRecursion(objSchema, detail["properties"].(schemaMap))

		default:
			docSchema[name] = detail["example"]
		}
	}
	return docSchema
}
