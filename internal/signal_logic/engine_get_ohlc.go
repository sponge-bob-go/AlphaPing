package signal_logic

func (c *OHLCServces) GetCandle(coinSymbol string, interval string) ([]OHLCStruct, error) {
	ohlc, err := c.MarketsEngine.GetOHLC(coinSymbol, interval)
	if err != nil {
		return nil, err
	}
	return ohlc, nil
}
