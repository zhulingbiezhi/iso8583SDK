package iso8583SDK

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSale(t *testing.T) {
	iso := CreateISO8583("")
	//iso.AddField(0, Attr_n, Type_Fix, "0100")
	iso.AddField(2, Attr{19, Len_VarLL, Format_n}, "5413330089020029")
	iso.AddField(3, Attr{6, Len_Fix, Format_n}, "000000")
	iso.AddField(4, Attr{12, Len_Fix, Format_n}, "000000002000")
	iso.AddField(11, Attr{6, Len_Fix, Format_n}, "000004")
	iso.AddField(14, Attr{4, Len_Fix, Format_n}, "2512")
	iso.AddField(22, Attr{3, Len_Fix, Format_n}, "012")
	iso.AddField(24, Attr{3, Len_Fix, Format_n}, "028")
	iso.AddField(25, Attr{2, Len_Fix, Format_n}, "00")
	iso.AddField(41, Attr{8, Len_Fix, Format_ans}, "63150002")
	iso.AddField(42, Attr{15, Len_Fix, Format_ans}, "549915204000099")
	iso.AddField(60, Attr{999, Len_VarLLL, Format_ans}, "000078")
	iso.AddField(62, Attr{999, Len_VarLLL, Format_ans}, "000004")
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

func TestDefaultSale(t *testing.T) {
	iso := CreateISO8583("")
	iso.AddFieldDefault(2, "5413330089020029")
	iso.AddFieldDefault(3, "000000")
	iso.AddFieldDefault(4, "000000002000")
	iso.AddFieldDefault(11, "000004")
	iso.AddFieldDefault(14, "2512")
	iso.AddFieldDefault(22, "012")
	iso.AddFieldDefault(24, "028")
	iso.AddFieldDefault(25, "00")
	iso.AddFieldDefault(41, "63150002")
	iso.AddFieldDefault(42, "549915204000099")
	iso.AddFieldDefault(60, "000078")
	iso.AddFieldDefault(62, "000004")
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
	fmt.Println(ret)
}

func TestBitMap(t *testing.T) {
	bitmap := []byte{
		0x70,
		0x24,
		0x05,
		0x80,
		0x00,
		0xC0,
		0x00,
		0x14,
	}
	for i, v := range bitmap {
		for j := 0; j < 8; j++ {
			if (v>>(7-uint(j)))&0x01 == 1 {
				fmt.Println(i*8 + j + 1)
			}
		}
	}
}

func TestUnpack(t *testing.T) {
	iso := CreateISO8583("")
	str := "7024058000C0001416541333008902002900000000000000200000000425120012002800363331353030303235343939313532303430303030393900063030303037380006303030303034"
	b, _ := hex.DecodeString(str)
	err := iso.unpack(b)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTransaction(t *testing.T) {
	iso := CreateISO8583("0200")
	iso.AddFieldDefault(2, "5413330089020029")
	iso.AddFieldDefault(3, "000000")
	iso.AddFieldDefault(4, "000000002000")
	iso.AddFieldDefault(11, "000004")
	iso.AddFieldDefault(14, "2512")
	iso.AddFieldDefault(22, "012")
	iso.AddFieldDefault(24, "028")
	iso.AddFieldDefault(25, "00")
	iso.AddFieldDefault(41, "63150002")
	iso.AddFieldDefault(42, "549915204000099")
	iso.AddFieldDefault(60, "000078")
	iso.AddFieldDefault(62, "000004")
	b, err := iso.pack()
	if err != nil {
		t.Error(err)
	}
	r := fmt.Sprintf("%X", b)
	exp := "7024058000C0001416541333008902002900000000000000200000000425120012002800363331353030303235343939313532303430303030393900063030303037380006303030303034"
	if r != exp {
		fmt.Println(exp)
		fmt.Println(r)
		t.Error("make bytes error !")
		return
	}
	fmt.Println(r)

	key, _ := hex.DecodeString(defaultConfig.TMK)
	encb := encrypt(b[8:], key)
	dstMsg := make([]byte, 0)
	dstMsg = append(dstMsg, 0x00, 0x00) // len
	tpdu, _ := hex.DecodeString(defaultConfig.TPDU)
	dstMsg = append(dstMsg, tpdu...)
	eds, _ := hex.DecodeString(defaultConfig.EDS)
	dstMsg = append(dstMsg, eds...)
	MTI, _ := hex.DecodeString(iso.MessageType)
	dstMsg = append(dstMsg, MTI...)
	dstMsg = append(dstMsg, b[:8]...)
	dstMsg = append(dstMsg, encb...)
	dstMsg[2+5+4] = byte(len(encb) >> 8)
	dstMsg[2+5+5] = byte(len(encb) & 0x00FF)

	dstMsg[0] = byte((len(dstMsg) - 2) >> 8)
	dstMsg[1] = byte((len(dstMsg) - 2) & 0x00FF)
	fmt.Printf("final msg:%x\n", dstMsg)
	ret, err := send(dstMsg, &defaultConfig)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%x", ret)
	err = iso.unpack(ret[7+2:])
	if err != nil {
		t.Error(err)
		return
	}
}

var defaultConfig = Config{}
