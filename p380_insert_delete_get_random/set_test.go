package p380_insert_delete_get_random

import (
	"fmt"
	"testing"
)

/*
Explanation
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.
*/
func TestName(t *testing.T) {
	/*

		["RandomizedSet","insert","remove","insert","remove","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom"]
		[[],[0],[0],[-1],[0],[],[],[],[],[],[],[],[],[],[]]

	*/
	randomizedSet := Constructor()
	randomizedSet.Insert(0)
	randomizedSet.Remove(0)
	randomizedSet.Insert(-1)
	randomizedSet.Remove(0)
	res := randomizedSet.GetRandom()
	fmt.Printf("res: %d\n", res)
	//randomizedSet := Constructor()
	//var res bool
	//var ress int
	//res = randomizedSet.Insert(1) // Inserts 1 to the set. Returns true as 1 was inserted successfully.
	//fmt.Printf("randomizedSet.Insert(1) = %t\n", res)
	//
	//res = randomizedSet.Remove(2) // Returns false as 2 does not exist in the set.
	//fmt.Printf("randomizedSet.Remove(2) = %t\n", res)
	//
	//res = randomizedSet.Insert(2) // Inserts 2 to the set, returns true. Set now contains [1,2].
	//fmt.Printf("randomizedSet.Insert(2) = %t\n", res)
	//
	//ress = randomizedSet.GetRandom() // getRandom() should return either 1 or 2 randomly.
	//fmt.Printf("randomizedSet.GetRandom() = %d\n", ress)
	//
	//res = randomizedSet.Remove(1) // Removes 1 from the set, returns true. Set now contains [2].
	//fmt.Printf("randomizedSet.Remove(1) = %t\n", res)
	//
	//res = randomizedSet.Insert(2) // 2 was already in the set, so return false.
	//fmt.Printf("randomizedSet.Insert(2) = %t\n", res)
	//
	//ress = randomizedSet.GetRandom() // Since 2 is the only number in the set, getRandom() will always return 2.
	//fmt.Printf("randomizedSet.GetRandom() = %d\n", ress)

}
