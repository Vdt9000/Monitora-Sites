// go run main.go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	intro()
	for {
		exibeMenu()
		comando := LeituraComando()

		switch comando {
		case 1:
			initMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida")
			os.Exit(-1)

		}

	}
}

func intro() {
	nome := "Victor"
	versao := 1.2
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)

}

func exibeMenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}
func LeituraComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)

	return comando
}

func initMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://www.alura.com.br", "https://chat.openai.com/", "https://discord.com/"}

	for i, site := range sites {
		fmt.Println(sites[i])
		fmt.Println("Testando site", i, ":", site)
		testSite(site)
	}

}

func testSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

}
