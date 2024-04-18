// hello.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	//"net/http"
)

func main() {
	exibeNomes()
	//intro()
	for {
		//	exibeMenu()
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

//nome, idade := devolveNomeEIdade()
//fmt.Println(nome, "E tenho ", idade, "anos")

/*if comando == 1 {
	fmt.Println("Monitorando...")
} else if comando == 2 {
	fmt.Println("Exibindo logs")
} else if comando == 0 {
	fmt.Println("Saindo do programa...")
} else {
	fmt.Println("Opção inválida")
}*/

/*func devolveNomeEIdade() (string, int) {
	nome := "Victor"
	idade := 19
	return nome, idade
}*/

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
	var sites [2]string
	sites[0] = "https://youtube.com.br/"
	sites[1] = "https://pkg.go.dev/std"

	fmt.Println(sites)
	site := "https://dev.conectahub360.com.br/portal/services/in-progress"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!!!")
	} else {
		fmt.Print("O site:", site, "está com problemas e não foi possível completar a request. Status Code:", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{
		"Victor", "Vitinho", "Vini", "Nick"}
	nomes = append(nomes, "Luan")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("Seu slice tem", len(nomes))
}
