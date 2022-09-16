package hotelview

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
landing.view.background_hotel_top.uri=https://images.habbo.com/c_images/reception/pridefestiv21_background_hotel_top.png
landing.view.background_left.uri=https://images.habbo.com/c_images/reception/pridefestiv21_background_left.png
landing.view.background_right.uri=${image.library.url}reception/background_right_easter2016.png
	`

	m := matchRegex(expected)

	assert.Equal("background_right_easter2016.png", m[2][3])
	assert.Equal("pridefestiv21_background_hotel_top.png", m[0][3])
}

func TestGetAllImages(t *testing.T) {

	c := client.NewClient()
	d := downloader.NewDownloader(c)
	d.SetDomain("com")
	p := NewHotelView(*d, nil, nil)

	i := p.GetAllImages()
	for _, v := range i {
		fmt.Println(v)
	}
}
