package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func runCommand(command string) {
	// Execute commands in terminal
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func startServices() {
	clearScreen()
	fmt.Println("Starting The Nginx with MySQL...")
	fmt.Println("--------------------------------->")
	runCommand("sudo systemctl enable nginx mysql.service")
	runCommand("sudo systemctl start nginx mysql.service")
	fmt.Println("---------------------------------")
	runCommand("sudo systemctl status nginx")
	fmt.Println("---------------------------------")
	runCommand("sudo systemctl status mysql.service")
}

func stopServices() {
	clearScreen()
	fmt.Println("Stopping The Nginx and MySQL...")
	fmt.Println("--------------------------------->")
	runCommand("sudo systemctl stop nginx mysql.service")
	runCommand("sudo systemctl disable nginx mysql.service")
	fmt.Println("---------------------------------")
	runCommand("sudo systemctl status nginx")
	fmt.Println("---------------------------------")
	runCommand("sudo systemctl status mysql.service")
}

func sitesEnableEnginx() {
		clearScreen() // Clear screen every time you enter the menu
		fmt.Println("Checking Sites-Enable Config...")
		fmt.Println("---------------------------------")
		time.Sleep(500 * time.Millisecond)
		runCommand("ls -al /etc/nginx/sites-enabled")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\nChecking Preps...")
		fmt.Println("---------------------------------")
		runCommand("sudo nginx -t")
		fmt.Println("---------------------------------")
		fmt.Println(" ")
		
		var subChoice string

		for {
			fmt.Print("\033[F Press 'q' to go back to menu: ")
			fmt.Scan(&subChoice)

			if subChoice == "q" {
				break
			} else {
				fmt.Println("\033[F Invalid, Please Try Again....")
				time.Sleep(700 * time.Millisecond)
				}
			}
}

func checkServices() {
	clearScreen()
	fmt.Println("Checking The Nginx and MySQL...")
	time.Sleep(1700 * time.Millisecond)
	fmt.Println("---------------------------------")
	runCommand("sudo systemctl status nginx.service")
	time.Sleep(800 * time.Millisecond)
	fmt.Println("---------------------------------")
	runCommand("sudo systemctl status mysql.service")
}

func monitorNginx() {
	clearScreen()
	fmt.Println("Monitoring Nginx...")
	time.Sleep(1000 * time.Millisecond)
	runCommand("sudo multitail /var/log/nginx/access.log /var/log/nginx/error.log")
}

func main() {
	for {
		clearScreen()
		fmt.Println("+===========+")
		fmt.Println("Nginx, Mysql Panel")
		fmt.Println("-------------")
		fmt.Println("EMP-Menu")
		fmt.Println("+===========+")
		fmt.Println("1. Turn On Engine")
		fmt.Println("2. Turn Off Engine")
		fmt.Println("3. Check Preps")
		fmt.Println("4. Check Engine")
		// If you want to use option number 4,
		// make sure you have previously used option number 1, or 1 then 2
		// or option number 4 will look like it's not working properly
		fmt.Println("5. View Logs")
		fmt.Println("0. Bail Out")
		fmt.Println("-------------")
		fmt.Println("Always press 'q' to back menu")
		fmt.Println("+===========+")

		var choice string
		fmt.Print("Choose an option (1/2/3/4/5/0): ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			startServices()
		case "2":
			stopServices()
		case "3":
			sitesEnableEnginx()
		case "4":
			checkServices()
		case "5":
			monitorNginx()
		case "0":
			fmt.Println("Exiting the program.")
			time.Sleep(400 * time.Millisecond)
			clearScreen()
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
			time.Sleep(400 * time.Millisecond)
		}
	}
}

