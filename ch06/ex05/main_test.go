package ex05

import (
	"math"
	"testing"
)

func TestCopy(t *testing.T) {
	var i IntSet

	i.Add(1)
	i.Add(3)

	i2 := i.Copy()

	if i.String() != i2.String() {
		t.Errorf("copy error! expected %s, got %s\n", i.String(), i2.String())
	}

	if &i == &i2 {
		t.Errorf("not a copy! totally identical!")
	}
}
func TestClear(t *testing.T) {
	type dataset struct {
		intset IntSet
		result string
	}

	const bitSize = 64.0
	var tds []dataset

	var i IntSet
	i.Add(1)
	i.Add(3)
	i.Add(4)
	i.Add(6)
	i.Add(68)
	i.Add(356)
	i.Clear()

	var i2 IntSet
	i2.Add(3018)
	i.Clear()

	tds = append(tds, dataset{i, "{}"})
	tds = append(tds, dataset{i2, "{}"})

	for _, td := range tds {
		td.intset.Clear()
		if td.intset.String() != td.result {
			t.Errorf("error! expected %s, got %s\n", td.result, td.intset.String())
		}
	}
}

func TestReverse(t *testing.T) {
	type dataset struct {
		intset IntSet
		result string
	}

	const bitSize = 64.0
	var tds []dataset

	var i IntSet
	i.Add(1)
	i.Add(3)
	i.Add(4)
	i.Add(6)
	i.Add(68)
	i.Add(356)
	i.Remove(356)
	i.Remove(68)
	i.Remove(6)

	var i2 IntSet
	i2.Add(3064)
	i2.Add(601819)
	i2.Remove(3064)
	i2.Remove(1) //存在しない値をRemoveしても何も起きない

	tds = append(tds, dataset{i, "{1 3 4}"})
	tds = append(tds, dataset{i2, "{601819}"})

	for _, td := range tds {
		if td.intset.String() != td.result {
			t.Errorf("error! expected %s, got %s\n", td.result, td.intset.String())
		}
	}

}

func TestLen(t *testing.T) {
	type dataset struct {
		intset IntSet
		result int
	}

	const bitSize = 32 << (^uint(0) >> 63)
	var tds []dataset
	var i IntSet

	i.Add(1)
	i.Add(3)
	i.Add(4)
	i.Add(6)
	i.Add(68)
	i.Add(356)
	tds = append(tds, dataset{i, int(math.Ceil(356.0 / bitSize))})

	var i2 IntSet
	i2.Add(68)

	tds = append(tds, dataset{i2, int(math.Ceil(68.0 / bitSize))})

	// FIXME: this test fails with 32bit operator
	for _, td := range tds {
		if td.intset.Len() != td.result {
			t.Errorf("error! expected %d, got %d\n", td.result, td.intset.Len())
		}
	}
}

func TestElems(t *testing.T) {
	for _, td := range []struct {
		expected []int
	}{
		{[]int{1, 2, 3}},
		{[]int{1, 65, 89, 9888}},
	} {
		var a IntSet
		a.AddAll(td.expected...)

		elems := a.Elems()

		for _, v := range td.expected {
			if !a.Has(v) {
				t.Errorf("error! couldn't find %v in %v", v, elems)
			}
		}
	}
}
