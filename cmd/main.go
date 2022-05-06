package main

import (
	"fmt"

	refectoring "github.com/janael-pinheiro/CCCA6/refactoring"
)

func main() {

	fake_cpf := "111.444.777-05"

	result := refectoring.Validate(fake_cpf)

	fmt.Println(result)
}
