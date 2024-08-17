package main

import (
	"bufio"
	"ci_go/calc"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Even:", calc.IsEven(num))
	fmt.Println("Odd:", calc.IsOdd(num))
}
