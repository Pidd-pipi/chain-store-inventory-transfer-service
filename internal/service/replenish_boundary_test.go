package service

import (
	"log/slog"
	"os"
	"testing"

	"gorm.io/gorm"

	"github.com/ld/storeinventory/internal/constants"
	"github.com/ld/storeinventory/internal/model"
	"github.com/ld/storeinventory/internal/util"
)

type replenishInvRepo struct {
	alerts []model.StoreInventory
}

func (m *replenishInvRepo) Create(inv *model.StoreInventory) error { return nil }
func (m *replenishInvRepo) CreateTx(tx *gorm.DB, inv *model.StoreInventory) error { return nil }
func (m *replenishInvRepo) FindByStoreAndSKU(storeID, skuID uint) (*model.StoreInventory, error) {
	return nil, util.ErrNotFound
}
func (m *replenishInvRepo) FindByStoreAndSKUTx(tx *gorm.DB, storeID, skuID uint) (*model.StoreInventory, error) {
	return nil, util.ErrNotFound
}
func (m *replenishInvRepo) FindByID(id uint) (*model.StoreInventory, error) { return nil, util.ErrNotFound }
func (m *replenishInvRepo) List(page, pageSize int, storeID, skuID uint) ([]model.StoreInventory, int64, error) {
	return m.alerts, int64(len(m.alerts)), nil
}
func (m *replenishInvRepo) ListAlerts() ([]model.StoreInventory, error) { return m.alerts, nil }
func (m *replenishInvRepo) Update(inv *model.StoreInventory) error     { return nil }
func (m *replenishInvRepo) AdjustQuantity(storeID, skuID uint, delta int) error { return nil }
func (m *replenishInvRepo) AdjustQuantityTx(tx *gorm.DB, storeID, skuID uint, delta int) error {
	return nil
}

type replenishStoreRepo struct{}

func (m *replenishStoreRepo) Create(store *model.Store) error { return nil }
func (m *replenishStoreRepo) FindByID(id uint) (*model.Store, error) { return &model.Store{ID: id}, nil }
func (m *replenishStoreRepo) FindByCode(code string) (*model.Store, error) { return nil, util.ErrNotFound }
func (m *replenishStoreRepo) List(page, pageSize int) ([]model.Store, int64, error) { return nil, 0, nil }
func (m *replenishStoreRepo) ListAll() ([]model.Store, error) {
	return []model.Store{{ID: 1, Name: "北京朝阳门店"}}, nil
}
func (m *replenishStoreRepo) Update(store *model.Store) error { return nil }
func (m *replenishStoreRepo) Delete(id uint) error            { return nil }

type replenishRecordRepo struct{}

func (m *replenishRecordRepo) Create(record *model.StockRecord) error { return nil }
func (m *replenishRecordRepo) CreateTx(tx *gorm.DB, record *model.StockRecord) error { return nil }
func (m *replenishRecordRepo) List(page, pageSize int, storeID, skuID uint, recordType constants.StockRecordType) ([]model.StockRecord, int64, error) {
	return nil, 0, nil
}
func (m *replenishRecordRepo) CreateStocktake(st *model.Stocktake) error { return nil }
func (m *replenishRecordRepo) CreateStocktakeTx(tx *gorm.DB, st *model.Stocktake) error { return nil }
func (m *replenishRecordRepo) ListStocktakes(page, pageSize int, storeID uint) ([]model.Stocktake, int64, error) {
	return nil, 0, nil
}
func (m *replenishRecordRepo) MonthlySalesBySKU(storeID, skuID uint) (int, error) { return 0, nil }

func newReplenishRecordService() StockRecordService {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	invRepo := &replenishInvRepo{alerts: []model.StoreInventory{{ID: 1, StoreID: 1, SKUID: 10, Quantity: 20, SafetyStock: 80}}}
	return NewStockRecordService(&replenishRecordRepo{}, invRepo, &replenishStoreRepo{}, &mockSKURepo{}, nil, nil, logger)
}

func TestInventoryStatusBoundary(t *testing.T) {
	if got := util.InventoryStatusText(50, 50); got != "正常" {
		t.Fatalf("InventoryStatusText(50,50) = %s, want 正常", got)
	}
}

func TestCalculateSuggestQtyBoundary(t *testing.T) {
	if got := util.CalculateSuggestQty(20, 80); got != 100 {
		t.Fatalf("CalculateSuggestQty(20,80) = %d, want 100", got)
	}
}

func TestIsSlowMovingBoundary(t *testing.T) {
	if got := util.IsSlowMoving(500, 10); !got {
		t.Fatalf("IsSlowMoving(500,10) = false, want true")
	}
}

func TestReplenishBoundary(t *testing.T) {
	svc := newReplenishRecordService()
	suggestions, err := svc.ReplenishSuggestions()
	if err != nil {
		t.Fatalf("ReplenishSuggestions failed: %v", err)
	}
	if len(suggestions) != 1 {
		t.Fatalf("suggestions len = %d, want 1", len(suggestions))
	}
	if suggestions[0].SuggestQty != 100 {
		t.Fatalf("SuggestQty = %d, want 100", suggestions[0].SuggestQty)
	}
	if !suggestions[0].SlowMoving {
		t.Fatalf("SlowMoving = false, want true")
	}
}
