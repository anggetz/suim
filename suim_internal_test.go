package suim

import "testing"

func TestSetOnly(t *testing.T) {
	f := struct {
		Name string
	}{
		Name: "",
	}

	SetIfStruct(&f, "Name", f.Name == "", "Berubah")
	if f.Name != "Berubah" {
		t.Errorf("tidak berubah")
	}

	f.Name = ""
	SetIfStruct(&f, "Name", f.Name != "", "Saya")
	if f.Name == "Saya" {
		t.Errorf("harusnya tidak berubah")
	}

	d := 20
	SetIf(&d, d == 20, 15)
	if d != 15 {
		t.Errorf("harusnya 15")
	}

	SetIf(&d, d != 15, 30)
	if d != 15 {
		t.Errorf("harusnya masih 15")
	}

	SetIf(&d, d == 15, 30)
	if d != 30 {
		t.Errorf("harusnya sudah 30")
	}
}
