package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func menu() {
	versao := 1.0
	fmt.Println("------------------------------------------------")
	fmt.Println("Este programa atualmente está na versão:", versao)
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- exibir Logs")
	fmt.Println("0- sair do Programa")
	fmt.Println("Digite o comando desejado:")
}

func lerComando() int {
	var comando int
	fmt.Scanln(&comando)
	fmt.Println("Você digitou o comando:", comando)

	return comando
}

func monitoramento() {
	fmt.Println("Monitorando...")

	sites := lerAquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			fmt.Println("Testando o site:", site)
			testar(site)
		}
		fmt.Println("")
		fmt.Println("Aguarde 5 segundos para o próximo teste...")
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("Monitoramento finalizado!")

}

func logs() {
	fmt.Println("Exibindo logs...")
}

func testar(site string) {
	request, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
	}

	if request.StatusCode == 200 {
		fmt.Println("Site:", site, "está no ar!")
	} else {
		fmt.Println("Site:", site, "está fora do ar!", request.StatusCode)
	}
}

func lerAquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return sites
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		if linha != "" {
			sites = append(sites, linha)
		}

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func main() {

	for {
		menu()
		comando := lerComando()

		switch comando {
		case 1:
			monitoramento()
		case 2:
			logs()
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(255)
		}
	}
}
