package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	libcontainer "github.com/OneThing98/containerpkg"
	"github.com/OneThing98/namespaces"
)

func main() {
	var (
		configPath string
		cliCmd     string
	)

	flag.StringVar(&configPath, "config", "container.json", "Path to container config")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Please specify a command (exec, execin, net)")
		os.Exit(1)
	}

	cliCmd = flag.Arg(0)

	file, err := os.Open(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening config file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var container libcontainer.Container

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&container); err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding config: %v\n", err)
		os.Exit(1)
	}

	switch cliCmd {
	case "exec":
		if err := namespaces.ContainerExec(&container); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing container: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unsupported command: %s\n", cliCmd)
		os.Exit(1)
	}
}
