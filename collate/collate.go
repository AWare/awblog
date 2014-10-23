package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type post struct {
	filename string
	picture  string
	date     time
	content  string
}

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal("ERROR FILE HWAT")
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
