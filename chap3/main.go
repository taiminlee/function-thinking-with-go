package main

import (
	"chap3/fetch"
	"fmt"
)

func main() {
	// action fetch
	// data subs/coupons
	var subs = fetch.Subscribers()
	var coupons = fetch.Coupons()
	// calculations
	var email = planListOfEmails(subs, coupons)
	fmt.Println(email)
}

type EmailCoupons struct {
	email string
	code  string
}

func planListOfEmails(subs []fetch.Subscriber, coupons []fetch.Coupon) []EmailCoupons {
	var emails []EmailCoupons
	var codeByRank = couponCodeByRank(coupons)
	for _, sub := range subs {
		emails = append(emails, EmailCoupons{email: sub.Email, code: codeByRank[recCountCoupons(sub.RecCount)]})
	}
	return emails
}

func recCountCoupons(recCount int) string {
	switch {
	case recCount < 15:
		return "BAD"
	case recCount < 25:
		return "GOOD"
	default:
		return "BEST"
	}
}

func couponCodeByRank(coupons []fetch.Coupon) map[string]string {
	var nameMap = make(map[string]string)
	for _, coupon := range coupons {
		nameMap[coupon.Rank] = coupon.Code
	}
	return nameMap
}
