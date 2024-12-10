package main

import (
	"fmt"
	"net/http"
	"os"
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
	websites := []string{"http://youtube.com.br", "http://www.google.com", "http://www.facebook.com"}
	for _, website := range websites {
		monitorSite(website)
		time.Sleep(5 * time.Second)
	}
}

func monitorSite(website string) {
	resp, _ := http.Get(website)
	if resp.StatusCode == 200 {
		fmt.Println(website, "is online")
	} else {
		fmt.Println(website, "is offline")
		fmt.Println("Status code:", resp.StatusCode)
	}
}
