package talib

import (
	"math"
)

// SuperTrend — ATR-based trend line.
//
// Returns (line, direction, breakout) with the same length as close.
// direction: +1 (uptrend) or -1 (downtrend) at every bar after warmup.
// breakout:  +1 (bullish flip) or -1 (bearish flip) only at trend change bars, 0 otherwise.
// indices i <= period use 0 for all outputs (warmup).
func SuperTrend(high, low, close []float64, period int, mult float64) ([]float64, []float64, []float64) {
	n := len(close)
	out := make([]float64, n)
	dir := make([]float64, n)
	breakout := make([]float64, n)

	if n == 0 || period < 1 || mult <= 0 {
		return out, dir, breakout
	}

	atr := Atr(high, low, close, period)

	finalUpper := make([]float64, n)
	finalLower := make([]float64, n)

	p := period
	hl2 := (high[p] + low[p]) * 0.5
	finalUpper[p] = hl2 + mult*atr[p]
	finalLower[p] = hl2 - mult*atr[p]

	// seed out[p] as internal reference only — not real output
	out[p] = finalLower[p]

	for i := p + 1; i < n; i++ {
		hl2 := (high[i] + low[i]) * 0.5
		upperBasic := hl2 + mult*atr[i]
		lowerBasic := hl2 - mult*atr[i]

		// Final upper band — only tightens downward
		if upperBasic < finalUpper[i-1] || close[i-1] > finalUpper[i-1] {
			finalUpper[i] = upperBasic
		} else {
			finalUpper[i] = finalUpper[i-1]
		}

		// Final lower band — only tightens upward
		if lowerBasic > finalLower[i-1] || close[i-1] < finalLower[i-1] {
			finalLower[i] = lowerBasic
		} else {
			finalLower[i] = finalLower[i-1]
		}

		// Direction: close > previous supertrend line = bullish
		if close[i] > out[i-1] {
			dir[i] = 1
			out[i] = finalLower[i]
		} else {
			dir[i] = -1
			out[i] = finalUpper[i]
		}

		// Breakout: only when trend flips
		if dir[i-1] != 0 && dir[i] != dir[i-1] {
			breakout[i] = dir[i]
		}
	}

	out[p] = 0

	return out, dir, breakout
}

// Vwap computes the Volume Weighted Average Price.
// A new session is indicated by passing a boolean slice `newSession`
// where newSession[i] == true means reset the VWAP at bar i.
// If newSession is nil, no reset is ever applied (continuous VWAP).
//
// Inputs : high, low, close []float64  — OHLC bars
//
//	volume            []float64  — bar volume
//	newSession        []bool     — true at the first bar of each session
//
// Outputs: vwap
// []float64  — VWAP line
func Vwap(open, high, low, close, volume []float64, src PriceSource) []float64 {
	n := len(close)

	if len(high) < n {
		n = len(high)
	}
	if len(low) < n {
		n = len(low)
	}
	if len(open) < n {
		n = len(open)
	}
	if len(volume) < n {
		n = len(volume)
	}

	out := make([]float64, n)

	var cumPV, cumVol float64

	for i := 0; i < n; i++ {

		price := GetPrice(open, high, low, close, i, src)
		vol := volume[i]

		if vol <= 0 {
			if i > 0 {
				out[i] = out[i-1]
			}
			continue
		}

		cumPV += price * vol
		cumVol += vol

		if cumVol != 0 {
			out[i] = cumPV / cumVol
		}
	}

	return out
}

// / OBV calculates On-Balance Volume.
// Returns cumulative OBV line.
func OBV(close, volume []float64) []float64 {
	n := len(close)

	if len(volume) < n {
		n = len(volume)
	}

	out := make([]float64, n)

	if n == 0 {
		return out
	}

	obv := volume[0]
	out[0] = obv

	for i := 1; i < n; i++ {

		if close[i] > close[i-1] {
			obv += volume[i]
		} else if close[i] < close[i-1] {
			obv -= volume[i]
		}

		out[i] = obv
	}

	return out
}

