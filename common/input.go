package common

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

func Input() []string {
	inputData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Fatalf("Input file not found")
	}

	return strings.Split(string(inputData), "\n")
}
