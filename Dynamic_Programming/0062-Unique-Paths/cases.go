package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	m        int
	n        int
	expected int
}{
	{3, 2, 3},
	{7, 3, 28},
	{7, 1, 1},
	{1, 7, 1},
	{10, 10, -1},
}