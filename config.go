package parsley

type bufferType int

const (
	none bufferType = iota
	fixed
	relative
)

type Filter struct {
	Field  string
	Filter []Filter
}

type config struct {
	btype  bufferType
	bsize  int
	filter []Filter
}

// Use a filter to decide at runtime which fields should be encoded/decoded
func UseFilter(filter []Filter) config {
	return config{
		filter: filter,
	}
}

// Use a fixed size extra buffer space for encoding.
func UseFixedBuffer(bytes int) config {
	return config{
		btype: fixed,
		bsize: bytes,
	}
}

// Use a relative size extra buffer space for encoding.
func UseRelativeBuffer(percentage int) config {
	return config{
		btype: relative,
		bsize: percentage,
	}
}

// Merges configs into one config.
func MergeConfigs(cfgs ...config) config {
	cfg := config{
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
