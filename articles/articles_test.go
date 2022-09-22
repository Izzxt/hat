package articles

import (
	"fmt"
	"regexp"
	"sync"
	"testing"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/stretchr/testify/assert"
)

var wg sync.WaitGroup
var mu sync.Mutex

func TestFetchAll(t *testing.T) {

	assert := assert.New(t)
	ch := make(chan Result)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	a := NewArticles(&wg, *d, &mu)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go a.FetchAll(fmt.Sprintf("all_%d.html", i), ch)
	}

	for i := 1; i <= 3; i++ {
		select {
		case msg := <-ch:
			assert.Regexp(regexp.MustCompile("section"), string(msg.Response))
		case code := <-ch:
			assert.Equal(200, code.Code)
		}
	}

	wg.Wait()
}

func TestGetMaxPage(t *testing.T) {

	assert := assert.New(t)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	a := NewArticles(&wg, *d, &mu)

	p := a.GetMaxPage()

	assert.Equal(127, p)

}

func TestGetAllImages(t *testing.T) {

	assert := assert.New(t)
	ch := make(chan Result)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	a := NewArticles(&wg, *d, &mu)
	p := a.GetMaxPage()

	for i := 1; i <= p; i++ {
		wg.Add(1)
		go a.FetchAll(fmt.Sprintf("all_%d.html", i), ch)
	}

	for i := 1; i <= p; i++ {
		select {
		case msg := <-ch:
			assert.Regexp(regexp.MustCompile("section"), string(msg.Response))
		case code := <-ch:
			assert.Equal(200, code.Code)
		}
	}

	wg.Wait()
}

func TestGetAllImagesWithRegex(t *testing.T) {

	var data []string

	assert := assert.New(t)
	ch := make(chan Result)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	a := NewArticles(&wg, *d, &mu)
	p := a.GetMaxPage()

	for i := 1; i <= p; i++ {
		wg.Add(1)
		go a.FetchAll(fmt.Sprintf("all_%d.html", i), ch)
	}

	for i := 1; i <= p; i++ {
		select {
		case msg := <-ch:
			assert.Regexp(regexp.MustCompile("section"), string(msg.Response))
			data = append(data, string(msg.Response))
		case code := <-ch:
			assert.Equal(200, code.Code)
		}
	}

	// fmt.Printf("%#v ", data)

	wg.Wait()
}

func TestRegexMatch(t *testing.T) {
	assert := assert.New(t)
	html := `<article class="news-header"> <a href="/community/article/30955/we-want-your-feedback-on-the-nft-project" class="news-header__link news-header__banner"> <figure class="news-header__viewport"> <img src="https://images.habbo.com/web_images/habbo-web-articles/lpromo_gen15_07_thumb.png" alt="We want your feedback on the NFT project!" class="news-header__image news-header__image--thumbnail"> </figure> </a> <a href="/community/article/30955/we-want-your-feedback-on-the-nft-project" class="news-header__link news-header__wrapper"> <h2 class="news-header__title">We want your feedback on the NFT project!</h2> </a> <aside class="news-header__wrapper news-header__info"> <time class="news-header__date">{{ 1662114512000 | date: 'mediumDate' }}</time> <ul class="news-header__categories"> <li class="news-header__category"> <a href="/community/category/campaigns-activities" class="news-header__category__link" translate="NEWS_CATEGORY_CAMPAIGNS_ACTIVITIES"></a> </li> <li class="news-header__category"> <a href="/community/category/technical-updates" class="news-header__category__link" translate="NEWS_CATEGORY_TECHNICAL_UPDATES"></a> </li> </ul> </aside> <p class="news-header__wrapper news-header__summary">Our NFT project turns 1 years old soon! We’ve set up a poll in the Welcome Lounge to collect your feedback about it.</p> </article> <article class="news-header"><a href="/community/article/30888/war-of-the-seasons" class="news-header__link news-header__banner"><figure class="news-header__viewport"><img src="https://images.habbo.com/web_images/habbo-web-articles/lpromo_WOTS0922_thumb.png" alt="War of the Seasons!" class="news-header__image news-header__image--thumbnail"></figure></a><a href="/community/article/30888/war-of-the-seasons" class="news-header__link news-header__wrapper"> <h2 class="news-header__title">War of the Seasons!</h2> </a> <aside class="news-header__wrapper news-header__info"> <time class="news-header__date">{{ 1662546543000 | date: 'mediumDate' }}</time> <ul class="news-header__categories"> <li class="news-header__category"> <a href="/community/category/campaigns-activities" class="news-header__category__link" translate="NEWS_CATEGORY_CAMPAIGNS_ACTIVITIES"></a> </li> </ul> </aside> <p class="news-header__wrapper news-header__summary">Keep reading to find out more</p> </article>`

	rg, _ := regexp.Compile(`([\w])*\.png`)

	m := rg.FindAllString(html, -1)

	expected := []string(
		[]string{
			"lpromo_gen15_07_thumb.png",
			"lpromo_WOTS0922_thumb.png",
		},
	)

	assert.Equal(expected, m)
}

