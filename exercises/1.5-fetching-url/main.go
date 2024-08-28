package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		url = addProtocolPrefix(url)
		resp, err := http.Get(url)
		fmt.Printf("\nurl is: %s\n\n", url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "copy error %s: %v\n", url, err)
			_ = resp.Body.Close()
			os.Exit(1)
		}
		_ = resp.Body.Close()

		fmt.Printf("\nstatus code is: %s\n", resp.Status)
		if resp.StatusCode != http.StatusOK {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func addProtocolPrefix(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}

	return "https://" + url
}
