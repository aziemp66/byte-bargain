package test

import (
	"testing"

	httpCommon "github.com/aziemp66/byte-bargain/common/http"
)

func TestGetRecommendedProduct(t *testing.T) {
	productsDummy := []httpCommon.OrderProduct{
		{
			OrderProductID: "1",
			OrderID:        "1",
			ProductID:      "ID3",
			Qty:            2,
		},
		{
			OrderProductID: "2",
			OrderID:        "1",
			ProductID:      "ID2",
			Qty:            4,
		},
		{
			OrderProductID: "3",
			OrderID:        "2",
			ProductID:      "ID3",
			Qty:            1,
		},
		{
			OrderProductID: "4",
			OrderID:        "2",
			ProductID:      "ID2",
			Qty:            3,
		},
		{
			OrderProductID: "5",
			OrderID:        "3",
			ProductID:      "ID3",
			Qty:            1,
		},
		{
			OrderProductID: "6",
			OrderID:        "3",
			ProductID:      "ID1",
			Qty:            3,
		},
		{
			OrderProductID: "7",
			OrderID:        "4",
			ProductID:      "ID1",
			Qty:            2,
		},
		{
			OrderProductID: "8",
			OrderID:        "4",
			ProductID:      "ID1",
			Qty:            1,
		},
		{
			OrderProductID: "9",
			OrderID:        "5",
			ProductID:      "ID3",
			Qty:            1,
		},
		{
			OrderProductID: "10",
			OrderID:        "5",
			ProductID:      "ID1",
			Qty:            1,
		},
		{
			OrderProductID: "11",
			OrderID:        "6",
			ProductID:      "ID1",
			Qty:            5,
		},
	}

	type productOccurrence struct {
		ProductID  string
		Occurrence int
	}

	// count the occurrence of product id
	var bufferArray []productOccurrence
	for _, v := range productsDummy {
		isExist := false
		for i, v2 := range bufferArray {
			if v2.ProductID == v.ProductID {
				bufferArray[i].Occurrence++
				isExist = true
			}
		}

		if !isExist {
			bufferArray = append(bufferArray, productOccurrence{
				ProductID:  v.ProductID,
				Occurrence: 1,
			})
		}
	}

	t.Logf("bufferArray: %v", bufferArray)

	// sort the most ordered product by counting the most frequent product id
	for i := range bufferArray {
		for j := range bufferArray {
			if bufferArray[i].Occurrence > bufferArray[j].Occurrence {
				bufferArray[i], bufferArray[j] = bufferArray[j], bufferArray[i]
			}
		}
	}

	t.Logf("bufferArray: %v", bufferArray)
}
