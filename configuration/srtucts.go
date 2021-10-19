package configuration


type ServerAns struct {
	Name string `json:"name"`
	Cod int `json:"cod"`
	Main map[string] float64 `json:"main"`
}

type ConfigFromYaml struct {
	Name string `yaml:"name"`
	Units string `yaml:"units"`
}

type ConfigFromFlags struct {
	T string // Town name
	U string // Units (imperial, standard, metric)
}

type RequestConfig struct {
	TownName string
	UnitSystem string
}