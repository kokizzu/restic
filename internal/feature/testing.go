package feature

import (
	"fmt"
	"testing"
)

// TestSetFlag temporarily sets a feature flag to the given value until the
// returned function is called.
//
// Usage
// ```
// defer TestSetFlag(t, features.Flags, features.ExampleFlag, true)()
// ```
func TestSetFlag(t *testing.T, f *FlagSet, flag FlagName, value bool) func() {
	current := f.Enabled(flag)

	if err := f.Apply(fmt.Sprintf("%s=%v", flag, value)); err != nil {
		// not reachable
		panic(err)
	}

	return func() {
		if err := f.Apply(fmt.Sprintf("%s=%v", flag, current)); err != nil {
			// not reachable
			panic(err)
		}
	}
}
