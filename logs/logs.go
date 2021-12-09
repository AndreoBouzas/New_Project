package Logs

//Import dos Pkgs necessários
import (
	"fmt"
	"log"
	"os"
)

//Definição da variável de caminho do arquivo de Logs
var filepath = "files/log.txt"

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
	//retorna nulo
	return nil
}
