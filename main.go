package main

import (
	"fmt"
	"os"

	"github.com/den-is/cex/cmd"
)

func main() {
	if err := cmd.CexCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
