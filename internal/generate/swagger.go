package generate

type Swagger struct {
	Swagger     string                            `json:"swagger"`
	Info        interface{}                       `json:"info"`
	Paths       interface{}                       `json:"paths"`
	Definitions map[string]map[string]interface{} `json:"definitions"`
}
