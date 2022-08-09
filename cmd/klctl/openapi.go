package main

import (
	"bytes"
	"encoding/json"

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
	swaggerFileByte, err := ioutil.ReadFile(swaggerFile)
	if err != nil {
		panic(err)
	}
	y, err := yaml.JSONToYAML(swaggerFileByte)
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
