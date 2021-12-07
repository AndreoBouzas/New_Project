package Reader

import (
	"fmt"
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
	wordscounter := `\S+`
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
func ReadTimer(patern string) string {

	x := 300
	y := len(Reader(patern))
	var tempo string
	if y > x {
		timer := y / x
		if y%x != 0 {
			timerstring := fmt.Sprintf("%#v", timer)
			tempo = "Serão necessários " + timerstring + " minuto(s) em média, para ler este texto!"
			return tempo
		}
	} else if y == x {
		tempo := "Será necessário 1 minuto em média, para ler este texto!"
		return tempo
	} else {
		tempo := "Serão necessários alguns segundos para ler este texto!"
		return tempo
	}
	return tempo
}
