package goexpr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type scenario [4]interface{}

func TestBinaryComparisons(t *testing.T) {
	now := time.Now()
	scenarios := []scenario{
		scenario{"=", nil, nil, true},
		scenario{"=", nil, false, true},
		scenario{"=", nil, true, false},
		scenario{"=", nil, 0, true},
		scenario{"=", nil, 1, false},
		scenario{"=", nil, 0.0, true},
		scenario{"=", nil, 0.1, false},
		scenario{"=", nil, "", true},
		scenario{"=", nil, "something", false},
		scenario{"=", nil, time.Time{}, true},
		scenario{"=", nil, now, false},
		scenario{"=", false, nil, true},
		scenario{"=", false, false, true},
		scenario{"=", false, true, false},
		scenario{"=", false, 0, true},
		scenario{"=", false, 1, false},
		scenario{"=", false, 0.0, true},
		scenario{"=", false, 0.1, false},
		scenario{"=", false, "", true},
		scenario{"=", false, "something", false},
		scenario{"=", false, time.Time{}, true},
		scenario{"=", false, now, false},
		scenario{"=", 0, nil, true},
		scenario{"=", 0, false, true},
		scenario{"=", 0, true, false},
		scenario{"=", 0, 0, true},
		scenario{"=", 0, 1, false},
		scenario{"=", 0, 0.0, true},
		scenario{"=", 0, 0.1, false},
		scenario{"=", 0, "0", true},
		scenario{"=", 0, "0.0", true},
		scenario{"=", 0, time.Time{}, false},
		scenario{"=", 0, now, false},
		scenario{"=", now.String(), now, true},

		scenario{"!=", nil, nil, false},
		scenario{"!=", nil, false, false},
		scenario{"!=", nil, true, true},
		scenario{"!=", nil, 0, false},
		scenario{"!=", nil, 1, true},
		scenario{"!=", nil, 0.0, false},
		scenario{"!=", nil, 0.1, true},
		scenario{"!=", nil, "", false},
		scenario{"!=", nil, "something", true},
		scenario{"!=", nil, time.Time{}, false},
		scenario{"!=", nil, now, true},
		scenario{"!=", false, nil, false},
		scenario{"!=", false, false, false},
		scenario{"!=", false, true, true},
		scenario{"!=", false, 0, false},
		scenario{"!=", false, 1, true},
		scenario{"!=", false, 0.0, false},
		scenario{"!=", false, 0.1, true},
		scenario{"!=", false, "", false},
		scenario{"!=", false, "something", true},
		scenario{"!=", false, time.Time{}, false},
		scenario{"!=", false, now, true},
		scenario{"!=", 0, nil, false},
		scenario{"!=", 0, false, false},
		scenario{"!=", 0, true, true},
		scenario{"!=", 0, 0, false},
		scenario{"!=", 0, 1, true},
		scenario{"!=", 0, 0.0, false},
		scenario{"!=", 0, 0.1, true},
		scenario{"!=", 0, "0", false},
		scenario{"!=", 0, "0.0", false},
		scenario{"!=", 0, time.Time{}, true},
		scenario{"!=", 0, now, true},
		scenario{"!=", now.String(), now, false},

		scenario{"<", nil, nil, false},
		scenario{"<", nil, false, false},
		scenario{"<", nil, true, true},
		scenario{"<", nil, 0, false},
		scenario{"<", nil, 1, true},
		scenario{"<", nil, 0.0, false},
		scenario{"<", nil, 0.1, true},
		scenario{"<", nil, "", false},
		scenario{"<", nil, "something", true},
		scenario{"<", nil, time.Time{}, false},
		scenario{"<", nil, now, true},
		scenario{"<", false, nil, true},
		scenario{"<", false, false, true},
		scenario{"<", false, true, false},
		scenario{"<", false, 0, true},
		scenario{"<", false, 1, false},
		scenario{"<", false, 0.0, true},
		scenario{"<", false, 0.1, false},
		scenario{"<", false, "", true},
		scenario{"<", false, "something", false},
		scenario{"<", false, time.Time{}, true},
		scenario{"<", false, now, false},
		scenario{"<", 0, nil, true},
		scenario{"<", 0, false, true},
		scenario{"<", 0, true, false},
		scenario{"<", 0, 0, true},
		scenario{"<", 0, 1, false},
		scenario{"<", 0, 0.0, true},
		scenario{"<", 0, 0.1, false},
		scenario{"<", 0, "0", true},
		scenario{"<", 0, "0.0", true},
		scenario{"<", 0, time.Time{}, false},
		scenario{"<", 0, now, false},
		scenario{"<", now.String(), now, true},
	}
	// vals := [][]interface{}{
	// 	[]interface{}{nil, 1},
	// 	[]interface{}{false, true},
	// 	[]interface{}{1, 2},
	// 	[]interface{}{"x", "y"},
	// 	[]interface{}{now, now.Add(5 * time.Second)},
	// }
	// var scenarios []scenario
	// for _, vals := range vals {
	// 	scenarios = append(scenarios, scenario{"=", vals[0], vals[0], true})
	// 	scenarios = append(scenarios, scenario{"=", vals[0], vals[1], false})
	// 	scenarios = append(scenarios, scenario{"=", vals[1], vals[0], false})
	// 	scenarios = append(scenarios, scenario{"==", vals[0], vals[0], true})
	// 	scenarios = append(scenarios, scenario{"==", vals[0], vals[1], false})
	// 	scenarios = append(scenarios, scenario{"==", vals[1], vals[0], false})
	// 	scenarios = append(scenarios, scenario{"LIKE", vals[0], vals[0], true})
	// 	scenarios = append(scenarios, scenario{"LIKE", vals[0], vals[1], false})
	// 	scenarios = append(scenarios, scenario{"LIKE", vals[1], vals[0], false})
	// 	scenarios = append(scenarios, scenario{"<>", vals[0], vals[0], false})
	// 	scenarios = append(scenarios, scenario{"<>", vals[0], vals[1], true})
	// 	scenarios = append(scenarios, scenario{"<>", vals[1], vals[0], true})
	// 	scenarios = append(scenarios, scenario{"!=", vals[0], vals[0], false})
	// 	scenarios = append(scenarios, scenario{"!=", vals[0], vals[1], true})
	// 	scenarios = append(scenarios, scenario{"!=", vals[1], vals[0], true})
	// 	scenarios = append(scenarios, scenario{"<=", vals[0], vals[0], true})
	// 	scenarios = append(scenarios, scenario{"<=", vals[0], vals[1], true})
	// 	scenarios = append(scenarios, scenario{"<=", vals[1], vals[0], false})
	// 	scenarios = append(scenarios, scenario{">=", vals[0], vals[0], true})
	// 	scenarios = append(scenarios, scenario{">=", vals[0], vals[1], false})
	// 	scenarios = append(scenarios, scenario{">=", vals[1], vals[0], true})
	// 	scenarios = append(scenarios, scenario{"<", vals[0], vals[0], false})
	// 	scenarios = append(scenarios, scenario{"<", vals[0], vals[1], true})
	// 	scenarios = append(scenarios, scenario{"<", vals[1], vals[0], false})
	// 	scenarios = append(scenarios, scenario{">", vals[0], vals[0], false})
	// 	scenarios = append(scenarios, scenario{">", vals[0], vals[1], false})
	// 	scenarios = append(scenarios, scenario{">", vals[1], vals[0], true})
	// }

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
		scenario{"+", 4, "string", 0},
		scenario{"-", 4, 2, 2},
		scenario{"-", nil, 2, 0},
		scenario{"-", 4, nil, 0},
		scenario{"-", 4, "string", 0},
		scenario{"*", 4, 2, 8},
		scenario{"*", nil, 2, 0},
		scenario{"*", 4, nil, 0},
		scenario{"*", 4, "string", 0},
		scenario{"/", 4, 2, 2},
		scenario{"/", nil, 2, 0},
		scenario{"/", 4, nil, 0},
		scenario{"/", 4, 0, 0},
		scenario{"/", 4, "string", 0},
	}

	for _, scenario := range scenarios {
		params := MapParams{"a": scenario[1]}
		e, err := Binary(scenario[0].(string), Param("a"), Constant(scenario[2]))
		if assert.NoError(t, err, "Unable to create Binary expression for %v", scenario) {
			assert.Equal(t, scenario[3], e.Eval(params), "Evaluation failed for %v", scenario)
		}
	}
}
