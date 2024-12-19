package main

import (
	"ToDo-CLI/todofile"
	"ToDo-CLI/work"
	"encoding/json"
	"flag"
	"os"
)

const (
	ADD   = "add"
	CLEAR = "clear"
	LIST  = "ls"
)

func main() {

	//flag를 파싱한다.
	action := flag.String("action", "123", "a String")
	todoName := flag.String("work", "X", "a string")
	flag.Parse()

	//fmt.Println("main - 1", *action)
	//s := os.Args[len(os.Args)-1]
	//action 플래그에 따라 다른 다른 행동을 한다.
	switch *action {
	case ADD:
		work.Works = work.AppendWork(work.Works, *todoName)
	case CLEAR:
		work.Works = work.ClearWork(work.Works, *todoName)
	case LIST:
		work.ListWork(work.Works)
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
