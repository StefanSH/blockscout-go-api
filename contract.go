package blockscout

func (b BlockScoutClient) GetABI(addr string) string {
	url := b.Url + moduleValues["contract"] +
		contractActions["getabi"] +
		"&address=" + addr

	var response Response
	sendApiRpcRequest(url, &response)

	return response.Result.(string)
}

func (b BlockScoutClient) GetContractInfo(addr string) (*[]ContractInfo, error) {
	url := b.Url + moduleValues["contract"] +
		contractActions["getsourcecode"] +
		"&address=" + addr

	var result []ContractInfo

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	//fmt.Println("Solc", result[0].CompilerVersion)

	return &result, nil
}
