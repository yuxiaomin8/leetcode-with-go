package Offer
func fib(n int)int{
	const mod int = 1e9 + 7
	f0,f1:=0,1
	if n < 2 {
		return n
	}
	for i:=1;i<n;i++{
		f0,f1=f1,(f1+f0)%mod
	}
	return f1
}