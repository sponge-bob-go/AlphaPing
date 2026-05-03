package signal_logic

import (
	"errors"
	"strconv"
	"time"
)

type bybitOhlcResp struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category string     `json:"category"`
		Symbol   string     `json:"symbol"`
		List     [][]string `json:"list"`
	} `json:"result"`
	Time int64 `json:"time"`
}

func ParseOHLCFromRaw(list [][]string) ([]OHLCStruct, error) {

	candles := make([]OHLCStruct, 0, len(list))

	for _, c := range list {

		ts, err := strconv.ParseInt(c[0], 10, 64)
		if err != nil {
			continue
		}

		candles = append(candles, OHLCStruct{
			TimeOpen: time.UnixMilli(ts),
			Open:     parse(c[1]),
			High:     parse(c[2]),
			Low:      parse(c[3]),
			Close:    parse(c[4]),
			Volume:   parse(c[5]),
		})
	}

	if len(candles) == 0 {
		return nil, errors.New("no valid candles")
	}

	return candles, nil
}
func parse(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
