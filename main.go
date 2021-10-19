package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"weatherInformer/configuration"

)

func GetRequest(c *configuration.RequestConfig) (s configuration.ServerAns, err error) {
	var (
		req *http.Request
		res *http.Response
		jsonFile configuration.ServerAns
		result []byte
	)

	req, err = http.NewRequest(configuration.Method, configuration.UrlAddress, nil)
	if err != nil {
		err = errors.New("ошибка формирования запроса")
		return s, err
	}
	query := req.URL.Query()
	query.Add(configuration.ApiField, configuration.ApiKey)
	query.Add(configuration.Town, c.TownName)
	query.Add(configuration.Units, c.UnitSystem)

	req.URL.RawQuery = query.Encode()
	res, err = http.Get(req.URL.String())
	if err != nil {
		err = errors.New("ошибка исполнения запроса")
		return s, err
	}

	result, err = io.ReadAll(res.Body)
	if err != nil {
		err = errors.New("ошибка чтения запроса")
		return s, err
	}

	if err = json.Unmarshal(result, &jsonFile); err != nil {
		err = errors.New("ошибка обработки запроса")
		return s, err
	}
	s.Name = jsonFile.Name
	s.Cod = jsonFile.Cod
	s.Main = jsonFile.Main
	return s, nil
}


func main()  {
	var temperature float64
	var config configuration.RequestConfig
	var yamlOptions *configuration.ConfigFromYaml
	y, err := yamlOptions.GetOptions() // запрашиваем опции, прописанные yaml файле, r имеет тип RequestConfig
	if err != nil {
		fmt.Printf("ошибка конфигурации, %s\n", err)
		os.Exit(1)
	}

	var flagOptions configuration.ConfigFromFlags
	f, err := flagOptions.GetOptions()
	if err != nil{
		fmt.Printf("ошибка чтения параметров из командной строки, %s\n", err)
		os.Exit(1)
	}

	config.ConfigSelector(&y, &f)
	serverAns, err := GetRequest(&config)
	if err != nil{
		fmt.Printf("ошибка, %s\n", err)
		os.Exit(1)
	}

	for k, v := range serverAns.Main {
		if k == "temp" {
			temperature = v
			break
		}
	}
	fmt.Printf("cod: %v\ntown: %v\ntemp: %.2f\n",serverAns.Cod, serverAns.Name, temperature)

}
