package talib

type PriceSource int

const (
	Open PriceSource = iota
	High
	Low
	Close
	HL2
	HLC3
	OHLC4
)

func GetPrice(open, high, low, close []float64, i int, src PriceSource) float64 {
	switch src {

	case Open:
		return open[i]

	case High:
		return high[i]

	case Low:
		return low[i]

	case Close:
		return close[i]

	case HL2:
		return (high[i] + low[i]) / 2.0

	case OHLC4:
		return (open[i] + high[i] + low[i] + close[i]) / 4.0

	case HLC3:
		fallthrough

	default:
		return (high[i] + low[i] + close[i]) / 3.0
	}
}
