package utilities

import (
	"fmt"
	"ubersnap/static"
	"time"
)

func TransformDate(t time.Time, lang string) string {
	if lang == "en" {
		return t.Format("02 Jan 2006 15:04:05")
	}
	month := static.INDONESIAN_MONTH[int(t.Month())]
	h, i, s := t.Clock()
	return fmt.Sprintf("%d %s %d %d:%d:%d", t.Day(), month, t.Year(), h, i, s)
}
