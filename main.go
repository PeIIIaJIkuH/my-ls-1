package main

import (
	"fmt"
	"log"
	"my-ls-1/parse"
	"os"
)

func exit(err error) {
	log.Fatalf("error: %v\n", err)
}

func main() {
	args := os.Args[1:]

	directoryNames, _, err := parse.Args(args)
	if err != nil {
		exit(err)
	}

	err = parse.Directories(directoryNames)
	if err != nil {
		exit(err)
	}
	//for _, directoryName := range directoryNames {
	//	dirs, _ := os.ReadDir(directoryName)
	//	for _, dir := range dirs {
	//		fmt.Println(dir.Name(), dir.Type())
	//	}
	//}
	dirs, _ := os.ReadDir(".")
	for _, dir := range dirs {
		fmt.Println(dir.Name())
		info, _ := os.Stat(dir.Name())
		fmt.Printf("%+v\n\n", info.Mode())
	}
}
