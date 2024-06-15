package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var availableCommandsMap = map[string]func([]string){
	"exit": exit,
	"echo": echo,
	"type": typeFunc,
}

var availableCommands = []string{
	"exit",
	"echo",
	"type",
}

func typeFunc(args []string) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, arg := range args {
		found := false

		if slices.Contains(availableCommands, arg) {
			fmt.Printf("%s is a shell builtin\n", arg)
			continue
		}

		for _, path := range paths {
			fp := filepath.Join(path, arg)
			if _, err := os.Stat(fp); err == nil {
				fmt.Printf("%s is %s\n", arg, fp)
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("%s: not found\n", arg)
		}
	}
}

func echo(args []string) {
	if len(args) == 0 {
		fmt.Println("")
	} else {
		fmt.Println(strings.Join(args, " "))
	}
}

func exit(args []string) {
	if strings.Join(args, "") != "0" {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func evaluateInput(input string) {
	inputSplitted := strings.Split(input, " ")

	command, args := inputSplitted[0], inputSplitted[1:]

	functionFounded, ok := availableCommandsMap[command]

	if ok {
		functionFounded(args)
	} else {
		fmt.Printf("%s: command not found\n", strings.TrimRight(input, "\n"))
	}
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimRight(input, "\n")

		if input != "" {
			evaluateInput(input)
		}

	}
}
