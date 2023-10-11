package licensemanager

import (
	"testing"
)

func TestLicenseManager(t *testing.T) {
	t.Run("Acquire", func(t *testing.T) {
		lm := NewLicenseManager(2)

		if err := lm.Acquire(); err != nil {
			t.Errorf("failed to acquire license: %v", err)
		}

		if lm.UsedLicenses() != 1 {
			t.Errorf("expected 1 used license, got %d", lm.UsedLicenses())
		}
	})

	t.Run("Release", func(t *testing.T) {
		lm := NewLicenseManager(2)
		_ = lm.Acquire()

		if err := lm.Release(); err != nil {
			t.Errorf("failed to release license: %v", err)
		}

		if lm.UsedLicenses() != 0 {
			t.Errorf("expected 0 used license, got %d", lm.UsedLicenses())
		}
	})

	t.Run("Exceed", func(t *testing.T) {
		lm := NewLicenseManager(1)
		_ = lm.Acquire()

		if err := lm.Acquire(); err == nil {
			t.Errorf("expected error when exceeding licenses, got nil")
		}
	})

	t.Run("Release Without Acquire", func(t *testing.T) {
		lm := NewLicenseManager(1)

		if err := lm.Release(); err == nil {
			t.Errorf("expected error when releasing without acquiring, got nil")
		}
	})
}
