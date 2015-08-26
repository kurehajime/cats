// cats_test.go
package main

import (
	"testing"
)

func TestCats(t *testing.T) {
	str1 := "東京特許許可局局長は隣の客とよく柿食う客だ"
	str2 := "庭庭庭庭にとがいるような気がしないでもない気がするような気がしたけど気のせいかもしれない。"
	COL = 80
	result := cats([]string{str1, str2})
	ans := `東京特許許可局局長は隣の客とよく柿食う |庭庭庭庭にとがいるような気がしないでも 
客だ                                   |ない気がするような気がしたけど気のせい 
                                       |かもしれない。                         
`
	if result != ans {
		t.Errorf("cat(x,y) =\n%x\n, want \n%x", result, ans)
		return
	}
}
