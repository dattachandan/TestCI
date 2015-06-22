// Copyright (c) 2013 Chandan Datta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

package TestCI

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestIsItTested(t *testing.T) {
	word := "Tested"
	foo, err := IsItTestedOrNot(word)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, foo)
}
