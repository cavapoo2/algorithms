package main

import "testing"

func TestDiag(t *testing.T) {
	data := board{1}
	res := noDiags(data)
	if res != true {
		t.Error("error (1) one value should be ok")
	}
	data = append(data, 0)
	res = noDiags(data)
	if res != false {
		t.Error("this is on diagonal not allowed")
	}
	data = data[:len(data)-1] //remove from end
	data = append(data, 3)
	res = noDiags(data)
	t.Logf("new data %v", data)
	if res != true {
		t.Error("error (2) one value should be ok")
	}
	data = append(data, 0)
	res = noDiags(data)
	t.Logf("new data %v", data)
	if res != true {
		t.Error("error (3) one value should be ok")
	}
	data = append(data, 2)
	res = noDiags(data)
	t.Logf("new data %v", data)
	if res != true {
		t.Error("error (3) one value should be ok")
	}

}
