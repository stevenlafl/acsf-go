package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"./api"
)

func displayGlobalHelp() {
	fmt.Printf("Usage: %s [OPTION]... COMMAND [ARGUMENT]...\n", os.Args[0])
	fmt.Println("Options: ")
	fmt.Println("  --environment=[dev*,test,prod]")
	fmt.Println("    Select an environment to act on")
	fmt.Println("Available commands:")
	fmt.Println("  stacks")
	fmt.Println("    Show available stacks for the environment")
	fmt.Println("  vcs [STACK ID]")
	fmt.Println("    Show available vcs refs and currently deployed vcs ref for the environment")
	fmt.Println("  updates")
	fmt.Println("    Show a list of updates for the environment, previous and currently running")
	fmt.Println("  status [UPDATE ID]")
	fmt.Println("    Show update status of a specific update for the environment")

}

func main() {

	environmentPtr := flag.String("environment", "dev", "The Acquia environment")
	proxyPtr := flag.String("proxy", "", "The SOCKS5 Proxy Address")
	apiUserNamePtr := flag.String("api_user", "", "The API Username")
	apiPasswordPtr := flag.String("api_pass", "", "The API Password")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 || args[0] == "help" {
		if len(args) > 1 {
			switch args[1] {
			case "help":
				fmt.Println("Displays the help")
				break
			default:
				fmt.Printf("Command %s not found\n", args[1])
			}

			os.Exit(1)
		}
		displayGlobalHelp()
		os.Exit(1)
	}

	// @todo replace with a running dynamic list

	var APIUrls map[string]string = api.GetAPIUrls()
	if _, ok := APIUrls[*environmentPtr]; !ok {
		fmt.Printf("No such environment: %s\n", *environmentPtr)
		os.Exit(1)
	}

	switch args[0] {
	case "stacks":
		stackCommand(*environmentPtr, args)
		break
	case "vcs":
		vcsCommand(*environmentPtr, args)
		break
	case "updates":
		updateCommand(*environmentPtr, args)
		break
	case "status":
		statusCommand(*environmentPtr, args)
		break
	default:
		fmt.Printf("Command %s not found\n", args[0])
	}

	os.Exit(1)

	// @todo Make these work
	fmt.Printf("%s\n", args[0])
	fmt.Printf("%s\n", *proxyPtr)
	fmt.Printf("%s\n", *apiUserNamePtr)
	fmt.Printf("%s\n", *apiPasswordPtr)

}

func stackCommand(environment string, args []string) {
	var stacks []api.Stack = api.GetStacks(environment)

	for _, element := range stacks {
		fmt.Printf("[Stack %d] %s\n", element.StackID, element.StackName)
	}
}
func vcsCommand(environment string, args []string) {

	if len(args) < 2 {
		fmt.Println("status: missing [STACK ID]")
		os.Exit(1)
	}

	id, _ := strconv.Atoi(args[1])

	var vcs api.VCS = api.GetVCS(environment, id)
	fmt.Printf("Available VCS Refs on Stack %d:\n", id)

	for _, element := range vcs.Available {
		fmt.Printf("  %s\n", element)
	}

	fmt.Printf("Current VCS Ref: %s\n", vcs.Current)
}
func updateCommand(environment string, args []string) {
	var updates []api.Update = api.GetUpdates(environment)

	for _, element := range updates {
		fmt.Printf("[Update %d] Added %d. Status Code: %d\n", element.UpdateID, element.Added, element.Status)
	}
}
func statusCommand(environment string, args []string) {

	if len(args) < 2 {
		fmt.Println("status: missing [UPDATE ID]")
		os.Exit(1)
	}

	id, _ := strconv.Atoi(args[1])

	var updateStatus api.UpdateStatus = api.GetUpdateStatus(environment, id)
	fmt.Printf("[Update %d] %d%% - %s - Started: %d. Ended: %d\n", updateStatus.UpdateID, updateStatus.Percentage, updateStatus.Message, updateStatus.StartTime, updateStatus.EndTime)
}
