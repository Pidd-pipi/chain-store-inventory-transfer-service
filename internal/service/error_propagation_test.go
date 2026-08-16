package service

import (
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/ld/storeinventory/internal/util"
)

func newSlog() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
}

func TestEnsureCreatesInventoryWhenAbsent(t *testing.T) {
	svc := NewStoreInventoryService(newMockInventoryRepo(), &mockSKURepo{}, nil, newSlog())
	inv, err := svc.Ensure(1, 10, 100)
	if err != nil {
		t.Fatalf("Ensure should create inventory, got err: %v", err)
	}
	if inv == nil || inv.Quantity != 100 {
		t.Fatalf("Ensure inventory = %+v, want quantity 100", inv)
	}
}

func TestCheckSufficientReturnsStockNotEnough(t *testing.T) {
	svc := NewStoreInventoryService(newMockInventoryRepo(), &mockSKURepo{}, nil, newSlog())
	if _, err := svc.Ensure(1, 10, 100); err != nil {
		t.Fatalf("seed ensure failed: %v", err)
	}
	err := svc.CheckSufficient(1, 10, 101)
	if err == nil {
		t.Fatal("CheckSufficient should fail when stock is insufficient")
	}
	if !errors.Is(err, util.ErrStockNotEnough) {
		t.Fatalf("CheckSufficient err = %v, want errors.Is ErrStockNotEnough", err)
	}
}

func TestAsAppErrorMapsStockNotEnough(t *testing.T) {
	app := util.AsAppError(util.ErrStockNotEnough)
	if app == nil {
		t.Fatal("AsAppError returned nil")
	}
	if app.HTTPStatus != 400 {
		t.Fatalf("AsAppError HTTPStatus = %d, want 400", app.HTTPStatus)
	}
	if app.Code != 40000 {
		t.Fatalf("AsAppError Code = %d, want 40000", app.Code)
	}
}
