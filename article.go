/**
 * functions related to article
 */

package article

import (
	"fmt"
	"strings"
	"time"
)

type Article struct {
	content string
	date    time.Time
	title   string
}

func (art Article) toWordVector() map[string]int {
	res := make(map[string]int)
	var ok bool
	var count int
	text := strings.ToLower(art.content)
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
	return fmt.Sprintf(
		"(%v) %v\n%v\n",
		art.date,
		art.title[:80],
		art.content[:160])
}
