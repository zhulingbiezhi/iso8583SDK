package iso8583SDK

import (
	"fmt"
	"log"
)

type FormatType int

const (
	Format_a   FormatType = 1
	Format_an  FormatType = 2
	Format_as  FormatType = 3
	Format_ns  FormatType = 4
	Format_ans FormatType = 5
	Format_z   FormatType = 6
	Format_n   FormatType = 7
	Format_b   FormatType = 8
	Format_s   FormatType = 9
)

type LenType string

const (
	Len_Fix    LenType = "FIX"
	Len_VarL   LenType = "VARL"
	Len_VarLL  LenType = "VARLL"
	Len_VarLLL LenType = "VARLLL"
	Len_TLV    LenType = "TLV"
)

type Attr struct {
	Len     int
	LenType LenType
	Format  FormatType
}

type ISO8583 struct {
	bInited     bool
	FieldsArray []int
	ValueMap    map[int]interface{}
	AttrMap     map[int]Attr
}

func CreateISO8583() *ISO8583 {
	return &ISO8583{
		ValueMap: make(map[int]interface{}),
		AttrMap:  make(map[int]Attr),
		bInited:  true,
	}
}

func (iso *ISO8583) AddField(field int, attr Attr, value string) {
	if !iso.valid() {
		log.Println("error: the iso8583 is not init !")
		return
	}
	if field < 2 || field > 63 {
		panic("the field must in 1 < x < 64")
	}
	iso.FieldsArray = append(iso.FieldsArray, field)
	iso.ValueMap[field] = value
	iso.AttrMap[field] = attr
}

func (iso *ISO8583) AddFieldDefault(field int, value string) {
	if !iso.valid() {
		log.Println("error: the iso8583 is not init !")
		return
	}
	if field < 2 || field > 63 {
		panic("the field must in 1 < x < 64")
	}
	if attr, ok := defaultAttrMap[field]; ok {
		iso.ValueMap[field] = value
		iso.AttrMap[field] = attr
		iso.FieldsArray = append(iso.FieldsArray, field)
	} else {
		panic(fmt.Sprintf("the default field %d not exist !", field))
	}
}

func (iso *ISO8583) AddSubField(parentField, field int, attr Attr, value string) {
	if !iso.valid() {
		log.Println("error: the iso8583 is not init, please use createISO8583 first !")
		return
	}
	if iso.AttrMap[parentField].Len == 0 {
		log.Println("error: the parent field is empty !")
		return
	}
	subIso := new(ISO8583)
	subIso.FieldsArray = append(subIso.FieldsArray, field)
	subIso.AttrMap[field] = attr
	iso.ValueMap[parentField] = subIso
}

func (iso *ISO8583) DeleteField(field int) {
	delete(iso.ValueMap, field)
	delete(iso.AttrMap, field)
}

func (iso *ISO8583) DeleteSubField(parentField, field int) {
	if iso.AttrMap[parentField].Len == 0 {
		log.Println("error: the parent field is empty !")
		return
	}
	if subISO, OK := iso.ValueMap[parentField].(ISO8583); OK {
		subISO.DeleteField(field)
	} else {
		log.Printf("the parentField: %d not exist subFiled: %d \n", parentField, field)
	}
}

func (iso *ISO8583) valid() bool {
	if !iso.bInited {
		return false
	}
	return true
}

func (iso *ISO8583) ToJson() []byte {
	return nil
}

func (iso *ISO8583) ToXml() []byte {
	return nil
}

func ParseBytes(dataBytes, formatBytes []byte) *ISO8583 {
	return nil
}

func LoadFromJson(dataBytes []byte) *ISO8583 {
	return nil
}

func LoadFromXml(dataBytes []byte) *ISO8583 {
	return nil
}

func Marshal(v interface{}) ([]byte, error) {

	return nil, nil
}

func Unmarshal(data []byte, v interface{}) error {
	return nil
}
