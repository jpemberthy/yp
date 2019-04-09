package main

import (
	"io/ioutil"

	"github.com/kr/pretty"
	"gopkg.in/yaml.v2"
)

const filePath = "sample.yaml"

type Package struct {
	Hash       string
	Normalized string
}

func main() {
	source, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	config := map[string]map[string]Package{}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}

	pretty.Println(config)
}
