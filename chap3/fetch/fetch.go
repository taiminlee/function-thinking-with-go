package fetch

import (
	"encoding/json"
	"fmt"
	"os"
)

type Subscriber struct {
	Email    string `json:"email"`
	RecCount int    `json:"rec_count"`
}

type Coupon struct {
	Code string `json:"code"`
	Rank string `json:"rank"`
}

func Subscribers() []Subscriber {
	jsonFile, err := os.Open("./subscriber.json")

	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	defer jsonFile.Close()

	var subs []Subscriber
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&subs)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	return subs

	//for _, sub := range subs {
	//	fmt.Println(sub.Email, sub.RecCount)
	//}
	//return nil
	//return jsonFile
}

func Coupons() []Coupon {
	jsonFile, err := os.Open("./coupon.json")
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	defer jsonFile.Close()

	var coupons []Coupon
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&coupons)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	return coupons
}
