package handlers

import "github.com/fahmyabdul/golibs"

// Parameter :
type Handlers struct {
}

// Handle :
func (p *Handlers) Handle() error {
	golibs.Log.Println("| Hello World")
	return nil
}
