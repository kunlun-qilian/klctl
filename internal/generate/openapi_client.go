package generate

import (
    "fmt"
    "github.com/deepmap/oapi-codegen/pkg/codegen"
    "github.com/deepmap/oapi-codegen/pkg/util"
    "io/fs"
    "io/ioutil"
    "os"
)

func createDir(path string) {
    _, err := os.Stat(path)
    if err != nil {
        err = os.Mkdir(path, fs.ModePerm)
        if err != nil {
            panic(err)
        }
    }
}

func clientDir(path, packageName string) string {
    return fmt.Sprintf("%s/client_%s", path, packageName)
}

func packName(packageName string) string {
    return fmt.Sprintf("client_%s", packageName)
}

func GenerateOpenapiClient(packageName, path, url string) {

    createDir(clientDir(path, packageName))

    co := codegen.Configuration{
        PackageName: packName(packageName),
        Generate: codegen.GenerateOptions{
            Client:       true,
            Models:       true,
            EmbeddedSpec: true,
        },
    }

    swagger, err := util.LoadSwagger(url)
    if err != nil {
        panic(err)
    }

    code, err := codegen.Generate(swagger, co)
    if err != nil {
        fmt.Println("err: ", err)
        panic(err)
    }

    err = ioutil.WriteFile(fmt.Sprintf("%s/%s.go", clientDir(path, packageName), packName(packageName)), []byte(code), fs.ModePerm)
    if err != nil {
        panic(fmt.Sprintf("error writing generated code to file: %s", err))
    }
}
