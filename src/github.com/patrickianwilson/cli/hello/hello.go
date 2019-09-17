package main

import (
	"fmt"
	"github.com/patrickianwilson/cli/cmd"
	"github.com/patrickianwilson/cli/common"
	"github.com/patrickianwilson/cli/messages"
	"github.com/patrickianwilson/cli/properties"
	"log"
	"os"
	"path/filepath"
)

const (
	SUCCESS          int = iota
	INVALID_ARGS     int = iota
	INVALID_CMD_ARGS int = iota
	UNKNOWN_ERROR    int = iota
)

const UserDirEnvKey = "HOME"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(messages.GENERIC_USAGE_MESSAGE)
		os.Exit(INVALID_ARGS)
	}
	args := os.Args[1:]
	context := &common.Context{Config: properties.NewHardCodedConfigProperties()}

	profileName := findProfile(args)

	//attempt to load config from userdir
	userDir := os.Getenv(UserDirEnvKey)
	context.Config.OverlayConfigFromFile(filepath.FromSlash(
		fmt.Sprintf("%s/.config/%s.config", userDir, profileName)))

	var command common.Cmd = nil
	group := args[0]
	cmdName := ""
	justPrintHelp := false
	if group == "help" {
		if len(args) < 2 {
			//only generic help possible
			printGenericHelp()
			os.Exit(SUCCESS)
		}
		group = args[1]
		if len(args) < 3 {
			//generic help for a command group
			printHelpForGroup(group)
			os.Exit(SUCCESS)
		}
		group = args[1]
		cmdName = args[2]
		justPrintHelp = true

	}

	if !justPrintHelp {
		cmdName = args[1]
	}
	//find command
	for _, cmd := range provideCommands() {
		if cmd.GetCmdGroup() == group && cmd.GetCmdName() == cmdName {
			command = cmd
		}
	}

	if command == nil {
		//invalid command or group
		fmt.Printf("Unable to find a command for group %s and name %s.\n", group, cmdName)
		printHelpForGroup(group)
		os.Exit(INVALID_ARGS)
	}

	initErr := command.Init(args[2:], context)

	if initErr != nil {
		fmt.Printf("Hmm, There was a problem initializing the command: %s. This is probably a bug!", initErr.Error())
		os.Exit(UNKNOWN_ERROR)
	}

	if justPrintHelp {
		command.PrintHelp()
		os.Exit(SUCCESS)
	}

	validateErr := command.Validate()

	if validateErr != nil {
		fmt.Printf("Invalid command! %s\n\n", validateErr.Error())
		command.PrintHelp()
		os.Exit(INVALID_CMD_ARGS)
	}

	err := command.DoAction()

	if err != nil {
		log.Fatalf("%s\n\n", err.Error())
		command.PrintHelp()
	}
}

func provideCommands() []common.Cmd {
	return []common.Cmd{
		cmd.NewSampleCmd(),
	}
}
func findProfile(args []string) string {
	profile := "default"

	for i, arg := range args {
		if arg == "--profile" || arg == "-p" {
			if len(args) < i+1 {
				//malformed
				fmt.Printf("Invalid Argument!  --profile requires a value!")
			}
			profile = args[i+1]
			break
		}
	}
	return profile
}

func printHelpForGroup(group string) {
	//todo fill me in
	fmt.Printf("Help for Group %s\n", group)
}

func printGenericHelp() {
	//todo fill me in
	fmt.Printf("Generic Help\n")
}
