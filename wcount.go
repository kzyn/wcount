package main

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"os"
	"sort"
	"runtime"
)

type Entry struct {
	key string
	val int
}
type List []Entry

func main () {
	var seq string
	wordMap := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		if runtime.GOOS == "windows" {
			seq = "\r\n"
		} else {
			seq = "\n"
		}
		for _, line := range strings.Split(string(data), seq) {
			for _, word := range strings.Split(line, " ") {
				wordMap[word]++
			}
		}
	}

	list := List{}
	for key, val := range wordMap {
		add := Entry{key, val}
		list = append(list, add)
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].val == list[j].val {
			return (list[i].key < list[j].key)
		} else {
			return (list[i].val > list[j].val)
		}
	})

	for _, data := range list {
		if len(data.key) == 0 {
			continue
		}
		fmt.Printf("%v : %v \n", data.key, data.val)
	}
}
