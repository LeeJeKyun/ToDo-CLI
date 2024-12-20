package work

import (
	"ToDo-CLI/todofile"
	"encoding/json"
	"fmt"
)

type Work struct {
	Todo  string
	Clear bool
}

var Works *[]Work

func init() {
	//파일을 읽어온 []byte를 json형태로 언마샬링 후 Works 변수에 저장
	Works = JsonFileUnMarshal(todofile.Myfile)
}

func AppendWork(Todo string) *[]Work {
	w := *Works
	result := append(w, Work{Todo, false})
	return &result
}

func ClearWork(Todo string) *[]Work {
	w := *Works
	for i, v := range w {
		if v.Todo == Todo {
			w[i].Clear = true
		}
	}
	return &w
}

func ListWork() {
	for i, n := range *Works {
		var ox string
		if n.Clear {
			ox = "V"
		} else {
			ox = " "
		}
		fmt.Printf("%d). [%s] %s\n", i+1, ox, n.Todo)
	}
}

func DeleteWork(Todo string) {
	w := *Works
	for i, v := range w {
		if v.Todo == Todo {
			w = RemoveIndex(w, i)
		}
	}
	Works = &w
}

func JsonFileUnMarshal(file []byte) *[]Work {
	//myfile 내용물을 담기 위한 빈 Slice 생성
	result := []Work{}
	//result2 := structs.Work{}
	//file이라는 []byte 데이터를 result 슬라이스에 언마샬링
	unmarshalErr := json.Unmarshal(file, &result)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}
	return &result
}

func RemoveIndex(s []Work, index int) []Work {
	return append(s[:index], s[index+1:]...)
}
