package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var path string = "/home/andreo/Área de Trabalho/estudos_golang/The-Project/files/contatos.txt"

func initFile(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %s", err)
	}

	strifile := string(file)

	return strifile
}

func main() {

	fmt.Printf("%#v", initFile(path))

}
