package Reader

//Import dos Pkgs necessários
import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	//Import do Pkg Utils do projeto
	Utils "github.com/AndreoBouzas/NewProject/utils"
)

//Declaração da vaiável de caminho do arquivo
var path string = "files/teste.txt"

//Função InitFile recebe um path, procura e inicializa o arquivo contido no path informado
func InitFile(path string) string {
	//Declarando a variável que recebe o arquivo lido
	file, err := ioutil.ReadFile(path)
	//Verificando se houve algum erro durante a leitura do arquivo
	if err != nil {
		//Caso ocorra algum erro, é gravado um log do erro!
		log.Fatalf("Erro ao abrir o arquivo: %s", err)
		//Print Visual para informar a ocorrência do erro!
		fmt.Printf("Erro ao abrir o arquivo: %s", err)
	}
	//Conversão do file para o tipo String
	strifile := string(file)

	//Retorno da variável contendo o texto convertido em string
	return strifile
}

//Função Reader faz a leitura e aplica o regexp baseado no padrão de pesquisa informado na chamada da mesma
func Reader(patern string) [][]string {

	var (
		//Declaração da variável de padrão de pesquisa para palavras
		wordscounter = `\S+`
		//Declaraçãod da variável de padrão de pesquisa para dígitos
		digitscounter = `\d`
		//Declaração da variável de retorno da função com seu devido tipo
		newfile [][]string
	)

	//Verificação do padrão de pesquisa informado na chamada da função
	//Caso tenha sido passado o padrão "dígitos"
	if patern == "digitos" {
		//Declaração da variável que recebe o retorno da função InitFile baseada no path padrão
		thisfile := InitFile(path)
		//Declaração da variável que recebe o retorno da função FindInText recebendo como argumento
		//A variável thisfile e como padrão de regexp a variável digitscounter
		newfile = Utils.FindInText(thisfile, digitscounter)
		//Caso tenha sido passado o padrão "palavras"
	} else if patern == "palavras" {
		//Declaração da variável que recebe o retorno da função InitFile baseada no path padrão
		thisfile := InitFile(path)
		//Declaração da variável que recebe o retorno da função FindInText recebendo como argumento
		//A variável thisfile e como padrão de regexp a variável wordscounter
		newfile = Utils.FindInText(thisfile, wordscounter)
		//Caso o padrão de pesquisa esteja divergente dos casos esperados
	} else {
		//Gera um log de erro de padrão de pesquisa
		log.Fatalf("padrão de pesquisa incorreto: %s", patern)
		//Print Visual do erro de padrão de pesquisa
		fmt.Printf("padrão de pesquisa incorreto: %s", patern)
	}
	//retorno da variável preenchida com o texto requerido de acordo com o padrão identificado
	return newfile

}

//A função ReadTimer faz  uma contagem de tempo estimado de leitura baseada na quantidade de palavras
//Contidas no texto, em comparação com o tempo médio de leitura (300 palavras por minuto)
func ReadTimer(patern string) string {
	//Definição da variável com a quantidade média de palavras lidas em um minuto
	mediaPadrao := 300.0
	//Definição da variável com a quantidade de palavras contidas no texto analisado
	palavrasDoTexto := float64(len(Reader(patern)))
	//Definição da vairável que recebera o tempo estimado
	var tempo string
	//Se a quantidade de palavras for maior que a quantidade média padrão
	if palavrasDoTexto > mediaPadrao {
		//A variável de timer recebe o resultado da divisão de palavras do texto por média padrão
		timer := palavrasDoTexto / mediaPadrao
		//tranforma em string o valor todo
		timerstring := fmt.Sprintf("%2.2f", timer)
		//divide o valor string pelo ponto "."
		splittimed := strings.Split(timerstring, ".")
		//guarda a informação dos segundos
		timersegundos := fmt.Sprint(splittimed[1])
		//guarda a informação dos minutos
		mintimer := fmt.Sprint(splittimed[0])
		//adiciona um zero string no início da variável dos segundos
		timersegundoscomzero := "0." + timersegundos
		//converte para float a variável dos segundos
		floatsegtimer, _ := strconv.ParseFloat(timersegundoscomzero, 64)
		//converte para float a variável dos minutos
		floatmintimer, _ := strconv.ParseFloat(mintimer, 64)
		//divide por 60 a variável dos segundos
		calculatedsegs := floatsegtimer / 0.60
		//calcula o tempo
		timer = floatmintimer + calculatedsegs
		//Guarda na vairável o valor de tempo formatado
		finaltimer := fmt.Sprintf("%0.2f", timer)
		//A variável "tempo" recebe a formatação com as devidas informações
		tempo = "Serão necessários " + finaltimer + " minuto(s) em média, para ler este texto!"
		//retorna a variável tempo
		return tempo
		//Se a quantidade de palavras for igual a quantidade média padrão
	} else if palavrasDoTexto == mediaPadrao {
		//A variável "tempo" recebe a formatação com as devidas informações
		tempo := "Será necessário 1 minuto em média, para ler este texto!"
		//Retorna a variável tempo
		return tempo
		//Se a quantidade de palavras for menor que a quantidade média padrão
	} else {
		//A variável "tempo" recebe a formatação com as devidas informações
		tempo := "Serão necessários alguns segundos para ler este texto!"
		//Retorna a variável tempo
		return tempo
	}
}
