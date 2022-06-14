package convert

import "testing"

func TestDoSomthing(t *testing.T) {
	want := true
	got := DoSomethingWrong()
	if got != want {
		t.Errorf("DoSomething() = %t, want %t", got, want)
	}
}
