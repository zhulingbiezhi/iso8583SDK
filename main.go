package iso8583SDK

import "log"

type Attr int

const (
	Attr_a   Attr = 1
	Attr_an  Attr = 2
	Attr_ans Attr = 3
	Attr_s   Attr = 4
)

type LenType string

const (
	Type_Fix    LenType = "FIX"
	Type_VarL   LenType = "VARL"
	Type_VarLL  LenType = "VARLL"
	Type_VarLLL LenType = "VARLLL"
)

type ISO8583 struct {
	bInited    bool
	Fields     map[int]bool
	ValueMap   map[int]interface{}
	AttrMap    map[int]Attr
	LenTypeMap map[int]LenType
}

func CreateISO8583() *ISO8583 {
	return &ISO8583{
		ValueMap:   make(map[int]interface{}, 1),
		AttrMap:    make(map[int]Attr, 1),
		LenTypeMap: make(map[int]LenType, 1),
		bInited:    true,
	}
}

func (iso *ISO8583) AddField(field int, attr Attr, lenType LenType, value string) {
	if !iso.valid() {
		log.Println("error: the iso8583 is not init !")
		return
	}
	iso.Fields[field] = true
	iso.ValueMap[field] = value
	iso.LenTypeMap[field] = lenType
}

func (iso *ISO8583) AddSubField(parentField, field int, attr Attr, lenType LenType, value string) {
	if !iso.valid() {
		log.Println("error: the iso8583 is not init, please use createISO8583 first !")
		return
	}
	if iso.LenTypeMap[parentField] == "" {
		log.Println("error: the parent field is empty !")
		return
	}
	subIso := new(ISO8583)
	subIso.Fields[field] = true
	subIso.AttrMap[field] = attr
	subIso.ValueMap[field] = value
	iso.ValueMap[parentField] = subIso
}

func (iso *ISO8583) DeleteField(field int) {
	delete(iso.Fields, field)
	delete(iso.ValueMap, field)
	delete(iso.AttrMap, field)
}

func (iso *ISO8583) DeleteSubField(parentField, field int) {
	if iso.LenTypeMap[parentField] == "" {
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
	if iso.bInited {
		return false
	}
	return true
}

func (iso *ISO8583) ToJson() []byte {

}

func (iso *ISO8583) ToXml() []byte {

}

func ParseBytes(dataBytes, formatBytes []byte) *ISO8583 {

}

func LoadFromJson(dataBytes []byte) *ISO8583 {

}

func LoadFromXml(dataBytes []byte) *ISO8583 {

}

func Marshal(v interface{}) ([]byte, error) {

}

func Unmarshal(data []byte, v interface{}) error {

}
