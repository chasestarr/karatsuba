package karatsuba

import "testing"

func TestSub(t *testing.T) {
	a := "5678"
	b := "1234"
	correct := "4444"
	if prod := sub(a, b); prod != correct {
		t.Fatal(a, "*", b, "is incorrect, expected:", correct, "received:", prod)
	}
}

func Test2Digit(t *testing.T) {
	a := "12"
	b := "43"
	correct := "516"
	if prod := Multiply(a, b); prod != correct {
		t.Fatal(a, "*", b, "is incorrect, expected:", correct, "received:", prod)
	}
}

func Test4Digit(t *testing.T) {
	a := "1234"
	b := "5678"
	correct := "7006652"
	if prod := Multiply(a, b); prod != correct {
		t.Fatal(a, "*", b, "is incorrect, expected:", correct, "received:", prod)
	}
}

func Test4DigitSimple(t *testing.T) {
	a := "1234"
	b := "4321"
	correct := "5332114"
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
