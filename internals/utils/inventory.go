package utils

func QuantityMap(inv Inventory) map[string]int {
	quantity := make(map[string]int)

	for _, item := range inv.Assets {
		_, exists := quantity[item.Classid]

		if exists {
			quantity[item.Classid]++
		} else {
			quantity[item.Classid] = 1
		}
	}

	return quantity
}
