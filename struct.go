package iso8583SDK

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type BEA struct {
	MessageType         string
	TransId             string            `json:"trans_id" iso8583:"field:11;format:a11"`               //流水号--------11
	AcquireTransID      string            `json:"acq_trans_id"`                                         //收单行交易号---37
	TransDate           string            `json:"trans_date"`                                           //交易日期------13
	TransTime           string            `json:"trans_time"`                                           //交易时间------12
	Amount              string            `json:"amount"`                                               //授权金额------04
	TipAmount           string            `json:"tip"`                                                  //消费金额
	Pin                 string            `json:"pin"`                                                  //联机PINBLOCK--52
	Pan                 string            `json:"pan" iso8583:"field:2;format:a..19"`                   //主账号--------02
	NumberOfInstallment int               `json:"number_of_installment"`                                //分期
	PanSeqNo            string            `json:"pan_seq_no,omitempty"`                                 //卡片序列号-----
	CardExpireDate      string            `json:"card_exp_date"`                                        //有效期--------14
	Track1              string            `json:"track1,omitempty"`                                     //磁道一--------
	Track2              string            `json:"track2"`                                               //磁道二--------35
	PosEntryMode        int               `json:"pos_entry_mode"`                                       //刷卡方式------22
	IccRelatedData      map[string]string `json:"icc_related_data" iso8583:"field:55;format:ans...999"` //IC卡相关数据--
	AuthCode            string            `json:"auth_code"`                                            //授权码-------38
	ResponseCode        int               `json:"response_code"`                                        //响应码-------39
	Invoice             string            `json:"invoice,omitempty"`                                    //发票号-------62
	BatchNumber         string            `json:"batch_number,omitempty"`                               //settlement批次------60
	OriginalAmount      string            `json:"origin_amount,omitempty"`                              //原交易金额----
	OriginalTransType   int               `json:"origin_trans_type,omitempty"`                          //原交易金额----
	TransType           int               `json:"trans_type,omitempty"`                                 //交易类型------
}

func parseISO8583FromStruct(v interface{}) (*ISO8583, error) {
	iso := CreateISO8583("")
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	iso.MessageType = val.FieldByName("MessageType").String()

	for i := 0; i < typ.NumField(); i++ {
		fd := typ.Field(i)
		vv := val.Field(i)
		if !isZero(vv){
			fmt.Println(fd.Name)
		}
		tag := fd.Tag.Get("iso8583")
		if tag == "" {
			continue
		}
		ts := strings.Split(tag, ";")
		var attr Attr
		var field int
		for _, t := range ts {
			tt := strings.Split(t, ":")
			if len(tt) != 2 {
				return nil, fmt.Errorf("the tag error !")
			}
			switch tt[0] {
			case "field":
				d, _ := strconv.ParseInt(tt[1], 10, 64)
				field = int(d)
			case "format":
				attr = parseFormat(tt[1])
			default:
				return nil, fmt.Errorf("not support flag %s", tt[0])
			}
		}
		var iv string
		switch vv.Kind() {
		case reflect.Bool:
			tre, _ := strconv.ParseBool(vv.String())
			iv = "0"
			if tre {
				iv = "1"
			}
		case reflect.Map:
			for _, v := range vv.MapKeys() {
				iv = iv + v.String() + vv.MapIndex(v).String()
			}
		default:
			iv = vv.String()
		}

		if attr.valid() {
			iso.AddField(field, attr, iv)
		} else {
			return nil, fmt.Errorf("the field %s parse error", fd.Name)
		}
	}
	return iso, nil
}

func parseFormat(fm string) (attr Attr) {
	var ll int
	var strEnd, numStart int
	i := strings.Index(fm, ".")
	if i == -1 {
		attr.LenType = Len_Fix
		for i, v := range fm {
			if v > '0' && v < '9' {
				strEnd = i
				numStart = i
				break
			}
		}
	} else {
		e := strings.LastIndex(fm, ".")
		strEnd = i
		numStart = e + 1
		ll = numStart - strEnd
		switch ll {
		case 1:
			attr.LenType = Len_VarL
		case 2:
			attr.LenType = Len_VarLL
		case 3:
			attr.LenType = Len_VarLLL
		default:
			fmt.Println("%s not support the len %s", fm, strings.Repeat("L", ll))
		}
	}
	//fmt.Println(fm, strEnd, numStart)
	attr.Format = matchFmt(strings.TrimSpace(fm[:strEnd]))
	attr.Len, _ = strconv.Atoi(strings.TrimSpace(fm[numStart:]))
	return
}

func matchFmt(fm string) FormatType {
	switch fm {
	case "a":
		return Format_a
	case "an":
		return Format_an
	case "as":
		return Format_as
	case "ns":
		return Format_ns
	case "ans":
		return Format_ans
	case "z":
		return Format_z
	case "n":
		return Format_n
	case "b":
		return Format_b
	case "s":
		return Format_s
	default:
		fmt.Printf("unsupport format type %s\n", fm)
		return Format_unknow
	}
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}
