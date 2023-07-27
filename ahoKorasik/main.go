package main

import (
	"errors"
	"fmt"
)

var ErrorNotFoundTree = errors.New("Tree is not found on parent")

type Bor struct {
	id        rune
	fullName  []rune
	children  []*Bor
	parent    *Bor
	suffix    *Bor
	finalLink *Bor
	final     bool
}

// Печатает бор деревьев для визуальной наглядности.
// Существенной пользы не приносит
func printBor(bor *Bor, prefix string) {
	fmt.Println(prefix, "id", string(bor.id))
	fmt.Println(prefix, "fullName", string(bor.fullName))
	if bor.parent != nil {
		if bor.parent.id == 0 {
			fmt.Println(prefix, "parent", "<root>")
		} else {
			fmt.Println(prefix, "parent", string(bor.parent.id))
		}
	} else {
		fmt.Println(prefix, "parent", "<root>")
	}
	if bor.suffix != nil {
		fmt.Println(prefix, "suffix", string(bor.suffix.id))
	}
	fmt.Println(prefix, "final", bor.final)
	if bor.finalLink != nil {
		fmt.Println(prefix, "finalLink", string(bor.finalLink.fullName))
	} else {
		fmt.Println(prefix, "finalLink", "<root>")
	}
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
	for _, str := range findAll(*bor, "CARABASS BARABAS") {
		fmt.Println(string(str))
	}
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
				fullName := tree.fullName
				fullName = append(fullName, last.id)
				last.fullName = fullName
				tree.children = append(tree.children, last)
			}
			tree = last
		}
		last.final = true
	}
	var suffixBor []*Bor
	for _, b := range bor.children {
		b.suffix = &bor
		suffixBor = append(suffixBor, b.children...)
	}
	createSuffix(suffixBor, &bor)
	createFinalLink(suffixBor)
	return &bor
}

func createSuffix(bor []*Bor, mainBor *Bor) {
	var newBor []*Bor
	for _, b := range bor {
		childParentSuffix := b.parent.suffix.children
		for _, cps := range childParentSuffix {
			if b.id == cps.id {
				b.suffix = cps
				break
			}
			b.suffix = mainBor
		}
		newBor = append(newBor, b.children...)
	}
	if len(newBor) == 0 {
		return
	}
	createSuffix(newBor, mainBor)
}

func createFinalLink(bor []*Bor) {
	var newBor []*Bor
	for _, b := range bor {
		b.finalLink = findFinalLink(b.suffix)
		newBor = append(newBor, b.children...)
	}
	if len(newBor) == 0 {
		return
	}
	createFinalLink(newBor)
}

func findFinalLink(tree *Bor) *Bor {
	if tree.final {
		return tree
	}
	if tree.suffix == nil {
		return nil
	}
	return findFinalLink(tree.suffix)
}

func findAll(bor Bor, text string) [][]rune {
	var res [][]rune
	var tree *Bor
	tree = &bor
	for _, t := range text {
		tree = findCurrentTree(tree, t)
		if tree.final {
			res = append(res, tree.fullName)
		}
		if tree.finalLink != nil {
			res = append(res, resFinalLink(tree.finalLink)...)
		}
	}
	return res
}

func findCurrentTree(bor *Bor, t rune) *Bor {
	for _, c := range bor.children {
		if c.id == t {
			return c
		}
	}
	if bor.suffix == nil {
		return bor
	}
	return findCurrentTree(bor.suffix, t)
}

func resFinalLink(tree *Bor) [][]rune {
	var res [][]rune
	res = append(res, tree.fullName)
	if tree.finalLink != nil {
		res = append(res, resFinalLink(tree.finalLink)...)
	}
	return res
}
