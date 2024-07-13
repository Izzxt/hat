package furnitures

import (
	"fmt"
	"testing"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/stretchr/testify/assert"
)

func TestGetFurniIcons(t *testing.T) {
	assert := assert.New(t)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	d.SetDomain("com")

	f := NewFurnitures(*d, nil, nil)
	i := f.GetIcons()

	fmt.Print(len(i))
	fmt.Printf("%+v\n", i[0])
	assert.Equal("shelves_norja_icon.png", i[0].Name)
	assert.Equal("61856", i[0].Revision)
}

func TestGetFurnis(t *testing.T) {
	assert := assert.New(t)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	d.SetDomain("com")

	f := NewFurnitures(*d, nil, nil)
	i := f.GetFurnis()

	fmt.Print(len(i))
	fmt.Printf("%+v\n", i[0])
	assert.Equal("shelves_norja.swf", i[0].Name)
	assert.Equal("61856", i[0].Revision)
}
