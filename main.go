package main

import (
	"GoTestWork20.02.2024/Database"
	db "GoTestWork20.02.2024/Database"
	"fmt"
	_ "github.com/lib/pq"
	"sort"
	"strconv"
	"strings"
)

func arrStringToArrInt(strArr []string) []int {
	arrInt := make([]int, len(strArr), len(strArr))
	for i := 0; i < len(strArr); i++ {
		num, err := strconv.Atoi(strArr[i])
		if err != nil {
			panic(err)
		}
		arrInt[i] = num
	}
	return arrInt
}

func mergeMap(firstMap map[int]string, secondMap map[int]string) map[int]string {
	for key, value := range secondMap {
		firstMap[key] = value
	}
	return firstMap
}

type ExtendedProductInfo struct {
	Database.ProductInfo
	OrderId int
}

func getOrdersProducts(inputArr []string) map[int][]ExtendedProductInfo {
	db, _ := db.DatabaseConnect()
	defer db.Close()
	ordersArr := arrStringToArrInt(inputArr)
	shelfIdArr := make(map[int]string)
	arrInfo := make(map[int][]Database.ProductInfo)
	for i := 0; i < len(ordersArr); i++ {
		orderInfo := db.GetOrderInfo(ordersArr[i])
		productInf, tmpShelfId := db.GetProductInfo(orderInfo)
		arrInfo[ordersArr[i]] = productInf
		shelfIdArr = mergeMap(shelfIdArr, tmpShelfId)
	}

	shelving := make(map[int][]ExtendedProductInfo)
	for key, _ := range shelfIdArr {
		for keyArrInf, valueArrInf := range arrInfo {
			for i := 0; i < len(valueArrInf); i++ {
				if key == valueArrInf[i].ShelfId {
					newExtProductInfo := ExtendedProductInfo{valueArrInf[i], keyArrInf}
					shelving[key] = append(shelving[key], newExtProductInfo)
				}
			}
		}
	}
	return shelving
}

func main() {
	db, _ := db.DatabaseConnect()
	defer db.Close()
	if err := db.ExecuteQueryFile("Sql/CreateDb.sql"); err != nil {
		panic(err)
	}

	//if err := db.ExecuteQueryFile("Sql/InsertInfo.sql"); err != nil {
	//	panic(err)
	//}

	var inputStr string
	//inputStr = "10,11,14,15"
	fmt.Scan(&inputStr)
	numberArr := strings.Split(inputStr, ",")
	fmt.Println("=+=+=+=")
	fmt.Println("Страница сборки заказов", inputStr)
	shelfData := getOrdersProducts(numberArr)
	keys := make([]int, 0, len(shelfData))
	for key := range shelfData {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	// вывод дополнительно отсортировал так как не во всех версиях go есть сортировка вывода map
	for _, key := range keys {
		value := shelfData[key]
		fmt.Println("===Стеллаж", value[0].ShelfName)
		for i := 0; i < len(value); i++ {
			fmt.Printf("%s (id=%d)\n", value[i].ProductName, value[i].ProductId)
			fmt.Printf("заказ %d, %d шт\n", value[i].OrderId, value[i].Count)
			if len(value[i].AddShelfId) != 0 {
				fmt.Printf("доп стеллаж: %s\n", strings.Join(value[i].AddShelf, ", "))
			}
			fmt.Println()
		}
	}
}
