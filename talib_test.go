/*
Copyright 2016 Mark Chenoweth
Copyright 2018 Alessandro Sanino
Licensed under terms of MIT license (see LICENSE)
*/

package talib

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func ok(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n", filepath.Base(file), line, err.Error())
		t.FailNow()
	}
}

func equals(t *testing.T, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\tgo: %#v\n\tpy: %#v\n", filepath.Base(file), line, exp, act)
		t.FailNow()
	}
}

var (
	testOpen   = []float64{202.21, 200.04, 198.0, 197.35, 199.89, 202.23, 200.28, 199.99, 195.61, 197.55, 194.75, 198.31, 197.43, 199.87, 201.63, 200.57, 198.87, 200.04, 196.33, 196.51, 196.01, 198.9, 199.8, 200.72, 202.38, 200.63, 201.72, 202.43, 203.69, 204.84, 205.17, 205.42, 205.18, 205.24, 206.68, 206.85, 207.38, 207.24, 206.99, 206.52, 207.19, 206.15, 206.36, 205.19, 203.54, 202.53, 201.14, 201.11, 202.59, 202.53, 203.49, 203.2, 205.71, 206.39, 207.09, 206.52, 205.76, 201.71, 201.88, 203.7, 203.98, 203.12, 202.36, 202.12, 204.57, 204.26, 204.49, 205.89, 206.54, 205.54, 206.72, 206.7, 205.63, 205.75, 207.33, 206.68, 206.82, 208.31, 208.97, 207.4, 207.04, 206.55, 206.08, 207.88, 207.69, 206.24, 204.63, 207.54, 208.22, 206.29, 207.14, 207.89, 209.07, 208.88, 209.86, 209.77, 209.34, 209.66, 209.03, 207.9, 208.97, 209.01, 208.58, 207.68, 208.64, 207.73, 206.62, 206.32, 205.15, 206.05, 208.13, 207.3, 205.33, 205.62, 207.25, 207.96, 209.12, 209.57, 209.79, 209.38, 208.77, 207.96, 205.75, 204.97, 205.43, 205.77, 203.49, 204.67, 204.14, 204.75, 205.0, 206.68, 207.4, 208.4, 209.53, 209.94, 210.4, 210.08, 208.6, 209.19, 207.97, 204.65, 205.49, 207.16, 207.84, 209.08, 208.13, 207.38, 208.12, 207.96, 205.86, 206.97, 206.66, 204.82, 206.42, 206.13, 206.4, 207.94, 206.78, 204.23, 199.5, 185.42, 193.27, 189.96, 194.84, 196.31, 195.92, 190.98, 192.47, 194.09, 190.72, 193.77, 197.12, 192.41, 193.22, 194.77, 194.44, 196.62, 197.81, 194.55, 195.28, 192.73, 192.96, 191.01, 193.49, 190.65, 187.16, 189.24, 190.94, 188.65, 195.3, 197.14, 197.72, 197.77, 200.19, 200.23, 199.46, 199.0, 198.9, 201.63, 201.3, 201.65, 202.41, 201.78, 206.02, 206.07, 204.98, 205.78, 207.12, 207.82, 207.09, 208.73, 210.1, 209.19, 208.5, 208.07, 206.28, 207.64, 205.28, 203.14, 201.12, 204.77, 204.82, 207.36, 208.21, 208.14, 206.64, 208.26, 208.19, 208.51, 208.2, 209.37, 207.59, 204.39, 207.99, 205.27, 204.97, 204.2, 202.15, 200.87, 203.49, 205.15, 207.17, 202.77, 201.41, 202.72, 204.69, 205.72, 204.86, 206.51, 207.11, 205.13}
	testHigh   = []float64{202.7, 200.24, 198.62, 198.62, 201.99, 202.25, 200.46, 201.33, 197.03, 197.93, 197.74, 198.62, 199.54, 202.09, 201.93, 201.4, 199.99, 200.16, 198.21, 198.08, 197.95, 200.71, 201.23, 202.13, 203.05, 201.48, 202.93, 203.26, 204.76, 205.6, 206.07, 205.97, 206.17, 207.06, 206.94, 207.76, 207.95, 207.43, 207.3, 207.77, 207.76, 206.23, 206.54, 205.7, 204.57, 202.63, 201.35, 202.99, 203.73, 204.47, 204.21, 207.0, 206.21, 207.68, 207.77, 207.07, 206.03, 203.1, 202.69, 205.3, 204.8, 203.15, 203.7, 205.15, 205.45, 205.21, 205.87, 206.76, 207.29, 206.39, 207.7, 207.64, 205.91, 206.92, 207.52, 207.51, 208.58, 208.61, 209.11, 208.15, 207.94, 207.02, 207.43, 208.66, 208.11, 206.6, 206.06, 208.5, 208.53, 207.29, 207.87, 208.96, 209.24, 210.02, 210.19, 210.39, 210.36, 210.16, 209.54, 209.61, 209.22, 209.06, 208.98, 208.83, 209.3, 208.5, 207.24, 206.5, 205.79, 208.06, 208.73, 208.13, 206.13, 207.02, 207.97, 209.96, 209.21, 210.24, 210.09, 209.82, 208.91, 208.25, 207.51, 205.03, 205.73, 205.97, 205.35, 205.87, 204.47, 205.06, 205.68, 207.58, 208.72, 208.94, 209.95, 210.2, 210.82, 210.39, 209.43, 209.31, 208.04, 205.25, 207.18, 208.71, 208.69, 209.11, 208.2, 207.93, 208.97, 208.09, 206.04, 208.34, 207.15, 206.83, 207.23, 207.19, 208.26, 208.35, 207.69, 205.99, 201.68, 195.3, 193.29, 192.64, 197.21, 197.63, 196.93, 192.62, 193.3, 195.86, 191.72, 195.42, 197.26, 195.04, 194.64, 194.83, 196.79, 198.19, 200.65, 197.5, 196.51, 193.31, 193.52, 192.31, 193.85, 190.77, 188.62, 190.7, 191.35, 193.88, 197.56, 197.8, 198.65, 200.36, 200.71, 200.57, 200.96, 199.68, 201.16, 202.09, 202.17, 202.63, 202.58, 204.29, 206.72, 206.14, 205.78, 207.74, 208.03, 208.2, 209.37, 210.41, 210.25, 209.73, 209.08, 208.25, 207.37, 207.7, 205.83, 203.46, 204.47, 205.82, 207.66, 207.81, 208.88, 208.74, 208.59, 208.5, 208.56, 208.65, 209.57, 209.75, 207.91, 208.73, 208.49, 207.06, 207.45, 206.2, 202.93, 201.85, 204.89, 207.16, 207.25, 202.93, 201.88, 203.85, 206.07, 206.33, 205.26, 207.79, 207.21, 205.89}
	testLow    = []float64{200.05, 197.28, 194.84, 196.82, 199.87, 199.4, 197.84, 196.46, 194.56, 194.86, 194.54, 196.12, 196.88, 198.24, 200.67, 199.73, 197.66, 195.87, 194.66, 195.1, 193.86, 198.45, 199.4, 200.63, 200.78, 200.01, 200.54, 201.67, 202.79, 204.54, 204.87, 205.11, 205.01, 204.51, 206.22, 206.5, 206.95, 206.39, 206.34, 206.46, 205.83, 204.83, 205.61, 202.91, 203.35, 200.79, 200.27, 201.05, 200.44, 201.7, 202.8, 202.44, 204.8, 206.17, 206.67, 205.43, 202.45, 200.89, 201.65, 203.68, 203.09, 201.27, 202.15, 201.96, 203.96, 203.8, 203.91, 205.65, 205.72, 204.8, 206.62, 206.47, 203.73, 205.65, 205.92, 205.59, 206.68, 207.77, 207.2, 206.01, 206.28, 204.33, 205.96, 207.76, 205.42, 203.48, 204.23, 207.44, 207.18, 205.31, 206.42, 207.57, 208.5, 208.8, 209.32, 209.13, 209.14, 209.54, 206.87, 207.42, 208.28, 207.48, 207.28, 206.94, 207.98, 206.43, 205.67, 205.09, 204.4, 205.98, 207.85, 206.36, 204.5, 205.41, 206.04, 207.29, 208.03, 209.3, 209.23, 208.14, 207.45, 206.85, 203.06, 203.01, 204.28, 204.52, 203.26, 201.85, 201.99, 202.51, 202.68, 206.63, 207.33, 207.72, 209.24, 209.46, 209.86, 209.05, 208.56, 207.43, 205.3, 203.98, 204.51, 207.0, 207.1, 207.84, 206.34, 206.49, 207.41, 205.35, 204.58, 206.97, 205.46, 203.09, 205.71, 205.96, 205.86, 207.38, 205.06, 201.65, 195.34, 180.38, 184.85, 186.29, 193.05, 195.73, 194.83, 188.62, 190.29, 192.8, 189.49, 193.01, 192.2, 192.1, 192.38, 193.27, 193.79, 196.22, 197.08, 193.81, 194.06, 191.42, 191.77, 189.43, 190.68, 186.53, 185.82, 188.32, 188.7, 188.0, 195.17, 195.83, 196.31, 197.42, 199.39, 199.72, 198.87, 197.76, 198.46, 200.73, 200.93, 201.35, 200.46, 200.66, 205.08, 205.34, 204.57, 204.99, 206.98, 206.51, 206.94, 208.46, 208.48, 207.85, 207.23, 205.73, 205.96, 206.43, 203.61, 201.24, 200.98, 203.67, 204.77, 206.97, 207.62, 207.29, 206.18, 207.77, 207.62, 207.33, 207.87, 207.0, 203.54, 204.39, 205.97, 204.56, 202.97, 203.93, 200.32, 198.77, 201.67, 203.59, 203.63, 199.83, 200.09, 201.55, 204.58, 205.42, 203.94, 206.47, 205.76, 203.87}
	testClose  = []float64{201.28, 197.64, 195.78, 198.22, 201.74, 200.12, 198.55, 197.99, 196.8, 195.0, 197.55, 197.97, 198.97, 201.93, 200.83, 201.3, 198.64, 196.09, 197.91, 195.42, 197.84, 200.7, 199.93, 201.95, 201.39, 200.49, 202.63, 202.75, 204.7, 205.54, 205.86, 205.88, 205.73, 206.97, 206.94, 207.53, 207.35, 207.11, 206.4, 207.7, 206.85, 205.98, 206.2, 203.3, 204.15, 200.84, 200.37, 202.91, 201.67, 204.36, 203.76, 206.2, 205.26, 207.08, 206.67, 205.51, 202.5, 202.02, 202.48, 204.95, 203.16, 202.44, 203.17, 204.54, 204.0, 204.68, 205.59, 206.71, 205.78, 206.17, 207.1, 207.04, 204.66, 206.52, 206.28, 207.29, 207.81, 208.3, 207.43, 208.09, 207.23, 205.16, 207.38, 207.97, 205.59, 204.74, 205.56, 208.27, 207.27, 206.65, 206.69, 208.85, 209.07, 209.72, 209.65, 209.51, 210.12, 209.62, 207.36, 209.33, 209.09, 207.79, 208.22, 208.01, 208.56, 206.8, 206.45, 205.18, 205.15, 207.62, 208.28, 206.68, 205.79, 206.92, 207.25, 209.41, 208.48, 209.55, 209.71, 208.18, 207.54, 207.5, 203.15, 203.57, 205.21, 205.03, 204.43, 205.71, 202.27, 202.63, 205.19, 207.44, 208.35, 208.28, 209.95, 210.12, 210.24, 209.42, 209.03, 207.86, 205.7, 204.5, 207.02, 208.44, 208.49, 208.17, 207.47, 207.06, 207.75, 206.05, 205.65, 208.24, 206.35, 206.61, 206.35, 207.1, 208.26, 207.66, 206.02, 201.71, 195.64, 187.4, 185.2, 192.31, 197.07, 197.08, 195.48, 189.65, 193.25, 193.39, 190.46, 195.25, 192.64, 193.68, 194.56, 193.84, 196.26, 197.97, 197.52, 194.29, 195.3, 192.75, 192.45, 191.76, 191.73, 186.9, 187.01, 190.5, 190.99, 193.85, 197.3, 196.62, 198.23, 200.02, 200.14, 200.33, 199.07, 198.11, 201.15, 202.07, 202.17, 201.91, 200.66, 204.05, 206.28, 205.78, 205.38, 207.71, 207.59, 206.7, 209.15, 209.75, 209.12, 208.91, 208.8, 206.85, 207.33, 206.51, 203.63, 201.34, 204.4, 204.25, 207.5, 207.32, 208.07, 207.83, 208.11, 208.08, 208.32, 207.46, 209.43, 207.3, 204.39, 208.38, 207.12, 205.73, 204.13, 204.65, 200.69, 201.7, 203.82, 206.8, 203.65, 200.02, 201.67, 203.5, 206.02, 205.68, 205.21, 207.4, 205.93, 203.87}
	testVolume = []float64{121465900, 169632600, 209151400, 125346700, 147217800, 158567300, 144396100, 214553300, 192991100, 176613900, 211879600, 130991100, 122942700, 174356000, 117516800, 92009700, 134044600, 168514300, 173585400, 197729700, 163107000, 124212900, 134306700, 97953200, 125672000, 87219000, 96164200, 91087800, 97545900, 93670400, 76968200, 80652900, 91462500, 140896400, 74411100, 72472300, 73061700, 72697900, 108076000, 87491400, 110325800, 114497200, 76873000, 188128000, 89818900, 157121300, 110145700, 93993500, 162410900, 136099200, 94510400, 228808500, 117917300, 177715100, 71784500, 77805300, 159521700, 153067200, 118939000, 96180400, 126768700, 137303600, 86900900, 114368200, 81236300, 89351900, 85548900, 72722900, 74436600, 75099900, 99529300, 68934900, 191113200, 92189500, 72559800, 78264600, 102585900, 61327400, 79358100, 86863500, 125684900, 161304900, 103399700, 70927200, 113326200, 135060200, 88244900, 155877300, 75708100, 119727600, 94667900, 95934000, 76510100, 74549700, 72114600, 76857500, 64764600, 57433500, 124308600, 93214000, 74974600, 124919600, 93338800, 91531000, 87820900, 151882800, 121704700, 89063300, 105034700, 134551300, 73876400, 135382400, 124384200, 85308200, 126708600, 165867900, 130478700, 70696000, 68476800, 92307300, 97107400, 104174800, 202621300, 182925100, 135979900, 104373700, 117975400, 173820200, 164020100, 144113100, 129456900, 106069400, 81709600, 97914100, 106683300, 89030000, 70446800, 77965000, 88667900, 90509100, 117755000, 132361100, 123544800, 105791300, 91304400, 103266900, 113965700, 81820800, 85786800, 116030800, 117858000, 80270700, 126081400, 172123700, 89383300, 72786500, 79072600, 71692700, 172946000, 194327900, 346588500, 507244300, 369833100, 339257000, 274143900, 160414400, 163298800, 256000400, 160269300, 152087800, 207081000, 116025700, 149347700, 158611100, 119691200, 79452000, 113806200, 99581600, 276046600, 223657500, 105726200, 153890900, 92790600, 159378800, 155054800, 178515900, 159045600, 163452000, 131079000, 211003300, 126320800, 110274500, 124307300, 153055200, 107069200, 56395600, 88038700, 99106200, 134142200, 109692900, 76523900, 78448500, 102038000, 174911700, 144442300, 69033000, 77905800, 135906700, 90525500, 131076900, 86270800, 95246100, 96224500, 78408700, 110471500, 131008700, 75874600, 67846000, 121315200, 153577100, 117645200, 121123700, 121342500, 88220500, 94011500, 64931200, 98874400, 51980100, 37317800, 112822700, 97858400, 108441300, 166224200, 192913900, 102027100, 103372400, 162401500, 116128900, 211173300, 182385200, 154069600, 197017000, 173092500, 251393500, 99094300, 111026200, 110987200, 48542200, 65899900, 92640700, 63317700, 114877900}
	testRand   = []float64{0.42422904963267427, 0.16755615298728432, 0.5946077386900349, 0.17611040890583352, 0.29152918200482136, 0.27807733751955355, 0.7177400699036796, 0.5036012923358724, 0.1629504791237938, 0.6483065114032258, 0.5703588423748475, 0.7161845737507714, 0.6942714038794598, 0.42176699339445745, 0.7884431075157385, 0.24584359985404292, 0.7480158197252457, 0.2651217282085182, 0.4437589032368914, 0.9845738324910773, 0.5590040804528499, 0.25521017265864154, 0.1372114571360159, 0.1218701299153161, 0.25511876291008395, 0.7483943425884052, 0.076845841747889, 0.5389677976892574, 0.9015900382854415, 0.13503746751073498, 0.17237105554803778, 0.022111455150970016, 0.4735780024560894, 0.694458845807901, 0.5530772348613145, 0.3444350790493579, 0.6468662907768967, 0.6359557337589957, 0.5650572127602662, 0.621587087190788, 0.5634446451263618, 0.6967583014608363, 0.3366771423506647, 0.8920892600559512, 0.00029418556385873984, 0.1664001753124047, 0.2032534540019577, 0.30597531513267284, 0.4581883332445693, 0.4877258346021447}

	// testCrossunder1 = []float64{1, 2, 3, 4, 8, 6, 7}
	// testCrossunder2 = []float64{1, 1, 10, 9, 5, 3, 7}

	// testNothingCrossed1 = []float64{1, 2, 3, 4, 8, 6, 7}
	// testNothingCrossed2 = []float64{1, 4, 5, 9, 5, 3, 7}

	// testCrossover1 = []float64{1, 3, 2, 4, 8, 6, 7}
	// testCrossover2 = []float64{1, 5, 1, 4, 5, 6, 7}
)