func TestRegexReplace(t *testing.T) {
	var after []string
	assert := assert.New(t)
	html := `<article class="news-header"> <a href="/community/article/30955/we-want-your-feedback-on-the-nft-project" class="news-header__link news-header__banner"> <figure class="news-header__viewport"> <img src="https://images.habbo.com/web_images/habbo-web-articles/lpromo_gen15_07_thumb.png" alt="We want your feedback on the NFT project!" class="news-header__image news-header__image--thumbnail"> </figure> </a> <a href="/community/article/30955/we-want-your-feedback-on-the-nft-project" class="news-header__link news-header__wrapper"> <h2 class="news-header__title">We want your feedback on the NFT project!</h2> </a> <aside class="news-header__wrapper news-header__info"> <time class="news-header__date">{{ 1662114512000 | date: 'mediumDate' }}</time> <ul class="news-header__categories"> <li class="news-header__category"> <a href="/community/category/campaigns-activities" class="news-header__category__link" translate="NEWS_CATEGORY_CAMPAIGNS_ACTIVITIES"></a> </li> <li class="news-header__category"> <a href="/community/category/technical-updates" class="news-header__category__link" translate="NEWS_CATEGORY_TECHNICAL_UPDATES"></a> </li> </ul> </aside> <p class="news-header__wrapper news-header__summary">Our NFT project turns 1 years old soon! We’ve set up a poll in the Welcome Lounge to collect your feedback about it.</p> </article> <article class="news-header"><a href="/community/article/30888/war-of-the-seasons" class="news-header__link news-header__banner"><figure class="news-header__viewport"><img src="https://images.habbo.com/web_images/habbo-web-articles/lpromo_WOTS0922_thumb.png" alt="War of the Seasons!" class="news-header__image news-header__image--thumbnail"></figure></a><a href="/community/article/30888/war-of-the-seasons" class="news-header__link news-header__wrapper"> <h2 class="news-header__title">War of the Seasons!</h2> </a> <aside class="news-header__wrapper news-header__info"> <time class="news-header__date">{{ 1662546543000 | date: 'mediumDate' }}</time> <ul class="news-header__categories"> <li class="news-header__category"> <a href="/community/category/campaigns-activities" class="news-header__category__link" translate="NEWS_CATEGORY_CAMPAIGNS_ACTIVITIES"></a> </li> </ul> </aside> <p class="news-header__wrapper news-header__summary">Keep reading to find out more</p> </article>`

	rg, _ := regexp.Compile(`([\w])*\.png`)

	m := rg.FindAllString(html, -1)

	expected := []string(
		[]string{
			"lpromo_gen15_07.png",
			"lpromo_WOTS0922.png",
		},
	)

	rgm := regexp.MustCompile(`_thumb.png`)

	for _, v := range m {
		r := rgm.ReplaceAllString(v, ".png")
		after = append(after, r)
	}

	assert.Equal(expected, after)
}

func TestRegexGetAllAndReplace(t *testing.T) {
	var data []string
	var after []string

	assert := assert.New(t)
	ch := make(chan Result)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	a := NewArticles(&wg, *d, &mu)
	p := a.GetMaxPage()

	for i := 1; i <= p-2; i++ {
		wg.Add(1)
		go a.FetchAll(fmt.Sprintf("all_%d.html", i), ch)
	}

	for i := 1; i <= p-2; i++ {
		select {
		case msg := <-ch:
			assert.Regexp(regexp.MustCompile("section"), string(msg.Response))
			data = append(data, string(msg.Response))
		case code := <-ch:
			assert.Equal(200, code.Code)
		}
	}

	defer wg.Wait()

	rg, _ := regexp.Compile(`([\w!@#$%^&*+-])*\.png`)
	rgmt := regexp.MustCompile(`_thumb.png`)

	for _, d := range data {
		s := rg.FindAllString(d, -1)
		for _, tr := range s {
			r := rgmt.ReplaceAllString(tr, ".png")
			after = append(after, r)
		}
	}

	for _, v := range after {
		assert.Regexp(regexp.MustCompile(".png"), v)
	}
}

// func TestFetch(t *testing.T) {
// assert := assert.New(t)
// ch := make(chan Result)
// c := client.NewClient()
// d := downloader.NewDownloader(c)
// a := articles.NewArticles(&wg, *d, &mu)
//
// i := 1
//
// defer wg.Wait()
//
// for {
// 	wg.Add(1)
// 	go func(i int, ch chan Result) {
// 		a.FetchAll(fmt.Sprintf("all_%d.html", i), ch)
// 	}(i, ch)
//
// 	go func(i int, ch chan Result) {
// 		fmt.Println(i)
// 		select {
// 		case c := <-ch:
// 			if c.Code == 404 {
// 				break
// 			}
// 		}
// 	}(i, ch)
// 	i++
// }
//
// }
