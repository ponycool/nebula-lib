package main

import (
	"github.com/ponycool/nebula-lib/uuid"
	"testing"
)

func TestUuid(t *testing.T) {
	if ans := uuid.GetUuid(); !uuid.Check(ans) {
		t.Errorf("expected be uuid string, but %s got", ans)
	}
}
