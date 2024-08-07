package downloader

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/Izzxt/hat/client"
)

type Downloader struct {
	client        client.Client
	pathUrl       string
	output        string
	domain        string
	fileName      string
	production    string
	fileExtension string
	revision      string
	isOther       bool
	isGordon      bool
	isFurni       bool
	isImages      bool
	isXml         bool
	isTxt         bool
	isSwf         bool
	url           string
}

type Furni struct {
	RoomItemTypes FurniTypeS `json:"roomitemtypes"`
	WallItemTypes FurniTypeS `json:"wallitemtypes"`
}

type FurniTypeS struct {
	FurniType []Data `json:"furnitype"`
}

type Data struct {
	ClassName string `json:"classname"`
	Revision  int64  `json:"revision"`
}

func NewDownloader(c client.Client) *Downloader {
	return &Downloader{
		client:   c,
		isGordon: false,
		isFurni:  false,
		isImages: false,
		isXml:    false,
		isTxt:    false,
		isSwf:    false,
		isOther:  false,
	}
}

func (g *Downloader) Fetch() ([]byte, int) {
	var url string

	if g.isGordon {
		url = fmt.Sprintf("https://images.habbo.%s/gordon/%s%s%s", g.domain, g.production, g.pathUrl, g.fileName)
	} else if g.isFurni {
		url = fmt.Sprintf("https://images.habbo.%s/dcr/hof_furni/%s/%s", g.domain, g.revision, g.fileName)
	} else if g.isImages {
		url = fmt.Sprintf("https://images.habbo.%s/c_images%s/%s", g.domain, g.pathUrl, g.fileName)
	} else if g.isOther {
		url = fmt.Sprintf("https://images.habbo.%s%s/%s", g.domain, g.pathUrl, g.fileName)
	} else {
		url = fmt.Sprintf("https://www.habbo.%s/gamedata%s", g.domain, g.pathUrl)
	}

	g.url = url

	rg := regexp.MustCompile("(_json)")

	if g.isXml {
		url = rg.ReplaceAllString(url, "_xml")
	}

	if g.isTxt {
		url = rg.ReplaceAllString(url, "")
	}

	resp := g.client.Get(url)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bodyBytes, resp.StatusCode
}

func (g *Downloader) Download() int {
	var linkUrl string
	var fileName string

	if g.isGordon {
		linkUrl = fmt.Sprintf("https://images.habbo.%s/gordon/%s%s/%s", g.domain, g.production, g.pathUrl, g.fileName)
	} else if g.isFurni {
		linkUrl = fmt.Sprintf("https://images.habbo.%s/dcr/hof_furni/%s/%s", g.domain, g.revision, g.fileName)
	} else if g.isImages {
		linkUrl = fmt.Sprintf("https://images.habbo.%s/c_images%s/%s", g.domain, g.pathUrl, g.fileName)
		fmt.Println(linkUrl)
	} else if g.isOther {
		linkUrl = fmt.Sprintf("https://images.habbo.%s%s/%s", g.domain, g.pathUrl, g.fileName)
	} else {
		linkUrl = fmt.Sprintf("https://www.habbo.%s/gamedata%s", g.domain, g.pathUrl)
	}

	g.url = linkUrl

	rg := regexp.MustCompile("(_json)")

	if g.isXml {
		linkUrl = rg.ReplaceAllString(linkUrl, "_xml")
	}

	if g.isTxt {
		linkUrl = rg.ReplaceAllString(linkUrl, "")
	}

	fileURL, err := url.Parse(linkUrl)

	if err != nil {
		log.Fatal(err)
	}

	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]

	if len(g.fileName) > 0 {
		fileName = g.fileName
	}

	if _, err := os.Stat(g.GetOutput()); os.IsNotExist(err) {
		err := os.MkdirAll(g.GetOutput(), 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", g.output, fileName))
	if err != nil {
		log.Fatal(err)
	}

	resp := g.client.Get(linkUrl)

	defer resp.Body.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return resp.StatusCode
}

func (g *Downloader) GetCurrentProduction() string {
	url := fmt.Sprintf("https://www.habbo.%s/gamedata/external_variables/0", g.domain)

	resp := g.client.Get(url)

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	r, _ := regexp.Compile(`(\\w+)*(flash-assets-PRODUCTION-[\d\w]+-[\d\w]+)`)

	m := r.FindString(string(bodyBytes))

	return m
}

func (g *Downloader) GetOutput() string {
	return g.output
}

func (g *Downloader) GetPath() string {
	return g.pathUrl
}

func (g *Downloader) GetUrl() string {
	return g.url
}

func (g *Downloader) GetProduction() string {
	return g.production
}

func (g *Downloader) GetDomain() string {
	return g.domain
}

func (g *Downloader) SetRevision(revision string) {
	g.revision = revision
}

func (g *Downloader) SetPath(pathUrl string) {
	g.pathUrl = pathUrl
}

func (g *Downloader) SetOutput(output string) {
	g.output = output
}

func (g *Downloader) SetXml() {
	g.isXml = true
}

func (g *Downloader) SetOther() {
	g.isOther = true
}

func (g *Downloader) SetTxt() {
	g.isTxt = true
}

func (g *Downloader) SetGordon() {
	g.isGordon = true
}

func (g *Downloader) SetFurni() {
	g.isFurni = true
}

func (g *Downloader) SetImages() {
	g.isImages = true
}

func (g *Downloader) SetSwf() {
	g.isSwf = true
}

func (g *Downloader) SetProduction(production string) {
	g.production = production
}

func (g *Downloader) SetDomain(domain string) {
	g.domain = domain
}

func (g *Downloader) SetFileName(name string) {
	g.fileName = name
}
