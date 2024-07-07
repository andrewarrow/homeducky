package browser

import "github.com/andrewarrow/feedback/wasm"

type Product struct {
	Id string
}

func handleAsins() {
	for _, a := range Document.Id("top").SelectAllByClass("voter") {
		p := Product{a.Id}
		a.EventWithId(p.click)
	}
}

func (p *Product) click() {
	m := map[string]any{}
	go wasm.DoPost("/core/asin/"+p.Id[5:], m)
}
