package purchaseService

import (
	"auto_order/internal/models"
	"auto_order/internal/repo"
	"auto_order/pkg/client/sqlite"
	"fmt"
	"strings"
	"time"
)

type CatalogImporter struct {
	prodRepo *repo.ProductRepo
}

func NewSearcher(r *repo.ProductRepo) *CatalogImporter {
	return &CatalogImporter{prodRepo: r}
}

func (i *CatalogImporter) Import(path string) {
	tt := time.Now()
	f, err := xlsx.OpenFile(path)
	if err != nil {
		panic(err)
	}

	db, err := sqlite.NewDB()
	if err != nil {
		panic(err)
	}

	repo = NewProductRepo(db)

	for i := 0; i < 2; i++ {
		sh := f.Sheets[i]
		rowsCount := sh.MaxRow
		colsCount := sh.MaxCol
		fmt.Printf("%s: rows:%d cols:%d\n", sh.Name, rowsCount, colsCount)

		proceedSheet(sh)
	}
	fmt.Println("Done in ", time.Since(tt))
}

func proceedSheet(sh *xlsx.Sheet) {
	err := sh.ForEachRow(rowReader)
	if err != nil {
		fmt.Printf("%s: row foreach err:%s\n", sh.Name, err)
	}

	fmt.Println("Skipped ", skippedCols)
}

func rowReader(r *xlsx.Row) error {
	//if r.GetCoordinate() > 25 {
	//	return nil
	//}

	return parseRow(r)
}

func parseRow(r *xlsx.Row) error {

	kng := r.GetCell(0)
	planet := r.GetCell(1)
	progress := r.GetCell(2)
	essco := r.GetCell(3)

	name := r.GetCell(5)
	price := r.GetCell(6)

	if kng == nil ||
		planet == nil ||
		progress == nil ||
		essco == nil ||
		name == nil ||
		price == nil {
		fmt.Printf("Parse Row %d Error: got nil cell\n", r.GetCoordinate())
		return nil
	}

	if isEmpty(kng.String()) &&
		isEmpty(planet.String()) &&
		isEmpty(progress.String()) &&
		isEmpty(essco.String()) &&
		isEmpty(price.String()) {
		fmt.Printf("Parse Row %d Possible chapter? Name: %s\n", r.GetCoordinate(), name.String())
		return nil
	}

	p := RowProduct{
		KngCode:      kng.String(),
		PlanetCode:   planet.String(),
		ProgressCode: progress.String(),
		EsscoCode:    essco.String(),
		Name:         name.String(),
	}

	_, err := repo.CreateProduct(p)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func isEmpty(s string) bool {
	if len(strings.TrimSpace(s)) == 0 {
		return true
	}
	return false
}

var skippedCols = make(map[int]struct{})

func cellReader(c *xlsx.Cell) error {

	x, _ := c.GetCoordinates()
	switch x {
	case 0:
		_, err := c.Int64()
		if err != nil {
			fmt.Println("Cell 0 Error: ", err)
		}
	case 1:
		_, err := c.Int64()
		if err != nil {
			fmt.Println("Cell 1 Error: ", err)
		}
	case 2:
		_, err := c.Int64()
		if err != nil {
			fmt.Println("Cell 2 Error: ", err)
		}
	case 3:
		_, err := c.Int64()
		if err != nil {
			fmt.Println("Cell 3 Error: ", err)
		}
	case 5:
		_ = c.String()

	default:
		skippedCols[x] = struct{}{}
	}
	return nil
}
