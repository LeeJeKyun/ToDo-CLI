package work

import (
	"ToDo-CLI/todofile"
	"encoding/json"
	"fmt"
)

type Work struct {
	Name  string
	Clear bool
}

var Works *[]Work

func init() {
	//파일을 읽어온 []byte를 json형태로 언마샬링 후 Works 변수에 저장
	Works = JsonFileUnMarshal(todofile.Myfile)
}

func AppendWork(works *[]Work, name string) *[]Work {
	w := *works
	result := append(w, Work{name, false})
	return &result
}

func ClearWork(works *[]Work, name string) *[]Work {
	w := *works
	for i, v := range w {
		if v.Name == name {
			w[i].Clear = true
		}
	}
	return &w
}

func ListWork(works *[]Work) {
	for i, n := range *works {
		var ox string
		if n.Clear {
			ox = "O"
		} else {
			ox = "X"
		}
		fmt.Println(i+1, "번째 할 일:", n.Name, "[", ox, "]")
	}
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
