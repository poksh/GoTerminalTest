package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
)

func main() {
	fmt.Printf("IsTerminal=%v\n", logrus.IsTerminal())
}
