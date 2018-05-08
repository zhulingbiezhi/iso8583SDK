package iso8583SDK

var defaultAttrMap = map[int]Attr{

	//Bitmap ---b 64
	1: {64, Len_Fix, Format_b},

	//Primary account number (PAN) ---n ..19
	2: {19, Len_VarLL, Format_n},

	//Processing code ---n 6
	3: {6, Len_Fix, Format_n},

	//Amount, transaction ---n 12
	4: {12, Len_Fix, Format_n},

	//Amount, settlement ---n 12
	5: {12, Len_Fix, Format_n},

	//Amount, cardholder billing ---n 12
	6: {12, Len_Fix, Format_n},

	//Transmission date & time ---n 10
	7: {10, Len_Fix, Format_n},

	//Amount, cardholder billing fee ---n 8
	8: {8, Len_Fix, Format_n},

	//Conversion rate, settlement ---n 8
	9: {8, Len_Fix, Format_n},

	//Conversion rate, cardholder billing ---n 8
	10: {8, Len_Fix, Format_n},

	//System trace audit number (STAN) ---n 6
	11: {6, Len_Fix, Format_n},

	//Local transaction time (hhmmss) ---n 6
	12: {6, Len_Fix, Format_n},

	//Local transaction date (MMDD) ---n 4
	13: {4, Len_Fix, Format_n},

	//Expiration date ---n 4
	14: {4, Len_Fix, Format_n},

	//Settlement date ---n 4
	15: {4, Len_Fix, Format_n},

	//Currency conversion date ---n 4
	16: {4, Len_Fix, Format_n},

	//Capture date ---n 4
	17: {4, Len_Fix, Format_n},

	//Merchant type, or merchant category code ---n 4
	18: {4, Len_Fix, Format_n},

	//Acquiring institution (country code) ---n 3
	19: {3, Len_Fix, Format_n},

	//PAN extended (country code) ---n 3
	20: {3, Len_Fix, Format_n},

	//	Forwarding institution (country code) ---n 3
	21: {3, Len_Fix, Format_n},

	//	Point of service entry mode ---n 3
	22: {3, Len_Fix, Format_n},

	//Application PAN sequence number ---n 3
	23: {3, Len_Fix, Format_n},

	//	Function code (ISO 8583:1993), or network international identifier (NII) ---n 3
	24: {3, Len_Fix, Format_n},

	//Point of service condition code ---n 2
	25: {2, Len_Fix, Format_n},

	//	Point of service capture code ---n 2
	26: {2, Len_Fix, Format_n},

	//Authorizing identification response length ---n 1
	27: {1, Len_Fix, Format_n},

	//	Amount, transaction fee ---n 8
	28: {8, Len_Fix, Format_n},

	//Amount, settlement fee ---n 8
	29: {8, Len_Fix, Format_n},

	//Amount, transaction processing fee ---n 8
	30: {8, Len_Fix, Format_n},

	//Amount, settlement processing fee ---n 8
	31: {8, Len_Fix, Format_n},

	//	Acquiring institution identification code ---n ..11
	32: {11, Len_VarLL, Format_n},

	//	Forwarding institution identification code ---n ..11
	33: {11, Len_VarLL, Format_n},

	//	Primary account number, extended ---ns ..28
	34: {28, Len_VarLL, Format_ns},

	//	Track 2 data ---z ..37
	35: {37, Len_VarLL, Format_z},

	//	Track 3 data ---n ...104
	36: {104, Len_VarLLL, Format_n},

	//Retrieval reference number ---an 12
	37: {12, Len_Fix, Format_an},

	//	Authorization identification response ---an 6
	38: {6, Len_Fix, Format_an},

	//Response code ---an 2
	39: {2, Len_Fix, Format_an},

	//	Service restriction code ---an 3
	40: {3, Len_Fix, Format_an},

	//Card acceptor terminal identification ---ans 8
	41: {8, Len_Fix, Format_ans},

	//Card acceptor identification code ---ans 15
	42: {15, Len_Fix, Format_ans},

	//Card acceptor name/location (1-23 street address, 24-36 city, 37-38 state, 39-40 country) ---ans 40
	43: {40, Len_Fix, Format_ans},

	//	Additional response data ---an ..25
	44: {25, Len_VarLL, Format_an},

	//	Track 1 data ---an ..76
	45: {76, Len_VarLL, Format_an},

	//Additional data (ISO) ---an ...999
	46: {999, Len_VarLLL, Format_an},

	//Additional data (national) ---an ...999
	47: {999, Len_VarLLL, Format_an},

	//Additional data (private) ---an ...999
	48: {999, Len_VarLLL, Format_an},

	//Currency code, transaction---n 3
	49: {3, Len_Fix, Format_n},

	//	Currency code, settlement ---n 3
	50: {3, Len_Fix, Format_n},

	//Currency code, cardholder billing ---n 3
	51: {3, Len_Fix, Format_n},

	//	Personal identification number data ---b 8
	52: {8, Len_Fix, Format_b},

	//Security related control information ---n 16
	53: {16, Len_Fix, Format_n},

	//Additional amounts ---an ...120
	54: {120, Len_VarLLL, Format_ans},

	//ICC data – EMV having multiple tags ---ans ...999
	55: {999, Len_VarLLL, Format_ans},

	//Reserved (ISO) ---ans ...999
	56: {999, Len_VarLLL, Format_ans},

	//Reserved (national) ---ans ...999
	57: {999, Len_VarLLL, Format_ans},

	//Reserved (national) ---ans ...999
	58: {999, Len_VarLLL, Format_ans},

	//Reserved (national) ---ans ...999
	59: {999, Len_VarLLL, Format_ans},

	//Reserved (national) ---ans ...999
	60: {999, Len_VarLLL, Format_ans},

	//Reserved (private) ---ans ...999
	61: {999, Len_VarLLL, Format_ans},

	//	Reserved (private)  ---ans ...999
	62: {999, Len_VarLLL, Format_ans},

	//	Reserved (private)  ---ans ...999
	63: {999, Len_VarLLL, Format_ans},

	//	Message authentication code (MAC) ---b 16
	64: {16, Len_Fix, Format_b},
}
