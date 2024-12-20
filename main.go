package main

import (
	"ToDo-CLI/todofile"
	"ToDo-CLI/work"
	"encoding/json"
	"flag"
	"os"
)

const (
	ADD    = "add"
	CLEAR  = "cl"
	LIST   = "ls"
	DELETE = "rm"
)

func main() {
	//flag를 파싱한다.
	//action := flag.String("action", "default", "a string")
	todoName := flag.String("todo", "defaultTodo", "Something to do")
	flag.Parse()
	//fmt.Println(*todoName)

	action := os.Args[len(os.Args)-1]
	//action 플래그에 따라 다른 다른 행동을 한다.
	switch action {
	case ADD:
		work.Works = work.AppendWork(*todoName)
	case CLEAR:
		work.Works = work.ClearWork(*todoName)
	case LIST:
		work.ListWork()
	case DELETE:
		work.DeleteWork(*todoName)
	}

	marshal, err := json.Marshal(*work.Works)
	if err != nil {
		panic(err)
	}

	savefile, err := os.Create(todofile.Filepath)
	if err != nil {
		panic(err)
	}
	defer savefile.Close()
	if savefile != nil {
		_, writeErr := savefile.Write(marshal)
		if writeErr != nil {
			panic(writeErr)
		}
	}

}
