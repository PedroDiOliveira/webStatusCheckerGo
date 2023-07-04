package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 1
const delay = 5 //segundos

func exibeIntroducao() {
	nome := "Pedro"
	fmt.Println(" Ola sr.", nome)
}
func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}
func exibeMenu() {
	fmt.Println("1-Iniciar Monitoramento")
	fmt.Println("2-Exibir Logs")
	fmt.Println("0-Sair do Programa")
}
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//var sites = []string{"http://github.com", "https://www.youtube.com", "https://openai.com"}
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i+1, "...")
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro")
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "esta com problema, o StatusCode:", resp.Status)
		registraLog(site, false)
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
		if err == io.EOF { /// EOF = end of file, entao acontecera um erro logo apos a ultima linha de do arquivo, sendo assim break
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
}
func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}

func main() {

	exibeIntroducao()
	registraLog("Site Falso", false)

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
			fmt.Println(" ")
		case 2:
			fmt.Println("Exibindo Logs...")
			fmt.Println(" ")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco esse comando")
			os.Exit(-1)
		}

	}
}
