package hpdf

type ICacheColorText interface {
	ICacheContent
	equal(obj ICacheColorText) bool
}
