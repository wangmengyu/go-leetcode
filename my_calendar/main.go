package main

import "fmt"

/**
设计一个MyCalendar类， 提供一个函数， 表示对在自己的calendar上某一个时间段预定时间
public boolean book(int start, int end)
每次预定成功， 表示start <= x < end的时间被预定了。
同时， 要求时间是不能overlap的。 如果新输入的时间段和之前的时间段有冲突的话，就返回false
*/
type MyCalendar struct {
	BookMap map[int]int
}

func (mc *MyCalendar) Book(start, end int) bool {
	for i := start; i < end; i++ {
		if _, ex := mc.BookMap[i]; ex {
			return false
		}
		mc.BookMap[i] = i
	}
	return true
}
func main() {

	mc := new(MyCalendar)
	mc.BookMap = make(map[int]int)
	fmt.Println(mc.Book(10, 20))
	fmt.Println(mc.Book(15, 25))
	fmt.Println(mc.Book(20, 30))
	/**
	MyCalendar.book(10, 20); // returns true
	MyCalendar.book(15, 25); // returns false
	MyCalendar.book(20, 30); // returns true
	*/

}
