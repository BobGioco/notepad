package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	CREATED        = "[OK] The note was successfully created"
	CLEARED        = "[OK] All notes were successfully deleted"
	UNKNOWN        = "[Error] Unknown command"
	MISSING_ERROR  = "[Error] Missing note argument"
	EMPTY          = "[Info] Notepad is empty"
	FULL_ERROR     = "[Error] Notepad is full"
	EXIT           = "[Info] Bye!"
	POSITION_ERROR = "[Error] Missing position argument"
	INVALID_ERROR  = "[Error] Invalid position:"
	UPDATE_ERROR = "[Error] There is nothing to update"
	DELETE_ERROR = "[Error] There is nothing to delete"
)

var (
	maxSize int
)

func updateNote(position int8, note string, list *strings.Builder) {
	lines := strings.Split(list.String(), "\n")
	if list.Len() == 0 {
		fmt.Println(UPDATE_ERROR)
	} else if position > 0 && position <= int8(len(lines)) {
		lines[position-1] = note
		fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
	} else {
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", position, len(lines))
	}

	list.Reset()
	list.WriteString(strings.Join(lines, "\n"))

}

func deleteNote(position int8, list *strings.Builder) {
	lines := strings.Split(list.String(), "\n")
	if list.Len() == 0 {
		fmt.Println(DELETE_ERROR)
	} else if position > 0 && position <= int8(len(lines)) {
		lines = append(lines[:position-1], lines[position:]...)
		fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
	} else {
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", position, len(lines))
	}

	list.Reset()
	list.WriteString(strings.Join(lines, "\n"))

}

func checkNumber(num string) (int8, error) {
	value, err := strconv.ParseInt(num, 10, 8)
	return int8(value), err
}

func main() {
	fmt.Print("Enter the maximum number of notes: ")
	fmt.Scan(&maxSize)

	var builder strings.Builder
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a command and data: ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)

		switch parts[0] {
		case "create":
			if builder.Len() == maxSize {
				fmt.Println(FULL_ERROR)
			} else {
				if len(parts) > 1 {
					builder.WriteString(parts[1])
					builder.WriteString("\n")
					fmt.Println(CREATED)
				} else {
					fmt.Println(MISSING_ERROR)
				}
			}
		case "update":
			if len(parts) > 1 {
				updateParts := strings.SplitN(parts[1], " ", 2)

				num, err := checkNumber(updateParts[0])
				if err != nil {
					fmt.Printf("%s %s\n", INVALID_ERROR, updateParts[0])
				} else if len(updateParts) != 2 {
					fmt.Println(MISSING_ERROR)
				} else {
					updateNote(num, updateParts[1], &builder)
				}
			} else {
				fmt.Println(POSITION_ERROR)
			}
		case "delete":
			if len(parts) > 1 {
				num, err := checkNumber(parts[1])
				if err != nil {
					fmt.Printf("%s %s\n", INVALID_ERROR, parts[1])
				} else {
					deleteNote(num, &builder)
				}
			} else {
				fmt.Println(POSITION_ERROR)
			}
		case "list":
			if builder.Len() == 0 {
				fmt.Println(EMPTY)
			}
			lines := strings.Split(builder.String(), "\n")
			for i, line := range lines {
				if line != "" {
					fmt.Printf("[Info] %d: %s\n", i+1, line)
				}
			}
		case "clear":
			builder.Reset()
			fmt.Println(CLEARED)
		case "exit":
			fmt.Println(EXIT)
		default:
			fmt.Println(UNKNOWN)
		}
	}
}