package main

type Cart struct {
	name  string
	price float32
}

var shoppingCart []Cart
var shoppingCartTotal float32 = 0

func addItemToCart(name string, price float32) {
	shoppingCart = append(shoppingCart, Cart{
		name:  name,
		price: price,
	})
	shoppingCartTotal = calcCartTotal(shoppingCart)
}

func calcCartTotal(shoppingCart []Cart) float32 {
	var total float32 = 0
	for _, cart := range shoppingCart {
		total += cart.price
	}

	return total
}

func update_shipping_icons() {

}

func main() {
	addItemToCart("88D", 3.3)
	addItemToCart("jpx", 0.9)
}
