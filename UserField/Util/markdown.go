package Util

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"

	"github.com/russross/blackfriday"
)

// 迭代
func MarkToHtml() {

	input := []byte(`
**你的慢漫报告**  
)`)
	unsafe := blackfriday.MarkdownCommon(input)
	output := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Println(string(output))
}
