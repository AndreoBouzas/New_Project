package main

import (
	"fmt"
	"log"

	Logs "github.com/AndreoBouzas/NewProject/logs"
	Reader "github.com/AndreoBouzas/NewProject/reader"
)

func main() {
	logthis := Logs.InitLogs()
	fmt.Println("informe o padrão de pesquisa do texto:")
	var patern string
	fmt.Scanln(&patern)
	//fmt.Println(Reader.Reader())
	if patern == "palavras" {
		verifiedtext := Reader.Reader(patern)
		for i := 0; i < len(verifiedtext); i++ {
			log.Println((verifiedtext[i]), logthis)
		}
		log.Println(Reader.ReadTimer(patern), logthis)
		fmt.Println(Reader.ReadTimer(patern))
		fmt.Println("o mesmo contém ", len(verifiedtext), patern)
	}

}
