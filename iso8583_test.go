package iso8583SDK

import (
	"fmt"
	"testing"
)

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

func TestSale(t *testing.T) {
	iso := CreateISO8583()
	//iso.AddField(0, Attr_n, Type_Fix, "0100")
	iso.AddField(2, Attr_n, Type_VarLL, "5413330089020029")
	iso.AddField(3, Attr_n, Type_Fix, "000000")
	iso.AddField(4, Attr_n, Type_Fix, "000000002000")
	iso.AddField(11, Attr_n, Type_Fix, "000004")
	iso.AddField(14, Attr_n, Type_Fix, "2512")
	iso.AddField(22, Attr_n, Type_Fix, "012")
	iso.AddField(24, Attr_n, Type_Fix, "028")
	iso.AddField(25, Attr_n, Type_Fix, "00")
	iso.AddField(41, Attr_ans, Type_Fix, "63150002")
	iso.AddField(42, Attr_ans, Type_Fix, "549915204000099")
	iso.AddField(60, Attr_ans, Type_VarLLL, "000078")
	iso.AddField(62, Attr_ans, Type_VarLLL, "000004")
	b, err := iso.pack()
	if err != nil {
		t.Error(err)
	}
	ret := fmt.Sprintf("%X", b)
	exp := "7024058000C0001416541333008902002900000000000000200000000425120012002800363331353030303235343939313532303430303030393900063030303037380006303030303034"
	if ret != exp {
		fmt.Println(exp)
		fmt.Println(ret)
		t.Error("make bytes error !")
	}
}
