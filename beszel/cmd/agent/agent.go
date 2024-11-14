package main

import (
	"beszel"
	"beszel/internal/agent"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// handle flags / subcommands
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v":
			fmt.Println(beszel.AppName+"-agent", beszel.Version)
		case "update":
			agent.Update()
		}
		os.Exit(0)
	}

	var pubKey []byte
	if pubKeyEnv, exists := os.LookupEnv("KEY"); exists {
		pubKey = []byte(pubKeyEnv)
	} else if pubKeyFile, exists := os.LookupEnv("KEY_FILE"); exists {
		var err error
		pubKey, err = os.ReadFile(pubKeyFile)
		if err != nil {
			log.Fatal("Cannot read file given by KEY_FILE")
		}
	} else {
		log.Fatal("Neither KEY nor KEY_FILE environment variable is set")
	}

	addr := ":45876"
	if portEnvVar, exists := os.LookupEnv("PORT"); exists {
		// allow passing an address in the form of "127.0.0.1:45876"
		if !strings.Contains(portEnvVar, ":") {
			portEnvVar = ":" + portEnvVar
		}
		addr = portEnvVar
	}

	agent.NewAgent().Run(pubKey, addr)
}
