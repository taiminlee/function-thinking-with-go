package main

type Affiliate struct {
	sales      float32
	commission float32
	bankCode   int
}

func sendPayout(code int, owned float32) bool {
	// action
	return true
}

func figurePayout(affiliate Affiliate) {
	var owned float32 = affiliate.sales * affiliate.commission
	if owned > 100 {
		sendPayout(affiliate.bankCode, owned)
	}
}

func affiliatePayout(affiliates []Affiliate) {
	for _, affiliate := range affiliates {
		figurePayout(affiliate)
	}
}

func main() {
	var affiliates = []Affiliate{
		{
			sales:      5,
			commission: 5,
			bankCode:   123,
		},
		{
			sales:      25,
			commission: 20,
			bankCode:   123,
		},
	}
	affiliatePayout(affiliates)
}
