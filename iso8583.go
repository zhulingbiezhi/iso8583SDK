package iso8583SDK

import (
	"encoding/hex"
	"fmt"
	"sort"
)

func (iso *ISO8583) pack() ([]byte, error) {
	b, err := iso.packBytes()
	if err != nil {
		return nil, err
	}
	bitMapBytes := make([]byte, 8)
	for _, fd := range iso.FieldsArray {
		if fd%8 == 0 && fd != 0 {
			bitMapBytes[(fd/8)-1] |= byte(1)
		} else {
			bitMapBytes[fd/8] |= byte(1 << (8 - (uint(fd) % 8)))
		}
	}
	for key, value := range bitMapBytes {
		fmt.Printf("byte %d: %08b\n", key, value)
	}
	return append(bitMapBytes, b...), nil
}

func (iso *ISO8583) unpack([]byte) error {
	return nil
}

func (iso *ISO8583) packBytes() ([]byte, error) {
	msg := make([]byte, 0)
	sort.Ints(iso.FieldsArray)
	for _, fd := range iso.FieldsArray {
		fmt.Println("key:", fd, "value:", iso.ValueMap[fd])
		var data []byte
		var varLen int
		var err error

		switch t := iso.ValueMap[fd].(type) {
		case string:
			varLen = len(t)
		case *ISO8583: //处理subField
			data, err = t.packBytes()
			if err != nil {
				return nil, err
			}
		}
		if data != nil { //处理subField
			varLen = len(data) + 2 //subField 的tag占2个字节
		}
		switch iso.AttrMap[fd].LenType {
		case Len_Fix:
			if data != nil {
				return nil, fmt.Errorf("the sub field is not support Type_Fix")
			}
		case Len_VarL:
			l := byte(varLen)
			msg = append(msg, l)
		case Len_VarLL:
			ll := byte(((varLen / 10) << 4) | (varLen % 10))
			msg = append(msg, ll)
			fmt.Printf("%x\n", msg)
		case Len_VarLLL:
			lllHigh := byte(varLen / 100)
			lllLow := byte((((varLen % 100) / 10) << 4) | (varLen % 10))
			msg = append(msg, lllHigh, lllLow)
		default:
			return nil, fmt.Errorf("not support type by field: %d", fd)
		}
		if data == nil {
			v := iso.ValueMap[fd].(string)
			switch iso.AttrMap[fd].Format {
			case Format_a, Format_an, Format_ans:
				msg = append(msg, []byte(v)...)
			case Format_z:
				if (len(v) % 2) != 0 {
					v = v + "F"
				}
				b, _ := hex.DecodeString(v)
				msg = append(msg, b...)
			case Format_n:
				if (len(v) % 2) != 0 {
					v = "0" + v
				}
				b, _ := hex.DecodeString(v)
				msg = append(msg, b...)
			case Format_b:
				b, _ := hex.DecodeString(v)
				msg = append(msg, b...)
			default:
				return nil, fmt.Errorf("not support attr by field: %d attr: %s", fd, iso.AttrMap[fd])
			}
		} else { //处理subField
			b, _ := hex.DecodeString(fmt.Sprintf("%02d", fd))
			msg = append(msg, b...)
			msg = append(msg, data...)
		}
	}
	return msg, nil
}
