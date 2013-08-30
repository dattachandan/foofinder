// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package foofinder

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestIsItFoo(t *testing.T) {
	word := "foo"
	foo, err := IsItFoo(word)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, foo)
}
