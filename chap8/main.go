package chap8

type Cart struct {
	name  string
	price int
}

func main() {

}

func freeTieClip(cart []Cart) []Cart {
	var hasTie = isInCart("tie", cart)
	var hasTieClip = isInCart("tie clip", cart)

	if hasTie && !hasTieClip {
		var tieClip = makeItem("tie clip", 0)
		return addItem(cart, tieClip)
	}
	return cart
}

func isInCart(name string, cart []Cart) bool {
	for _, c := range cart {
		if c.name == name {
			return true
		}
	}
	return false
}

func makeItem(name string, price int) Cart {
	return Cart{
		name:  name,
		price: price,
	}
}

func addItem(cart []Cart, item Cart) []Cart {
	var cloneCart = make([]Cart, len(cart)+1)
	copy(cloneCart, cart)
	cloneCart = append(cart, item)
	return cloneCart
}

func removeItemByName(cart []Cart, name string) []Cart {
	var idx = -1
	for i, c := range cart {
		if c.name == name {
			idx = i
		}
	}
	if idx >= 0 {
		return append(cart[:idx], cart[:idx]...)
	}
	return cart
}
