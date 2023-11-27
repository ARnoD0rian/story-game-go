package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	name_quest, steps := read_configuration("configure.txt")
	fmt.Println(steps)
	var name string
	fmt.Printf("привет путешевственник, добро пожаловать в приклучение %s, скажи мне, как тебя зовут? \n", name_quest)
	fmt.Scan(&name)

	var command string
	fmt.Println("Напиши start, если готов начать наше приключение, если хочешь выйти из игры, напиши end")

	is_continue := true
	for is_continue {
		fmt.Scan(&command)
		switch command {
		case "start":
			start(name, steps)
			is_continue = false
		case "end":
			is_continue = false
		default:
			fmt.Println("неизвестная команда")
		}

	}

	fmt.Printf("До встречи, %s", name)

}

func start(name string, steps [][]string) {

	if len(steps) == 0 {
		return
	}
	var next_step string
	last_step := 0
	next_step = game_text("1")
	for next_step != "end" {
		index, err := strconv.Atoi(next_step)
		if err != nil {
			fmt.Println("ошибка игры:", err)
			return
		}
		fmt.Println(last_step, index-1)
		next_step = game_text(steps[last_step][index-1])
		last_step, err = strconv.Atoi(steps[last_step][index-1])
		last_step -= 1

	}

}

func game_text(directory string) string {
	file, err := os.Open("location/" + directory + ".txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}
	fmt.Print("\n")

	var command string
	fmt.Scan(&command)

	return command
}

func read_configuration(directory string) (string, [][]string) {
	file, err := os.Open(directory)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var name_quest string
	var N int
	if scanner.Scan() {
		name_quest = scanner.Text()
	}
	if scanner.Scan() {
		N, err = strconv.Atoi(scanner.Text())
	}
	steps := make([][]string, N)

	for scanner.Scan() {
		step := strings.Split(scanner.Text(), "-")
		index, err := strconv.Atoi(step[0])
		step = strings.Split(step[1], ",")
		if err != nil {
			fmt.Println("Ошибка преобразования индекса в число:", err)
			continue
		}
		steps[index-1] = append(steps[index-1], step...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return name_quest, steps
}
