package goexpr

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type scenario [4]interface{}

func (s scenario) String() string {
	return fmt.Sprintf("%s %s %s %v", s[1], s[0], s[2], s[3])
}

func TestBinaryComparisons(t *testing.T) {
	now := time.Now()
	scenarios := []scenario{
		scenario{"==", nil, nil, true},
		scenario{"==", nil, false, false},
		scenario{"==", nil, true, false},
		scenario{"==", nil, 0, false},
		scenario{"==", nil, 1, false},
		scenario{"==", nil, 0.0, false},
		scenario{"==", nil, 0.1, false},
		scenario{"==", nil, "", false},
		scenario{"==", nil, "something", false},
		scenario{"==", nil, time.Time{}, false},
		scenario{"==", nil, now, false},
		scenario{"==", false, nil, false},
		scenario{"==", false, false, true},
		scenario{"==", false, true, false},
		scenario{"==", false, 0, true},
		scenario{"==", false, 1, false},
		scenario{"==", false, 0.0, true},
		scenario{"==", false, 0.1, false},
		scenario{"==", false, "false", true},
		scenario{"==", false, "true", false},
		scenario{"==", false, "", true},
		scenario{"==", true, "", false},
		scenario{"==", false, "something", false},
		scenario{"==", false, time.Time{}, false},
		scenario{"==", false, now, false},
		scenario{"==", 0, 0, true},
		scenario{"==", 0, 1, false},
		scenario{"==", 0, 0.0, true},
		scenario{"==", 0, 0.1, false},
		scenario{"==", 0, "0", true},
		scenario{"==", 0, "0.0", true},
		scenario{"==", 0, "0.1", false},
		scenario{"==", 0, time.Time{}, false},
		scenario{"==", 0, now, false},
		scenario{"==", now, now, true},
		scenario{"==", now, zeroTime, false},
		scenario{"==", now, now.String(), true},
		scenario{"==", now, zeroTime.String(), false},

		scenario{"!=", nil, nil, false},
		scenario{"!=", nil, false, true},
		scenario{"!=", nil, true, true},
		scenario{"!=", nil, 0, true},
		scenario{"!=", nil, 1, true},
		scenario{"!=", nil, 0.0, true},
		scenario{"!=", nil, 0.1, true},
		scenario{"!=", nil, "", true},
		scenario{"!=", nil, "something", true},
		scenario{"!=", nil, time.Time{}, true},
		scenario{"!=", nil, now, true},
		scenario{"!=", false, nil, true},
		scenario{"!=", false, false, false},
		scenario{"!=", false, true, true},
		scenario{"!=", false, 0, false},
		scenario{"!=", false, 1, true},
		scenario{"!=", false, 0.0, false},
		scenario{"!=", false, 0.1, true},
		scenario{"!=", false, "false", false},
		scenario{"!=", false, "true", true},
		scenario{"!=", false, "something", false},
		scenario{"!=", false, time.Time{}, false},
		scenario{"!=", false, now, false},
		scenario{"!=", 0, 0, false},
		scenario{"!=", 0, 1, true},
		scenario{"!=", 0, 0.0, false},
		scenario{"!=", 0, 0.1, true},
		scenario{"!=", 0, "0", false},
		scenario{"!=", 0, "0.0", false},
		scenario{"!=", 0, "0.1", true},
		scenario{"!=", 0, time.Time{}, false},
		scenario{"!=", 0, now, false},
		scenario{"!=", now, now, false},
		scenario{"!=", now, zeroTime, true},
		scenario{"!=", now, now.String(), false},
		scenario{"!=", now, zeroTime.String(), true},

		scenario{"<", nil, nil, false},
		scenario{"<", nil, false, true},
		scenario{"<", nil, true, true},
		scenario{"<", nil, 0, true},
		scenario{"<", nil, 1, true},
		scenario{"<", nil, 0.0, true},
		scenario{"<", nil, 0.1, true},
		scenario{"<", nil, "", true},
		scenario{"<", nil, "something", true},
		scenario{"<", nil, time.Time{}, true},
		scenario{"<", nil, now, true},
		scenario{"<", false, nil, false},
		scenario{"<", false, false, false},
		scenario{"<", false, true, true},
		scenario{"<", false, 0, false},
		scenario{"<", false, 1, true},
		scenario{"<", false, 0.0, false},
		scenario{"<", false, 0.1, true},
		scenario{"<", false, "false", false},
		scenario{"<", false, "true", true},
		scenario{"<", false, "something", false},
		scenario{"<", false, time.Time{}, false},
		scenario{"<", false, now, false},
		scenario{"<", 0, 0, false},
		scenario{"<", 0, 1, true},
		scenario{"<", 0, 0.0, false},
		scenario{"<", 0, 0.1, true},
		scenario{"<", 0, "0", false},
		scenario{"<", 0, "0.0", false},
		scenario{"<", 0, "0.1", true},
		scenario{"<", 0, time.Time{}, false},
		scenario{"<", 0, now, false},
		scenario{"<", now, now, false},
		scenario{"<", zeroTime, now, true},
		scenario{"<", now, now.String(), false},
		scenario{"<", zeroTime.String(), now, true},

		scenario{"<=", nil, nil, true},
		scenario{"<=", nil, false, true},
		scenario{"<=", nil, true, true},
		scenario{"<=", nil, 0, true},
		scenario{"<=", nil, 1, true},
		scenario{"<=", nil, 0.0, true},
		scenario{"<=", nil, 0.1, true},
		scenario{"<=", nil, "", true},
		scenario{"<=", nil, "something", true},
		scenario{"<=", nil, time.Time{}, true},
		scenario{"<=", nil, now, true},
		scenario{"<=", false, nil, false},
		scenario{"<=", false, false, true},
		scenario{"<=", false, true, true},
		scenario{"<=", false, 0, true},
		scenario{"<=", false, 1, true},
		scenario{"<=", false, 0.0, true},
		scenario{"<=", false, 0.1, true},
		scenario{"<=", false, "false", true},
		scenario{"<=", false, "true", true},
		scenario{"<=", false, "something", false},
		scenario{"<=", false, time.Time{}, false},
		scenario{"<=", false, now, false},
		scenario{"<=", 0, 0, true},
		scenario{"<=", 0, 1, true},
		scenario{"<=", 0, 0.0, true},
		scenario{"<=", 0, 0.1, true},
		scenario{"<=", 0, "0", true},
		scenario{"<=", 0, "0.0", true},
		scenario{"<=", 0, "0.1", true},
		scenario{"<=", 0, time.Time{}, false},
		scenario{"<=", 0, now, false},
		scenario{"<=", now, now, true},
		scenario{"<=", zeroTime, now, true},
		scenario{"<=", now, now.String(), true},
		scenario{"<=", zeroTime.String(), now, true},
	}

	// For == and !=, add the equivalent operators
	var equivalentScenarios []scenario
	for _, s := range scenarios {
		if s[0] == "==" {
			equivalentScenarios = append(equivalentScenarios, scenario{"=", s[1], s[2], s[3]})
		} else if s[0] == "!=" {
			equivalentScenarios = append(equivalentScenarios, scenario{"<>", s[1], s[2], s[3]})
		}
	}
	scenarios = append(scenarios, equivalentScenarios...)

	// Add the inverse of the scenarios
	var inverseScenarios []scenario
	for _, s := range scenarios {
		if s[0] == "==" || s[0] == "!=" || s[0] == "=" || s[0] == "<>" {
			inverseScenarios = append(inverseScenarios, scenario{s[0], s[2], s[1], s[3]})
		} else if s[0] == "<" {
			inverseScenarios = append(inverseScenarios, scenario{">", s[2], s[1], s[3]})
		}
	}
	scenarios = append(scenarios, inverseScenarios...)

	for _, scenario := range scenarios {
		params := MapParams{"a": scenario[1]}
		e, err := Binary(scenario[0].(string), Param("a"), Constant(scenario[2]))
		if assert.NoError(t, err, "Unable to create Binary expression for %v", scenario) {
			e = msgpacked(t, e)
			if !assert.Equal(t, scenario[3], e.Eval(params), "Evaluation failed for %v", scenario) {
				t.Logf("%v : %v", e.(*binaryExpr).Left, e.(*binaryExpr).Right)
			}
		}
	}
}

