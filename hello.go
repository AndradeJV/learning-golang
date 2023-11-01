package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	//niceToMeetYou()

	for {
		showMenu()
		chose := readCommand()
	
		switch chose {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("EXIBINDO LOGS")
		case 3:
			fmt.Println("SESSÃO FINALIZADA")
	
			os.Exit(0)
		default:
			fmt.Println("Comando não existe!")
	
			os.Exit(-1)
		}
	}

}

func niceToMeetYou() {
	var name string

	fmt.Println("Olá, seja bem vindo. Digite seu nome: ")
	fmt.Scan(&name)

	fmt.Println("Bem vindo", name)
}

func readCommand() int {
	var chose int

	fmt.Scan(&chose)
	fmt.Println("The commnad chosen was:", chose)

	return chose
}

func showMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Finalizar sessão")
}


func startMonitoring(){
	fmt.Println("MONITORAMENTO INICIADO")

	site := "https://www.mercadolivre.com.br/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, "está apresentando problemas. Erro: ", resp.StatusCode)
	}
}