package vars

import (
	"testing"
	"time"
)

func TestNewCacheTtlPeriod(t *testing.T) {
	if got := GetTtlShortPeriod(); got != 3*time.Hour {
		t.Errorf("GetTtlShortPeriod() = %v, want 3h", got)
	}

	if got := GetTtlMediumPeriod(); got != 24*time.Hour {
		t.Errorf("GetTtlMediumPeriod() = %v, want 24h", got)
	}

	if got := GetTtlLongPeriod(); got != 72*time.Hour {
		t.Errorf("GetTtlLongPeriod() = %v, want 72h", got)
	}
}
