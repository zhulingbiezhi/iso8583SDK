package iso8583SDK

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	bea := BEA{
		TransId:     "000006",
		MessageType: "0020",
		Pan:         "1234567890123456",
		IccRelatedData: map[string]string{
			"4F": "123",
			"65": "145",
			"6E": "140KL",
		},
	}
	iso, err := parseISO8583FromStruct(bea)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v\n", iso)
}

func TestStructTransaction(t *testing.T) {

}
