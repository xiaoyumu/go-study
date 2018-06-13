package main

type RamlDefiniton struct {
	Title   string                                `json:"title"`
	BaseURI string                                `json:"baseUri"`
	Types   *map[string]*RamlEntityTypeDefinition `json:"types"` // Key is entity type name
}

type RamlEntityTypeDefinition struct {
	Type       string                             `json:"type"`
	Properties *map[string]RamlPropertyDefinition `json:"properties"`
}

type RamlPropertyDefinition struct {
	Type     string `json:"type"`
	Required bool   `json:"required"`
}

type RamlMethodDefinition struct {
	Body      *map[string]RamlBasicType   `json:"body"`      // Key is content-type, such as 'application/json'
	Responses *map[int]RamlBodyDefinition `json:"responses"` // Key is status code, such as 200
}

type RamlBodyDefinition struct {
	Body *map[string]RamlBasicType `json:"body"` // Key is content-type, such as 'application/json'
}

type RamlBasicType struct {
	Type string `json:"type"`
}
