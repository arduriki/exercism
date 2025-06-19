package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	}
	// Check if the item already exists
	bill[item] += value
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Return false if the given item is not in the bill
	if _, exists := bill[item]; !exists {
		return false
	}

	// Return false if the given unit is not in the units map
	unitValue, exists := units[unit]
	if !exists {
		return false
	}

	// Calculate new quantity after removing the unit amount
	newQuantity := bill[item] - unitValue

	// Return false if the new quantity would be less than 0
	if newQuantity < 0 {
		return false
	}

	// If the new quantity is 0, completely remove the item from the bill
	if newQuantity == 0 {
		delete(bill, item)
		return true
	} else {
		// Otherwise, reduce the quantity of the item
		bill[item] = newQuantity
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	// Return 0 and false if the item is not in the bill
	itemBill, exists := bill[item]
	if !exists {
		return 0, false
	} else {
		// Otherwise, return the quantity of the item in the bill and true
		return itemBill, true
	}
}
