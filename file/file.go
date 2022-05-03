package file

import (
	"genius-lyrics-parser/log"
	"os"
)

func Write(filename string, content string) {
	f, createErr := os.Create(filename)
	os.Chmod(filename, 0777)
	log.CheckError(createErr)

	_, writeErr := f.WriteString(content)
	log.CheckError(writeErr)

	defer f.Close()
}
