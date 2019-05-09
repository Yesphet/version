package version

type Format int

const (
	FormatSimple  Format = iota
	FormatComplex
)

func New(format Format) Version {

	switch format {
	case FormatComplex:
		return newComplex()
	default:
		fallthrough
	case FormatSimple:
		return newSimple()
	}

}
