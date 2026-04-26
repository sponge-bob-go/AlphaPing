package signal_logic

type OHLCServces struct {
	MarketsEngine CryptoMarketsEngine
}

func CreateOHLCServces(MarketsEngine CryptoMarketsEngine) *OHLCServces {
	return &OHLCServces{
		MarketsEngine: MarketsEngine,
	}
}
