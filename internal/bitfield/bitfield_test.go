
package bitfield

import "testing"

func TestHighBit(t *testing.T) {
	b := New(128)
	b.Set(70)
	if !b.Get(70) {
		t.Fatal("bit 70 not set")
	}
	if b.Get(0) {
		t.Fatal("bit 0 should be clear")
	}
}
