package browser

import (
	"encoding/json"
	"fmt"

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
		votes := fmt.Sprintf("%0.0f", m["votes"])
		Document.Id("vote-total-"+asin).Set("innerHTML", votes)

	}()
}