type PivotLevels struct {
	R1    float64
	R2    float64
	R3    float64
	R4    float64
	R5    float64
	S1    float64
	S2    float64
	S3    float64
	S4    float64
	S5    float64
	P     float64
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// Pivot computes pivot levels for the selected pivot type.
// Timeframe (daily/weekly/etc.) is decided by the OHLC inputs fed to this method.
func Pivot(open, high, low, close, currOpen float64, pivotType PivotType) PivotLevels {
	p := (high + low + close) / 3
	p2 := (high + low + 2*currOpen) / 4

	x := 0.0
	if open < close {
		x = 2*high + low + close
	} else if open > close {
		x = high + 2*low + close
	} else {
		x = high + low + 2*close
	}

	switch pivotType {
	case PivotTypeClassic:
		return PivotLevels{
			R1:    roundFloat((2*p)-low, 2),
			R2:    roundFloat(p+(high-low), 2),
			R3:    roundFloat(p+2*(high-low), 2),
			R4:    roundFloat(p+3*(high-low), 2),
			R5:    roundFloat(p+4*(high-low), 2),
			S1:    roundFloat((2*p)-high, 2),
			S2:    roundFloat(p-(high-low), 2),
			S3:    roundFloat(p-2*(high-low), 2),
			S4:    roundFloat(p-3*(high-low), 2),
			S5:    roundFloat(p-4*(high-low), 2),
			P:     roundFloat(p, 2),
		}
	case PivotTypeFibonacci:
		return PivotLevels{
			R1:    roundFloat(p+0.382*(high-low), 2),
			R2:    roundFloat(p+0.618*(high-low), 2),
			R3:    roundFloat(p+(high-low), 2),
			R4:    roundFloat(p+1.382*(high-low), 2),
			R5:    roundFloat(p+1.618*(high-low), 2),
			S1:    roundFloat(p-0.382*(high-low), 2),
			S2:    roundFloat(p-0.618*(high-low), 2),
			S3:    roundFloat(p-(high-low), 2),
			S4:    roundFloat(p-1.382*(high-low), 2),
			S5:    roundFloat(p-1.618*(high-low), 2),
			P:     roundFloat(p, 2),
		}
	case PivotTypeCamarilla:
		return PivotLevels{
			R1:    roundFloat(close+(1.1*(high-low))/12, 2),
			R2:    roundFloat(close+(1.1*(high-low))/6, 2),
			R3:    roundFloat(close+(1.1*(high-low))/4, 2),
			R4:    roundFloat(close+(1.1*(high-low))/2, 2),
			R5:    roundFloat(close+1.1*(high-low), 2),
			S1:    roundFloat(close-(1.1*(high-low))/12, 2),
			S2:    roundFloat(close-(1.1*(high-low))/6, 2),
			S3:    roundFloat(close-(1.1*(high-low))/4, 2),
			S4:    roundFloat(close-(1.1*(high-low))/2, 2),
			S5:    roundFloat(close-1.1*(high-low), 2),
			P:     roundFloat(close, 2),
		}
	case PivotTypeWoodie:
		return PivotLevels{
			R1:    roundFloat((2*p2)-low, 2),
			R2:    roundFloat(p2+(high-low), 2),
			R3:    roundFloat(high+2*(p2-low), 2),
			S1:    roundFloat((2*p2)-high, 2),
			S2:    roundFloat(p2-(high-low), 2),
			S3:    roundFloat(low-2*(high-p2), 2),
			P:     roundFloat(p2, 2),
		}
	case PivotTypeDeMark:
		return PivotLevels{
			P:     roundFloat(x/4, 2),
			R1:    roundFloat((x/2)-low, 2),
			S1:    roundFloat((x/2)-high, 2),	
		}
	default:
		return PivotLevels{}
	}
}

// PivotLevelSeries returns the selected pivot level for each bar.
// At index i >= 1, prior bar OHLC (i-1) and current open (i) feed scalar Pivot.
// out[0] is 0 (warmup). Output length matches the minimum OHLC slice length.
// Supports Classic, Fibonacci, Camarilla, Woodie, and DeMark via scalar Pivot.
func PivotLevelSeries(open, high, low, close []float64, pivotType PivotType, level PivotLevel) []float64 {
	n := len(close)
	if len(high) < n {
		n = len(high)
	}
	if len(low) < n {
		n = len(low)
	}
	if len(open) < n {
		n = len(open)
	}

	out := make([]float64, n)
	if n < 2 {
		return out
	}

	for i := 1; i < n; i++ {
		levels := Pivot(open[i-1], high[i-1], low[i-1], close[i-1], open[i], pivotType)
		out[i] = PivotLevelValue(levels, level)
	}

	return out
}
