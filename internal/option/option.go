// Package option implements an Option struct to conifigure the parameters for
// a search in Sieve.
package option

type caseSearch int

const (
	Sensitive caseSearch = iota
	Insensitive
	Smart
)

// Option implements the configuration for a search in Sieve.
type Option struct {
	Case caseSearch
	
}
