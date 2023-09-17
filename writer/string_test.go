package writer

// func Test_StringLen(t *testing.T) {
// 	tests := []struct {
// 		Name     string
// 		In       string
// 		Fixed    int
// 		Volatile int
// 	}{
// 		{
// 			Name: "Normal String",
// 			In:   "Test String", Fixed: 13, Volatile: 11,
// 		},
// 		{
// 			Name: "Empty String",
// 			In:   "", Fixed: 2, Volatile: 0,
// 		},
// 		{
// 			Name: "Escaped String",
// 			In:   "Some \"String\" Value", Fixed: 21, Volatile: 19,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			fix, vol := StringLen(test.In)
// 			if fix != test.Fixed {
// 				t.Errorf("got fixed length %d, want %d", fix, test.Fixed)
// 			}
// 			if vol != test.Volatile {
// 				t.Errorf("got volatile length %d, want %d", vol, test.Volatile)
// 			}
// 		})
// 	}
// }

// func Test_StringpLen(t *testing.T) {
// 	tests := []struct {
// 		Name     string
// 		In       *string
// 		Fixed    int
// 		Volatile int
// 	}{
// 		{
// 			Name: "Nil String",
// 			In:   nil, Fixed: 4, Volatile: 0,
// 		},
// 		{
// 			Name: "Normal String",
// 			In:   newp("Test String"), Fixed: 13, Volatile: 11,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			fix, vol := StringpLen(test.In)
// 			if fix != test.Fixed {
// 				t.Errorf("got fixed length %d, want %d", fix, test.Fixed)
// 			}
// 			if vol != test.Volatile {
// 				t.Errorf("got volatile length %d, want %d", vol, test.Volatile)
// 			}
// 		})
// 	}
// }

// func TestStringsLen(t *testing.T) {
// 	tests := []struct {
// 		Name     string
// 		In       []string
// 		Fixed    int
// 		Volatile int
// 	}{
// 		{
// 			Name: "Nil",
// 			In:   nil, Fixed: 4, Volatile: 0,
// 		},
// 		{
// 			Name: "Empty",
// 			In:   []string{}, Fixed: 2, Volatile: 0,
// 		},
// 		{
// 			Name: "One Value",
// 			In:   []string{"one"}, Fixed: 7, Volatile: 3,
// 		},
// 		{
// 			Name: "Two Values",
// 			In:   []string{"one", "true"}, Fixed: 14, Volatile: 7,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			fix, vol := StringsLen(test.In)
// 			if fix != test.Fixed {
// 				t.Errorf("got fixed length %d, want %d", fix, test.Fixed)
// 			}
// 			if vol != test.Volatile {
// 				t.Errorf("got volatile length %d, want %d", vol, test.Volatile)
// 			}
// 		})
// 	}
// }

// func Test_WriteRaw(t *testing.T) {
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
// 			In:   "Test", Out: []byte("Test"),
// 			Space: 10, Stored: 0, Cur: 4,
// 		},
// 		{
// 			Name: "Split Buffer",
// 			In:   "Test", Out: []byte("Test"),
// 			Space: 2, Stored: 1, Cur: 2,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.Raw(test.In)
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

// func Test_WriteByte(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		In     byte
// 		Out    []byte
// 		Space  int
// 		Stored int
// 		Cur    int
// 		Err    error
// 	}{
// 		{
// 			Name: "Free Space",
// 			In:   'X', Out: []byte("X"),
// 			Space: 10, Stored: 0, Cur: 1,
// 		},
// 		{
// 			Name: "No Free Space",
// 			In:   'X', Out: []byte("X"),
// 			Space: 0, Stored: 1, Cur: 1,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.Byte(test.In)
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

// func Test_WriteString(t *testing.T) {
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
