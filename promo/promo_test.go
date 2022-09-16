package promo

import (
	"fmt"
	"testing"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/stretchr/testify/assert"
)

func TestMatchRegex(t *testing.T) {
	assert := assert.New(t)
	expected := `
21EGG02.image=web_promo_small/spromo_bunnyvillage.png
	21aprilfb1.image=web_promo_small/spromo_aprilmakeuproom.png
21aprilfb2.image=web_promo_small/spromo_fsgamemuseum.png
21aprilfb3.image=web_promo_small/spromo_backstageendgame.png
		21aug01.image=web_promo_small/spromo_vaporwavegame1.png
	`

	m := matchRegex(expected)

	assert.Equal("spromo_bunnyvillage.png", m[0][1])
	assert.Equal("spromo_aprilmakeuproom.png", m[1][1])
}

func TestGetAllImages(t *testing.T) {

	c := client.NewClient()
	d := downloader.NewDownloader(c)
	p := NewPromo(*d, nil, nil)

	i := p.GetAllImages()
	for _, v := range i {
		fmt.Println(v)
	}
}
