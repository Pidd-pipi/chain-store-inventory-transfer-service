package service

import (
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/ld/storeinventory/internal/model"
	"github.com/ld/storeinventory/internal/repository"
	"github.com/ld/storeinventory/internal/util"
)

type conflictStoreRepo struct{}

func (m *conflictStoreRepo) Create(store *model.Store) error { return repository.ErrDuplicate }
func (m *conflictStoreRepo) FindByID(id uint) (*model.Store, error) { return nil, util.ErrNotFound }
func (m *conflictStoreRepo) FindByCode(code string) (*model.Store, error) { return nil, util.ErrNotFound }
func (m *conflictStoreRepo) List(page, pageSize int) ([]model.Store, int64, error) { return nil, 0, nil }
func (m *conflictStoreRepo) ListAll() ([]model.Store, error) { return nil, nil }
func (m *conflictStoreRepo) Update(store *model.Store) error { return nil }
func (m *conflictStoreRepo) Delete(id uint) error           { return nil }

type conflictSKURepo struct{}

func (m *conflictSKURepo) Create(sku *model.SKU) error { return repository.ErrDuplicate }
func (m *conflictSKURepo) FindByID(id uint) (*model.SKU, error) { return nil, util.ErrNotFound }
func (m *conflictSKURepo) FindByCode(code string) (*model.SKU, error) { return nil, util.ErrNotFound }
func (m *conflictSKURepo) List(page, pageSize int, category, keyword string) ([]model.SKU, int64, error) {
	return nil, 0, nil
}
func (m *conflictSKURepo) BatchCreate(skus []*model.SKU) (int, error) { return 0, nil }
func (m *conflictSKURepo) Update(sku *model.SKU) error                 { return nil }
func (m *conflictSKURepo) Delete(id uint) error                        { return nil }

func newConflictLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
}

func TestStoreCreateDuplicateReturnsConflict(t *testing.T) {
	svc := NewStoreService(&conflictStoreRepo{}, newConflictLogger())
	_, err := svc.Create("S1", "门店", "", nil)
	if err == nil {
		t.Fatal("StoreService.Create duplicate should return error")
	}
	if !errors.Is(err, util.ErrConflict) {
		t.Fatalf("err = %v, want errors.Is util.ErrConflict", err)
	}
}

func TestSKUCreateDuplicateReturnsConflict(t *testing.T) {
	svc := NewSKUService(&conflictSKURepo{}, newConflictLogger())
	_, err := svc.Create("SKU-1", "商品", "", "", "", "")
	if err == nil {
		t.Fatal("SKUService.Create duplicate should return error")
	}
	if !errors.Is(err, util.ErrConflict) {
		t.Fatalf("err = %v, want errors.Is util.ErrConflict", err)
	}
}
