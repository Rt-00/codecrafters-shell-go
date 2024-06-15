package main

import (
	"bufio"
	"fmt"
	"os"
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
	for _, arg := range args {
		if slices.Contains(availableCommands, arg) {
			fmt.Printf("%s is a shell builtin\n", arg)
			continue
		}
		fmt.Printf("%s: not found\n", arg)
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

		evaluateInput(input)

	}
}
