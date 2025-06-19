package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgMinutesPerLayer int) int {
	if avgMinutesPerLayer == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * avgMinutesPerLayer
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	// for each noodle layer, you will need 50 grams of noodles
	// for each sauce layer, you will need 0.2 liters of sauce
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}
		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	// Create a slice with the same len as quantities
	// make(typeOfSlice, lenQuantitiesSlice)
	scaled := make([]float64, len(quantities))
	// Copy the slice to the new variable
	// copy(slice, len)
	copy(scaled, quantities)

	for i := 0; i < len(quantities); i++ {
		scaled[i] *= (float64(numberOfPortions) / 2)
	}
	
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
