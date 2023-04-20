package parsley

type bufferType int

const (
	none bufferType = iota
	fixed
	relative
)

// Filter specifies a filter structure with a field name (key) and a nested
// filter list used for nested objects.
type Filter struct {
	Field  string
	Filter []Filter
}

// Config describes optional settings that alter the behavior of the encoder/decoder.
type Config struct {
	btype  bufferType
	bsize  int
	filter []Filter
}

// UseFilter creates a config from a filter list. Only fields that appear in the
// filter list will be processed by the encoder/decoder If the filter is nil,
// all elements get processed.
//
// Example:
//    var filter = []parsley.Filter{
//        {"name": []parsley.Filter{
//            {"firstName": nil},
//            {"lastName": nil},
//        }},
//        {"address": []parsley.Filter{
//            {"country": nil},
//            {"city": nil},
//        }},
//        {"age": nil},
//    }
func UseFilter(filter []Filter) Config {
	return Config{
		filter: filter,
	}
}

// UseFixedBuffer creates a config to add extra buffer space for the encoder.
// The function adds a fixed number of bytes space.
//
// A larger buffer can speed up encoding by reducing required allocations.
// Extra buffer space is only recommended if the object contains strings with
// characters that should be escaped.
func UseFixedBuffer(bytes int) Config {
	return Config{
		btype: fixed,
		bsize: bytes,
	}
}

// UseRelativeBuffer creates a config to add extra buffer space for the encoder.
// The function calculates the length of all string fields and uses a percentage
// of that as extra buffer space.
//
// A larger buffer can speed up encoding by reducing required allocations.
// Extra buffer space is only recommended if the object contains strings with
// characters that should be escaped.
func UseRelativeBuffer(percentage int) Config {
	return Config{
		btype: relative,
		bsize: percentage,
	}
}

// MergeConfigs takes multiple config options and merges them into a single object.
func MergeConfigs(cfgs ...Config) Config {
	cfg := Config{
		btype:  fixed,
		bsize:  0,
		filter: nil,
	}

	for _, e := range cfgs {
		if e.btype != none {
			cfg.btype = e.btype
			cfg.bsize = e.bsize
		}
		if e.filter != nil {
			cfg.filter = e.filter
		}
	}

	return cfg
}
