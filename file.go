package main

import (
	"bufio"
	"io"
	"os"
)

func createnewFile(name string, message string) error {
	file, error := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if error != nil {
		return error
	}

	defer file.Close() // tutup agar tidak memori leak
	file.WriteString(message)

	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	// read
	reader := bufio.NewReader(file)
	var result string
	for {
		line,_, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		result += string(line) + "\n"
	}

	return result, nil
}

func addToFIle(name string, message string) error {
	file, error := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if error != nil {
		return error
	}

	defer file.Close() // tutup agar tidak memori leak
	file.WriteString(message)

	return nil
}

func main() {
	// createnewFile("test.log", "Hello World")

	// result, _ := readFile("test.log")
	// fmt.Println(result)

	addToFIle("test.log", "\nHello Rizki adding message\n")
}