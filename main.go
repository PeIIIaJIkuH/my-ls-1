package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my-ls-1/parse"
	"os"
	"strings"
)

func exit(err error) {
	log.Fatalf("error: %v\n", err)
}

func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
}

func main() {
	args := os.Args[1:]

	filenames, flags, err := parse.Args(args)
	if err != nil {
		exit(err)
	}

	err = parse.Filenames(filenames)
	if err != nil {
		exit(err)
	}

	_, err = parse.Entities(filenames, "", flags)
	fmt.Println(err)
	//PrettyPrint(entities)
	//for _, entity := range entities {
	//	for _, child := range entity.Children {
	//		fmt.Println(child.Name)
	//		fmt.Println(child.Children)
	//	}
	//}
	//fmt.Printf("%+v\n", entities)

	envColors := strings.Split(os.Getenv("LS_COLORS"), ":")
	PrettyPrint(envColors)
}
