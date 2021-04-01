package main

import (
	"fmt"
	"github.com/namhyun-gu/brick/pkg/cmd/factory"
	"github.com/namhyun-gu/brick/pkg/cmd/root"
	"log"
)

func main() {
	cmdFactory, err := factory.New()
	if err != nil {
		panic(fmt.Errorf("can't initialize factory (cause: %s)", err.Error()))
	}

	rootCmd := root.NewCmdRoot(cmdFactory)
	if _, err := rootCmd.ExecuteC(); err != nil {
		log.Fatal(err)
	}
}
