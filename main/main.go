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

	rann := Reader.Reader(patern)
	for i := 0; i < len(rann); i++ {
		log.Println((rann[i]), logthis)
	}
	fmt.Println(len(rann))

}
