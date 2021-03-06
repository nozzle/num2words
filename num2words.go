// Package num2words implements numbers to words converter.
package num2words

import "math"

// how many digit's groups to process
const groupsNumber int = 9

var smallNumbersWords = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen",
	"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}
var tensWords = []string{
	"", "", "twenty", "thirty", "forty", "fifty",
	"sixty", "seventy", "eighty", "ninety",
}
var scaleNumbersWords = []string{
	"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion",
}

type digitGroup int

// Convert converts number into the words representation.
func Convert(number int) string {
	// Zero rule
	if number == 0 {
		return smallNumbersWords[0]
	}

	// Divide into three-digits group
	var groups [groupsNumber]digitGroup
	positive := math.Abs(float64(number))

	// Form three-digit groups
	for i := 0; i < groupsNumber; i++ {
		groups[i] = digitGroup(math.Mod(positive, 1000))
		positive /= 1000
	}

	var textGroup [groupsNumber]string
	for i := 0; i < groupsNumber; i++ {
		textGroup[i] = digitGroup2Text(groups[i])
	}
	combined := textGroup[0]
	appendAnd := groups[0] > 0 && groups[0] < 100

	for i := 1; i < groupsNumber; i++ {
		if groups[i] != 0 {
			prefix := textGroup[i] + " " + scaleNumbersWords[i]

			if len(combined) != 0 {
				if appendAnd && i == 1 {
					prefix += " and "
				} else {
					prefix += " "
				}
			}

			combined = prefix + combined
		}
	}

	if number < 0 {
		combined = "minus " + combined
	}

	return combined
}

func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

func digitGroup2Text(group digitGroup) (ret string) {
	hundreds := group / 100
	tensUnits := intMod(int(group), 100)

	if hundreds != 0 {
		ret += smallNumbersWords[hundreds] + " hundred"

		if tensUnits != 0 {
			ret += " "
		}
	}

	tens := tensUnits / 10
	units := intMod(tensUnits, 10)

	if tens >= 2 {
		ret += tensWords[tens]

		if units != 0 {
			ret += "-" + smallNumbersWords[units]
		}
	} else if tensUnits != 0 {
		ret += smallNumbersWords[tensUnits]
	}

	return
}
