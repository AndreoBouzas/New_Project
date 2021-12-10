package Logs

//Import dos Pkgs necessários
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

//Definição da variável de caminho do arquivo de Logs
var filepath = "files/" + date + ".log.txt"
var year, month, day = time.Now().Date()
var date = strconv.Itoa(day) + "." + strconv.Itoa(int(month)) + "." + strconv.Itoa(year)

//A função InitLogs  é responsável por inicializar e proparar o arquivo de log.
func InitLogs() error {
	//A variável filelog recebe o arquivos contido  no caminho informado.
	//caso este arquivo não exista o mesmo é criado com as devidas permissões

	filelog, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	//verifica se o erro é diferente de nulo
	if err != nil {
		//Caso sim da um Print do erro
		fmt.Println("erro ao criar o arquivo de log", err)
	}
	//define como saída padrão de todos os logas o arquivo declarado
	log.SetOutput(filelog)

	//Log de inicialiação da análise
	log.Print("\n")
	log.Println("|======================|")
	log.Println("|      *Novo Log*      |")
	log.Println("|======================|")
	log.Print("\n")
	//retorna nulo
	return nil
}

func LogPersonalizado(patern string) error {
	filelog, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0777)
	//verifica se o erro é diferente de nulo
	if err != nil {
		//Caso sim da um Print do erro
		fmt.Println("Erro ao buscar o arquivo de log", err)
	}
	//define como saída padrão de todos os logas o arquivo declarado
	log.SetOutput(filelog)

	if patern == "palavras" {
		log.Print("\n")
		log.Println("|=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=|=-=-=-=-=-=-=-=-=-=-=-=-|")
		log.Println("|        *Padrão de pesquisa*   |       (Palavras)       |")
		log.Println("|=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=|=-=-=-=-=-=-=-=-=-=-=-=-|")
		log.Print("\n")

	} else if patern == "digitos" {
		log.Print("\n")
		log.Println("|=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=|=-=-=-=-=-=-=-=-=-=-=|")
		log.Println("|        *Padrão de pesquisa*     |      (Dígitos)      |")
		log.Println("|=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=|=-=-=-=-=-=-=-=-=-=-=|")
		log.Print("\n")

	} else {
		log.Print("\n")
		log.Println("|------------------------------------------|")
		log.Println("|        *Padrão não identificado*         |")
		log.Println("|------------------------------------------|")
		log.Print("\n")

	}

	return err
}
