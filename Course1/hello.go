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

func main() {
	readFile()
	for vali == false {
		handleIntro()
		showMenu()

		command := handleScan()

		switch command {
		case 1:
			startCustomMonitoring()
		case 2:
			startMonitoring()
		case 3:
			file, _ := ioutil.ReadFile("log.txt")
			if file != nil {
				fmt.Println("Exibindo Logs")
				fmt.Println()
			}
			logs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}
}

func contM() int {
	var command int
	fmt.Scan(&command)
	return command
}

func handleIntro() {
	version := 1.2
	fmt.Println("Este programa está na versao:", version)
	fmt.Println()
}

func handleScan() int {
	var command int
	fmt.Scan(&command)
	return command
}

func showMenu() {

	fmt.Println("1- Iniciar monitoramento com URL personalizada")
	fmt.Println("2- Iniciar monitoramento predefinido")
	file, _ := ioutil.ReadFile("log.txt")
	if file != nil {
		fmt.Println("3- Exibir Logs")
	}
	fmt.Println("0- Sair do programa")
}

func startCustomMonitoring() {
	fmt.Println("Digite uma url: ")
	var webSite string
	fmt.Scan(&webSite)
	res, _ := http.Get("http://" + webSite)
	fmt.Println("Monitorando...")
	if res.StatusCode == 200 {
		fmt.Println()
		fmt.Println(webSite, "foi carregado com sucesso")
		regLog(webSite, true)
	} else if res.StatusCode != 200 {
		fmt.Println()
		fmt.Println(webSite, "está com problemas. Status Code:", res.StatusCode)
		regLog(webSite, false)
	} else {
		os.Exit(-1)
	}
	otherMonitorin()
}

func startMonitoring() {
	webSites := readFile()

	for i := 0; i < len(webSites); i++ {
		res, err := http.Get("http://" + webSites[i])
		if err != nil {
			fmt.Println("Ocorreu um erro: ", err)
			// os.Exit(-1)
		}
		fmt.Println()
		fmt.Println("Monitorando...")
		if res.StatusCode == 200 {
			fmt.Println()
			fmt.Println(webSites[i], "foi carregado com sucesso")
			regLog(webSites[i], true)
		} else if res.StatusCode != 200 {
			fmt.Println()
			fmt.Println(webSites[i], "está com problemas. Status Code:", res.StatusCode)
			regLog(webSites[i], false)
		} else {
			os.Exit(-1)
		}
	}
	otherMonitorin()
}

var vali bool

func otherMonitorin() {
	fmt.Println()
	fmt.Println("1- Retornar ao menu principal")
	fmt.Println("0- Sair do programa")
	var otherCommand int
	fmt.Scan(&otherCommand)
	if otherCommand == 0 {
		fmt.Println("Saindo do programa...")
		vali = true
	}
}

func readFile() []string {

	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		// fmt.Println(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func regLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		os.Exit(-1)
	}
	currentTime := time.Now().Format("02/01/2006 15:04:05")

	file.WriteString(site + " online: " + strconv.FormatBool(status) + " " + currentTime + "\n")

}

func logs() {
	//condi
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		os.Exit(-1)
	}

	fmt.Println(string(file))
}
