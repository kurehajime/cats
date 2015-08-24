// cats.go
package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"math"
	"strings"
	"github.com/mattn/go-runewidth" 
)
const (
	UTF8      = 65001
	SHIFT_JIS = 932
)
const (
	COL      = 80
)
func main() {
	texts, err := readFileByArg(os.Args[1:])
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	fmt.Println(cats(texts))
}

func readFileByArg(pathes []string) ([]string, error) {
    sarr := make([]string, 0, 0)
	for i:=0;i<len(pathes);i++{
		content, err := ioutil.ReadFile(pathes[i])
		if err != nil {
			return sarr, err
		}
		sarr= append(sarr,string(content))
	}
	return sarr, nil
}

func cats(texts_o []string)string {
	width:=int(math.Floor(float64((COL-len(texts_o))/len(texts_o))))
	texts:=make([]string,0,0)
	maxline:=0
	for i:=0;i<len(texts_o);i++{
		lines:=strings.Split(runewidth.Wrap(texts_o[i],width),"\n")
		str:=""
		for l :=0;l<len(lines);l++{
			str+= runewidth.Truncate(lines[l]+strings.Repeat("x",0)  ,width,"")+"\n"
			maxline=int(math.Max(float64(maxline) ,float64(l)))
		}
		texts=append(texts,str)
	}
	rtnstr:=""
	for i:=0;i<=maxline;i++{
		for c:=0;c<len(texts);c++{
			if  i<len(strings.Split(texts[c],"\n")){
				rtnstr+=runewidth.Truncate(strings.Split(texts[c],"\n")[i]+
				strings.Repeat(" ",width)  ,width,"")

			}else{
				rtnstr+=strings.Repeat(" ",width)
			}
			rtnstr+="|"
		}
		rtnstr=rtnstr[:len(rtnstr)-1]
		rtnstr+="\n"
	}
	return rtnstr
}
