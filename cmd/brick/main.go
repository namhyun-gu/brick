package main

import (
	"github.com/namhyun-gu/brick/pkg/cmd/factory"
	"github.com/namhyun-gu/brick/pkg/cmd/root"
	"log"
)

func main() {
	cmdFactory := factory.New()

	rootCmd := root.NewCmdRoot(cmdFactory)
	if _, err := rootCmd.ExecuteC(); err != nil {
		log.Fatal(err)
	}
}
