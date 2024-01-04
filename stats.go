package blockscout

import (
	"math/big"
)

func (b BlockScoutClient) GetTokenTotalSupply(addr string) (*big.Int, error) {
	url := b.Url +
		moduleValues["stats"] +
		statsActions["tokensupply"] +
		"&contractaddress=" + addr

	var response Response
	err := sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	res, err := decStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b BlockScoutClient) GetTotalSupplyNativeCoin() (*big.Float, error) {
	url := b.Url +
		moduleValues["stats"] +
		statsActions["ethsupplyexchange"]

	var response Response
	err := sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	res, err := decStringToBigFloat(response.Result.(string))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetCoinPriceUSD() (*CoinPrice, error) {
	url := b.Url +
		moduleValues["stats"] +
		statsActions["coinprice"]

	var result CoinPrice

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
