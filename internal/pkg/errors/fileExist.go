package errors

import "fmt"

var FileIsExist = fmt.Errorf("file exists")

type FileExist struct {
	FileName string
}

func (f FileExist) Error() string {
	return FileIsExist.Error()
}
