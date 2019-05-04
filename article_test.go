package greenshark

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

type BowData map[string]map[string]int //Id:{ map[string] int}

func GetArticles(filename string) []Article {
	var articles []Article
	cont, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(cont, &articles)
	if err != nil {
		log.Fatal(err)
	}
	return articles
}

func GetBow(filename string) BowData {
	var data BowData
	cont, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(cont, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func TestBow(t *testing.T) {
	articles := GetArticles("testdata/articles.json")
	bows := GetBow("testdata/bow_data.json")
	for _, smpl := range articles {
		got := smpl.ToWordVector()
		expected := bows[smpl.Id]

		for k, _ := range expected {
			if _, ok := got[k]; !ok {
				t.Errorf("word %q expected.", k)
			}
		}
		for k, v := range got {
			expected_value, ok := expected[k]
			if !ok {
				t.Errorf("word %q not in  expected words", k)
			}
			if expected_value != v {
				t.Errorf("key %q: expected %d, got %d", k, expected_value, v)
			}
		}
	}
}

func TestHtmlArticle(t *testing.T) {
	input, err := ioutil.ReadFile("testdata/manu.hbrt.eu.xml")
	if err != nil {
		log.Fatal(err)
	}

    //TODO compare HtmlToArticle(input) and test data
	HtmlToArticle(input)
}
