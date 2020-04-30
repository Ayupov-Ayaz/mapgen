package models

import "testing"

const (
	s     = 7
	count = 4
)

func BenchmarkGetP(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if k := Lp.GetP(s, count); k == 0 {
			b.Fatal("=(")
		}
	}
}

func BenchmarkMapPayout_Get(b *testing.B) {
	p := NewMapPayout(Lp)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if k := p.Get(s, count); k == 0 {
			b.Fatal("=(")
		}
	}
}

func BenchmarkGet(b *testing.B) {
	lp := NewSwitchPayout(LinePayout{})
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if k := lp.Get(s, count); k == 0 {
			b.Fatal("=(")
		}
	}
}

func BenchmarkLinePayout_Get(b *testing.B) {
	lp := LinePayout{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if k := lp.Get(s, count); k == 0 {
			b.Fatal("=(")
		}
	}
}
