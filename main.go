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
	s := pkg.ReadFile(linuxPath)
	switch {
	case *allApp:
		runAllCommand(s)
	case *changeApp != "":
		runOneCommand(s, *changeApp)
	}
}

func runOneCommand(s []pkg.App, goal string) {
	fmt.Printf("start deploy %s service\n", goal)
	for i := range s {
		if s[i].AppName == goal {
			runCommand(s[i].Command, s[i].AppName)
			break
		}

	}
	fmt.Printf("%s override", goal)
}

func runAllCommand(s []pkg.App) {
	fmt.Println("start deploy all service")
	for i := range s {
		runCommand(s[i].Command, s[i].AppName)
	}
	fmt.Println("All override")
}

func runCommand(command []string, appName string) {
	barName := fmt.Sprintf(appName + ".bar")
	for _, elem := range command {
		commandForOverride := fmt.Sprintf(override + " -b " + barName + " -m " + elem + " -k " + appName)
		fmt.Println(commandForOverride)
		cmd := exec.Command("bash", "-c", commandForOverride)
		err := cmd.Run()
		fmt.Println(err)
	}
}