func a2s(a []float64) string { // go float64 array to python list initializer string
	return strings.Replace(fmt.Sprintf("%f", a), " ", ",", -1)
}

func round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}

func compare(t *testing.T, goResult []float64, taCall string) {
	pyprog := fmt.Sprintf(`import talib,numpy
testOpen = numpy.array(%s)
testHigh = numpy.array(%s)
testLow = numpy.array(%s)
testClose = numpy.array(%s)
testVolume = numpy.array(%s)
testRand = numpy.array(%s)
%s
print(' '.join([str(p) for p in result]).replace('nan','0.0'))`,
		a2s(testOpen), a2s(testHigh), a2s(testLow), a2s(testClose), a2s(testVolume), a2s(testRand), taCall)

	//fmt.Println(pyprog)
	pyOut, err := exec.Command("python", "-c", pyprog).Output()
	ok(t, err)

	var pyResult []float64
	strResult := strings.Fields(string(pyOut))
	for _, arg := range strResult {
		if n, err := strconv.ParseFloat(arg, 64); err == nil {
			pyResult = append(pyResult, n)
		}
	}

	equals(t, len(goResult), len(pyResult))

	for i := 0; i < len(goResult); i++ {

		if (goResult[i] < -0.00000000000001) || (goResult[i] < 0.00000000000001) {
			goResult[i] = 0.0
		}
		if (pyResult[i] < -0.00000000000001) || (pyResult[i] < 0.00000000000001) {
			pyResult[i] = 0.0
		}

		var s1, s2 string
		if (goResult[i] > -1000000) && (goResult[i] < 1000000) {
			s1 = fmt.Sprintf("%.6f", goResult[i])
		} else {
			s1 = fmt.Sprintf("%.1f", round(goResult[i])) // reduce precision for very large numbers
		}

		if (pyResult[i] > -1000000) && (pyResult[i] < 1000000) {
			s2 = fmt.Sprintf("%.6f", pyResult[i])
		} else {
			s2 = fmt.Sprintf("%.1f", round(pyResult[i])) // reduce precision for very large numbers
		}
		//equals(t, s1, s2)
		if s1[:len(s1)-2] != s2[:len(s2)-2] {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf("%s:%d:\n\tgo!: %#v\n\tpy!: %#v\n", filepath.Base(file), line, s1, s2)
			t.FailNow()
		}

	}
}

