package iso8583SDK

import "testing"

func TestCreateISO8583(t *testing.T) {
	iso := CreateISO8583()
	iso.AddField(1, Attr_n, Type_Fix, "123")
	iso.AddField(2, Attr_n, Type_Fix, "123")
	iso.AddField(3, Attr_n, Type_Fix, "123")
	iso.AddField(4, Attr_n, Type_Fix, "123")
	iso.AddField(60, Attr_n, Type_Fix, "123")
	iso.AddField(63, Attr_n, Type_Fix, "123")
	_, err := iso.pack()
	if err != nil {
		t.Error(err)
	}
}
