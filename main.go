package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pshvedko/textgame/student"
)

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
	initGame()
	for {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		command, err := r.ReadString('\n')
		if err != nil {
			break
		}
		answer := handleCommand(command)
		fmt.Println(answer)
	}
}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все локации
		если что-то было - оно корректно перезатирается
	*/
	g = student.New()
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	return g.HandleCommand(command)
}

type Game interface {
	HandleCommand(string) string
}

var g Game
