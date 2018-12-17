package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/arjanvaneersel/lava/pkg/repl"
)

func main() {
	var usr string
	u, err := user.Current()
	if err != nil {
		usr = "anonymous user"
	} else {
		usr = u.Name
	}

	fmt.Printf("Hello %s, welcome to the Lava REPL\n\n", usr)
	repl.Start(os.Stdin, os.Stdout)
}
