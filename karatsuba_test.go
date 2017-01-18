package karatsuba

import "testing"

func Test4Digit(t *testing.T) {
	a := "1234"
	b := "5678"
	correct := "7006652"
	if prod := Multiply(a, b); prod != correct {
		t.Fatal(a, "*", b, "is incorrect, expected:", correct, "received:", prod)
	}
}

func Test64Digit(t *testing.T) {
	a := "3141592653589793238462643383279502884197169399375105820974944592"
	b := "2718281828459045235360287471352662497757247093699959574966967627"
	correct := "8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184"
	if prod := Multiply(a, b); prod != correct {
		t.Fatal(a, "*", b, "is incorrect, expected:", correct, "received:", prod)
	}
}

func TestUnevenDigits(t *testing.T) {
	a := "123"
	b := "4567"
	correct := "561741"
	if prod := Multiply(a, b); prod != correct {
		t.Fatal(a, "*", b, "is incorrect, expected:", correct, "received:", prod)
	}
}
