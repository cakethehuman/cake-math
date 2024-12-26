package Mylogger

import (
	_ "fmt"
	"log"
)

func Loginfo(message string) {
	log.Printf("Hello, %v ", message)
}
