package spec

import "fmt"

type depreciation struct {
	old, new string
}

// ErrDeprecated is a non-fatal error that occurs when deprecated fields are
// used in the spec.json
type ErrDeprecated []depreciation

func (e ErrDeprecated) Error() string {
	buf := ""
	for _, d := range e {
		buf += fmt.Sprintf("Warning: `%s` is deprecated, use `%s` instead.\n", d.old, d.new)
	}
	return buf
}

type ErrMistypedField struct {
	name string
	t    interface{}
}

func (e ErrMistypedField) Error() string {
	return fmt.Sprintf("`%s` is of type %T but should be string", e.name, e.t)
}

type ErrNoSpec struct {
	name string
}

func (e ErrNoSpec) Error() string {
	return fmt.Sprintf("unable to find a spec.json for environment `%s`.\nRefer to https://tanka.dev/directory-structure#environments for instructions", e.name)
}
