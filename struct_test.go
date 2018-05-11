package iso8583SDK

import "testing"

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
	parseISO8583FromStruct(bea)
}
