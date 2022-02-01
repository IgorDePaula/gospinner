package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type actionDone struct{}
type actionStep struct{}

func runAction() <-chan tea.Msg {
	ch := make(chan tea.Msg, 0)

	go func() {
		time.Sleep(time.Second * 1)
		ch <- actionStep{}
		time.Sleep(time.Second * 2)
		ch <- actionStep{}
		time.Sleep(time.Second * 3)
		ch <- actionStep{}

		time.Sleep(time.Second * 5)
		ch <- actionDone{}
		close(ch)
	}()

	return ch
}
