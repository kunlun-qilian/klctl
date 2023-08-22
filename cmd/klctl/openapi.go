package main

import (
	"bytes"
	"encoding/json"
	"github.com/kunlun-qilian/klctl/internal/generate"
	"os"
	"strings"

	"github.com/ghodss/yaml"

	"io/fs"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var swaggerFile string

var cmdSwagger = &cobra.Command{
	Use:     "openapi",
	Aliases: []string{"swagger"},
	Short:   "scan current project and generate openapi.json",
	Run: func(cmd *cobra.Command, args []string) {

		SwaggerToOpenapi()

	},
}

func init() {
	cmdSwagger.Flags().
		StringVarP(&swaggerFile, "file", "f", "./swagger.json", "(required) if convert swagger json")

	cmdRoot.AddCommand(cmdSwagger)
}

func SwaggerToOpenapi() {
	swaggerFileByte, err := os.ReadFile(swaggerFile)
	if err != nil {
		panic(err)
	}
	newSwaggerFileByte := changeEnum(swaggerFileByte)

	y, err := yaml.JSONToYAML(newSwaggerFileByte)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(y)
	resp, err := http.Post("https://converter.swagger.io/api/convert", "application/yaml", reader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	openapiYamlByte, _ := ioutil.ReadAll(resp.Body)

	openapi3Obj := make(map[string]interface{})

	err = json.Unmarshal(openapiYamlByte, &openapi3Obj)
	if err != nil {
		panic(err)
	}

	openapiJson, err := json.MarshalIndent(openapi3Obj, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./openapi.json", openapiJson, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

var (
	TYPESKEY_ENUM            = `enum`
	TYPESKEY_TYPE            = `type`
	TYPESKEY_x_ENUM_VARNAMES = `x-enum-varnames`
)

func changeEnum(swaggerFile []byte) []byte {
	swagger := generate.Swagger{}
	err := json.Unmarshal(swaggerFile, &swagger)
	if err != nil {
		panic(err)
	}
	for k, v := range swagger.Definitions {
		if strings.Contains(k, "types.") {
			v[TYPESKEY_TYPE] = "string"
			v[TYPESKEY_ENUM] = []interface{}{}
			for _, xEnumVarName := range v[TYPESKEY_x_ENUM_VARNAMES].([]interface{}) {
				enumName := xEnumVarName.(string)
				if strings.Contains(enumName, "__") {
					v[TYPESKEY_ENUM] = append(v[TYPESKEY_ENUM].([]interface{}), strings.Split(xEnumVarName.(string), "__")[1])
				}
			}
		}
	}

	swaggerJson, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		panic(err)
	}
	return swaggerJson
}
