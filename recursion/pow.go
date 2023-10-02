package recursion

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return (1 / x) * MyPow(x, n + 1)
	}
	return x * MyPow(x, n - 1)
}


func MyPow2(x float64, n int) float64 {
	if n < 0 {
		x , n = 1/x, -n
	}
	return myPowHelper(x ,n)
}

func myPowHelper(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n % 2 == 1 {
		return x * myPowHelper(x,n - 1)
	} else {
		return myPowHelper(x * x,n / 2)
	}
}