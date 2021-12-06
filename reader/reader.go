package Reader

import (
	"io/ioutil"
	"log"

	Utils "github.com/AndreoBouzas/NewProject/utils"
)

var path string = "files/teste.txt"

func InitFile(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %s", err)
	}

	strifile := string(file)

	return strifile
}

func Reader(patern string) [][]string {
	wordscounter := `\w+'*[[:word:]]*`
	digitscounter := `\d`
	var newfile [][]string
	if patern == "digitos" {
		thisfile := InitFile(path)
		newfile = Utils.FindInText(thisfile, digitscounter)
	} else if patern == "palavras" {
		thisfile := InitFile(path)
		newfile = Utils.FindInText(thisfile, wordscounter)
	} else {
		log.Fatalf("padrão de pesquisa incorreto: %s", patern)
	}

	return newfile

}
