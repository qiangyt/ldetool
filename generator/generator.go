package generator

import (
	"io"

	"github.com/sirkon/ldetool/token"
)

// Generator describes methods needed of data lookup and extraction
type Generator interface {
	// Data handlers
	AddField(name string, fieldType string, t *token.Token)
	RegGravity(name string)

	// Pass
	PassN(n int)

	//
	AtEnd()

	// Head
	HeadString(anchor string, ignore bool)
	HeadChar(char string, ignore bool)

	// Lookups
	LookupString(anchor string, lower, upper int, ignore bool)
	LookupFixedString(anchor string, offset int, ignore bool)
	LookupChar(anchor string, lower, upper int, ignore bool)
	LookupFixedChar(anchor string, offset int, ignore bool)

	// Takes
	// Take before anchor (string or character)
	TakeBeforeString(name, fieldType, anchor string, lower, upper int, expand bool)
	TakeBeforeChar(name, fieldType, char string, lower, upper int, expand bool)

	// Take the rest
	TakeRest(name, fieldType string)

	// Optionals
	OpenOptionalScope(name string, t *token.Token)
	CloseOptionalScope()

	// Stress set mismatch treatment as serious error
	Stress()

	// Relax move mismatch treatment into unserious error
	Relax()

	// UseRule ...
	UseRule(name string, t *token.Token)

	// Push is used to signal all the data for current parser was generated
	Push()

	// Generate code at the place
	Generate(pkgName string, dest io.Writer)
}
