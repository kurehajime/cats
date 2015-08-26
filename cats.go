// cats.go
package main

import (
	"flag"
	"fmt"
	"github.com/mattn/go-runewidth"
	"io/ioutil"
	"math"
	"os"
	"path"
	"strings"
	"sync"
)

var (
	COL = 80
)

func main() {
	var headFlag bool
	flag.BoolVar(&headFlag, "h", false, "add filename")
	flag.IntVar(&COL, "c", 80, "columns count")

	flag.Parse()
	texts, err := readFilesByArg(flag.Args(), headFlag)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	fmt.Println(cats(texts))

	fmt.Println(os.Getenv("COLUMNS"))
}

//read multi text files
func readFilesByArg(pathes []string, headFlag bool) ([]string, error) {
	sarr := make([]string, len(pathes), len(pathes))
	wg := new(sync.WaitGroup)
	wg.Add(len(pathes))
	for i := 0; i < len(pathes); i++ {
		go func(i0 int) {
			content, err := ioutil.ReadFile(pathes[i0])
			if err != nil {
				content = []byte(err.Error())
			}
			if headFlag {
				_, fn := path.Split(pathes[i0])
				sarr[i0] = fn + "\n" + string(content)
			} else {
				sarr[i0] = string(content)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return sarr, nil
}

//multi columns cat command
func cats(texts_o []string) string {
	width := int(math.Floor(float64((COL - len(texts_o)) / len(texts_o))))
	texts := make([]string, 0, 0)
	maxline := 0
	rtnstr := ""

	//wrap text
	for i := 0; i < len(texts_o); i++ {
		lines := strings.Split(runewidth.Wrap(texts_o[i], width), "\n")
		str := ""
		for l := 0; l < len(lines); l++ {
			str += runewidth.Truncate(lines[l]+strings.Repeat("x", 0), width, "") + "\n"
			maxline = int(math.Max(float64(maxline), float64(l)))
		}
		texts = append(texts, str)
	}

	//put columns
	for i := 0; i <= maxline; i++ {
		for c := 0; c < len(texts); c++ {
			if i < len(strings.Split(texts[c], "\n")) {
				rtnstr += runewidth.Truncate(strings.Split(texts[c], "\n")[i]+
					strings.Repeat(" ", width), width, "")

			} else {
				rtnstr += strings.Repeat(" ", width)
			}
			rtnstr += "|"
		}
		rtnstr = rtnstr[:len(rtnstr)-1]
		rtnstr += "\n"
	}
	return rtnstr
}
