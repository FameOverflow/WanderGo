package Comment

import (
	con "SparkForge/Config"
)

type HotComments []con.Comment
type NewComments []con.Comment

var hotComments HotComments
var newComments NewComments

func (a HotComments) Len() int {
	return len(a)
}
func (a HotComments) Less(i, j int) bool {
	return a[i].StarCnt > a[j].StarCnt
}
func (a HotComments) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a NewComments) Len() int {
	return len(a)
}
func (a NewComments) Less(i, j int) bool {
	return a[i].Data > a[j].Data
}
func (a NewComments) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}