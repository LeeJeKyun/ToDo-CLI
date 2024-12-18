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

func init() {
	file = jsonFileRead()
	works = jsonFileUnMarshal(file)
}

func main() {
	s := os.Args[1]
	switch s {
	case ADD:
		todoName := flag.String("name", "X", "Todo Name")
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

	savefile, err := os.Create("/Users/jekyunlee/json/todoData.json")
	if err != nil {
		panic(err)
	}
	defer savefile.Close()
	savefile.Write(marshal)

}

func jsonFileRead() []byte {
	readFile, err := os.ReadFile("/Users/jekyunlee/json/todoData.json")
	if err != nil {
		os.Create("/Users/jekyunlee/json/todoData.json")
	}
	return readFile
}

func jsonFileUnMarshal(file []byte) *[]structs.Work {
	works := new([]structs.Work)
	json.Unmarshal(file, &works)
	return works
}

func appendWork(works *[]structs.Work, name string) *[]structs.Work {
	w := *works
	result := append(w, structs.Work{name, false})
	return &result
}
