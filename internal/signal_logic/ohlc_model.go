package signal_logic

import "time"

type OHLCStruct struct {
	TimeOpen time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
}
