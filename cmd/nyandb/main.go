package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline"
	"github.com/duynhanf/nyandb"
)

func main() {

	l, err := readline.NewEx(&readline.Config{
		Prompt:          "# ",
		HistoryFile:     "/tmp/gosql.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})

	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Println("Welcome to nyan sql.")
repl:
	for {
		fmt.Print("# ")
		line, err := l.Readline()

		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue repl
			}
		} else if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading line:", err)
			continue repl
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "quit" || trimmed == "exit" || trimmed == "\\q" {
			fmt.Println("Bye!")
			break
		}

		parser := nyandb.NewParser()

		parser.Parse(line)
	}
}
