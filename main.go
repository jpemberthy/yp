package main

import (
	"fmt"
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
		fmt.Errorf(err.Error())
	}

	config := map[string]map[string]Package{}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	pretty.Println(config)
}
