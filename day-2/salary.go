package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readStr(r *bufio.Reader, delim byte) string {
	input, err := r.ReadString(delim)
	if err != nil {
		return ""
	}
	return strings.Trim(input, string(delim))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buffTxt string

	fmt.Print("Salary: ")
	buffTxt = readStr(reader, '\n')
	salary, _ := strconv.ParseFloat(buffTxt, 64)

	fmt.Print("Months: ")
	buffTxt = readStr(reader, '\n')
	months, _ := strconv.Atoi(buffTxt)

	fmt.Println("Total amount is", float64(months)*salary)
}
