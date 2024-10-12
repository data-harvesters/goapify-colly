package goapifycolly

import (
	"github.com/data-harvesters/goapify"
	"github.com/gocolly/colly/v2"
)

type Colly struct {
	controller *colly.Collector

	actor *goapify.Actor
}

func NewColly(actor *goapify.Actor, options ...colly.CollectorOption) *Colly {
	c := colly.NewCollector(options...)

	return &Colly{
		controller: c,

		actor: actor,
	}
}

func (c *Colly) Visit(url string) {
	c.controller.Visit(url)
}

func (c *Colly) HtmlHook(goquerySelector string, f colly.HTMLCallback) {
	c.controller.OnHTML(goquerySelector, f)
}

func (c *Colly) RequestHook(f colly.RequestCallback) {
	c.controller.OnRequest(f)
}

func (c *Colly) ResponseHook(f colly.ResponseCallback) {
	c.controller.OnResponse(f)
}

func (c *Colly) Controller() *colly.Collector {
	return c.controller
}

// Proxied returns a proxied Colly if available if not just return Colly
func (c *Colly) Proxied() *Colly {
	if c.actor.ProxyConfiguration == nil {
		return c
	}
	proxyUrl, err := c.actor.ProxyConfiguration.Proxy()
	if err != nil {
		return c
	}

	co := c
	err = co.controller.SetProxy(proxyUrl.String())
	if err != nil {
		return c
	}

	return co
}
