package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const monitoring = 3
const delay = 5

func main() {
	niceToMeetYou()

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
	fmt.Println("")

	return chose
}

func showMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Finalizar sessão")
}

func startMonitoring() {
	fmt.Println("MONITORAMENTO INICIADO")

	// sites := []string{"https://www.mercadolivre.com.br", "https://www.mercadolivre.com.br/naoacessa"}

	sites := readSitesFile()

	for i := 0; i < monitoring; i++ {
		for index, site := range sites {
			fmt.Println("Posição: ", index, "Site: ", site)

			siteTest(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func siteTest(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, "está apresentando problemas. Erro: ", resp.StatusCode)
	}
}

func readSitesFile() []string {
	var sites []string

	file, err := os.Open("./data/sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	for {
		reader := bufio.NewReader(file)
		line, err := reader.ReadString('\n')

		sites = append(sites, line)

		fmt.Println(line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}