// Ensure that python and talib are installed and in the PATH
func TestMain(m *testing.M) {
	pyout, _ := exec.Command("python", "-c", "import talib; print('success')").Output()
	if string(pyout[0:7]) != "success" {
		fmt.Println("python and talib must be installed to run tests")
		os.Exit(-1)
	}
	os.Exit(m.Run())
}

// Test all the functions

func TestSma(t *testing.T) {
	result := Sma(testClose, 20)
	compare(t, result, "result = talib.SMA(testClose,20)")
}

func TestEma(t *testing.T) {
	result := Ema(testClose, 5)
	compare(t, result, "result = talib.EMA(testClose,5)")
	result = Ema(testClose, 20)
	compare(t, result, "result = talib.EMA(testClose,20)")
	result = Ema(testClose, 50)
	compare(t, result, "result = talib.EMA(testClose,50)")
	result = Ema(testClose, 100)
	compare(t, result, "result = talib.EMA(testClose,100)")
}

func TestRsi(t *testing.T) {
	result := Rsi(testClose, 10)
	compare(t, result, "result = talib.RSI(testClose,10)")
}

func TestAdd(t *testing.T) {
	result := Add(testHigh, testLow)
	compare(t, result, "result = talib.ADD(testHigh,testLow)")
}

func TestDiv(t *testing.T) {
	result := Div(testHigh, testLow)
	compare(t, result, "result = talib.DIV(testHigh,testLow)")
}

func TestMax(t *testing.T) {
	result := Max(testClose, 10)
	compare(t, result, "result = talib.MAX(testClose,10)")
}

func TestMaxIndex(t *testing.T) {
	result := MaxIndex(testClose, 10)
	compare(t, result, "result = talib.MAXINDEX(testClose,10)")
}

func TestMin(t *testing.T) {
	result := Min(testClose, 10)
	compare(t, result, "result = talib.MIN(testClose,10)")
}

func TestMinIndex(t *testing.T) {
	result := MinIndex(testClose, 10)
	compare(t, result, "result = talib.MININDEX(testClose,10)")
}

func TestMult(t *testing.T) {
	result := Mult(testHigh, testLow)
	compare(t, result, "result = talib.MULT(testHigh,testLow)")
}

func TestSub(t *testing.T) {
	result := Sub(testHigh, testLow)
	compare(t, result, "result = talib.SUB(testHigh,testLow)")
}

func TestRocp(t *testing.T) {
	result := Rocp(testClose, 10)
	compare(t, result, "result = talib.ROCP(testClose,10)")
}

