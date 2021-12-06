package Reader

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/AndreoBouzas/New_Project/utils"
)

var path string = "/home/andreo/Área de Trabalho/estudos_golang/New_Project/files/contatos.txt"

func InitFile(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %s", err)
	}

	strifile := string(file)

	return strifile
}

func Reader() {

	thisfile := InitFile(path)
	patern := `m`
	newfile := utils.FindInText(thisfile, patern)
	fmt.Println(newfile)

}
