package service

import (
	"log/slog"
	"os"
	"testing"


	"github.com/ld/storeinventory/internal/model"
	"github.com/ld/storeinventory/internal/util"
)

type batchMockSKURepo struct {
	created []*model.SKU
}

func (m *batchMockSKURepo) Create(sku *model.SKU) error { return nil }
func (m *batchMockSKURepo) FindByID(id uint) (*model.SKU, error) { return &model.SKU{ID: id}, nil }
func (m *batchMockSKURepo) FindByCode(code string) (*model.SKU, error) { return nil, util.ErrNotFound }
func (m *batchMockSKURepo) List(page, pageSize int, category, keyword string) ([]model.SKU, int64, error) {
	return nil, 0, nil
}
func (m *batchMockSKURepo) BatchCreate(skus []*model.SKU) (int, error) {
	m.created = append([]*model.SKU(nil), skus...)
	return len(skus), nil
}
func (m *batchMockSKURepo) Update(sku *model.SKU) error { return nil }
func (m *batchMockSKURepo) Delete(id uint) error        { return nil }

func newBatchSKUService(repo *batchMockSKURepo) SKUService {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	return NewSKUService(repo, logger)
}

func TestBatchImportKeepsDistinctCodes(t *testing.T) {
	repo := &batchMockSKURepo{}
	svc := newBatchSKUService(repo)
	items := []SKUImportItem{
		{Code: "SKU-A", Name: "A"},
		{Code: "SKU-B", Name: "B"},
		{Code: "SKU-C", Name: "C"},
	}
	n, err := svc.BatchImport(items)
	if err != nil {
		t.Fatalf("BatchImport failed: %v", err)
	}
	if n != 3 {
		t.Fatalf("BatchImport n = %d, want 3", n)
	}
	if len(repo.created) != 3 {
		t.Fatalf("created len = %d, want 3", len(repo.created))
	}
	want := []string{"SKU-A", "SKU-B", "SKU-C"}
	for i, code := range want {
		if repo.created[i].Code != code {
			t.Fatalf("created[%d].Code = %s, want %s", i, repo.created[i].Code, code)
		}
	}
}