func TestObv(t *testing.T) {
	result := Obv(testClose, testVolume)
	compare(t, result, "result = talib.OBV(testClose,testVolume)")
}

func TestAtr(t *testing.T) {
	result := Atr(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.ATR(testHigh,testLow,testClose,14)")
}

func TestNatr(t *testing.T) {
	result := Natr(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.NATR(testHigh,testLow,testClose,14)")
}

func TestTRange(t *testing.T) {
	result := TRange(testHigh, testLow, testClose)
	compare(t, result, "result = talib.TRANGE(testHigh,testLow,testClose)")
}

func TestAvgPrice(t *testing.T) {
	result := AvgPrice(testOpen, testHigh, testLow, testClose)
	compare(t, result, "result = talib.AVGPRICE(testOpen,testHigh,testLow,testClose)")
}

func TestMedPrice(t *testing.T) {
	result := MedPrice(testHigh, testLow)
	compare(t, result, "result = talib.MEDPRICE(testHigh,testLow)")
}

func TestTypPrice(t *testing.T) {
	result := TypPrice(testHigh, testLow, testClose)
	compare(t, result, "result = talib.TYPPRICE(testHigh,testLow,testClose)")
}

func TestWclPrice(t *testing.T) {
	result := WclPrice(testHigh, testLow, testClose)
	compare(t, result, "result = talib.WCLPRICE(testHigh,testLow,testClose)")
}

func TestAcos(t *testing.T) {
	result := Acos(testRand)
	compare(t, result, "result = talib.ACOS(testRand)")
}

func TestAsin(t *testing.T) {
	result := Asin(testRand)
	compare(t, result, "result = talib.ASIN(testRand)")
}

func TestAtan(t *testing.T) {
	result := Atan(testRand)
	compare(t, result, "result = talib.ATAN(testRand)")
}

func TestCeil(t *testing.T) {
	result := Ceil(testClose)
	compare(t, result, "result = talib.CEIL(testClose)")
}

func TestCos(t *testing.T) {
	result := Cos(testRand)
	compare(t, result, "result = talib.COS(testRand)")
}

func TestCosh(t *testing.T) {
	result := Cosh(testRand)
	compare(t, result, "result = talib.COSH(testRand)")
}

func TestExp(t *testing.T) {
	result := Exp(testRand)
	compare(t, result, "result = talib.EXP(testRand)")
}

func TestFloor(t *testing.T) {
	result := Floor(testClose)
	compare(t, result, "result = talib.FLOOR(testClose)")
}

func TestLn(t *testing.T) {
	result := Ln(testClose)
	compare(t, result, "result = talib.LN(testClose)")
}

func TestLog10(t *testing.T) {
	result := Log10(testClose)
	compare(t, result, "result = talib.LOG10(testClose)")
}

func TestSin(t *testing.T) {
	result := Sin(testRand)
	compare(t, result, "result = talib.SIN(testRand)")
}

func TestSinh(t *testing.T) {
	result := Sinh(testRand)
	compare(t, result, "result = talib.SINH(testRand)")
}

func TestSqrt(t *testing.T) {
	result := Sqrt(testClose)
	compare(t, result, "result = talib.SQRT(testClose)")
}

func TestTan(t *testing.T) {
	result := Tan(testRand)
	compare(t, result, "result = talib.TAN(testRand)")
}

func TestTanh(t *testing.T) {
	result := Tanh(testRand)
	compare(t, result, "result = talib.TANH(testRand)")
}

func TestSum(t *testing.T) {
	result := Sum(testClose, 10)
	compare(t, result, "result = talib.SUM(testClose,10)")
}

func TestVar(t *testing.T) {
	result := Var(testClose, 10)
	compare(t, result, "result = talib.VAR(testClose,10)")
}

func TestTsf(t *testing.T) {
	result := Tsf(testClose, 10)
	compare(t, result, "result = talib.TSF(testClose,10)")
}

func TestStdDev(t *testing.T) {
	result := StdDev(testRand, 10, 1.0)
	compare(t, result, "result = talib.STDDEV(testRand,10,1.0)")
}

func TestLinearRegSlope(t *testing.T) {
	result := LinearRegSlope(testClose, 10)
	compare(t, result, "result = talib.LINEARREG_SLOPE(testClose,10)")
}

func TestLinearRegIntercept(t *testing.T) {
	result := LinearRegIntercept(testClose, 10)
	compare(t, result, "result = talib.LINEARREG_INTERCEPT(testClose,10)")
}

func TestLinearRegAngle(t *testing.T) {
	result := LinearRegAngle(testClose, 10)
	compare(t, result, "result = talib.LINEARREG_ANGLE(testClose,10)")
}

func TestLinearReg(t *testing.T) {
	result := LinearReg(testClose, 10)
	compare(t, result, "result = talib.LINEARREG(testClose,10)")
}

func TestCorrel(t *testing.T) {
	result := Correl(testHigh, testLow, 10)
	compare(t, result, "result = talib.CORREL(testHigh,testLow,10)")
}

func TestBeta(t *testing.T) {
	result := Beta(testHigh, testLow, 5)
	compare(t, result, "result = talib.BETA(testHigh,testLow,5)")
}

func TestHtDcPeriod(t *testing.T) {
	result := HtDcPeriod(testClose)
	compare(t, result, "result = talib.HT_DCPERIOD(testClose)")
}

func TestHtPhasor(t *testing.T) {
	result1, result2 := HtPhasor(testClose)
	compare(t, result1, "result,_ = talib.HT_PHASOR(testClose)")
	compare(t, result2, "_,result = talib.HT_PHASOR(testClose)")
}

func TestHtSine(t *testing.T) {
	result1, result2 := HtSine(testClose)
	compare(t, result1, "result,_ = talib.HT_SINE(testClose)")
	compare(t, result2, "_,result = talib.HT_SINE(testClose)")
}

func TestHtTrendline(t *testing.T) {
	result := HtTrendline(testClose)
	compare(t, result, "result = talib.HT_TRENDLINE(testClose)")
}

func TestHtTrendMode(t *testing.T) {
	result := HtTrendMode(testClose)
	compare(t, result, "result = talib.HT_TRENDMODE(testClose)")
}

func TestWillR(t *testing.T) {
	result := WillR(testHigh, testLow, testClose, 9)
	compare(t, result, "result = talib.WILLR(testHigh,testLow,testClose,9)")
}

func TestAdx(t *testing.T) {
	result := Adx(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.ADX(testHigh,testLow,testClose,14)")
}

func TestAdxR(t *testing.T) {
	result := AdxR(testHigh, testLow, testClose, 5)
	compare(t, result, "result = talib.ADXR(testHigh,testLow,testClose,5)")
}

func TestCci(t *testing.T) {
	result := Cci(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.CCI(testHigh,testLow,testClose,14)")
}

func TestRoc(t *testing.T) {
	result := Roc(testClose, 10)
	compare(t, result, "result = talib.ROC(testClose,10)")
}

func TestRocr(t *testing.T) {
	result := Rocr(testClose, 10)
	compare(t, result, "result = talib.ROCR(testClose,10)")
}

func TestRocr100(t *testing.T) {
	result := Rocr100(testClose, 10)
	compare(t, result, "result = talib.ROCR100(testClose,10)")
}

func TestMom(t *testing.T) {
	result := Mom(testClose, 10)
	compare(t, result, "result = talib.MOM(testClose,10)")
}

func TestBBands(t *testing.T) {
	upper, middle, lower := BBands(testClose, 5, 2.0, 2.0, SMA)
	compare(t, upper, "result,upper,lower = talib.BBANDS(testClose,5,2.0,2.0)")
	compare(t, middle, "upper,result,lower = talib.BBANDS(testClose,5,2.0,2.0)")
	compare(t, lower, "upper,middle,result = talib.BBANDS(testClose,5,2.0,2.0)")
}

func TestDema(t *testing.T) {
	result := Dema(testClose, 10)
	compare(t, result, "result = talib.DEMA(testClose,10)")
}

func TestTema(t *testing.T) {
	result := Tema(testClose, 10)
	compare(t, result, "result = talib.TEMA(testClose,10)")
}

func TestWma(t *testing.T) {
	result := Wma(testClose, 10)
	compare(t, result, "result = talib.WMA(testClose,10)")
}

func TestMa(t *testing.T) {
	result := Ma(testClose, 10, DEMA)
	compare(t, result, "result = talib.MA(testClose,10,talib.MA_Type.DEMA)")
}

func TestTrima(t *testing.T) {
	result := Trima(testClose, 10)
	compare(t, result, "result = talib.TRIMA(testClose,10)")
	result = Trima(testClose, 11)
	compare(t, result, "result = talib.TRIMA(testClose,11)")
}

func TestMidPoint(t *testing.T) {
	result := MidPoint(testClose, 10)
	compare(t, result, "result = talib.MIDPOINT(testClose,10)")
}

func TestMidPrice(t *testing.T) {
	result := MidPrice(testHigh, testLow, 10)
	compare(t, result, "result = talib.MIDPRICE(testHigh,testLow,10)")
}

func TestT3(t *testing.T) {
	result := T3(testClose, 5, 0.7)
	compare(t, result, "result = talib.T3(testClose,5,0.7)")
}

func TestKama(t *testing.T) {
	result := Kama(testClose, 10)
	compare(t, result, "result = talib.KAMA(testClose,10)")
}

func TestMaVp(t *testing.T) {
	periods := make([]float64, len(testClose))
	for i := range testClose {
		periods[i] = 5.0
	}
	result := MaVp(testClose, periods, 2, 10, SMA)
	compare(t, result, "result = talib.MAVP(testClose,numpy.full(len(testClose),5.0),2,10,talib.MA_Type.SMA)")
}

func TestMinusDM(t *testing.T) {
	result := MinusDM(testHigh, testLow, 14)
	compare(t, result, "result = talib.MINUS_DM(testHigh,testLow,14)")
}

func TestPlusDM(t *testing.T) {
	result := PlusDM(testHigh, testLow, 14)
	compare(t, result, "result = talib.PLUS_DM(testHigh,testLow,14)")
}

func TestSar(t *testing.T) {
	result := Sar(testHigh, testLow, 0.0, 0.0)
	compare(t, result, "result = talib.SAR(testHigh,testLow,0.0,0.0)")
}

func TestSarExt(t *testing.T) {
	result := SarExt(testHigh, testLow, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	compare(t, result, "result = talib.SAREXT(testHigh,testLow,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0)")
}

func TestMama(t *testing.T) {
	mama, fama := Mama(testClose, 0.5, 0.05)
	compare(t, mama, "result,fama = talib.MAMA(testClose,0.5,0.05)")
	compare(t, fama, "mama,result = talib.MAMA(testClose,0.5,0.05)")
}

func TestMinMax(t *testing.T) {
	min, max := MinMax(testClose, 10)
	compare(t, min, "result,max = talib.MINMAX(testClose,10)")
	compare(t, max, "min,result = talib.MINMAX(testClose,10)")
}

func TestMinMaxIndex(t *testing.T) {
	minidx, maxidx := MinMaxIndex(testClose, 10)
	compare(t, minidx, "result,maxidx = talib.MINMAXINDEX(testClose,10)")
	compare(t, maxidx, "minidx,result = talib.MINMAXINDEX(testClose,10)")
}

func TestApo(t *testing.T) {
	result := Apo(testClose, 12, 26, SMA)
	compare(t, result, "result = talib.APO(testClose,12,26,talib.MA_Type.SMA)")
	result = Apo(testClose, 26, 12, SMA)
	compare(t, result, "result = talib.APO(testClose,26,12,talib.MA_Type.SMA)")
}

func TestPpo(t *testing.T) {
	result := Ppo(testClose, 12, 26, SMA)
	compare(t, result, "result = talib.PPO(testClose,12,26,talib.MA_Type.SMA)")
	result = Ppo(testClose, 26, 12, SMA)
	compare(t, result, "result = talib.PPO(testClose,26,12,talib.MA_Type.SMA)")
}

func TestAroon(t *testing.T) {
	dn, up := Aroon(testHigh, testLow, 14)
	compare(t, dn, "result,up = talib.AROON(testHigh,testLow,14)")
	compare(t, up, "dn,result = talib.AROON(testHigh,testLow,14)")
}

func TestAroonOsc(t *testing.T) {
	result := AroonOsc(testHigh, testLow, 14)
	compare(t, result, "result = talib.AROONOSC(testHigh,testLow,14)")
}

func TestBop(t *testing.T) {
	result := Bop(testOpen, testHigh, testLow, testClose)
	compare(t, result, "result = talib.BOP(testOpen,testHigh,testLow,testClose)")
}

func TestCmo(t *testing.T) {
	result := Cmo(testClose, 14)
	compare(t, result, "result = talib.CMO(testClose,14)")
}

func TestDx(t *testing.T) {
	result := Dx(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.DX(testHigh,testLow,testClose,14)")
}

func TestMinusDI(t *testing.T) {
	result := MinusDI(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.MINUS_DI(testHigh,testLow,testClose,14)")
}

func TestPlusDI(t *testing.T) {
	result := PlusDI(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.PLUS_DI(testHigh,testLow,testClose,14)")
}

func TestMfi(t *testing.T) {
	result := Mfi(testHigh, testLow, testClose, testVolume, 14)
	compare(t, result, "result = talib.MFI(testHigh,testLow,testClose,testVolume,14)")
}

func TestUltOsc(t *testing.T) {
	result := UltOsc(testHigh, testLow, testClose, 7, 14, 28)
	compare(t, result, "result = talib.ULTOSC(testHigh,testLow,testClose,7,14,28)")
}

func TestStoch(t *testing.T) {
	slowk, slowd := Stoch(testHigh, testLow, testClose, 5, 3, SMA, 3, SMA)
	compare(t, slowk, "result,slowd = talib.STOCH(testHigh,testLow,testClose,5,3,talib.MA_Type.SMA,3,talib.MA_Type.SMA)")
	compare(t, slowd, "slowk,result = talib.STOCH(testHigh,testLow,testClose,5,3,talib.MA_Type.SMA,3,talib.MA_Type.SMA)")
}

func TestStoch2(t *testing.T) {
	slowk, slowd := Stoch(testHigh, testLow, testClose, 12, 3, SMA, 3, SMA)
	compare(t, slowk, "result,slowd = talib.STOCH(testHigh,testLow,testClose,12,3,talib.MA_Type.SMA,3,talib.MA_Type.SMA)")
	compare(t, slowd, "slowk,result = talib.STOCH(testHigh,testLow,testClose,12,3,talib.MA_Type.SMA,3,talib.MA_Type.SMA)")
}

func TestStoch3(t *testing.T) {
	slowk, slowd := Stoch(testHigh, testLow, testClose, 12, 3, SMA, 15, SMA)
	compare(t, slowk, "result,slowd = talib.STOCH(testHigh,testLow,testClose,12,3,talib.MA_Type.SMA,15,talib.MA_Type.SMA)")
	compare(t, slowd, "slowk,result = talib.STOCH(testHigh,testLow,testClose,12,3,talib.MA_Type.SMA,15,talib.MA_Type.SMA)")
}

func TestStochF(t *testing.T) {
	fastk, fastd := StochF(testHigh, testLow, testClose, 5, 3, SMA)
	compare(t, fastk, "result,fastd = talib.STOCHF(testHigh,testLow,testClose,5,3,talib.MA_Type.SMA)")
	compare(t, fastd, "fastk,result = talib.STOCHF(testHigh,testLow,testClose,5,3,talib.MA_Type.SMA)")
}

func TestStochRsi(t *testing.T) {
	fastk, fastd := StochRsi(testClose, 14, 5, 2, SMA)
	compare(t, fastk, "result,fastd = talib.STOCHRSI(testClose,14,5,2,talib.MA_Type.SMA)")
	compare(t, fastd, "fastk,result = talib.STOCHRSI(testClose,14,5,2,talib.MA_Type.SMA)")
}

func TestMacdExt(t *testing.T) {
	macd, macdsignal, macdhist := MacdExt(testClose, 12, SMA, 26, SMA, 9, SMA)
	compare(t, macd, "result, macdsignal, macdhist = talib.MACDEXT(testClose,12,talib.MA_Type.SMA,26,talib.MA_Type.SMA,9,talib.MA_Type.SMA)")
	compare(t, macdsignal, "macd, result, macdhist = talib.MACDEXT(testClose,12,talib.MA_Type.SMA,26,talib.MA_Type.SMA,9,talib.MA_Type.SMA)")
	compare(t, macdhist, "macd, macdsignal, result = talib.MACDEXT(testClose,12,talib.MA_Type.SMA,26,talib.MA_Type.SMA,9,talib.MA_Type.SMA)")
}

func TestTrix(t *testing.T) {
	result := Trix(testClose, 5)
	compare(t, result, "result = talib.TRIX(testClose,5)")
	result = Trix(testClose, 30)
	compare(t, result, "result = talib.TRIX(testClose,30)")
}

func TestMacd(t *testing.T) {
	macd, macdsignal, macdhist := Macd(testClose, 12, 26, 9)
	unstable := 100
	compare(t, macd[unstable:], fmt.Sprintf("result, macdsignal, macdhist = talib.MACD(testClose,12,26,9); result = result[%d:]", unstable))
	compare(t, macdsignal[unstable:], fmt.Sprintf("macd, result, macdhist = talib.MACD(testClose,12,26,9); result = result[%d:]", unstable))
	compare(t, macdhist[unstable:], fmt.Sprintf("macd, macdsignal, result = talib.MACD(testClose,12,26,9); result = result[%d:]", unstable))
}

func TestMacdFix(t *testing.T) {
	macd, macdsignal, macdhist := MacdFix(testClose, 9)
	unstable := 100
	compare(t, macd[unstable:], fmt.Sprintf("result, macdsignal, macdhist = talib.MACDFIX(testClose,9); result = result[%d:]", unstable))
	compare(t, macdsignal[unstable:], fmt.Sprintf("macd, result, macdhist = talib.MACDFIX(testClose,9); result = result[%d:]", unstable))
	compare(t, macdhist[unstable:], fmt.Sprintf("macd, macdsignal, result = talib.MACDFIX(testClose,9); result = result[%d:]", unstable))
}

func TestAd(t *testing.T) {
	result := Ad(testHigh, testLow, testClose, testVolume)
	compare(t, result, "result = talib.AD(testHigh,testLow,testClose,testVolume)")
}

func TestAdOsc(t *testing.T) {
	result := AdOsc(testHigh, testLow, testClose, testVolume, 3, 10)
	compare(t, result, "result = talib.ADOSC(testHigh,testLow,testClose,testVolume,3,10)")
}

func TestSuperTrend(t *testing.T) {
	///The below results are for 'NSE:NIFTY50-INDEX' for 1min timestamp
	// 261 | 24344.0335 | 1 | 0
	// 262 | 24344.0335 | 1 | 0
	// 263 | 24345.4668 | 1 | 0
	// 264 | 24346.6626 | 1 | 0
	// 265 | 24351.8388 | 1 | 0
	// 266 | 24351.8388 | 1 | 0
	// 267 | 24351.8388 | 1 | 0
	// 268 | 24351.8388 | 1 | 0
	// 269 | 24351.8388 | 1 | 0
	// 270 | 24383.6651 | -1 | -1
	// 271 | 24375.9461 | -1 | 0
	// 272 | 24369.1240 | -1 | 0
	// 273 | 24363.5916 | -1 | 0
	// 274 | 24363.5916 | -1 | 0
	// 275 | 24363.5916 | -1 | 0
	// 276 | 24363.5916 | -1 | 0
	// 277 | 24363.5916 | -1 | 0
	// 278 | 24363.5916 | -1 | 0
	// 279 | 24363.5916 | -1 | 0
	// 280 | 24363.5916 | -1 | 0
	// 281 | 24361.7171 | -1 | 0
	// 282 | 24359.9854 | -1 | 0
	// 283 | 24359.9854 | -1 | 0
	// 284 | 24357.6474 | -1 | 0
	// 285 | 24355.3827 | -1 | 0
	// 286 | 24355.1919 | -1 | 0
	// 287 | 24348.6877 | -1 | 0
	// 288 | 24341.8539 | -1 | 0
	// 289 | 24341.3685 | -1 | 0
	// 290 | 24340.1242 | -1 | 0
	// 291 | 24340.1242 | -1 | 0
	// 292 | 24339.9361 | -1 | 0
	// 293 | 24334.5200 | -1 | 0
	// 294 | 24334.5200 | -1 | 0
	// 295 | 24334.5200 | -1 | 0
	// 296 | 24334.5200 | -1 | 0
	// 297 | 24334.5200 | -1 | 0
	// 298 | 24334.5200 | -1 | 0
	// 299 | 24334.5200 | -1 | 0

	///The below results are Candles for  'NSE:NIFTY50-INDEX' for 1min timestamp at "time":"2026-05-07T11:24:49+05:30",

	high := []float64{
		24050.95, 24048.4, 24044.35, 24044.65, 24041.55, 24038.75, 24041.2, 24043.45, 24036.15, 24028.4, 24025.35, 24022.55, 24015.1, 24017.95, 24020.1, 24019.2, 24023.45, 24023.05, 24014.75, 24013.3, 24014.15, 24017.05, 24024.1, 24023.6, 24023.45, 24023.5, 24017.2, 24013.05, 24015.25, 24005.2, 24007.15, 24021.5, 24027.85, 24035.45, 24039.4, 24040.15, 24041.9, 24040.8, 24042.2, 24044.15, 24050, 24059.45, 24060.1, 24068.55, 24068.3, 24067.65, 24071.35, 24078.6, 24084.4, 24085.8, 24082.7, 24082.45, 24078.35, 24081.2, 24079.3, 24083.75, 24081.85, 24091.2, 24093, 24090.9, 24091.85, 24096.75, 24105.1, 24109.95, 24107.4, 24108.1, 24102, 24102.05, 24103.9, 24102.25, 24101.55, 24105.1, 24110.4, 24106.5, 24102.9, 24113.55, 24115.95, 24113.5, 24115.5, 24112.2, 24120.25, 24118.3, 24117.6, 24114.55, 24108.7, 24108.55, 24113.65, 24114.4, 24117.1, 24115.7, 24115, 24113.35, 24112.05, 24112.95, 24110.55, 24113.35, 24117.05, 24118.8, 24132.5, 24133.35, 24145.1, 24147.7, 24151.1, 24150.4, 24146, 24193.05, 24214.45, 24237.15, 24255.65, 24254.2, 24277.1, 24297.85, 24304.15, 24332.05, 24323.4, 24334.4, 24326.85, 24323.7, 24313.8, 24308.3, 24308.7, 24295.8, 24296.75, 24306.8, 24304.2, 24321.75, 24326.65, 24334.6, 24335.9, 24337.8, 24334.65, 24336.4, 24334.5, 24345.35, 24354.55, 24356.5, 24353.45, 24345.75, 24339.4, 24336.1, 24333.4, 24334.15, 24340.4, 24348.15, 24349.65, 24353.8, 24333.95, 24350.45, 24348.8, 24344.25, 24344.75, 24343.95, 24340.6, 24335.05, 24322.15, 24329.3, 24329.3, 24330.95, 24327.6, 24328.15, 24322.7, 24325.6, 24324.8, 24332.2, 24335.6, 24331.15, 24326, 24323.05, 24326.65, 24331.15, 24333.9, 24334.4, 24338.05, 24346.45, 24346.6, 24423.35, 24380.1, 24378.9, 24363.9, 24334.35, 24325.6, 24337.05, 24325.7, 24333.25, 24336.3, 24336.4, 24322.6, 24329.45, 24343.75, 24346.75, 24336.55, 24352.4, 24359.1, 24367.95, 24359.2, 24368.75, 24372.4, 24372.55, 24369.95, 24354.6, 24369.9, 24382.7, 24396, 24401.8, 24402.15, 24401.4, 24394.35, 24382.6, 24380.65, 24375.3, 24354.45, 24363.55, 24357.9, 24357.2, 24338.15, 24351.55, 24351.25, 24344.4, 24347.8, 24354.15, 24357.45, 24373.45, 24378.7, 24377.35, 24375.75, 24374.45, 24366.85, 24350.95, 24335, 24330.4, 24335.65, 24337.45, 24318.15, 24320.75, 24325.85, 24329.55, 24329.8, 24338.65, 24350.8, 24350.6, 24351.35, 24353.5, 24354.75, 24355.95, 24353.85, 24359.5, 24366.55, 24368.9, 24365.6, 24362.75, 24371.7, 24372.6, 24370.75, 24373.55, 24375.55, 24368.25, 24362.95, 24368.1, 24367.25, 24365.35, 24368.25, 24366.95, 24370.5, 24374.85, 24373.95, 24379.05, 24378.1, 24370.9, 24369.75, 24371.3, 24360.8, 24352.1, 24345.35, 24341.2, 24352.35, 24353.05, 24350.15, 24342.85, 24344.1, 24342.1, 24343.75, 24340.4, 24340.6, 24340.25, 24338.3, 24335.05, 24336.5, 24331.05, 24322.25, 24322.05, 24321.95, 24321.5, 24322.15, 24315.2, 24318.65, 24328.55, 24332.25, 24324.4, 24327, 24325.25}

	low := []float64{
		24047.2, 24040.65, 24039.45, 24039.7, 24033.2, 24032.9, 24032.3, 24035.25, 24025.45, 24019.7, 24017.7, 24010.15, 24005.95, 24011.45, 24014.55, 24014.85, 24017.3, 24012.1, 23997.9, 23999.6, 24009.8, 24011.1, 24012.7, 24018.4, 24019.5, 24013.5, 24008.75, 24008.35, 24002.5, 23999.95, 23999.65, 24005.6, 24020.3, 24024.95, 24033.5, 24035.7, 24037.5, 24031, 24032.25, 24039.2, 24036.3, 24048.55, 24055.4, 24056.25, 24061.6, 24062.35, 24062.75, 24068.85, 24074.3, 24075.75, 24077.75, 24075.2, 24070.6, 24073.05, 24073.9, 24076.25, 24074.65, 24076.05, 24085.65, 24084.3, 24085.85, 24084.15, 24094.7, 24099.1, 24101.4, 24102.2, 24094.9, 24095.2, 24095.8, 24096.9, 24094.15, 24098.4, 24104.4, 24096.7, 24094.3, 24101.8, 24109.2, 24108.55, 24103.75, 24102.65, 24112.05, 24112.85, 24109.25, 24109.1, 24101.9, 24096.8, 24096.4, 24111.05, 24108.05, 24107.35, 24107.4, 24106.25, 24106, 24108.35, 24103.75, 24107.4, 24110, 24114.4, 24115.9, 24129, 24133.4, 24142.5, 24144.6, 24141.8, 24139.75, 24139.15, 24180.55, 24211.5, 24231.05, 24243.05, 24252.05, 24268.1, 24289.65, 24299.55, 24300.15, 24314.7, 24312.8, 24310.2, 24280.75, 24280.1, 24292.1, 24277.9, 24280.9, 24293.85, 24296.2, 24299.4, 24319.9, 24322.35, 24325.4, 24329.7, 24325.1, 24327.85, 24320.8, 24327.05, 24345.1, 24344.7, 24340.4, 24330.65, 24331.4, 24328.75, 24324.65, 24325.35, 24330.7, 24339.85, 24337.55, 24313.85, 24320, 24324.8, 24338.3, 24338.35, 24335.55, 24336.95, 24331.75, 24320.75, 24314.05, 24314.95, 24319.9, 24322.7, 24322.1, 24318.8, 24318.3, 24317.05, 24318.1, 24321.8, 24330.4, 24317.4, 24320.4, 24318.4, 24317.85, 24325.7, 24327.95, 24325.6, 24327.3, 24336.45, 24334.4, 24369.95, 24346.6, 24359.65, 24327.2, 24317.7, 24314.3, 24317.35, 24300.7, 24302, 24325.3, 24313.55, 24312.2, 24313.6, 24324.75, 24330.75, 24320.15, 24326.85, 24346.7, 24350.25, 24350, 24356.15, 24362.1, 24364, 24350.65, 24346.9, 24354.8, 24369.15, 24382.35, 24392.9, 24390.3, 24391.95, 24377.4, 24374.45, 24371.65, 24352.25, 24334.55, 24350.35, 24350.85, 24329.4, 24325.65, 24331.2, 24339.5, 24334.45, 24333.35, 24340.5, 24340.15, 24358.85, 24370.85, 24367.95, 24366.7, 24358.9, 24350.55, 24333.4, 24315.25, 24324, 24323.3, 24316.85, 24310, 24311.3, 24315, 24321.6, 24324.2, 24329.25, 24338.8, 24337.55, 24345.55, 24345.05, 24346.2, 24349, 24344.95, 24348.75, 24356.55, 24361.6, 24357.25, 24353.9, 24360.35, 24362.65, 24365.25, 24366.45, 24367.25, 24359.2, 24356.6, 24357.65, 24352.85, 24359.2, 24363.45, 24358.7, 24364.8, 24366.5, 24368.2, 24372.5, 24369.85, 24356.3, 24357.85, 24359.55, 24350.05, 24344.25, 24339.25, 24332.45, 24335.15, 24347.5, 24342.1, 24334.4, 24338.7, 24336.55, 24337.4, 24335.8, 24331.25, 24334.7, 24330.1, 24331.2, 24329.8, 24320.05, 24316.15, 24316.7, 24313.85, 24315.15, 24312.5, 24310.05, 24313.15, 24316.25, 24323.75, 24318.1, 24318.6, 24314,
	}

	close := []float64{
		24047.85, 24043.3, 24042.2, 24041.3, 24036.65, 24037.95, 24039.4, 24037.5, 24028.4, 24022.3, 24022.35, 24012.05, 24013.7, 24015.65, 24018.3, 24017.95, 24022.1, 24014.75, 24000.35, 24011.6, 24013.8, 24015.15, 24020.7, 24021.35, 24023.1, 24017.55, 24010.65, 24013.05, 24005.95, 24002.1, 24007.15, 24021.15, 24027.65, 24035.45, 24038.9, 24039.45, 24041.05, 24033.5, 24042.2, 24040.65, 24049.9, 24057.45, 24058.7, 24067.5, 24063.45, 24066.15, 24071.35, 24078.4, 24084.05, 24081.1, 24080.85, 24078.05, 24072.9, 24076.4, 24075.35, 24079.95, 24080.2, 24089.65, 24088, 24086.35, 24088.05, 24096.75, 24102.1, 24104, 24103.55, 24103, 24096.5, 24102.05, 24098.6, 24100.9, 24097.45, 24105.1, 24106.65, 24098.1, 24102.75, 24112.55, 24110.4, 24112.65, 24106.15, 24111.25, 24116.65, 24113.9, 24115.15, 24109.3, 24106.95, 24097.1, 24113.65, 24112.3, 24117.05, 24109.5, 24112.85, 24108, 24111.7, 24110.05, 24109.35, 24112.45, 24116.95, 24116.7, 24132.5, 24133.35, 24145.05, 24147.15, 24149.95, 24146.35, 24141.15, 24181.7, 24213.25, 24235.4, 24242.9, 24250.5, 24269.55, 24292.15, 24301.2, 24321.15, 24311.65, 24323.3, 24321.95, 24312.9, 24281.5, 24306.75, 24292.1, 24280.1, 24296.75, 24302.3, 24300.45, 24321.2, 24326.65, 24330.05, 24329.1, 24334.35, 24328, 24333.05, 24326.75, 24343.95, 24349.3, 24347.55, 24346.65, 24338.8, 24334.3, 24332.1, 24328.1, 24332.3, 24340.05, 24345.6, 24348.25, 24324.55, 24330.3, 24345.4, 24339, 24342.15, 24337.4, 24340.45, 24332.5, 24320.8, 24315.5, 24328.05, 24323.75, 24325.55, 24324.45, 24321.35, 24321, 24322.65, 24324.15, 24331.9, 24330.65, 24321.7, 24321.9, 24319.75, 24326.35, 24329.05, 24331.35, 24329, 24335.35, 24342.7, 24339.3, 24379.6, 24365.55, 24363.15, 24330.8, 24318.25, 24317.9, 24324.1, 24304.9, 24328.45, 24335.35, 24314.25, 24313.45, 24323.5, 24340.7, 24334.15, 24329.55, 24352, 24358.6, 24352.75, 24356.7, 24367.6, 24367.55, 24369, 24350.65, 24354.15, 24369.1, 24382.7, 24394.9, 24400.05, 24397.35, 24392.65, 24378.6, 24376.15, 24371.85, 24352.25, 24352.65, 24356.15, 24353.2, 24331.95, 24335.2, 24348.65, 24343.65, 24337.8, 24347.3, 24342.95, 24356.65, 24373.45, 24375.4, 24373.15, 24370.25, 24363.85, 24350.55, 24334.1, 24325.7, 24328.3, 24331.7, 24316.85, 24312.75, 24315.7, 24323.65, 24325.5, 24328.45, 24338.65, 24341.95, 24349.05, 24349.15, 24348.75, 24353, 24349.4, 24353.15, 24356.4, 24365.2, 24363.75, 24357.3, 24360.3, 24370.25, 24369.8, 24368.5, 24373, 24367.3, 24361.15, 24359.25, 24367.2, 24361.25, 24364.6, 24364.85, 24366.55, 24367.05, 24369.55, 24373.45, 24376.1, 24369.85, 24358, 24368.95, 24360.45, 24350.7, 24344.4, 24339.9, 24334.95, 24351.6, 24349.55, 24342.3, 24340.65, 24340.4, 24340.85, 24337.4, 24340.1, 24336, 24337.15, 24332.85, 24334.85, 24330.3, 24320.75, 24318.3, 24320, 24316.6, 24321.5, 24312.5, 24313.95, 24317.75, 24328.55, 24323.95, 24318.5, 24325.85, 24318.4,
	}

	period := 10
	mult := 3.0

	line, dir, breakout := SuperTrend(high, low, close, period, mult)

	// --- length checks ---
	if len(line) != len(close) {
		t.Fatalf("line length: got %d, want %d", len(line), len(close))
	}
	if len(dir) != len(close) {
		t.Fatalf("dir length: got %d, want %d", len(dir), len(close))
	}
	if len(breakout) != len(close) {
		t.Fatalf("breakout length: got %d, want %d", len(breakout), len(close))
	}

	// --- dynamic warmup (FIXED) ---
	warmup := period * 3
	if warmup >= len(close) {
		warmup = 0
	}

	// --- basic validity checks after warmup ---
	for i := warmup; i < len(close); i++ {

		// direction must be valid
		if dir[i] != 1 && dir[i] != -1 {
			t.Errorf("dir[%d]=%f invalid (must be 1 or -1)", i, dir[i])
		}

		// line must be non-zero after warmup
		if line[i] == 0 {
			t.Errorf("line[%d]=0 after warmup", i)
		}

		// breakout must be valid
		if breakout[i] != 0 && breakout[i] != 1 && breakout[i] != -1 {
			t.Errorf("breakout[%d]=%f invalid", i, breakout[i])
		}
	}

	// --- warmup sanity (soft check, not strict zero forcing) ---
	for i := 0; i < warmup && i < len(close); i++ {
		if dir[i] != 0 && dir[i] != 1 && dir[i] != -1 {
			t.Errorf("dir[%d]=%f invalid in warmup", i, dir[i])
		}
	}

	// --- breakout must match direction flip only ---
	for i := warmup + 1; i < len(close); i++ {

		if breakout[i] != 0 {

			// must be flip point
			if dir[i] == dir[i-1] {
				t.Errorf("breakout[%d]=%f but no direction flip", i, breakout[i])
			}

			// must match new direction
			if breakout[i] != dir[i] {
				t.Errorf("breakout[%d]=%f must equal dir[%d]=%f",
					i, breakout[i], i, dir[i])
			}
		}
	}

	// --- no consecutive breakouts ---
	for i := warmup + 1; i < len(close); i++ {
		if breakout[i] != 0 && breakout[i-1] != 0 {
			t.Errorf("consecutive breakouts at %d and %d", i-1, i)
		}
	}

	// --- structural sanity (NOT strict price comparison) ---
	for i := warmup; i < len(close); i++ {

		if dir[i] == 1 {
			// bullish trend line should be below or near price range
			if line[i] > high[i] {
				t.Errorf("bullish %d: line %.2f above high %.2f", i, line[i], high[i])
			}
		}

		if dir[i] == -1 {
			// bearish trend line should be above or near price range
			if line[i] < low[i] {
				t.Errorf("bearish %d: line %.2f below low %.2f", i, line[i], low[i])
			}
		}
	}

	// --- edge cases ---

	// empty input
	l, d, b := SuperTrend([]float64{}, []float64{}, []float64{}, period, mult)
	if len(l) != 0 || len(d) != 0 || len(b) != 0 {
		t.Error("empty input should return empty slices")
	}

	// invalid period
	l, d, b = SuperTrend(high, low, close, 0, mult)
	for _, v := range l {
		if v != 0 {
			t.Error("invalid period: expected zero output")
			break
		}
	}

	// invalid multiplier
	l, d, b = SuperTrend(high, low, close, period, -1)
	for _, v := range l {
		if v != 0 {
			t.Error("negative multiplier: expected zero output")
			break
		}
	}
	n := 40
	start := len(line) - n

	fmt.Println("LAST 20 VALUES:")
	for i := start; i < len(line); i++ {
		fmt.Printf("%4d | %.4f | %.0f | %.0f\n",
			i, line[i], dir[i], breakout[i])
	}
}