func TestLIKE(t *testing.T) {
	scenarios := []scenario{
		scenario{"LIKE", "ABc", "%a%", true},
		scenario{"LIKE", "ABc", "a%", true},
		scenario{"LIKE", "ABc", "%a", false},
		scenario{"LIKE", "ABc", "a", false},
		scenario{"LIKE", "ABc", "%b%", true},
		scenario{"LIKE", "ABc", "b%", false},
		scenario{"LIKE", "ABc", "%b", false},
		scenario{"LIKE", "ABc", "b", false},
		scenario{"LIKE", "ABc", "%c%", true},
		scenario{"LIKE", "ABc", "c%", false},
		scenario{"LIKE", "ABc", "%c", true},
		scenario{"LIKE", "ABc", "c", false},
		scenario{"LIKE", "ABc", "%ab%", true},
		scenario{"LIKE", "ABc", "ab%", true},
		scenario{"LIKE", "ABc", "%ab", false},
		scenario{"LIKE", "ABc", "ab", false},
		scenario{"LIKE", "ABc", "%bc%", true},
		scenario{"LIKE", "ABc", "bc%", false},
		scenario{"LIKE", "ABc", "%bc", true},
		scenario{"LIKE", "ABc", "bc", false},
		scenario{"LIKE", "ABc", "%abc%", true},
		scenario{"LIKE", "ABc", "abc%", true},
		scenario{"LIKE", "ABc", "%abc", true},
		scenario{"LIKE", "ABc", "ABc", true},
		scenario{"LIKE", "ABc", "%abcd%", false},
		scenario{"LIKE", "ABc", "abcd%", false},
		scenario{"LIKE", "ABc", "%abcd", false},
		scenario{"LIKE", "ABc", "abcd", false},
		scenario{"LIKE", "ABc", "%", true},
		scenario{"LIKE", "ABc", nil, false},
		scenario{"LIKE", nil, "abcd", false},
	}

	for _, scenario := range scenarios {
		params := MapParams{"a": scenario[1]}
		e, err := Binary(scenario[0].(string), Param("a"), Constant(scenario[2]))
		if assert.NoError(t, err, "Unable to create Binary expression for %v", scenario) {
			assert.Equal(t, scenario[3], e.Eval(params), "Evaluation failed for %v", scenario)
			ne, err := Binary("NOT "+scenario[0].(string), Param("a"), Constant(scenario[2]))
			if assert.NoError(t, err, "Unable to create NOT Binary expression for %v", scenario) {
				assert.NotEqual(t, scenario[3], ne.Eval(params), "Evaluation failed for %v", scenario)
			}
		}
	}
}

