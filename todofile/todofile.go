package todofile

import (
	"os"
)

var Myfile []byte
var Filepath string

func init() {
	//파일경로는 UserHomDir/todoData.json
	Filepath, _ = os.UserHomeDir()
	Filepath += "/todoData.json"
	//해당 경로의 파일 읽어온 후 전역변수인 file에 저장
	Myfile = JsonFileRead(Filepath)
}

func JsonFileRead(filePath string) []byte {
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
