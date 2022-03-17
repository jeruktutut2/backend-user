package exception

import "log"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func LogFatallnIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
