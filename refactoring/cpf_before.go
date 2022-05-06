package refectoring

import (
	"fmt"
	"strconv"
	"strings"
)

func Validate(str string) bool {

	if str != "" {
		if len(str) >= 11 || len(str) <= 14 {
			str = strings.Replace(str, ".", "", len(str))
			str = strings.Replace(str, "-", "", len(str))
			str = strings.Replace(str, " ", "", len(str))

			c := str[0:1]
			var equal bool = true
			for _, letter := range str {
				if string(letter) != c {
					equal = false
				}
			}

			if equal == true {
				return false
			}

			var nDigVerific, nDigResult string
			var d1, d2, nCount int
			var dg1, dg2, rest int
			d1 = 0
			d2 = 0
			dg1 = 0
			dg2 = 0
			rest = 0

			for nCount = 1; nCount < len(str)-1; nCount++ {
				digito, err := strconv.Atoi(str[nCount-1 : nCount])

				if err != nil {
					return false
				}

				d1 += (11 - nCount) * digito
				d2 += (12 - nCount) * digito
			}

			rest = (d1 % 11)

			if rest < 2 {
				dg1 = 0
			} else {
				dg1 = 11 - rest
			}

			d2 += 2 * dg1
			rest = (d2 % 11)

			if rest < 2 {
				dg2 = 0
			} else {
				dg2 = 11 - rest
			}
			nDigVerific = str[len(str)-2:]
			nDigResult = fmt.Sprintf("%d%d", dg1, dg2)
			return nDigVerific == nDigResult
		}
		return false
	}
	return false
}
