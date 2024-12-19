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

	//파일경로는 UserHomDir/todoData.json
	filePath, _ = os.UserHomeDir()
	filePath += "/todoData.json"
	//해당 경로의 파일 읽어온 후 전역변수인 file에 저장
	file = jsonFileRead()

	//파일을 읽어온 []byte를 json형태로 언마샬링 후 works 변수에 저장
	works = jsonFileUnMarshal(file)
}

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
		works = appendWork(works, *todoName)
	case CLEAR:
		works = clearWork(works, *todoName)
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
		_, writeErr := savefile.Write(marshal)
		if writeErr != nil {
			panic(writeErr)
		}
	}

}

func jsonFileRead() []byte {
	//filePath의 파일을 읽어서 반환
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		_, createErr := os.Create(filePath)
		if createErr != nil {
			panic(createErr)
		}
	}
	return readFile
}

func jsonFileUnMarshal(file []byte) *[]structs.Work {
	//file 내용물을 담기 위한 빈 Slice 생성
	result := []structs.Work{}
	//result2 := structs.Work{}
	//file이라는 []byte 데이터를 result 슬라이스에 언마샬링
	unmarshalErr := json.Unmarshal(file, &result)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}
	return &result
}

func appendWork(works *[]structs.Work, name string) *[]structs.Work {
	w := *works
	result := append(w, structs.Work{name, false})
	return &result
}

func clearWork(works *[]structs.Work, name string) *[]structs.Work {
	w := *works
	for i, v := range w {
		if v.Name == name {
			w[i].Clear = true
		}
	}

	return &w
}
