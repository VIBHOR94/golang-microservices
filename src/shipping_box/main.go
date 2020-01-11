package shippingbox

import "reflect"

type Product struct {
	Name int
	Len  int
	Wid  int
	Hei  int
}

type Box struct {
	Len int
	Wid int
	Hei int
}

var minRequiredBoxDimension []int

func getBestBox(availableBoxes []Box, products []Product) Box {
	allProductLenSum, allProductHeiSum, allProductWidSum := 0, 0, 0
	for _, product := range products {
		allProductLenSum += product.Len
		allProductHeiSum += product.Hei
		allProductWidSum += product.Wid
	}
	minRequiredBoxDimension = []int{allProductLenSum, allProductWidSum, allProductHeiSum}
	bestBox := Box{}
	for _, currentBox := range availableBoxes {
		if accomodationPossible(currentBox) && smallerBox(bestBox, currentBox) {
			bestBox = currentBox
		}
	}

	return bestBox
}

func accomodationPossible(box Box) bool {
	return (box.Len >= minRequiredBoxDimension[0] && box.Wid >= minRequiredBoxDimension[1] && box.Hei >= minRequiredBoxDimension[2])
}

func smallerBox(bestBox, currentBox Box) bool {
	if !reflect.DeepEqual(bestBox, Box{}) {
		return smallerThanBestBox(bestBox, currentBox)
	}
	return false
}

func smallerThanBestBox(bestBox, currentBox Box) bool {
	isSmaller := (bestBox.Len + bestBox.Wid + bestBox.Hei) > (currentBox.Len + currentBox.Wid + currentBox.Hei)
	return isSmaller
}
