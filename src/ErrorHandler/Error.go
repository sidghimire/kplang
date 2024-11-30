package errorhandler

import "log"

func CheckError(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}
func LogError(message string) {
	log.Fatal(message)
}
