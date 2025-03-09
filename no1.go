package main

import "fmt"

const PI = 3.14

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	*l = PI * r * r
	*k = 2 * PI * r
}

func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(lL, lP, kL, kP float64, totLuas, totkel *float64) {
	*totLuas = lL + lP
	*totkel = kL + kP
}

func main() {
	var r, s, lL, lP, kL, kP, totkel, totLuas float64
	for {
		fmt.Scan(&r, &s)
		if r == 0 && s == 0 {
			break
		}

		hitungLuasKelilingLingkaran(r, &lL, &kL)
		hitungLuasKelilingPersegi(s, &lP, &kP)
		hitungTotal(lL, lP, kL, kP, &totLuas, &totkel)

		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n",
			r, s, lL, lP, kL, kP, totLuas, totkel)

	}
}
