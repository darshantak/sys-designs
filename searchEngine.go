package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Document struct {
	Name    int
	Content string
}
type Engine struct {
	Datasets map[string][]Document
}

func newSearchEngine() *Engine {
	return &Engine{
		Datasets: make(map[string][]Document),
	}
}
func (se *Engine) createDataset(name string) {
	if _, exists := se.Datasets[name]; exists {
		fmt.Println("This doc already exists")
		return
	}

	se.Datasets[name] = make([]Document, 0)
	fmt.Printf("Document %s is added. \n", name)
}

func (se *Engine) insertDoc(ds string, content string, count *int) {
	dataset := se.Datasets[ds]
	newDoc := Document{
		Name:    *count,
		Content: content,
	}
	dataset = append(dataset, newDoc)
	fmt.Println("Document has added")
	*count++
	return
}

func (se *Engine) deleteDoc(ds string, docName int) {
	docArr := se.Datasets[ds]
	for i, r := range se.Datasets[ds] {
		if r.Name == docName {
			se.Datasets[ds] = append(docArr[:i], docArr[i+1:]...)
		}
	}
	fmt.Printf("Document %d deleted successfully \n", docName)
	return
}

func (se *Engine) searchDoc(ds string, searchStr string) {
	dataset := se.Datasets[ds]
	var result []int
	for _, doc := range dataset {
		parts := strings.Fields(doc.Content)
		for _, word := range parts {
			if word == searchStr {
				result = append(result, doc.Name)
			}
		}
	}
	fmt.Println(result)
}
func SearchEngine() {
	scanner := bufio.NewScanner(os.Stdin)
	newSE := newSearchEngine()
	count := 1
	for {
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		parts := strings.Fields(input)
		cmd, args := parts[0], parts[1:]

		switch cmd {
		case "CREATE_DATASET":
			newSE.createDataset(args[0])

		case "INSERT_DOC":
			var contentToAdd string
			for _, v := range args[1:] {
				contentToAdd += v
			}
			newSE.insertDoc(args[0], contentToAdd, &count)

		case "DELETE_DOC":
			val, _ := strconv.Atoi(args[1])
			newSE.deleteDoc(args[0], val)
		case "SEARCH_DOC":
			newSE.searchDoc(args[0], args[1])
		}
	}
}

// CREATE_DOC doc1
// CREATE_DOC doc2
// INSERT_DOC doc1 str1 str2
// INSERT_DOC doc2 str2 str3
// DELETE_DOC doc1
// SEARCH_DOC str1 doc2
