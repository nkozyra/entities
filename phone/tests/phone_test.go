package phonetest

import (
	"github.com/nkozyra/entities/phone"
	"testing"
)

var cases = []struct {
	in        string
	expected    string

}{
	{"+1 (555)  123     4567", "+15551234567"},
	{"(555)  123-4567", "15551234567"},
}

func Void(v interface{}) {

}

func RunTest(ph string) string {	

	rt := phone.New(ph)
	rt.Normalize()

	return rt.Normalized

}

func TestPhones(t *testing.T) {
	for _, test := range cases {
		observed := RunTest(test.in)
		if observed != test.expected {
			t.Fatalf("RunTest(%s) = %s, expected %s ", 	test.in, observed, test.expected)
		}
	}
}

func BenchmarkPhones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range cases {
			foo  := RunTest(test.in)
			Void(foo)
		}
	}
}