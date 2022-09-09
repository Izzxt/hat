package badges

import (
	"fmt"
	"testing"
)

func TestMatchRegex(t *testing.T) {
	// assert := assert.New(t)
	match := `
		badge_desc_JKFES FKES=Hello World
		badge_desc_JK|FES_FKES=Hello World
		badge_desc_JKFES-FKES!0=Hello World
		badge_desc_JKFES+FKES_0.=Hello World
		badge_desc_JKFES(FKES__=Hello World
		badge_desc_JKFES FKES_=Hello World
		badge_name_JKFES FKES=Hello World
		badge_name_JK-$FES_FKES=Hello World
		badge_name_JKFES-FKES!0=Hello World
		badge_name_$#JKFES+FKES_0.=Hello World
		badge_name_@#JKFES(FKES__=Hello World
		badge_name_ ()JKFES FKES_=Hello World
		badge_name_()JKFES FKES_=Hello World
	`

	m := matchRegex(match)

	for i, match := range m {
		fmt.Println(match[2], "found at index", i)
	}
}
