package signal_logic

type CryptoMarketsEngine interface {
	GetOHLC(coinSymbol string, interval string) ([]OHLCStruct, error)
}
