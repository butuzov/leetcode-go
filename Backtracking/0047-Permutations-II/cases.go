package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	nums     []int
	expected [][]int
}{
	{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
}
