package stubs

func CreatePaymentsOkResponse() []byte {
	return []byte(`
		{
			"error":"ok",
			"result":{
				"amount":"1.00000000",
				"address":"ZZZ",
				"dest_tag":"YYY",
				"txn_id":"XXX",
				"confirms_needed":"10",
				"timeout":9000,
				"checkout_url":"https:\/\/www.coinpayments.net\/index.php?cmd=checkout&id=XXX&key=ZZZ",
				"status_url":"https:\/\/www.coinpayments.net\/index.php?cmd=status&id=XXX&key=ZZZ",
				"qrcode_url":"https:\/\/www.coinpayments.net\/qrgen.php?id=XXX&key=ZZZ"
			}
}
`)
}
