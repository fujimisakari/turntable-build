package jsonschema

type (
	APISchema interface {
		GetRequestSchema() map[string]interface{}
		GetResponseSchema() map[string]interface{}
	}
)
