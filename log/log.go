package log

import baseLog "log"

func CheckError(err error) {
	if err != nil {
		baseLog.Fatalln(err.Error())
	}
}