func TestANDOR(t *testing.T) {
	scenarios := []scenario{
		scenario{"AND", true, true, true},
		scenario{"AND", false, false, false},
		scenario{"AND", false, true, false},
		scenario{"AND", true, false, false},
		scenario{"AND", 5, true, false},
		scenario{"AND", true, 5, false},
		scenario{"OR", true, true, true},
		scenario{"OR", false, false, false},
		scenario{"OR", false, true, true},
		scenario{"OR", true, false, true},
		scenario{"OR", 5, true, true},
		scenario{"OR", true, 5, true},
	}

	for _, scenario := range scenarios {
		params := MapParams{"a": scenario[1]}
		e, err := Binary(scenario[0].(string), Param("a"), Constant(scenario[2]))
		e = msgpacked(t, e)
		if assert.NoError(t, err, "Unable to create Binary expression for %v", scenario) {
			assert.Equal(t, scenario[3], e.Eval(params), "Evaluation failed for %v", scenario)
			assert.NotEqual(t, scenario[3].(bool), Not(e).Eval(params), "Evaluation failed for Not of %v", scenario)
		}
	}
}

func TestBinaryCalculations(t *testing.T) {
	scenarios := []scenario{
		scenario{"+", 4, 2, 6},
		scenario{"+", 4, "2", 6},
		scenario{"+", "4", 2, 6},
		scenario{"+", nil, 2, nil},
		scenario{"+", 4, nil, nil},
		scenario{"+", 4, "string", nil},

		scenario{"-", 4, 2, 2},
		scenario{"-", 4, "2", 2},
		scenario{"-", "4", 2, 2},
		scenario{"-", nil, 2, nil},
		scenario{"-", 4, nil, nil},
		scenario{"-", 4, "string", nil},

		scenario{"*", 4, 2, 8},
		scenario{"*", 4, "2", 8},
		scenario{"*", "4", 2, 8},
		scenario{"*", nil, 2, nil},
		scenario{"*", 4, nil, nil},
		scenario{"*", 4, "string", nil},

		scenario{"/", 4, 2, 2},
		scenario{"/", 4, "2", 2},
		scenario{"/", "4", 2, 2},
		scenario{"/", 4, 0, nil},
		scenario{"/", nil, 2, nil},
		scenario{"/", 4, nil, nil},
		scenario{"/", 4, "string", nil},
	}

	for _, scenario := range scenarios {
		params := MapParams{"a": scenario[1]}
		e, err := Binary(scenario[0].(string), Param("a"), Constant(scenario[2]))
		if assert.NoError(t, err, "Unable to create Binary expression for %v", scenario) {
			assert.EqualValues(t, scenario[3], e.Eval(params), "Evaluation failed for %v", scenario)
		}
	}
}
