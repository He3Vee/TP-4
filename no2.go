package main

import "fmt"

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm++
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk++
	}
}

func hitungJumGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
	*jg += g
	*jk += k
	*jsg = *jg - *jk
}

func hitungJumPoint(jm, jd int, jp *int) {
	*jp = (jm * 3) + jd
}

func main() {
	var n, g, k, jm, jd, jk, jg, jkg, jsg, jp int

	fmt.Print("Jumlah Pertandingan : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&g, &k)

		hitungMenang(g, k, &jm)
		hitungDraw(g, k, &jd)
		hitungKalah(g, k, &jk)
		hitungJumGolKegolanSelisih(g, k, &jg, &jkg, &jsg)
	}

	hitungJumPoint(jm, jd, &jp)

	fmt.Printf("%d %d %d %d %d %d %d %d\n", n, jm, jd, jk, jg, jkg, jsg, jp)

}
