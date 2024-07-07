package browser

import (
	"encoding/json"

	"github.com/andrewarrow/feedback/wasm"
)

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
	go func() {
		asin := p.Id[5:]
		js, _ := wasm.DoPost("/core/asin/"+asin, m)
		var m map[string]any
		json.Unmarshal([]byte(js), &m)
		Document.Id("vote-total-"+asin).Set("innerHTML", m["votes"])

	}()
}
