package greenshark

import (
	"testing"
	"time"
)

type ArticleTest struct {
	article Article
	bow     map[string]int //Expected result of Bag-of-word representation
}

var base_time = time.Date(2019, time.January, 10, 23, 0, 0, 0, time.UTC)
var samples = []ArticleTest{
	{
		article: Article{
			Date:    base_time,
			Title:   "Title",
			Content: "This article is really interresting.",
			Id:      "1a5b084f-ff4e-477f-924d-c75d60d4ac8b",
		},
		bow: map[string]int{
			"this":         1,
			"article":      1,
			"is":           1,
			"really":       1,
			"interresting": 1},
	},
	{
		article: Article{
			Id:      "aee74c41-b2f2-4452-924f-02b8a98a40cf",
			Date:    base_time.Add(19 * time.Minute),
			Title:   "Youhou",
			Content: "This article is not really interresting.",
		},
		bow: map[string]int{
			"this":         1,
			"not":          1,
			"article":      1,
			"is":           1,
			"really":       1,
			"interresting": 1},
	},
	{
		article: Article{
			Id:      "7046144e-9357-4d2d-8a75-026399d0214a",
			Date:    base_time.Add(19 * time.Minute),
			Title:   "Youhou",
			Content: "This word is duplicated duplicated, even Duplicated.",
		},
		bow: map[string]int{
			"this":       1,
			"word":       1,
			"duplicated": 3,
			"is":         1,
			"even":       1},
	},
}

func TestBow(t *testing.T) {
	for _, smpl := range samples {
		got := smpl.article.ToWordVector()

		for k, _ := range smpl.bow {
			if _, ok := got[k]; !ok {
				t.Errorf("word %q expected.", k)
			}
		}
		for k, v := range got {
			expected_value, ok := smpl.bow[k]
			if !ok {
				t.Errorf("word %q not in  expected words", k)
			}
			if expected_value != v {
				t.Errorf("key %q: expected %d, got %d", k, expected_value, v)
			}
		}
	}
}
