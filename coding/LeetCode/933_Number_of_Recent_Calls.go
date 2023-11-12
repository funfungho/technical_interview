package main

/*

https://leetcode.com/problems/number-of-recent-calls/description/

Easy, Queue

*/

func main() {

}

type RecentCounter struct {
	Queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

// ! dequeue to shrink the data size
func (this *RecentCounter) Ping(t int) int {
	this.Queue = append(this.Queue, t)
	for this.Queue[0] < t-3000 {
		// dequeue
		this.Queue = this.Queue[1:]
	}
	return len(this.Queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
