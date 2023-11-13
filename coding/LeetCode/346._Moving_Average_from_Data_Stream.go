package main

/*

https://leetcode.com/problems/moving-average-from-data-stream/description/

Easy, Queue

*/

type MovingAverage struct {
	Queue []int
	Size  int
	Sum   int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		Size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.Sum += val
	this.Queue = append(this.Queue, val)
	for len(this.Queue) > this.Size {
		this.Sum -= this.Queue[0]
		this.Queue = this.Queue[1:]
	}

	return float64(this.Sum) / float64(len(this.Queue))
}
