package talib

type PriceSource string

const (
	Open  PriceSource = "Open"
	High  PriceSource = "High"
	Low   PriceSource = "Low"
	Close PriceSource = "Close"
	HL2   PriceSource = "HL2"
	HLC3  PriceSource = "HLC3"
	OHLC4 PriceSource = "OHLC4"
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

type PivotType string

const (
	PivotTypeClassic   PivotType = "Classic"
	PivotTypeFibonacci PivotType = "Fibonacci"
	PivotTypeCamarilla PivotType = "Camarilla"
	PivotTypeWoodie    PivotType = "Woodie"
	PivotTypeDeMark    PivotType = "DeMark"
)
