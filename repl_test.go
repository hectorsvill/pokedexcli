package main

// import "testing"

// func TestCleanInput(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		expected []string
// 	}{
// 		{
// 			input:    "  hello  world  ",
// 			expected: []string{"hello", "world"},
// 		},
// 		{
// 			input:    " Abra ",
// 			expected: []string{"abra"},
// 		},
// 	}

// 	for _, c := range cases {
// 		actual := cleanInput(c.input)
// 		if len(actual) != len(c.expected) {
// 			t.Errorf("len error: Test: %v", actual)
// 			break
// 		}
// 		for i := range actual {
// 			word := actual[i]
// 			expectedWord := c.expected[i]
// 			if word != expectedWord {
// 				t.Errorf("word: %v is not actual: %v", word, expectedWord)
// 				break
// 			}
// 		}
// 	}
// }
