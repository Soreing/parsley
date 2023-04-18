package reader

import (
	"strconv"
	"testing"
)

// go test -bench=. -benchmem -benchtime=10000000x -cpu=1 ./reader

const str = "\"4rqxGqtk8DJZpMdTiWnqW9rlBMl5gHJ3jAIb4Ujib5JZerPCVZOEujdZF0OOIlnrsc83jb1F7c9n87RbELnQ1rTo1E4TlmNu73i7" +
	"qRp3FJDxUygcAUyovxHao92FxBzwv6wjpTTjcb7AciOkYwzCO86DHYBf3lxLknQbAENRKZhaVqVRVKhrpNCtfo9aTBwJkkEonpdx" +
	"V6y7dNss4MTVrNTWkrU1y4OxAl033jbWs4g9v0UnrepRpEhcDjLPNbhYlmmnr7mZMcZUA7v3MxpcodBDK6l8lOfRWDqIxqXkPQTp" +
	"HBvUfjmi14hndohLyF70Yy0ykyaNSEtMrMkmFOCKjsN2qcDiQ7bUKHKpi1R1jZbjtqGpuhrnPAdXoVS6Gp3UJU3Er7SUmE7SboyN" +
	"Rcp1vtERzuGc7h7hhO220vKYLUlM48VULoouOoVHqJCD0m6s9vOk2LRRTYrQPSMM6yAi18K1TsJyqUcIPfSBdJAAHtkAHA5VDmgH\""

const stre = "\"4rqxGqtk8DJZ\\npMdTiWnqW9rlBMl5gHJ3jAIb4Ujib5JZ\\nerPCVZOEujdZF0OOIlnrsc83jb1F\\n\\n7c9n87RbELnQ1rTo\\uDEAD1E4TlmNu73i7" +
	"qRp3FJDxUygcAUyovxHao\\n92FxBzwv6wjpTTjcb7AciOkYwzCO86D\\nHYBf3lxLknQbAENRKZhaVqVRVKhr\\n\\npNCtfo9aTBwJkkEonpdx" +
	"V6y7dNss4MTVrNTWkrU1y4OxAl\\n033jbWs4g9v0UnrepRpEhcDjLPNbhYl\\nmmnr7mZMcZUA7v3MxpcodBDK6l8l\\n\\nOfRWDqIxqXkPQTp" +
	"HBvUfjmi14hndohLyF70Yy0ykyaNSEt\\nMrMkmFOCKjsN2qcDiQ7bUKHKpi1R1jZ\\nbjtqGpuhrnPAdXoVS6Gp3UJU3Er7SU\\n\\nmE7SboyN" +
	"Rcp1vtERzuGc7h7hhO220vKYLUlM48VULoouO\\noVHqJCD0m6s9vOk2LRRTYrQPSMM6yA\\ni18K1TsJyqUcIPfSBdJAAHtkAHA5VDm\\n\\ngH\""

const strs = `["Cq0o7rS4uG","tKHXIMxgcX","yxlWAeGNcz","n9wzNjrYuh","VxAULoo7ly","hWMC25YOqc","xkg8C0lnb9","25LzgA1X6Z","JdZkNZPUYm","LsDPmZCR3w","8wDHAf5oCu","nozvWtKZ8W","VcamuftCS0","c22Jv7MHVu","DkvW10nHLk","IEtjjZ6LLb","9khV0aERmi","nEMwBl6dHb","3nyjXUPp9P","bhTgf8HzG3","wb75gwgR5A","nLySBTPE7W","BXczkTFKjF","G46IQYZ0xT","M43rPExOYb"]`

const i64 = `-6162958237523.345e-4`
const i64z = `-0.00006162958237523345e+15`
const ui64 = `6162958237523.345e-4`
const ui64z = `0.00006162958237523345e+15`

const flt64 = `723453.19034675134561064923e+129`
const fltfb = `723453.19034675134561064923e+350`

func Benchmark_ReadFloatFallback(b *testing.B) {
	r := NewReader([]byte(fltfb))
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if _, err := r.Float64(); err != nil {
			//b.Fatal(err)
		}
	}
}

func Benchmark_ReadFloatFallbackStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := strconv.ParseFloat(fltfb, 64); err != nil {
			//b.Fatal(err)
		}
	}
}

func Benchmark_ReadFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, _, _, _, _, _, _, _, ok := readFloat([]byte(flt64)); !ok {
			b.Fatal(ok)
		}
	}
}

func Benchmark_SkipNumber(b *testing.B) {
	r := NewReader([]byte(i64))
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if err := r.skipNumber(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadFloat64(b *testing.B) {
	r := NewReader([]byte(flt64))
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if _, err := r.Float64(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ParseFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := strconv.ParseFloat(flt64, 64); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadInt64(b *testing.B) {
	r := NewReader([]byte(i64))
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if _, err := r.Int64(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadInt64Zero(b *testing.B) {
	r := NewReader([]byte(i64z))
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if _, err := r.Int64(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadUInt64(b *testing.B) {
	r := NewReader([]byte(ui64))
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if _, err := r.UInt64(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadUInt64Zero(b *testing.B) {
	r := NewReader([]byte(ui64z))
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if _, err := r.UInt64(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadString(b *testing.B) {
	r := NewReader([]byte(str))
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if _, err := r.String(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadStringWithEscapes(b *testing.B) {
	r := NewReader([]byte(stre))
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if _, err := r.String(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadStringSlice(b *testing.B) {
	r := NewReader([]byte(strs))
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if _, err := r.Strings(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_SkipString(b *testing.B) {
	r := NewReader([]byte(str)[1:])
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if err := r.skipString(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_SkipStringWithEscapes(b *testing.B) {
	r := NewReader([]byte(stre)[1:])
	for i := 0; i < b.N; i++ {
		r.SetPosition(0)
		if err := r.skipString(); err != nil {
			b.Fatal(err)
		}
	}
}
