# Goapify-Colly
a goapify Colly module

## Usage
```go
c := goapifycolly.NewColly(actor, colly.AllowedDomains("go-colly.org"))
c = c.Proxied() // get a proxied client (if avaliable)

c.HtmlHook("a[href]", func(e *colly.HTMLElement) {
	e.Request.Visit(e.Attr("href"))
})

c.RequestHook(func(r *colly.Request) {
	fmt.Println("Visiting:", r.URL)
})

c.Visit("http://go-colly.org/")
```