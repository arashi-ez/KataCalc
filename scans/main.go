package main

import (
	"fmt"
	rom "github.com/brandenc40/romannumeral"
	"time"
)

func main() {
	var num int
	var timer = 5 * time.Second
	fmt.Println(
		`
----------------------------------------
Format: [number 1] [operator] [number 2]
----------------------------------------

Choose Numeral:
1 - Arab  Numeric.
2 - Roman Numeric.`)

	fmt.Scanln(&num)
	switch num {
	case 1:
		ArabCalc()
	case 2:
		RomanCalc()
	}
	time.Sleep(timer)
}

func RomanCalc() {
	fmt.Println(`-Roman Numeric-`)
	var (
		a string
		b string
		o string
		c int
	)

	fmt.Scanln(&a, &o, &b)
	integer, err := rom.StringToInt(a)
	integer2, err2 := rom.StringToInt(b)
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err)
	}
	if o == "+" {
		c = integer + integer2
	} else if o == "-" {
		c = integer - integer2
	} else if o == "*" {
		c = integer * integer2
	} else if o == "/" {
		c = integer / integer2
	} else if o == "%" {
		c = integer % integer2
	} else {
		fmt.Println(err)
	}
	roman, err := rom.IntToString(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(roman)

}

func ArabCalc() {
	fmt.Println(`-Arabic  Numeric-`)
	var (
		a    int
		b    int
		oper string
		//input string
		err = "Unknown operator or syntax"
	)

	fmt.Scanln(&a, &oper, &b)
	if oper == "+" {
		fmt.Printf("%d\n", a+b)
	} else if oper == "-" {
		fmt.Printf("%d\n", a-b)
	} else if oper == "*" {
		fmt.Printf("%d\n", a*b)
	} else if oper == "/" {
		fmt.Printf("%d\n", a/b)
	} else if oper == "%" {
		fmt.Printf("%d\n", a%b)
	} else {
		fmt.Println(err)
	}
}
