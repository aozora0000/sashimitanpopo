package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/mattn/go-tty"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var version = "unknown"
var revision = "unknown"

func main() {
	app := cli.NewApp()
	app.Usage = "Sashimi Tanpopo: You will be an automata"
	app.Version = version + "@" + revision
	app.EnableBashCompletion = true
	app.Action = func(context *cli.Context) error {
		if context.NArg() == 0 {
			return errors.New("required execute command")
		}
		command := context.Args().Get(0)
		body, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		t, err := tty.Open()
		if err != nil {
			return err
		}
		defer t.Close()

		buf := bytes.NewReader(body)
		scanner := bufio.NewScanner(buf)
		for scanner.Scan() {
			cmd := exec.Command(os.Getenv("SHELL"), []string{"-c", fmt.Sprintf(command, scanner.Text())}...)
			cmd.Stdout = os.Stdout
			fmt.Println(cmd.String())
			err = cmd.Run()
			if err != nil {
				return err
			}
			_, err = t.ReadRune()
			if err != nil {
				return err
			}
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
