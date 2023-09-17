package writer

func newp[T any](v T) *T {
	return &v
}
