package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/weining-da/goalgorithms/collections"
)

func main() {
	if err := GenReport(collections.CreateString()); err != nil {
		panic(err)
	}
}

func GenReport(msg string) error {
	_, err := fmt.Println(msg)
	return errors.Wrap(err, "Fail to generate report")
}
