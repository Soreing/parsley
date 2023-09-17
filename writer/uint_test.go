package writer

// func Test_WriteUInt(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		In     uint
// 		Out    []byte
// 		Space  int
// 		Stored int
// 		Cur    int
// 		Err    error
// 	}{
// 		{
// 			Name: "One Buffer",
// 			In:   1234, Out: []byte("1234"),
// 			Space: 10, Stored: 0, Cur: 4,
// 		},
// 		{
// 			Name: "Split Buffer",
// 			In:   1234, Out: []byte("1234"),
// 			Space: 2, Stored: 1, Cur: 4,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.UInt(test.In)
// 			res := w.Build()
// 			if string(res) != string(test.Out) {
// 				t.Errorf("got result %s, want %s", res, test.Out)
// 			}
// 			if w.Cursor != test.Cur {
// 				t.Errorf("got cursor %d, want %d", w.Cursor, test.Cur)
// 			}
// 			if len(w.Storage) != test.Stored {
// 				t.Errorf("got stored buffers %d, want %d", len(w.Storage), test.Stored)
// 			}
// 		})
// 	}
// }

// func Test_WriteUIntp(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		In     *uint
// 		Out    []byte
// 		Space  int
// 		Stored int
// 		Cur    int
// 		Err    error
// 	}{
// 		{
// 			Name: "One Buffer (nil)",
// 			In:   nil, Out: []byte("null"),
// 			Space: 10, Stored: 0, Cur: 4,
// 		},
// 		{
// 			Name: "Split Buffer (nil)",
// 			In:   nil, Out: []byte("null"),
// 			Space: 2, Stored: 1, Cur: 2,
// 		},
// 		{
// 			Name: "Pointer With Value",
// 			In:   newp(uint(1234)), Out: []byte("1234"),
// 			Space: 10, Stored: 0, Cur: 4,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.UIntp(test.In)
// 			res := w.Build()
// 			if string(res) != string(test.Out) {
// 				t.Errorf("got result %s, want %s", res, test.Out)
// 			}
// 			if w.Cursor != test.Cur {
// 				t.Errorf("got cursor %d, want %d", w.Cursor, test.Cur)
// 			}
// 			if len(w.Storage) != test.Stored {
// 				t.Errorf("got stored buffers %d, want %d", len(w.Storage), test.Stored)
// 			}
// 		})
// 	}
// }

// func Test_WriteUInts(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		In     []uint
// 		Out    []byte
// 		Space  int
// 		Stored int
// 		Cur    int
// 		Err    error
// 	}{
// 		{
// 			Name: "One Buffer (nil)",
// 			In:   nil, Out: []byte("null"),
// 			Space: 10, Stored: 0, Cur: 4,
// 		},
// 		{
// 			Name: "Split Buffer (nil)",
// 			In:   nil, Out: []byte("null"),
// 			Space: 2, Stored: 1, Cur: 2,
// 		},
// 		{
// 			Name: "One Buffer (empty)",
// 			In:   []uint{}, Out: []byte("[]"),
// 			Space: 10, Stored: 0, Cur: 2,
// 		},
// 		{
// 			Name: "Split Buffer (empty)",
// 			In:   []uint{}, Out: []byte("[]"),
// 			Space: 1, Stored: 1, Cur: 1,
// 		},
// 		{
// 			Name: "One Buffer (value)",
// 			In:   []uint{1234}, Out: []byte("[1234]"),
// 			Space: 30, Stored: 0, Cur: 6,
// 		},
// 		{
// 			Name: "One Buffer (values)",
// 			In:   []uint{1234, 56789, 1357}, Out: []byte("[1234,56789,1357]"),
// 			Space: 30, Stored: 0, Cur: 17,
// 		},
// 		{
// 			Name: "Split After [",
// 			In:   []uint{1234, 56789, 1357}, Out: []byte("[1234,56789,1357]"),
// 			Space: 1, Stored: 1, Cur: 16,
// 		},
// 		{
// 			Name: "Split Before ]",
// 			In:   []uint{1234, 56789, 1357}, Out: []byte("[1234,56789,1357]"),
// 			Space: 16, Stored: 1, Cur: 1,
// 		},
// 		{
// 			Name: "Split Mid Value",
// 			In:   []uint{1234, 56789, 1357}, Out: []byte("[1234,56789,1357]"),
// 			Space: 3, Stored: 1, Cur: 16,
// 		},
// 		{
// 			Name: "Split Before Comma",
// 			In:   []uint{1234, 56789, 1357}, Out: []byte("[1234,56789,1357]"),
// 			Space: 5, Stored: 1, Cur: 12,
// 		},
// 		{
// 			Name: "Split After Comma",
// 			In:   []uint{1234, 56789, 1357}, Out: []byte("[1234,56789,1357]"),
// 			Space: 6, Stored: 1, Cur: 11,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.UInts(test.In)
// 			res := w.Build()
// 			if string(res) != string(test.Out) {
// 				t.Errorf("got result %s, want %s", res, test.Out)
// 			}
// 			if w.Cursor != test.Cur {
// 				t.Errorf("got cursor %d, want %d", w.Cursor, test.Cur)
// 			}
// 			if len(w.Storage) != test.Stored {
// 				t.Errorf("got stored buffers %d, want %d", len(w.Storage), test.Stored)
// 			}
// 		})
// 	}
// }
