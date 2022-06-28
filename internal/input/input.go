package input

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func ReadPassword() (string, error) {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	defer signal.Stop(ch)

	state, err := terminal.GetState(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	go func() {
		<-ch
		terminal.Restore(int(syscall.Stdin), state)
		os.Exit(1)
	}()

	fmt.Print("type the password : ")
	data, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", err
	}
	fmt.Println("")

	return string(data), nil
}
