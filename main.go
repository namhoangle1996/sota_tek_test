package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// giai thua
	fmt.Printf("giai thua %v\n",Giaithua(0))
	fmt.Printf("giai thua %v\n",Giaithua(3))

	// so hoan hao
	fmt.Printf("so hoan hao %v\n",SoHoanHao(6))
	fmt.Printf("so hoan hao %v\n",SoHoanHao(7))

	// xac dinh chuoi ki tu
	fmt.Printf("vi tri %v\n",ChuoiKitu("231"))
	fmt.Printf("vi tri %v\n",ChuoiKitu("124"))

	// mua keo
	fmt.Printf("tong so keo %v\n", MuaKeo(10, 2, 5))
	fmt.Printf("tong so keo %v\n",MuaKeo(12 , 4, 4 ))
	fmt.Printf("tong so keo %v\n",MuaKeo(6 , 2, 2 ))

}


func Giaithua(input uint) uint  {
	var response uint = 1
	if input == 0 {
		return 1
	}

	for i:= 1 ; i <= int(input) ; i++{
		response =  uint(int(response) * i)
	}

	return response
}

func SoHoanHao(input uint) bool  {
	var sum = 0

	for i:= 1 ; i < int(input) ; i++{
		if int(input) % i == 0 {
			sum = sum + i
		}
	}

	if sum == int(input) {
		return true
	}

	return false
}


func ChuoiKitu(input string)  int64  {
	if len(input) < 1 {
		panic("do dai chuoi khong hop le , cần > 1")
	}

	var newInputInt []int64

	mapToArr := strings.Split(input,"")
	var maxNumber int64 = 0
	var position int64

	for i:=0;  i < len(mapToArr); i++ {
		n, err := strconv.ParseInt(mapToArr[i], 10, 64)
		if err != nil  {
			panic(err)
		}
		newInputInt = append(newInputInt, n )

		if newInputInt[i] >= maxNumber {
			maxNumber = newInputInt[i]
			position = int64(i) + 1
		}
	}

	return position
}

func MuaKeo(balance uint , pricePerCandy uint , shopPromotion uint  ) uint {
	initTotalCandyFromBalance := balance / pricePerCandy    // doi keo tu tien
	initTotalCandyFromSwap := initTotalCandyFromBalance / shopPromotion // doi keo tu giay goi keo


	initTotalCandyFromSwap = initTotalCandyFromSwap + initTotalCandyFromBalance / shopPromotion

	initTotalCandyFromBalance = initTotalCandyFromBalance+  initTotalCandyFromSwap

	return initTotalCandyFromBalance
}

