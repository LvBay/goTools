package parseForm_test

import (
	"goTools/parseForm"
	"net/url"
	"testing"
)

func TestParse(t *testing.T) {
	// test1
	var info, want parseForm.Input
	var vs url.Values
	vs = url.Values{"code": []string{"10"}, "Name": []string{"zz"}}
	parseForm.ParseForm(vs, &info)
	want = parseForm.Input{Id: 10, Name: "zz"}
	if want != info {
		t.Errorf("err1")
	}

	// test2
	vs = url.Values{"code": []string{"10"}, "Name": []string{"zz"}, "Cname": []string{"cc"}}
	parseForm.ParseForm(vs, &info)
	want = parseForm.Input{Id: 10, Name: "zz", Cg: parseForm.Change{Cname: "cc"}}
	if want != info {
		t.Errorf("err2")
	}
}
