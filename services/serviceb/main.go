package main

import (
	"fmt"
	"github.com/mohammadVatandoost/go-bazel-monorepo/services/serviceb/internal/utils"
)


func main() {
	fmt.Println("Hello world from serviceb")

	fmt.Printf("Sum 1 and 2 : %v \n", utils.Sum(1, 2))
}