package blockscout

func (b BlockScoutClient) GetToken(addr string) (*TokenInfo, error) {

	url := b.Url +
		moduleValues["token"] +
		tokenActions["getToken"] +
		"&contractaddress=" + addr

	var result TokenInfo

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
