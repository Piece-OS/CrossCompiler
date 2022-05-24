// Code generated by command: go run fe_amd64_asm.go -out ../fe_amd64.s -stubs ../fe_amd64.go -pkg field. DO NOT EDIT.

//go:build amd64 && gc && !purego

package field

// feMul sets out = a * b. It works like feMulGeneric.
//go:noescape
func feMul(out *Element, a *Element, b *Element)

// feSquare sets out = a * a. It works like feSquareGeneric.
//go:noescape
func feSquare(out *Element, a *Element)