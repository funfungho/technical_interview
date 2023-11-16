package main

/*

https://leetcode.com/problems/online-stock-span/

Medium, Stack, Monotonic Stack

*/

type StockSpanner struct {
	// 0: price, 1: ans
	Stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

// find the first element that's larger than price on its left
func (this *StockSpanner) Next(price int) int {
	// monotonically decreasing stack
	days := 1
	for len(this.Stack) > 0 && price >= this.Stack[len(this.Stack)-1][0] {
		// ! inherit the previous qualified days (the previous answer)
		days += this.Stack[len(this.Stack)-1][1]
		// pop
		this.Stack = this.Stack[:len(this.Stack)-1]
	}

	this.Stack = append(this.Stack, [2]int{price, days})
	return days
}

type StockSpannerMy struct {
	Stack  []int // * store index
	Prices []int
}

func ConstructorMy() StockSpannerMy {
	return StockSpannerMy{}
}

// find the first element that's larger than price on its left
func (this *StockSpannerMy) Next(price int) int {
	// monotonically decreasing stack (top is the smallest)
	// keep the price value
	this.Prices = append(this.Prices, price)
	currIndex := len(this.Prices) - 1
	// less or equal
	for len(this.Stack) > 0 && price >= this.Prices[this.Stack[len(this.Stack)-1]] {
		// pop
		this.Stack = this.Stack[:len(this.Stack)-1]
	}

	var span int
	if len(this.Stack) > 0 {
		span = currIndex - this.Stack[len(this.Stack)-1]
	} else {
		span = currIndex + 1
	}
	this.Stack = append(this.Stack, currIndex)
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
