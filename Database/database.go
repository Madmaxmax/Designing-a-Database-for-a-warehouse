package Database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

const (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)

type Database struct {
	db *sql.DB
}

func (d *Database) Close() error {
	return d.db.Close()
}

func DatabaseConnect() (*Database, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) ExecuteQueryFile(filename string) error {
	query, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	_, err = d.db.Exec(string(query))
	if err != nil {
		return err
	}
	return nil
}

type OrderInfo struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}

func (d *Database) GetOrderInfo(order int) []OrderInfo {
	query := "SELECT order_info FROM order_details WHERE order_id=$1"
	var productsStr string
	row := d.db.QueryRowContext(context.TODO(), query, order)
	err := row.Scan(&productsStr)
	if err != nil {
		panic(err)
	}
	var intermediate map[string]OrderInfo

	err = json.Unmarshal([]byte(productsStr), &intermediate)
	if err != nil {
		panic(err)
	}
	var arrProduct []OrderInfo
	for _, value := range intermediate {
		arrProduct = append(arrProduct, value)
	}
	return arrProduct
}

type ProductInfo struct {
	ProductId   int
	ProductName string
	ShelfId     int
	ShelfName   string
	IsMain      bool
	Count       int
	AddShelf    []string
	AddShelfId  []int
}

func (d *Database) GetProductInfo(arrProduct []OrderInfo) ([]ProductInfo, map[int]string) {
	query, err := os.ReadFile("Sql/GetShelving.sql")
	if err != nil {
		panic(err)
	}
	var data []ProductInfo
	shelfIdArr := make(map[int]string)
	for i := 0; i < len(arrProduct); i++ {
		var tmp []ProductInfo
		row, err := d.db.QueryContext(context.TODO(), string(query), arrProduct[i].Id)
		if err != nil {
			panic(err)
		}
		for row.Next() {
			p := ProductInfo{}
			err := row.Scan(&p.ProductId, &p.ShelfId, &p.IsMain, &p.Count, &p.ProductName, &p.ShelfName)
			if err != nil {
				panic(err)
			}
			if p.Count >= arrProduct[i].Count {
				p.Count = arrProduct[i].Count
			} else {
				p.Count = -1
			}
			shelfIdArr[p.ShelfId] = p.ShelfName
			tmp = append(tmp, p)
		}

		if len(tmp) > 1 {
			tmp = getTotalPrInfo(tmp, arrProduct[i].Count)
		}

		data = append(data, tmp[0])
	}
	return data, shelfIdArr
}

func getTotalPrInfo(product []ProductInfo, count int) []ProductInfo {
	totalProduct := product[0]
	for i := 0; i < len(product); i++ {
		if product[i].IsMain {
			totalProduct.IsMain = true
			totalProduct.ShelfId = product[i].ShelfId
			totalProduct.ShelfName = product[i].ShelfName
			totalProduct.Count = count
		} else {
			totalProduct.AddShelf = append(totalProduct.AddShelf, product[i].ShelfName)
			totalProduct.AddShelfId = append(totalProduct.AddShelfId, product[i].ShelfId)
		}
	}
	arr := make([]ProductInfo, 0)
	arr = append(arr, totalProduct)
	return arr
}
