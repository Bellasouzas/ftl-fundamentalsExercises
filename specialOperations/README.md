##09 - Special Operations


In previous exercises, we've tested just one function at a time, even when we tested it with multiple cases. This time, we're going to write three tests, one for calculator.Subtract, one for calculator.Multiply, and another for calculator.Divide.

To get you started, the calculator_test.go file contains three stub tests that you'll need to complete:

func TestSubtract(t *testing.T) {
	t.Fatal("Not implemented yet")
}

func TestMultiply(t *testing.T) {
	t.Fatal("Not implemented yet")
}

func TestDivide(t *testing.T) {
	t.Fatal("Not implemented yet")
}
TASK: Turn these into table-driven tests, with five test cases each, just like in the previous exercise, and make them pass. There are some special instructions for TestDivide, given below.

TestDivide: Handle With Care
For TestDivide, don't include any test cases that require dividing by zero. If we try to divide a number by zero in Go, we get a runtime error:

division by zero
We'll see how to handle this in the next exercise.

Because our calculator only handles integer arithmetic (that is, dealing with whole numbers), we'll also need to make sure that we don't include any divisions that result in a fractional part. For example, the result of dividing 2 by 3 can't be expressed as an integer. Stick to test cases that have an integer result, for now.

Done?
When all your test cases pass, you're done! Go on to the next exercise:

Planning to Fail