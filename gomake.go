package gomake

import (
	"errors"
	"fmt"
)

type Target interface {
	Name() string
}

type fileTarget struct {
	name string //full or relative path
}

type fakeTarget struct {
	name string //only a name
}

type Command func() error

type Rule struct {
	Target     Target
	Dependents []string
	Commands   []Command
}

var (
	allRules []Rule
	ruleMap  map[string]Rule = map[string]Rule{}
)

func (self fakeTarget) Name() string {
	return self.name
}

func (self fileTarget) Name() string {
	return self.name
}

func FakeTarget(name string) Target {
	t := fakeTarget{name: name}
	return t
}

func FileTarget(name string) Target {
	t := fileTarget{name: name}
	return t
}

func AddRule(target Target, dependents []string, commands []Command) {
	r := Rule{Target: target, Dependents: dependents, Commands: commands}
	ruleMap[target.Name()] = r
	allRules = append(allRules, r)
}

func Build(targetName string) error {
	fmt.Println("building", targetName, "...")
	rule, ok := ruleMap[targetName]
	if !ok {
		return errors.New("no such target defined")
	}
	for _, dependTargetName := range rule.Dependents {
		err := Build(dependTargetName)
		if err != nil {
			return err
		}
	}
	for _, cmd := range rule.Commands {
		err := cmd()
		if err != nil {
			return err
		}
	}
	fmt.Println("building", targetName, "done.")
	return nil
}

func compareTarget(target1, target2 string) int {
	return 0
}
