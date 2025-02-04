package p380_insert_delete_get_random

import "math/rand"

// https://leetcode.com/problems/insert-delete-getrandom-o1/description/?envType=study-plan-v2&envId=top-interview-150

type RandomizedSet struct {
	list map[int]int
	keys []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		list: make(map[int]int),
		keys: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.list[val]; ok {
		return false
	}

	this.keys = append(this.keys, val)

	this.list[val] = len(this.keys) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.list[val]; !ok {
		return false
	}

	// move list elem
	this.keys[this.list[val]] = this.keys[len(this.keys)-1]
	this.list[this.keys[this.list[val]]] = this.list[val]
	this.keys = this.keys[:len(this.keys)-1]

	delete(this.list, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.keys[rand.Intn(len(this.keys))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
