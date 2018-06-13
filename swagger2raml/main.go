package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/ghodss/yaml"
)

// The function GetRemoteSwaggerJson retrieve the remote Swagger Json from given url
func GetRemoteSwaggerJson(url string) *SwaggerAPI {
	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}
	var swaggerAPIObj SwaggerAPI

	deserializeErr := json.Unmarshal(body, &swaggerAPIObj)

	if deserializeErr != nil {
		panic(deserializeErr)
	}

	return &swaggerAPIObj
}

func main() {

	swaggerJSONURL := "http://10.16.75.24:3000/finance/abs/v1/invoice-scan/doc/v1/swagger.json"

	var swaggerAPIObj SwaggerAPI

	swaggerAPIObj = *GetRemoteSwaggerJson(swaggerJSONURL)

	for uri, endPoint := range *swaggerAPIObj.Paths {
		fmt.Println("------------------------------------")

		fmt.Printf("Endpoint URI: %s\n", uri)
		//var endPointStructure APIEndpoint
		//mapstructure.Decode(endPoint, &endPointStructure)

		//p := convertEndpoint(endPoint)
		for method, methodInfo := range endPoint {
			parameters := methodInfo.Parameters
			fmt.Printf("Method: %s\n", strings.ToUpper(method))

			for _, parameter := range parameters {
				fmt.Printf("    Parameter: %s [In:%s] [Type: %s] ",
					parameter.Name,
					parameter.In,
					parameter.Type)
				if parameter.Schema == nil {
					fmt.Println()
					continue
				}
				fmt.Printf("Ref: %s\n", (*parameter.Schema).Reference)
			}

			for responseCode, response := range *methodInfo.Responses {
				fmt.Printf("    ResponseCode: %s, %s ",
					responseCode,
					response.Description)

				if response.Schema != nil {
					fmt.Printf("Ref: %s", (*response.Schema).Reference)
				}

				fmt.Println()
			}

		}
	}

	for entityName, entityDef := range *swaggerAPIObj.Definitions {
		fmt.Printf("\nEntity Definition: %s [%s]\n", entityName, entityDef.Type)

		for propertyName, propertyDef := range *entityDef.Properties {
			fmt.Printf("    Property: %s \tType:%s \tFormat:%s \tReadonly:%t Items[%v] Ref[%s] AdditionalProperties[%v] \n",
				propertyName,
				propertyDef.Type,
				propertyDef.Format,
				propertyDef.ReadOnly,
				propertyDef.Items,
				propertyDef.Reference,
				propertyDef.AdditionalProperties)
		}
	}

	ramlBasicInfo, ramlEndPointList := convertToRaml(swaggerAPIObj)

	yamlDataBasic, err := yaml.Marshal(ramlBasicInfo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(string(yamlDataBasic))

	yamlDataEndpointList, _ := yaml.Marshal(ramlEndPointList)
	fmt.Println(string(yamlDataEndpointList))
}

func convertToRaml(api SwaggerAPI) (RamlDefiniton, *map[string]*map[string]RamlMethodDefinition) {
	def := RamlDefiniton{
		Title: api.Info.Title,
		Types: convertRamlTypeMapping(api),
	}

	endPoints := convertEndpoints(api)

	return def, endPoints
}

func convertRamlTypeMapping(api SwaggerAPI) *map[string]*RamlEntityTypeDefinition {
	mapping := map[string]*RamlEntityTypeDefinition{}

	for entityName, entityDef := range *api.Definitions {
		ramlTypeName := toRamlTypeName(entityName)    // string
		ramlEntityType := toRamlEntityType(entityDef) // RamlEntityTypeDefinition

		mapping[ramlTypeName] = ramlEntityType
	}

	return &mapping
}

func toRamlTypeName(swaggerEntityName string) string {
	if !strings.Contains(swaggerEntityName, "[") {
		return swaggerEntityName
	}
	name := strings.Replace(swaggerEntityName, "[", "_", 1)
	name = name[:len(name)-1]

	return name
}

func toRamlEntityType(swaggerEntityDef EntityDefinition) *RamlEntityTypeDefinition {
	var entityTypeDef RamlEntityTypeDefinition
	entityTypeDef.Type = swaggerEntityDef.Type

	propertiesMap := make(map[string]RamlPropertyDefinition)

	for propertyName, propertyDef := range *swaggerEntityDef.Properties {
		propertiesMap[propertyName] = RamlPropertyDefinition{
			Type:     propertyDef.Type,
			Required: false,
		}
	}

	entityTypeDef.Properties = &propertiesMap

	return &entityTypeDef
}

func convertEndpoints(api SwaggerAPI) *map[string]*map[string]RamlMethodDefinition {
	endpoints := map[string]*map[string]RamlMethodDefinition{}

	for path, endPointDefMapping := range *api.Paths {
		endpoints[path] = toRamlEndpointMethodMapping(api, &endPointDefMapping)
	}

	return &endpoints
}

func toRamlEndpointMethodMapping(api SwaggerAPI, method *map[string]APIEndpointMethod) *map[string]RamlMethodDefinition {
	return nil
}
