package main

//Import dos Pkgs necessários
import (
	"fmt"
	"log"
	"time"

	//Import do Pkg Logs e Reader do projeto
	Logs "github.com/AndreoBouzas/NewProject/logs"
	Reader "github.com/AndreoBouzas/NewProject/reader"
)

//A função Main, chama executa o projeto em si, além de coletar o padrão de pesquisa
func main() {
	//chamando a função de iicialização dos logs
	Logs.InitLogs()

	//Solicita ao usuário que informe o padrão de pesquisa
	fmt.Println("informe o padrão de pesquisa do texto:")
	//Log de coleta de input
	log.Println("| Coletando padrão de pesquisa informado pelo usuário  ")
	//cria a variável de padrão
	var patern string
	//A variável de padrão recebe o retorno do usuário
	fmt.Scanln(&patern)
	//chamando a função de log personalizado
	Logs.LogPersonalizado(patern)

	//se o padrão coletado for igual a "palavras"
	if patern == "palavras" {
		//A variável verifiedtext recebe o retorno da função Reader com o padrão informado pelo usuário
		verifiedtext := Reader.Reader(patern)
		//log de leitura do arquivo
		log.Println("|  Lendo arquivo na path informado!")
		//log de leitura do arquivo
		log.Println("|  Arquivo lido com sucesso!")
		//log de controle de regexp
		log.Println("|  Iniciando a coleta de dados baseda no padrão informado pelo usuário")
		//Abre um laço que roda enquanto o "i" for menor que a quantidade de palavras do texto analisado.
		var palavras [][]string
		for i := 0; i < len(verifiedtext); i++ {
			palavras = append(palavras, verifiedtext[i])
			//Faz log da palavra colatada no cliclo do laço
		}
		//Log de inicialização da contagem de palavras do texto
		log.Println("|  Iniciando o calculo de tempo estimado de leitura")
		log.Println("|  ", palavras)
		//Estrutura de laço simples de espera de 3 segundos
		for i := 0; i <= 5; i++ {
			//Print Visual para formatar a exibição do resultado
			time.Sleep(1 * time.Second)
			fmt.Print("\n")
			if i == 3 {
				//Print do resultado
				fmt.Println("|----------------------------------------------------------------------------------------------------------------|")
				fmt.Println("|   O texto analisado possui", len(verifiedtext), patern, "e", Reader.ReadTimer(patern), "   |")
				time.Sleep(1 * time.Second)
				fmt.Println("|----------------------------------------------------------------------------------------------------------------|")
				fmt.Print("\n")
			}
		}
		//log do resultado
		log.Println("|  O texto analisado possui", len(verifiedtext), patern, "e", Reader.ReadTimer(patern), "   |")

		//se o padrão de pesquisa for igual a "digitos"
	} else if patern == "digitos" {
		////a variável verifiedtext recebe o retorno da função Reader com o padrão informado pelo usuário
		verifiedtext := Reader.Reader(patern)
		//log de leitura do arquivo
		log.Printf("|  Lendo arquivo na path informado!")
		//Print Visual do resultado
		fmt.Println("Este texto contem", len(verifiedtext), patern, "sendo eles:")
		//log do resultado
		log.Println("Este texto contem", len(verifiedtext), patern, "sendo eles:")
		//Criando a variável que receberá os digitos
		var digitos [][]string
		//abre um laço que ocorre enquanto "i" for menor que a quantidade de digitos
		for i := 0; i < len(verifiedtext); i++ {
			// salvando na variável criada os dígitos coletados a cada loop do laço
			digitos = append(digitos, verifiedtext[i])
		}

		//Faz log da palavra colatada no cliclo do laço
		log.Println(digitos)
		//Print do resultado
		fmt.Printf("\n %v \n", digitos)

		//Se não for identificado nenhum padrão válido
	} else {
		//log do erro de padrão
		log.Println("Erro no padrão de analise informado", patern)
		//Print do erro de padrão
		fmt.Println("Erro no padrão de analise informado \"", patern, "\" Tente novamente!")
	}

}
