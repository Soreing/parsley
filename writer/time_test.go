package writer

// func Test_TimeLen() {

// }

// func Test_WriteTime(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		In     string
// 		Out    []byte
// 		Space  int
// 		Stored int
// 		Cur    int
// 		Err    error
// 	}{
// 		{
// 			Name: "One Buffer",
// 			In:   "Test", Out: []byte("\"Test\""),
// 			Space: 10, Stored: 0, Cur: 6,
// 		},
// 		{
// 			Name: "Split Before Opening Quote",
// 			In:   "Test", Out: []byte("\"Test\""),
// 			Space: 0, Stored: 1, Cur: 6,
// 		},
// 		{
// 			Name: "Split After Opening Quote",
// 			In:   "Test", Out: []byte("\"Test\""),
// 			Space: 1, Stored: 1, Cur: 5,
// 		},
// 		{
// 			Name: "Split Before Closing Quote",
// 			In:   "Test", Out: []byte("\"Test\""),
// 			Space: 5, Stored: 1, Cur: 1,
// 		},
// 		{
// 			Name: "Split Mid Value",
// 			In:   "Test", Out: []byte("\"Test\""),
// 			Space: 2, Stored: 1, Cur: 4,
// 		},
// 		{
// 			Name: "Escape Characters",
// 			In:   "\n\t\u0000", Out: []byte("\"\\n\\t\\u0000\""),
// 			Space: 20, Stored: 0, Cur: 12,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.String(test.In)
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

// func Test_WriteStringp(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		In     *string
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
// 			In:   newp("Test"), Out: []byte("\"Test\""),
// 			Space: 10, Stored: 0, Cur: 6,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.Stringp(test.In)
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

// func Test_WriteStrings(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		In     []string
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
// 			In:   []string{}, Out: []byte("[]"),
// 			Space: 10, Stored: 0, Cur: 2,
// 		},
// 		{
// 			Name: "Split Buffer (empty)",
// 			In:   []string{}, Out: []byte("[]"),
// 			Space: 1, Stored: 1, Cur: 1,
// 		},
// 		{
// 			Name: "One Buffer (value)",
// 			In:   []string{"one"}, Out: []byte(`["one"]`),
// 			Space: 30, Stored: 0, Cur: 7,
// 		},
// 		{
// 			Name: "One Buffer (values)",
// 			In:   []string{"one", "two", "three"}, Out: []byte(`["one","two","three"]`),
// 			Space: 30, Stored: 0, Cur: 21,
// 		},
// 		{
// 			Name: "Split After Opening Quote",
// 			In:   []string{"one", "two", "three"}, Out: []byte(`["one","two","three"]`),
// 			Space: 2, Stored: 1, Cur: 19,
// 		},
// 		{
// 			Name: "Split Befroe Closing Quote",
// 			In:   []string{"one", "two", "three"}, Out: []byte(`["one","two","three"]`),
// 			Space: 5, Stored: 1, Cur: 16,
// 		},
// 		{
// 			Name: "Split After [",
// 			In:   []string{"one", "two", "three"}, Out: []byte(`["one","two","three"]`),
// 			Space: 1, Stored: 1, Cur: 20,
// 		},
// 		{
// 			Name: "Split Before ]",
// 			In:   []string{"one", "two", "three"}, Out: []byte(`["one","two","three"]`),
// 			Space: 20, Stored: 1, Cur: 1,
// 		},
// 		{
// 			Name: "Split Mid Value",
// 			In:   []string{"one", "two", "three"}, Out: []byte(`["one","two","three"]`),
// 			Space: 3, Stored: 1, Cur: 18,
// 		},
// 		{
// 			Name: "Split Before Comma",
// 			In:   []string{"one", "two", "three"}, Out: []byte(`["one","two","three"]`),
// 			Space: 6, Stored: 1, Cur: 15,
// 		},
// 		{
// 			Name: "Split After Comma",
// 			In:   []string{"one", "two", "three"}, Out: []byte(`["one","two","three"]`),
// 			Space: 7, Stored: 1, Cur: 14,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.Strings(test.In)
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
