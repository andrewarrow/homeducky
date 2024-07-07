package browser

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
}
