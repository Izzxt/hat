package fs

import (
	"fmt"
	"log"
	"os"
	"regexp"
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
		fmt.Println(fmt.Sprintf("Skipped %s", r.ReplaceAllString(fileName, "")))
		return true
	} else {
		r := regexp.MustCompile(`\/`)
		fmt.Println(fmt.Sprintf("Download %s", r.ReplaceAllString(fileName, "")))
	}

	return false
}
