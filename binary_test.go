package goexpr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type scenario [4]interface{}

func TestBinaryComparisons(t *testing.T) {
	now := time.Now()
	vals := [][]interface{}{
		[]interface{}{nil, 1},
		[]interface{}{"false", "true"},
		[]interface{}{1, 2},
		[]interface{}{"x", "y"},
		[]interface{}{now, now.Add(5 * time.Second)},
	}
	var scenarios []scenario
	for _, vals := range vals {
		scenarios = append(scenarios, scenario{"=", vals[0], vals[0], true})
		scenarios = append(scenarios, scenario{"=", vals[0], vals[1], false})
		scenarios = append(scenarios, scenario{"=", vals[1], vals[0], false})
		scenarios = append(scenarios, scenario{"==", vals[0], vals[0], true})
		scenarios = append(scenarios, scenario{"==", vals[0], vals[1], false})
		scenarios = append(scenarios, scenario{"==", vals[1], vals[0], false})
		scenarios = append(scenarios, scenario{"LIKE", vals[0], vals[0], true})
		scenarios = append(scenarios, scenario{"LIKE", vals[0], vals[1], false})
		scenarios = append(scenarios, scenario{"LIKE", vals[1], vals[0], false})
		scenarios = append(scenarios, scenario{"<>", vals[0], vals[0], false})
		scenarios = append(scenarios, scenario{"<>", vals[0], vals[1], true})
		scenarios = append(scenarios, scenario{"<>", vals[1], vals[0], true})
		scenarios = append(scenarios, scenario{"!=", vals[0], vals[0], false})
		scenarios = append(scenarios, scenario{"!=", vals[0], vals[1], true})
		scenarios = append(scenarios, scenario{"!=", vals[1], vals[0], true})
		scenarios = append(scenarios, scenario{"<=", vals[0], vals[0], true})
		scenarios = append(scenarios, scenario{"<=", vals[0], vals[1], true})
		scenarios = append(scenarios, scenario{"<=", vals[1], vals[0], false})
		scenarios = append(scenarios, scenario{">=", vals[0], vals[0], true})
		scenarios = append(scenarios, scenario{">=", vals[0], vals[1], false})
		scenarios = append(scenarios, scenario{">=", vals[1], vals[0], true})
		scenarios = append(scenarios, scenario{"<", vals[0], vals[0], false})
		scenarios = append(scenarios, scenario{"<", vals[0], vals[1], true})
		scenarios = append(scenarios, scenario{"<", vals[1], vals[0], false})
		scenarios = append(scenarios, scenario{">", vals[0], vals[0], false})
		scenarios = append(scenarios, scenario{">", vals[0], vals[1], false})
		scenarios = append(scenarios, scenario{">", vals[1], vals[0], true})
	}

	for _, scenario := range scenarios {
		params := MapParams{"a": scenario[1]}
		e, err := Binary(scenario[0].(string), Param("a"), Constant(scenario[2]))
		if assert.NoError(t, err, "Unable to create Binary expression for %v", scenario) {
			assert.Equal(t, scenario[3], e.Eval(params), "Evaluation failed for %v", scenario)
		}
	}
}

func TestLIKE(t *testing.T) {
	scenarios := []scenario{
		scenario{"LIKE", "abc", "a", true},
		scenario{"LIKE", "abc", "b", true},
		scenario{"LIKE", "abc", "c", true},
		scenario{"LIKE", "abc", "ab", true},
		scenario{"LIKE", "abc", "bc", true},
		scenario{"LIKE", "abc", "abc", true},
		scenario{"LIKE", "abc", "abcd", false},
		scenario{"LIKE", "abc", nil, false},
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
		scenario{"OR", true, true, true},
		scenario{"OR", false, false, false},
		scenario{"OR", false, true, true},
		scenario{"OR", true, false, true},
	}

	for _, scenario := range scenarios {
		params := MapParams{"a": scenario[1]}
		e, err := Binary(scenario[0].(string), Param("a"), Constant(scenario[2]))
		if assert.NoError(t, err, "Unable to create Binary expression for %v", scenario) {
			assert.Equal(t, scenario[3], e.Eval(params), "Evaluation failed for %v", scenario)
			assert.NotEqual(t, scenario[3].(bool), Not(e).Eval(params), "Evaluation failed for Not of %v", scenario)
		}
	}
}

func TestBinaryCalculations(t *testing.T) {
	scenarios := []scenario{
		scenario{"+", 4, 2, 6},
		scenario{"+", nil, 2, 0},
		scenario{"+", 4, nil, 0},
		scenario{"-", 4, 2, 2},
		scenario{"-", nil, 2, 0},
		scenario{"-", 4, nil, 0},
		scenario{"*", 4, 2, 8},
		scenario{"*", nil, 2, 0},
		scenario{"*", 4, nil, 0},
		scenario{"/", 4, 2, 2},
		scenario{"/", nil, 2, 0},
		scenario{"/", 4, nil, 0},
		scenario{"/", 4, 0, 0},
	}

	for _, scenario := range scenarios {
		params := MapParams{"a": scenario[1]}
		e, err := Binary(scenario[0].(string), Param("a"), Constant(scenario[2]))
		if assert.NoError(t, err, "Unable to create Binary expression for %v", scenario) {
			assert.Equal(t, scenario[3], e.Eval(params), "Evaluation failed for %v", scenario)
		}
	}
}
