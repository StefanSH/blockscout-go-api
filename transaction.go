package blockscout

func (b BlockScoutClient) GetTxInfo(hash string) (*Transaction, error) {
	url := b.Url + moduleValues["transaction"] +
		transactionActions["gettxinfo"] +
		"&txhash=" + hash

	var result Transaction

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (b BlockScoutClient) GetTxReceiptStatus(hash string) (*Status, error) {
	url := b.Url + moduleValues["transaction"] +
		transactionActions["gettxreceiptstatus"] +
		"&txhash=" + hash

	var result Status

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (b BlockScoutClient) GetTxStatus(hash string) (*TxStatus, error) {
	url := b.Url + moduleValues["transaction"] +
		transactionActions["gettxinfo"] +
		"&txhash=" + hash

	var result TxStatus

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
