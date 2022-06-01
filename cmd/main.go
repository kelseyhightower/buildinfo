package main

import (
	"debug/buildinfo"
	"flag"
	"fmt"
	"log"
)

var inputBinaryFile string

func main() {
	flag.StringVar(&inputBinaryFile, "binary", "", "The go binary to use")
	flag.Parse()

	fmt.Printf("Getting build info for the %s binary...\n", inputBinaryFile)

	i, err := buildinfo.ReadFile(inputBinaryFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Go version: %v\n", i.GoVersion)
	fmt.Printf("Build Settings: \n")
	for _, s := range i.Settings {
		fmt.Printf("  %s: %s\n", s.Key, s.Value)
	}

	fmt.Println("Dependencies: ")
	for _, dep := range i.Deps {
		fmt.Printf("  %s: %s %s\n", dep.Path, dep.Version, dep.Sum)
	}

}
