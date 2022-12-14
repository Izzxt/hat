package badges

import (
	"fmt"
	"sync"
	"testing"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/stretchr/testify/assert"
)

func TestDownloadBadge(t *testing.T) {
	code := []string{
		"thx",
		"AC7_HHCA",
		"AC7",
	}

	assert := assert.New(t)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	d.SetOther()
	d.SetOutput("out")
	d.SetDomain("com")
	d.SetPath("/c_images/album1584")

	for _, v := range code {
		d.SetFileName(fmt.Sprintf("/%s.gif", v))
		d.Download()
	}

	assert.FileExists(fmt.Sprintf("./out/%s.gif", code[0]))
	assert.FileExists(fmt.Sprintf("./out/%s.gif", code[1]))
	assert.FileExists(fmt.Sprintf("./out/%s.gif", code[2]))
}

func TestGetAllCode(t *testing.T) {
	var wg *sync.WaitGroup
	var mu *sync.Mutex

	c := client.NewClient()
	d := downloader.NewDownloader(c)
	b := NewBadges(*d, wg, mu)

	code := b.GetAllCode()

	fmt.Printf("%d Badges", len(code))
}

func TestMatchRegex(t *testing.T) {
	assert := assert.New(t)
	match := `
badge_desc_JKFES FKES=Hello World
badge_desc_JK|FES_FKES=Hello World
badge_desc_JKFES-FKES!0=Hello World
badge_desc_JKFES+FKES_0.=Hello World
badge_desc_JKFES(FKES__=Hello World
badge_name_JK-$FES_FKES=Hello World
badge_name_JKFES-FKES!0=Hello World
badge_name_$#JKFES+FKES_0.=Hello World
badge_name_@#JKFES(FKES__=Hello World
badge_name_()JKFES FKES_=Hello World
badge_desc_KITG1=${badge_desc_KIT01}
notifications.text.achievement.reward.2=Recompensa: %activitypoints% Coracoes
notifications.text.achievement.unlocked=Você ganhou a conquista
notifications.text.achievement=Você ganhou a conquista "%badge_name%"
notifications.text.activitypoints.0=%change% Pixels recebidos, seu total agora é %count%.
	`

	m := matchRegex(match)

	assert.Equal("JKFES FKES", m[0][2])
	assert.Equal("JK|FES_FKES", m[1][2])
	assert.Equal("JKFES-FKES!0", m[2][2])
	assert.Equal("JKFES+FKES_0.", m[3][2])
	assert.Equal("JKFES(FKES__", m[4][2])
	assert.Equal("JK-$FES_FKES", m[5][2])
	assert.Equal("JKFES-FKES!0", m[6][2])
	assert.Equal("$#JKFES+FKES_0.", m[7][2])
	assert.Equal("@#JKFES(FKES__", m[8][2])
	assert.Equal("()JKFES FKES_", m[9][2])
	assert.Equal("KITG1", m[10][2])
}
