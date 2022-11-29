package main

import (
	"autoDeployESB/pkg"
	"flag"
	"fmt"
	"os/exec"
)

const (
	override    = "mqsiapplybaroverride"
	node        = "ACEDS"
	server      = "solfy"
	windowsPath = "D:\\goProject\\autoDeployESB\\properties.txt"
	linuxPath   = "properties.txt"
)

func main() {
	allApp := flag.Bool("all", false, "deploy all service")
	changeApp := flag.String("n", "", "deploy selected app by name")
	flag.Parse()

	switch {
	case *allApp:
		fmt.Println("test")
	case *changeApp == "":
		fmt.Println("Your app is")
	}

	fmt.Println("Hello, DigitalOcean!")

	s := pkg.ReadFile(linuxPath)
	fmt.Println(s[1].AppName)
	fmt.Println(s[1].Command)

	runCommand(s[1].Command, s[1].AppName)

}

func runAllCommand(s []pkg.App) {
	for range s {

	}
}

func runCommand(command []string, appName string) {

	barName := fmt.Sprintf(appName + ".bar")
	fmt.Println(command[0])
	for _, elem := range command {
		commandForOverride := fmt.Sprintf(override + " -b " + barName + " -m " + elem + " -k " + appName)
		cmd := exec.Command("bash", "-c", commandForOverride)
		err := cmd.Run()
		fmt.Println(err)
	}

	//fmt.Println(commandForOverride)

}
