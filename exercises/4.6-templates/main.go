package main

import (
	"os"
	"strings"
	"text/template"
	"time"
)

type Report struct {
	Title    string
	Items    []Item
	Total    float64
	Customer Customer
}

type Item struct {
	Name     string
	Quantity int
	Price    float64
}

type Customer struct {
	Name    string
	Address string
}

func main() {
	// Step 1: Parse template
	const tmpl = `
SALES REPORT: {{.Title | uppercase}}
Date: {{now | date}}

Customer Information:
Name: {{.Customer.Name}}
Address: {{.Customer.Address}}

Items:
{{range .Items}}
- {{.Name}}: {{.Quantity}}x @ ${{.Price}} = ${{multiply .Quantity .Price}}
{{end}}

{{if gt .Total 1000.0}}HIGH VALUE ORDER!{{end}}
Total: ${{printf "%.2f" .Total}}
`

	// Create function map
	funcMap := template.FuncMap{
		"uppercase": strings.ToUpper,
		"multiply": func(x int, y float64) float64 {
			return float64(x) * y
		},
		"now": time.Now,
		"date": func(t time.Time) string {
			return t.Format("2006-01-02")
		},
	}

	// Parse template with functions
	t := template.Must(template.New("report").Funcs(funcMap).Parse(tmpl))

	// Step 2: Execute with data
	report := Report{
		Title: "Monthly Sales",
		Items: []Item{
			{"Widget", 5, 9.99},
			{"Gadget", 2, 499.99},
		},
		Total: 1029.93,
		Customer: Customer{
			Name:    "John Doe",
			Address: "123 Main St",
		},
	}

	err := t.Execute(os.Stdout, report)
	if err != nil {
		panic(err)
	}
}
