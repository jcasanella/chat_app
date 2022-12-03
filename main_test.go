package main

import "testing"

func Test_sayHello(t *testing.T) {
	name := "Jordi"
	want := "Hello Jordi, you are my new buddy"

	if got := sayHello(name); got != want {
		t.Errorf("sayHello() = %q, want %q", got, want)
	}
}
