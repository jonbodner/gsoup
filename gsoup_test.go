package main

import "testing"

var vals = []struct {
	input  string
	result string
}{
	{
		`<html><head><title>Hello</title></head><body>Hello There</body></html>`,
		"Hello There",
	},
	{
		`<html><head><title>Hello</title></head><body></body></html>`,
		"",
	},
	{
		`<html><head><title>Hello</title></head><body>
		</body></html>`,
		"",
	},
}

func TestText(t *testing.T) {
	testTextInner(t, Root.Text)
}

func TestTextBetter(t *testing.T) {
	testTextInner(t, Root.TextBetter)
}

func testTextInner(t *testing.T, f func(Root)string) {
	for _, v := range vals {
		r := HTMLParse(v.input)
		b := r.Find("body")
		txt  := f(b)
		if txt != v.result {
			t.Errorf("Expected %v, got %v",v.result, txt)
		}
	}
}

func BenchmarkText(b *testing.B) {
	benchTextInner(b, Root.Text)
}

func BenchmarkTextBetter(b *testing.B) {
	benchTextInner(b, Root.TextBetter)
}

func benchTextInner(b *testing.B, f func(Root)string) {
	for i := 0;i<b.N;i++ {
		for _, v := range vals {
			b.StopTimer()
			r := HTMLParse(v.input)
			bdy := r.Find("body")
			b.StartTimer()
			txt := f(bdy)
			if txt != v.result {
				b.Errorf("Expected %v, got %v", v.result, txt)
			}
		}
	}
}

