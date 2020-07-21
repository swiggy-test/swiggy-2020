package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Emp struct {
	Id      int
	Name    string
	Salary  float64
	Address Address
}

type Address struct {
	Hno    string
	Street string
	City   string
	Pin    int
}

type Animal interface {
	Eat() string
	Sleep() string
	Breathe() string
}

func (e Emp) Eat() string {
	return "Employee eats"
}

func (e Emp) Sleep() string {
	return "Employee sleeps"
}

func (e Emp) Breathe() string {
	return "Employee breathes"
}

type DataCollector struct {
	Employees []Emp
	Reader    *bufio.Reader
	Trials    int
	counter   int
}

func ExitOnError(err error, txt string) {
	if err != nil {
		fmt.Println(txt)
		os.Exit(1)
	}
}

func (dc *DataCollector) ReadInput(txt string) string {
	fmt.Print(txt)
	buff, err := dc.Reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		os.Exit(1)
	}
	buff = strings.Trim(buff, "\n")
	return buff
}

func (dc *DataCollector) GetAddr() Address {
	fmt.Println("Please enter the address details of the employee")
	hno := dc.ReadInput("House Number:")
	street := dc.ReadInput("Street:")
	city := dc.ReadInput("City:")
	pin, err := strconv.Atoi(dc.ReadInput("Pincode:"))
	ExitOnError(err, "Invalid pincode")
	return Address{Hno: hno, Street: street, City: city, Pin: pin}
}

func (dc *DataCollector) GetEmp() Emp {
	var id, err = strconv.Atoi(dc.ReadInput("Employee ID:"))
	ExitOnError(err, "Invalid employee ID")
	name := dc.ReadInput("Employee Name:")
	var salary float64
	salary, err = strconv.ParseFloat(dc.ReadInput("Employee Salary:"), 64)
	ExitOnError(err, "Invalid salary")
	address := dc.GetAddr()
	return Emp{Id: id, Name: name, Salary: salary, Address: address}
}

func (dc *DataCollector) Init() {
	fmt.Println("Welcome!")
	fmt.Println("Enter the details of the employee.")
	dc.Employees[0] = dc.GetEmp()
	dc.counter = 1
}

func (dc *DataCollector) PrintEmps() {
	for _, emp := range dc.Employees {
		fmt.Println("###### Employee Details ######")
		fmt.Printf("%+v\n\n", emp)
	}
}

func (dc *DataCollector) NextAction() {
	opt := dc.ReadInput("Choose your next action: [C] Continue	[Q] Quit (default C):")
	if opt == "Q" {
		dc.PrintEmps()
		os.Exit(0)
	} else {
		if dc.counter < dc.Trials {
			dc.Employees[dc.counter] = dc.GetEmp()
			dc.counter += 1
		} else {
			dc.Employees = append(dc.Employees, dc.GetEmp())
		}
	}
}

func main() {
	trials := 5
	dataCollector := DataCollector{Employees: make([]Emp, trials), Trials: trials, Reader: bufio.NewReader(os.Stdin)}
	dataCollector.Init()
	for {
		dataCollector.NextAction()
	}
}
