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
			date:    base_time,
			title:   "Title",
			content: "This article is really interresting.",
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
			date:    base_time.Add(19 * time.Minute),
			title:   "Youhou",
			content: "This article is not really interresting.",
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
			date:    base_time.Add(19 * time.Minute),
			title:   "Youhou",
			content: "This word is duplicated duplicated, even Duplicated.",
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
