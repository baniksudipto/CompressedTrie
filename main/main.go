package main

import (
	"awesomeProject/models"
	"fmt"
	"reflect"
	"runtime"
	"sort"
)

func main() {
	PrintMemUsage()
	t := models.NewTrie()
	var s = []string{"achieve", "affect", "appoint", "assume", "begin", "break", "calculate", "campaign", "close", "collapse", "commit", "comprise", "concede", "concern", "confront", "continue", "defend", "define", "direct", "engage", "flash", "flee", "handle", "hurry", "improve", "injure", "involve", "leave", "lend", "locate", "ma", "mark", "market", "park", "perform", "permit", "raise", "range", "recall", "relieve", "replace", "seek", "spin", "stem", "stop", "suggest", "take", "transform", "transport", "vote", "wake"}
	PrintMemUsage()
	for _, example := range s {
		t.Add(example)
	}

	searchCheck(s, t)
	getStringsCheck(t, s)

	PrintMemUsage()
	t = nil

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()
}

func getStringsCheck(t *models.Trie, s []string) {
	res := t.GetStrings()
	sort.Strings(res)
	fmt.Println("GetStrings Works: ", reflect.DeepEqual(res, s))
}

func searchCheck(s []string, t *models.Trie) {
	searchWorks := true
	for _, e := range s {
		searchWorks = searchWorks && (t.Search(e) != nil)
	}
	fmt.Println("SearchString Works: ", searchWorks)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v KiB", bToKb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v KiB", bToKb(m.TotalAlloc))
	fmt.Printf("\tSys = %v KiB", bToKb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToKb(b uint64) uint64 {
	return b / 1024
}
