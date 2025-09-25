package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	const intervalo = 5
	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
		for {
			monitorar()
			time.Sleep(intervalo * time.Second) // Espera 5 segundos entre as verificações
		}
	case 2:
		fmt.Println("Exibindo Logs...")
		imprimeLogs()
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}
func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}
func monitorar() {
	sites := leSitesDoArquivo()
	for _, site := range sites {
		resposta, _ := http.Get(site)
		if resposta.StatusCode == 200 {
			registraLog(site, true)
		} else {
			registraLog(site, false)
		}
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

// restante do código omitido

func imprimeLogs() {

	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
