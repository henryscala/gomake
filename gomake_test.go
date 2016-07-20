package gomake

import (
	"testing"
)

func TestBuild(t *testing.T) {
	AddRule(FakeTarget("1"), []string{"2", "3"}, []Command{})
	AddRule(FakeTarget("2"), []string{}, []Command{})
	AddRule(FakeTarget("3"), []string{}, []Command{})
	Build("1")
}
