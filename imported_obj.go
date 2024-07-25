package hpdf

import (
	"io"
)

// ImportedObj : imported object
type ImportedObj struct { //impl IObj
	Data string
}

func (c *ImportedObj) init(funcGetRoot func() *HPdf) {

}

func (c *ImportedObj) getType() string {
	return "Imported"
}

func (c *ImportedObj) write(w io.Writer, objID int) error {
	if c != nil {
		io.WriteString(w, c.Data)
	}
	return nil
}
