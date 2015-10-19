package verboten


import (
	"github.com/reiver/go-strm/driver"
)


const (
	// HEAD is a (string) constant that this Beginner driver
	// is registered under.
	HEAD = "HEAD"

	defaultLimit = 5
)


func init() {
	strmDriver := newStrmer()

	strmdriver.RegisterStrmer(HEAD, strmDriver)
}


type internalStrmer struct{}


func newStrmer() strmdriver.Strmer {
	strmDriver := internalStrmer{

	}

	return &strmDriver
}



func (strmDriver *internalStrmer) Strm(src <-chan []interface{}, dst chan<- []interface{}, args ...interface{}) {

	// Parse args.
	if 1 < len(args) {
//@TODO: Throw something better.
		panic("Too many parameters.")
	}

	var limit int = defaultLimit
	if 1 == len(args) {
		arg0 := args[0]
		switch n := arg0.(type) {
		case int:
			limit = n
		case int8:
			limit = int(n)
		case int16:
			limit = int(n)
		case int32:
			limit = int(n)
		case int64:
			limit = int(n)
		case uint:
			limit = int(n)
		case uint8:
			limit = int(n)
		case uint16:
			limit = int(n)
		case uint32:
			limit = int(n)
		case uint64:
			limit = int(n)
		default:
//@TODO: Throw something better.
			panic("Wrong type for limit.")
		}
	}

	// Execute.
	i :=0
	for datum := range src {
		if limit <= i {
			break
		}

		dst <- datum

		i++
	}
	go func() { // Drain the rest of the data on the stream.
		for range src {
			// Nothing here.
		}
		close(dst)
	}()
}
