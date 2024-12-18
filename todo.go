package main

import (
	"ToDo-CLI/structs"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const (
	ADD   = "add"
	CLEAR = "clear"
	LIST  = "ls"
)

var file []byte
var works *[]structs.Work
var filePath string

func init() {
	filePath, _ = os.UserHomeDir()
	filePath += "/todoData.json"
	file = jsonFileRead()
	works = jsonFileUnMarshal(file)
}

func main() {

	action := flag.String("word", "123", "a String")
	todoName := flag.String("fork", "X", "a string")
	flag.Parse()

	fmt.Println(*action)
	//s := os.Args[len(os.Args)-1]
	switch *action {
	case ADD:
		works = appendWork(works, *todoName)
	case CLEAR:

	case LIST:
		for _, i := range *works {
			fmt.Println(i)
		}
	}

	marshal, err := json.Marshal(*works)
	if err != nil {
		panic(err)
	}

	savefile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer savefile.Close()
	if savefile != nil {
		savefile.Write(marshal)
	}

}

func jsonFileRead() []byte {
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		os.Create(filePath)
	}
	return readFile
}

func jsonFileUnMarshal(file []byte) *[]structs.Work {
	works := make([]structs.Work, 10)
	json.Unmarshal(file, works)
	return &works
}

func appendWork(works *[]structs.Work, name string) *[]structs.Work {
	w := *works
	result := append(w, structs.Work{name, false})
	return &result
}
