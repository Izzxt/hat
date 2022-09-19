package files

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fatih/color"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsFileExists(output string, fileName string) bool {
	exts, err := Exists(fmt.Sprintf("%s/%s", output, fileName))
	if err != nil {
		log.Fatal(err)
	}

	if exts {
		r := regexp.MustCompile(`\/`)
		fmt.Println(fmt.Sprintf("%s %s", color.RedString("➜"), r.ReplaceAllString(fileName, "")))
		return true
	} else {
		r := regexp.MustCompile(`\/`)
		fmt.Println(fmt.Sprintf("%s %s", color.GreenString("✓"), r.ReplaceAllString(fileName, "")))
	}

	return false
}
