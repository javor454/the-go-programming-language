package main

import (
	"html/template"
	"os"
	"time"
)

type Product struct {
	Name        string
	Price       float64
	Description string
	InStock     bool
	Categories  []string
	AddedDate   time.Time
}

func main() {
	const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Name}} - Product Details</title>
    <style>
        .product { max-width: 600px; margin: 20px auto; padding: 20px; }
        .out-of-stock { color: red; }
        .price { font-size: 1.2em; font-weight: bold; }
    </style>
</head>
<body>
    <div class="product">
        <h1>{{.Name}}</h1>
        
        <div class="price">${{printf "%.2f" .Price}}</div>
        
        <p>{{.Description}}</p>
        
        {{if .InStock}}
            <p>Available for purchase</p>
        {{else}}
            <p class="out-of-stock">Currently out of stock</p>
        {{end}}
        
        <h3>Categories:</h3>
        <ul>
        {{range .Categories}}
            <li>{{.}}</li>
        {{end}}
        </ul>
        
        <p>Added: {{.AddedDate.Format "January 2, 2006"}}</p>
        
        {{with .Description}}
            <div class="description">
                <h3>About this product:</h3>
                {{.}}
            </div>
        {{end}}
        
        {{/* This is a template comment */}}
        {{if ge .Price 100.0}}
            <p><strong>Premium Product</strong></p>
        {{end}}
    </div>
</body>
</html>`

	t := template.Must(template.New("product").Parse(tmpl))

	product := Product{
		Name:        "Gaming Laptop",
		Price:       1299.99,
		Description: "High-performance gaming laptop with RGB keyboard",
		InStock:     true,
		Categories:  []string{"Electronics", "Gaming", "Computers"},
		AddedDate:   time.Now(),
	}

	err := t.Execute(os.Stdout, product)
	if err != nil {
		panic(err)
	}
}
