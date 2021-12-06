package main

import (
	"fmt"

	Reader "github.com/AndreoBouzas/NewProject/reader"
)

func main() {
	fmt.Println("informe o padrão de pesquisa do texto:")
	var patern string
	fmt.Scanln(&patern)
	//fmt.Println(Reader.Reader())

	rann := Reader.Reader(patern)
	for i := 0; i < len(rann); i++ {
		fmt.Printf("%#v \n", rann[i])
	}
	fmt.Println(len(rann))

}
