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

func showMenu() {
	fmt.Println("1 - Monitor")
	fmt.Println("2 - Logs")
	fmt.Println("3 - Exit")
}

func main() {
	for {
		showMenu()

		var option int
		fmt.Print("Choose an option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Println("Monitoring...")
			monitor()
		case 2:
			fmt.Println("Showing logs...")
			showLogs()
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}

func monitor() {
	websites := readLinesFromFile("sites.txt")
	if websites == nil {
		websites, _ = createSitesFile()
	}
	for _, website := range websites {
		monitorSite(website)
		time.Sleep(5 * time.Second)
	}
}

func showLogs() {
	logs := readLinesFromFile("log.txt")
	if logs == nil {
		fmt.Println("No logs to show")
		return
	}
	for _, log := range logs {
		fmt.Println(log)
	}
	time.Sleep(5 * time.Second)
}

func monitorSite(website string) {
	resp, err := http.Get(website)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if resp.StatusCode == 200 {
		fmt.Println(website, "is online")
		registerLog(website + " was online at that time")
	} else {
		fmt.Println(website, "is offline")
		fmt.Println("Status code:", resp.StatusCode)
		registerLog(website + " was offline at that time")
	}
}

func readLinesFromFile(fileName string) []string {
	var sites []string

	file, err := os.Open(fileName)
	if err == os.ErrNotExist {
		return nil
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	err = nil
	for {

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		sites = append(sites, strings.TrimSpace(line))
	}

	return sites
}

func createSitesFile() ([]string, error) {
	defaultSites := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.reddit.com",
	}
	file, err := os.Create("sites.txt")

	defer file.Close()

	for _, site := range defaultSites {
		_, err = file.WriteString(site + "\n")
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}
	registerLog("Default sites file created")
	return defaultSites, nil
}

func registerLog(logMessage string) {
	log := "[" + time.Now().Local().Format(time.DateTime) + "] - " + logMessage

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	file.WriteString(log + "\n")
}
