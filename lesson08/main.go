package main

import (
	"fmt"
	"main/myconf"
	"os"
)

func main() {
	configMap := myconf.GetFromEnvs()
	if len(os.Args) > 1 {
		configMap = myconf.GetFromFlags() // Priority to Flags if some of them specified
	}
	config, err := myconf.Parser(configMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
}
