/**
 * functions related to article
 */

package greenshark

import (
	"fmt"
	"strings"
	"time"
)

type Article struct {
	Id      string
	Content string
	Date    time.Time
	Title   string
}

func (art Article) ToWordVector() map[string]int {
	res := make(map[string]int)
	var ok bool
	var count int
	var ponct = []string{".", ";", ","}
	text := strings.ToLower(art.Content)
	for _, char := range ponct {
		text = strings.ReplaceAll(text, char, " ")
	}
	for _, w := range strings.Fields(text) {
		count, ok = res[w]
		if ok {
			res[w] = count + 1
		} else {
			res[w] = 1
		}
	}
	return res
}

func (art Article) String() string {
	repr_title := art.Title
	repr_content := art.Content

	if len(art.Title) > 80 {
		repr_title = art.Title[:80]
	}
	if len(art.Content) > 160 {
		repr_content = art.Content[:160]
	}

	return fmt.Sprintf(
		"%v\t%v (%v)\n%v\n",
		art.Id,
		repr_title,
		art.Date,
		repr_content)
}
