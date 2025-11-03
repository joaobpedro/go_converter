package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

// check the return errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// check if a string is numeric
func isFloat(a string) bool {
	_, err := strconv.ParseFloat(a, 64)
	return err == nil
}

func arc_length_translate(a []string) []string {

	var b []string

	for _, word := range a {
		val, err := strconv.ParseFloat(word, 64)
		check(err)
		val = val + 19.4
		word_added := strconv.FormatFloat(val, 'f', -1, 64)
		b = append(b, word_added)
	}
	return b
}

func main() {
	args := os.Args
	fmt.Print("Loading the file: ", args[1], "\n")
	// open file in go, this is just a dump of the info into memory
	f, err := os.Open(args[1])
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var arclength []string
	var speed []string
	var dbuffer []string

	for scanner.Scan() {
		line := scanner.Text()

		if s.Contains(line, "Arclength") {
			arclength = s.Split(line, ",")[1:]
		}

		if s.Contains(line, "Correlation") {
			speed = s.Split(line, ",")[1:]
		}

		// transform the arclength value
		// add 19.4m to the input
		arc_length := arc_length_translate(arclength)

		line_list := s.Split(line, ",")
		if isFloat(line_list[0]) {
			for i := 0; i < len(arclength); i++ {
				row := line_list[0] + " " + arc_length[i] + " " + line_list[i+1] + " 0 0 0 " + speed[i] + "\n"
				dbuffer = append(dbuffer, row)
			}
		}
	}

	// output the csv file for a text file
	output_file := s.Replace(args[1], "csv", "txt", 1)
	file, err := os.Create(output_file)
	check(err)
	defer file.Close()

	writer := bufio.NewWriter(file)

	// write data line by line
	for _, line := range dbuffer {
		// fmt.Print(line)
		_, err := writer.WriteString(line)
		check(err)
	}
	writer.Flush()
	fmt.Print("File printed with success!\n")
	// fmt.Print("Ploting the data into a figure\n")
}
