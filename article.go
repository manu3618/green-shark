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
	id      string
	content string
	date    time.Time
	title   string
}

func (art Article) ToWordVector() map[string]int {
	res := make(map[string]int)
	var ok bool
	var count int
	var ponct = []string{".", ";", ","}
	text := strings.ToLower(art.content)
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
	repr_title := art.title
	repr_content := art.content

	if len(art.title) > 80 {
		repr_title = art.title[:80]
	}
	if len(art.content) > 160 {
		repr_content = art.content[:160]
	}

	return fmt.Sprintf(
		"%v\t%v (%v)\n%v\n",
		art.id,
		repr_title,
		art.date,
		repr_content)
}
