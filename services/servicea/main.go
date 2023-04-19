package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)


func power2(num int) int {
	return num*num
}


func main() {
	fmt.Println("Hello world from servicea")
	logrus.Info("Hello world from servicea with dependencies")
}