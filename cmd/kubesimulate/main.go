package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags
	simulateCMD := flag.NewFlagSet("simulate", flag.ExitOnError)
	nodeName := simulateCMD.String("node", "", "Name of the node to simulate a scenario on")

	// Verify a subcommand has been provided
	// os.Args[0] is the program name, os.Args[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments supplied")
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for the appropriate subcommand
	switch os.Args[1] {
	case "simulate":
		simulateCMD.Parse(os.Args[2:])
		if *nodeName == "" {
			simulateCMD.PrintDefaults()
			os.Exit(1)
		}
		fmt.Printf("Simulating scenario on node: %s/n", *nodeName)
	default:
		fmt.Println("Unknown subcommand")
		os.Exit(1)
	}

}
