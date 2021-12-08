package Utils

//Import do Pkg Regexp
import "regexp"

//Criação da função de procura no texto baseada no regexp
//Definie-se que a função receberá um alvo, um padrão de regexp
func FindInText(target, patern string) [][]string {
	//definindo o padrão informado como base de procura no texto
	paterntext := regexp.MustCompile(patern)
	//passando o regexp no texto e salvando o resultado em uma variável
	cleartext := paterntext.FindAllStringSubmatch(target, -1)
	//Definindo o retorno da função
	return cleartext
}
