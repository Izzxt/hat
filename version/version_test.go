package version

import (
	"testing"

	"github.com/Izzxt/hat/client"
)

func TestCheckForUpdate(t *testing.T) {
	c := client.NewClient()
	CheckForUpdate(c)
}
