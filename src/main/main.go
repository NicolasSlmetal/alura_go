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
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}

func monitor() {
	websites := readSitesFromFile()
	for _, website := range websites {
		monitorSite(website)
		time.Sleep(5 * time.Second)
	}
}

func monitorSite(website string) {
	resp, err := http.Get(website)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if resp.StatusCode == 200 {
		fmt.Println(website, "is online")
	} else {
		fmt.Println(website, "is offline")
		fmt.Println("Status code:", resp.StatusCode)
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	if err != os.ErrNotExist {
		sites, err = createSitesFile()
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		return sites
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

	return defaultSites, nil
}
