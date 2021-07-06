package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/weining-da/goalgorithms/collections"
	"github.com/weining-da/goalgorithms/resources"
)

func main() {
	if err := GenReport(collections.CreateString()); err != nil {
		panic(err)
	}
	if b, ok := resources.ReadTestFile("testdata/base/assets/model.mdl"); ok {
		fmt.Printf("%s\n", b)
	} else {
		panic(fmt.Errorf("can not read embeded test file"))
	}
	if b, ok := resources.ReadConfFile("config/conf.json"); ok {
		fmt.Printf("%s\n", b)
	} else {
		panic(fmt.Errorf("can not read embeded resource file"))
	}
}

func GenReport(msg string) error {
	_, err := fmt.Println(msg)
	return errors.Wrap(err, "Fail to generate report")
}
