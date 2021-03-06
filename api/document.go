package api

import (
	_ "strconv"

	_ "github.com/Sirupsen/logrus"
	"github.com/fujimisakari/turntable-build/jsonschema"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type dynamicMap map[string]interface{}

func GetDocumentData() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		schemaMappers := jsonschema.GetSchemaMapperForDocument()
		context := make(dynamicMap)

		// set URL Data
		urlGroups := []dynamicMap{}
		for _, mapperWithGroup := range schemaMappers {
			for groupName, mapper := range mapperWithGroup {
				mapper := mapper.(map[string]interface{})
				urls := make([]string, len(mapper))
				idx := 0
				for url, _ := range mapper {
					urls[idx] = url
					idx += 1
				}

				urlGroup := dynamicMap{"name": groupName, "urls": urls}
				urlGroups = append(urlGroups, urlGroup)
			}
		}
		context["urlGroups"] = urlGroups

		// set Schema Data
		schemaData := make(map[string]dynamicMap)
		for _, mapperWithGroup := range schemaMappers {
			for _, mapper := range mapperWithGroup {
				mapper := mapper.(map[string]interface{})
				for url, apiSchema := range mapper {
					apiSchema := apiSchema.(jsonschema.APISchema)
					reqSchema := apiSchema.GetRequestSchema()
					schemaData[url] = dynamicMap{
						"title":          reqSchema["title"],
						"description":    reqSchema["description"],
						"method":         reqSchema["method"],
						"requestSchema":  createRequestDocSchema(reqSchema),
						"responseSchema": createResponseDocSchema(apiSchema.GetResponseSchema()),
					}
				}
			}
		}
		context["schemaData"] = schemaData

		return c.JSON(fasthttp.StatusOK, context)
	}
}

func createRequestDocSchema(reqSchema dynamicMap) dynamicMap {
	properties := reqSchema["properties"].(map[string]interface{})
	requireds := reqSchema["required"].([]string)
	arguments := make([]map[string]string, len(properties))

	idx := 0
	for name, detail := range properties {
		detail := detail.(map[string]string)
		argument := make(map[string]string)
		argument["name"] = name
		argument["example"] = detail["example"]
		argument["description"] = detail["description"]

		required := "Optional"
		for _, arg := range requireds {
			if arg == name {
				required = "Required"
			}
		}
		argument["required"] = required
		arguments[idx] = argument
		idx++
	}

	docSchema := dynamicMap{
		"arguments": arguments,
	}
	return docSchema
}

func createResponseDocSchema(resSchema dynamicMap) dynamicMap {
	docSchema := make(dynamicMap)
	properties := resSchema["properties"].(map[string]interface{})
	docSchema = getPropertiesByRecursion(docSchema, properties)
	return docSchema
}

func getPropertiesByRecursion(docSchema dynamicMap, properties dynamicMap) dynamicMap {

	for name, detail := range properties {
		detail := detail.(map[string]interface{})

		switch detail["type"] {
		case "array":
			arraySchema := make(dynamicMap)
			itemSchema := detail["items"].(map[string]interface{})
			itemProperties := itemSchema["properties"].(map[string]interface{})
			item := getPropertiesByRecursion(arraySchema, itemProperties)
			schemaList := []dynamicMap{item}
			docSchema[name] = schemaList

		case "object":
			objSchema := make(dynamicMap)
			docSchema[name] = getPropertiesByRecursion(objSchema, detail["properties"].(map[string]interface{}))

		default:
			docSchema[name] = detail["example"]
		}
	}
	return docSchema
}
