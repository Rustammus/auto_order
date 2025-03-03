package models

import "strconv"

type BasketList struct {
	SupplierID int64
	TotalSum   float64
	TotalCount int64
	Items      []BasketItem
}

func (l *BasketList) Size() (rows int, cols int) {
	if l == nil {
		return 0, 0
	}
	cols = 5
	rows = len(l.Items)
	return
}

func (l *BasketList) Cell(row int, col int) string {
	if row >= len(l.Items) {
		return "RANGE"
	}

	item := l.Items[row]
	switch col {
	case 0:
		return item.ProductCode
	case 1:
		return item.Title
	case 2:
		return item.Article
	case 3:
		return strconv.Itoa(int(item.Count))
	case 4:
		return strconv.Itoa(int(item.Price))
	}
	return "RANGE"
}
