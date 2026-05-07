package talib

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
