/*
utils to tranform rss feeds to articles.
*/

package greenshark

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
    "time"
    "github.com/mattermost/html2text"
)

type RssArticles struct{
    XMLName xml.Name `feed`
    Articles []RssArticle `xml:"entry"`
}

type RssArticle struct {
	XMLName xml.Name `entry`
	Title   string   `xml:"title"`
	//Link string `xml:"link,attr"` //TODO
	DateCreated time.Time `xml:"published"`
	DateUpdated time.Time `xml:"updated"`
	Id          string `xml:"id"`
	Content     string `xml:"content"`
}

func HtmlToArticle(html []byte) []Article {
	var ret []Article
	var rss_articles RssArticles
	xml.Unmarshal(html, &rss_articles)
	for _, art := range rss_articles.Articles {
        var cur Article
        cur.Id = art.Id
        cur.Content, _ = html2text.FromString(art.Content) //TODO plain text
        cur.Title = art.Title

        // date
        cur.Date = art.DateCreated
        if art.DateUpdated.After(art.DateCreated) {
            cur.Date = art.DateUpdated
        }

		ret = append(ret, cur)
	}
	return ret
}

func getRss(url string) []Article {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	page, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return HtmlToArticle(page)
}
