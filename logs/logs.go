package Logs

import (
	"log"
	"os"
)

var filepath = "files/log.txt"

func InitLogs() error {
	filelog, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Println("erro ao criar o arquivo de log", err)
	}
	log.SetOutput(filelog)
	return nil
}
