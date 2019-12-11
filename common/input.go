package common

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func Input() string {
	inputData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Fatalf("Input file not found")
	}

	return string(inputData)
}
