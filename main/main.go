package main

import (
	"fmt"

	Logs "github.com/AndreoBouzas/NewProject/log"
	Reader "github.com/AndreoBouzas/NewProject/reader"
)

func main() {
	f := Logs.InitLogs()
	fmt.Println("informe o padrão de pesquisa do texto:")
	var patern string
	fmt.Scanln(&patern)
	//fmt.Println(Reader.Reader())

	rann := Reader.Reader(patern)
	for i := 0; i < len(rann); i++ {
		log.Fprintf("%#v \n", rann[i])
	}
	fmt.Println(len(rann))

}
