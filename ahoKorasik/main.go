package main

import (
	"errors"
	"fmt"
)

var ErrorNotFoundTree = errors.New("Tree is not found on parent")

type Bor struct {
	id       rune
	children []*Bor
	parent   *Bor
	suffics  *Bor
	final    bool
}

func printBor(bor *Bor, prefix string) {
	fmt.Println(prefix, string(bor.id))
	if bor.parent != nil {
		if bor.parent.id == 0 {
			fmt.Println(prefix, "<root>")
		} else {
			fmt.Println(prefix, string(bor.parent.id))
		}
	} else {
		fmt.Println(prefix, "<root>")
	}
	if bor.suffics != nil {
		fmt.Println(prefix, *bor.suffics)
	}
	fmt.Println(prefix, bor.final)
	if len(bor.children) > 0 {
		for _, c := range bor.children {
			printBor(c, prefix+prefix)
		}
	}
}

func main() {
	str := []string{
		"A",
		"ABA",
		"ARAB",
		"ASS",
		"BAR",
		"BASS",
		"CAR",
		"R",
		"RA",
		"RAB",
	}
	bor := createBor(str)
	printBor(bor, "-")
}

func findBor(s rune, tree *Bor) (*Bor, error) {
	for _, t := range tree.children {
		if t.id == s {
			return t, nil
		}
	}
	return nil, ErrorNotFoundTree
}

func createBor(findStr []string) *Bor {
	bor := Bor{}
	var last *Bor
	var err error
	for _, str := range findStr {
		tree := &bor
		for _, s := range str {
			last, err = findBor(s, tree)
			if err != nil {
				last = &Bor{id: s, parent: tree}
				tree.children = append(tree.children, last)
			}
			tree = last
		}
		last.final = true
	}
	return &bor
}
