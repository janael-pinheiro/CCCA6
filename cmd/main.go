package main

import (
	"fmt"

	refectoring "github.com/janael-pinheiro/CCCA6/pkg/refactoring"
)

func main() {

	fake_cpf := "111.444.777-05"

	result := refectoring.Validate(fake_cpf)

	fmt.Println(result)
}
