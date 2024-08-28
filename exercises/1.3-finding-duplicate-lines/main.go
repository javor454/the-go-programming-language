package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ReadLinesStreamingMode()
	// ReadLinesLoadToMemoryMode()
}

func ReadLinesLoadToMemoryMode() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func ReadLinesStreamingMode() {
	wordPerFile := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "No files to read\n")
		return
	} else {
		for _, arg := range files {
			countPerWord := make(map[string]int)

			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, countPerWord)
			wordPerFile[arg] = countPerWord
			_ = f.Close()
		}

	}

	for filename, counts := range wordPerFile {
		if len(counts) != 0 {
			fmt.Printf("File: %s\n", filename)
			for word, count := range counts {
				if count > 1 {
					fmt.Printf("%d\t%s\n", count, word)
				}
			}
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
