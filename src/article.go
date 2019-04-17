/**
 * functions related to article
 */

package article

import (
    "time"
    "fmt"
)

type Article struct {
    content string
    date time.Time
    title string
}

func (art Article) toWordVector() map[string]int{
	res :=make(map[string]int)
	var ok bool
	var count int
    var text:=art.content.Lower()
	for _, w := range strings.Fields(text) {
		count, ok = res[w]
		if ok {
			res[w] = count +1
		} else {
			res[w] = 1
		}
	}
	return res
}

func (art Article) String() string{
    return fmt.Sprintf("(%v) %v\n%v\n", date, title[:80], content[:160])
}
