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
		bitMapBytes[fd/8] |= byte(1 << (8 - (uint(fd) % 8)))
	}
	for key, value := range bitMapBytes {
		fmt.Printf("byte %d: %08b\n", key,value)
	}

	//for key, value := range bitMapBytes {
	//	fmt.Println(key, fmt.Sprintf("%08b",value))
	//}
	return append(bitMapBytes, b...), nil
}

func (iso *ISO8583) unpack([]byte) error {
	return nil
}

func (iso *ISO8583) packBytes() ([]byte, error) {
	msg := make([]byte, 0)
	fields := make([]int, 0)
	for k := range iso.Fields {
		fields = append(fields, k)
	}
	sort.Ints(fields)
	iso.FieldsArray = fields[:]
	for _, fd := range fields {
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
		switch iso.LenTypeMap[fd] {
		case Type_Fix:
			if data != nil {
				return nil, fmt.Errorf("the sub field is not support Type_Fix")
			}
		case Type_VarL:
			l := byte(varLen)
			msg = append(msg, l)
		case Type_VarLL:
			ll := byte(((varLen / 10) << 4) | (varLen % 10))
			msg = append(msg, ll)
		case Type_VarLLL:
			lllHigh := byte(varLen / 100)
			lllLow := byte((((varLen % 100) / 10) << 4) | (varLen % 10))
			msg = append(msg, lllHigh, lllLow)
		default:
			return nil, fmt.Errorf("not support type by field: %d", fd)
		}
		if data == nil {
			v := iso.ValueMap[fd].(string)
			switch iso.AttrMap[fd] {
			case Attr_a, Attr_an, Attr_ans:
				msg = append(msg, []byte(v)...)
			case Attr_z:
				if (len(v) % 2) != 0 {
					v = v + "F"
				}
				b, _ := hex.DecodeString(v)
				msg = append(msg, b...)
			case Attr_n:
				if (len(v) % 2) != 0 {
					v = "0" + v
				}
				b, _ := hex.DecodeString(v)
				msg = append(msg, b...)
			case Attr_b:
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
