package version

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/Izzxt/hat/client"
	"github.com/fatih/color"
	"github.com/hashicorp/go-version"
)

var Version string = "0.1.1"

type GithubReleases struct {
	TagName string `json:"tag_name"`
}

func CheckForUpdate(c client.Client) {
	current, err := version.NewVersion(Version)
	if err != nil {
		log.Fatal(err)
	}
	res := c.Get(fmt.Sprintf("https://api.github.com/repos/Izzxt/hat/releases/latest"))

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ghRelease GithubReleases

	if err := json.Unmarshal(bytes, &ghRelease); err != nil {
		log.Fatal(err)
	}

	latest, err := version.NewVersion(ghRelease.TagName)
	if err != nil {
		log.Fatal(err)
	}
	if current.LessThan(latest) {
		fmt.Println(fmt.Sprintf("%s New version avalaible, please considered to update to a new version. %s %s", color.YellowString("!!"), ghRelease.TagName, color.YellowString("!!")))
	}
}

func StartupMessage(c client.Client) {
	// cricket fonts
	fmt.Println("      ___ ___       _______       _______       _____        _____  ")
	fmt.Println("     |   Y   |     |   _   |     |       |     | _   |      | _   | ")
	fmt.Println("     |.  1   |     |.  1   |     |.|   | |     |.|   |  __  |.|   | ")
	fmt.Println("     |.  _   |     |.  _   |     `-|.  |-'     `-|.  | |__| `-|.  | ")
	fmt.Println("     |:  |   |     |:  |   |       |:  |         |:  |        |:  | ")
	fmt.Println("     |::.|:. |     |::.|:. |       |::.|         |::.|        |::.| ")
	fmt.Println("     `--- ---'     `--- ---'       `---'         `---'        `---' ")
	fmt.Println()
	fmt.Println("      -=- Discord @ Izzat#0333 -=-            -=- Version @ v1.1 -=-")
	fmt.Println()
	CheckForUpdate(c)
	fmt.Println()
	fmt.Println("-=- Initializing....")
	fmt.Println()
}
