package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	s        string
	expected bool
}{
	{"", true},
	{"()", true},
	{"{}", true},
	{"(]", false},
	{"(])", false},
	{"[]", true},
	{"[{}]", true},
	{"[{}]()", true},
	{"([{}]()", false},
	{"(([]){})", true},
}
