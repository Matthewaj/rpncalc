package main

import (
	"bufio"
	"fmt"
	"github.com/matthewaj/rpncalc/rpn"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	expr, err := readExpression()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to read expression"))
	}

	value, err := rpn.Parse(expr)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "cannot parse the given expression"))
	}

	fmt.Printf("= %.2f\n", value)
}

func readExpression() ([]byte, error) {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadBytes(byte('\n'))
	if err != nil {
		return nil, errors.Wrap(err, "read expression")
	}

	return data, nil
}
