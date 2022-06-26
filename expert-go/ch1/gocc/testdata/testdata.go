package testdata

func SimpleFunction(n int) {
	println(n)
}

func ComplexFunction(n int) {
	if n > 0 {
		println("morethanzero")
		if n > 1 {
			println("morethanone")
			if n > 2 {
				println("morethantwo")
				if n > 3 {
					println("morethanthree")
					if n > 4 {
						println("morethanfour")
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			println(i * j)
		}
	}
	for k := 0; k < n; k++ {
		for l := k; l < n; l++ {
			println(k * l)
		}
	}
}
