package main

// SwaggerAPI Swagger API
type SwaggerAPI struct {
	Swagger     string                                   `json:"swagger"`
	Info        *SwaggerAPIInfo                          `json:"info"`
	BasePath    string                                   `json:"basePath"`
	Paths       *map[string]map[string]APIEndpointMethod `json:"paths"`
	Definitions *map[string]EntityDefinition             `json:"definitions"`
}

// SwaggerAPIInfo Swagger API Info node
type SwaggerAPIInfo struct {
	Version        string      `json:"version"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	TermsOfService string      `json:"termsOfService"`
	Contact        *APIContact `json:"contact"`
}

// APIContact the Swagger API contact
type APIContact struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Email string `json:"email"`
}

// API Endpoint
type APIEndpointMethod struct {
	Tags        []string                `json:"tags"`
	Summary     string                  `json:"summary"`
	OperationID string                  `json:"operationId"`
	Consumes    []string                `json:"consumes"`
	Produces    []string                `json:"produces"`
	Parameters  []APIParameter          `json:"parameters"`
	Responses   *map[string]APIResponse `json:"responses"`
}

// API Parameter
type APIParameter struct {
	Name        string     `json:"name"`
	In          string     `json:"in"`
	Description string     `json:"description"`
	Required    bool       `json:"required"`
	Type        string     `json:"type"`
	Schema      *SchemaRef `json:"schema"`
}

type APIResponse struct {
	Description string     `json:"description"`
	Schema      *SchemaRef `json:"schema"`
}

type EntityDefinition struct {
	Type        string                               `json:"type"`
	Description string                               `json:"description"`
	Properties  *map[string]EntityPropertyDefinition `json:"properties"`
}

type EntityPropertyDefinition struct {
	Format               string                    `json:"format"`
	Type                 string                    `json:"type"`
	ReadOnly             bool                      `json:"readOnly"`
	Items                *SchemaRef                `json:"items"` // If the type is array, the items should be referenced to the object
	Reference            string                    `json:"$ref"`  // The json path to Definitions
	AdditionalProperties *EntityPropertyDefinition `json:"AdditionalProperties"`
}

type SchemaRef struct {
	Type      string `json:"type"` // If the SchemaReference is primitive, the type is provided
	Reference string `json:"$ref"` // The json path to Definitions
}
