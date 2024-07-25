package hpdf

import (
	"io"
)

// IObj inteface for all pdf object
type IObj interface {
	init(func() *HPdf)
	getType() string
	write(w io.Writer, objID int) error
}
