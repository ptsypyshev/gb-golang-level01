package main

import (
	"fmt"
	"main/myconf"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	var config myconf.MyConfig
	var err error
	configMap := myconf.GetFromEnvs()
	if len(os.Args) > 1 {
		configMap = myconf.GetFromFlags() // Priority to Flags if some of them specified
	}
	if fileJSON, ok := configMap["json"]; ok {
		configJSON := myconf.GetFromFile(fileJSON, myconf.C_JSON)
		config, err = configJSON.ParseToMyConfig()
	} else if fileYAML, ok := configMap["yaml"]; ok {
		configYAML := myconf.GetFromFile(fileYAML, myconf.C_YAML)
		config, err = configYAML.ParseToMyConfig()
	} else {
		config, err = myconf.ParseMap(configMap)
	}
	checkErr(err)
	fmt.Println(config)
}
