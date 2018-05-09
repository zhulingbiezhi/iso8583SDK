package iso8583SDK

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
)

func (iso *ISO8583) pack() ([]byte, error) { // the bytes behind messageType  bytes
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
	//for key, value := range bitMapBytes {
	//	fmt.Printf("byte %d: %08b\n", key, value)
	//}
	MTI, _ := hex.DecodeString(iso.MessageType)
	b = append(bitMapBytes, b...)
	return append(MTI, b...), nil
}

func (iso *ISO8583) unpack(msg []byte) error { // the bytes behind messageType  bytes
	msgType := string(msg[:2])
	respISO := CreateISO8583(msgType)
	bitmap := msg[2:10]
	for i, v := range bitmap {
		for j := 0; j < 8; j++ {
			if (v>>(7-uint(j)))&0x01 == 1 {
				respISO.FieldsArray = append(respISO.FieldsArray, i*8+j+1)
			}
		}
	}

	dataMsg := msg[10:]
	offset := 0
	sort.Ints(respISO.FieldsArray)
	for _, fd := range respISO.FieldsArray {
		var attr Attr
		if a, ok := iso.AttrMap[fd]; ok {
			attr = a
		} else {
			attr = defaultAttrMap[fd]
		}

		var vlen int
		switch attr.LenType {
		case Len_Fix:
			vlen = attr.Len
		case Len_VarL, Len_VarLL:
			strL := fmt.Sprintf("%02x", dataMsg[offset])
			vlen, _ = strconv.Atoi(strL)
			offset++
		case Len_VarLLL:
			strL := fmt.Sprintf("%04x", dataMsg[offset:offset+2])
			vlen, _ = strconv.Atoi(strL)
			offset += 2
		default:
			fmt.Printf("unknow attr len type %s\n", attr.LenType)
		}

		switch attr.Format {
		case Format_a, Format_an, Format_ans:
			respISO.ValueMap[fd] = string(dataMsg[offset : offset+vlen])
			offset += vlen
		case Format_z, Format_n:
			l := (vlen + 1) / 2
			v := hex.EncodeToString(dataMsg[offset : offset+l])
			if vlen%2 != 0 {
				v = v[1:]
			}
			respISO.ValueMap[fd] = v
			offset += l
		case Format_b:
			l := vlen / 8
			yu := vlen % 8
			if yu != 0 {
				l += 1
			}
			respISO.ValueMap[fd] = hex.EncodeToString(dataMsg[offset : offset+l])
			offset += l
		default:
			fmt.Printf("not support attr by attr format: %d", attr.Format)
		}

		fmt.Println("response---key:", fd, "value:", respISO.ValueMap[fd])
	}
	return nil
}

func (iso *ISO8583) packBytes() ([]byte, error) {
	msg := make([]byte, 0)
	sort.Ints(iso.FieldsArray)
	for _, fd := range iso.FieldsArray {
		fmt.Println("request---key:", fd, "value:", iso.ValueMap[fd])
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
		default:
			return nil, fmt.Errorf("not support value type %s", t)
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
