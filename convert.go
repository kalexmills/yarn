// Copyright 2021 Josh Deprez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yarn

import (
	"fmt"
	"math"
	"strconv"

	yarnpb "github.com/DrJosh9000/yarn/bytecode"
)

// convertToBool attempts conversion of the standard Yarn Spinner VM types
// (bool, number, string, null) to bool, in the same way Yarn Spinner does it
// (see AsBool in Value.cs).
func convertToBool(x interface{}) (bool, error) {
	if x == nil {
		return false, nil
	}
	switch x := x.(type) {
	case bool:
		return x, nil
	case float32:
		return !math.IsNaN(float64(x)) && x != 0, nil
	case float64:
		return !math.IsNaN(x) && x != 0, nil
	case int:
		return x != 0, nil
	case string:
		return x != "", nil
	default:
		return false, fmt.Errorf("%T %w to bool", x, ErrNotConvertible)
	}
}

// convertToInt attempts conversion of an arbitrary value to int. Right now it's
// only used by the VM to count function arguments. But it works in a way that
// should be compatible with how Yarn Spinner does it (see AsNumber in
// Value.cs).
func convertToInt(x interface{}) (int, error) {
	if x == nil {
		return 0, nil
	}
	switch t := x.(type) {
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	case int:
		return t, nil
	case string:
		return strconv.Atoi(t)
	default:
		if t == nil {
			return 0, nil
		}
		return 0, fmt.Errorf("%T %w to int", x, ErrNotConvertible)
	}
}

// convertToString converts a value to a string, in a way that matches what Yarn
// Spinner does. nil becomes "null", and booleans are title-cased.
func convertToString(x interface{}) string {
	if x == nil {
		return "null"
	}
	if x, ok := x.(bool); ok {
		if x {
			return "True"
		}
		return "False"
	}
	return fmt.Sprint(x)
}

// operandToInt is a helper for turning a number value into an int.
func operandToInt(op *yarnpb.Operand) (int, error) {
	if op == nil {
		return 0, ErrNilOperand
	}
	f, ok := op.Value.(*yarnpb.Operand_FloatValue)
	if !ok {
		return 0, fmt.Errorf("%w for operand [%T != Operand_FloatValue]", ErrWrongType, op.Value)
	}
	return int(f.FloatValue), nil
}
