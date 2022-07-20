package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func cmdCd(cmd []string) {
	if len(cmd) != 2 {
		fmt.Fprint(os.Stderr, "please insert a path\n")
	}
	err := os.Chdir(cmd[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
func cmdPwd(cmd []string) {
	if len(cmd) > 1 {
		fmt.Fprint(os.Stderr, "too many arguments\n")
	} else {
		path, err := os.Getwd()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Print(path)
	}
}
func cmdEcho(cmd []string) {
	if len(cmd) == 1 {
		fmt.Println()
	} else {
		for i := 1; i < len(cmd); i++ {
			cmd[i] = strings.Trim(cmd[i], "$")
			res := os.Getenv(cmd[i])
			fmt.Print(res, " ")
		}
		fmt.Println()
	}
}
func cmdKill(cmd []string) {
	pid, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	err = proc.Kill()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
func cmdPs(cmd []string) {
	sliceProc, _ := ps.Processes()

	for _, proc := range sliceProc {

		fmt.Printf("Process name: %v process id: %v\n", proc.Executable(), proc.Pid())

	}

}
func anotherCmd(cmd []string) {
	var comm *exec.Cmd
	if len(cmd) > 1 {
		comm = exec.Command(cmd[0], cmd[1:]...)
	} else {
		comm = exec.Command(cmd[0])
	}
	output, err := comm.Output()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Println(string(output))
}
func GoToExec(cmd []string) {
	for _, c := range cmd {
		partCmd := strings.Split(c, " ")
		switch partCmd[0] {
		case "cd":
			cmdCd(partCmd)
		case "pwd":
			cmdPwd(partCmd)
		case "echo":
			cmdEcho(partCmd)
		case "kill":
			cmdKill(partCmd)
		case "ps":
			cmdPs(partCmd)
		case "q":
			fmt.Fprint(os.Stdout, "quit\n")
			os.Exit(0)
		default:
			anotherCmd(partCmd)
		}
	}
}

//Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
//
//- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
//- pwd - показать путь до текущего каталога
//- echo <args> - вывод аргумента в STDOUT
//- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
//- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
//
//
//Так же требуется поддерживать функционал fork/exec-команд
//
//Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
//
//*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
//в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
//и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for {
		dir, err := os.Getwd()
		if err != nil {
			return
		}
		fmt.Print(dir, ":: ")
		if stdin.Scan() {
			cmd := stdin.Text()
			cmdSlice := strings.Split(cmd, "|")
			GoToExec(cmdSlice)
		}
	}
}
