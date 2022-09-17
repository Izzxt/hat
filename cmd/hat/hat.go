package main

import (
	"fmt"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/cmd"
	"github.com/Izzxt/hat/version"
)

func main() {
	ASCII()
	cmd.Execute()
}

func ASCII() {
	// cricket fonts
	fmt.Println("      ___ ___       _______       _______       _____        _______ ")
	fmt.Println("     |   Y   |     |   _   |     |       |     | _   |      |   _   |")
	fmt.Println("     |.  1   |     |.  1   |     |.|   | |     |.|   |  __  |.  |   |")
	fmt.Println("     |.  _   |     |.  _   |     `-|.  |-'     `-|.  | |__| |.  |   |")
	fmt.Println("     |:  |   |     |:  |   |       |:  |         |:  |      |:  1   |")
	fmt.Println("     |::.|:. |     |::.|:. |       |::.|         |::.|      |::.. . |")
	fmt.Println("     `--- ---'     `--- ---'       `---'         `---'      `-------'")
	fmt.Println()
	fmt.Println("      -=- Discord @ Izzat#0333 -=-            -=- Version @ v1.0 -=-")
	fmt.Println()
	checkForUpdate()
	fmt.Println()
	fmt.Println("-=- Initializing....")
	fmt.Println()
	fmt.Println()
}

func checkForUpdate() {
	c := client.NewClient()
	version.CheckForUpdate(c)
}
