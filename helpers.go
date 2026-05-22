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

// / Pivot related constants and helpers
type PivotType string

const (
	PivotTypeClassic   PivotType = "Classic"
	PivotTypeFibonacci PivotType = "Fibonacci"
	PivotTypeCamarilla PivotType = "Camarilla"
	PivotTypeWoodie    PivotType = "Woodie"
	PivotTypeDeMark    PivotType = "DeMark"
)

// PivotLevel selects which line to read from PivotLevels.
type PivotLevel string

const (
	PivotLevelP  PivotLevel = "P"
	PivotLevelR1 PivotLevel = "R1"
	PivotLevelR2 PivotLevel = "R2"
	PivotLevelR3 PivotLevel = "R3"
	PivotLevelR4 PivotLevel = "R4"
	PivotLevelR5 PivotLevel = "R5"
	PivotLevelS1 PivotLevel = "S1"
	PivotLevelS2 PivotLevel = "S2"
	PivotLevelS3 PivotLevel = "S3"
	PivotLevelS4 PivotLevel = "S4"
	PivotLevelS5 PivotLevel = "S5"
)

// PivotLevelValue returns the selected level from computed pivot levels.
// Unknown levels default to P.
func PivotLevelValue(levels PivotLevels, level PivotLevel) float64 {
	switch level {
	case PivotLevelP:
		return levels.P
	case PivotLevelR1:
		return levels.R1
	case PivotLevelS1:
		return levels.S1
	case PivotLevelR2:
		return levels.R2
	case PivotLevelS2:
		return levels.S2
	case PivotLevelR3:
		return levels.R3
	case PivotLevelS3:
		return levels.S3
	case PivotLevelR4:
		return levels.R4
	case PivotLevelS4:
		return levels.S4
	case PivotLevelR5:
		return levels.R5
	case PivotLevelS5:
		return levels.S5
	default:
		return levels.P
	}
}

// ParsePivotLevel maps a level string to PivotLevel. Empty or unknown values default to P.
func ParsePivotLevel(s string) PivotLevel {
	switch PivotLevel(s) {
	case PivotLevelP, PivotLevelR1, PivotLevelR2, PivotLevelR3, PivotLevelR4, PivotLevelR5,
		PivotLevelS1, PivotLevelS2, PivotLevelS3, PivotLevelS4, PivotLevelS5:
		return PivotLevel(s)
	default:
		return PivotLevelP
	}
}
