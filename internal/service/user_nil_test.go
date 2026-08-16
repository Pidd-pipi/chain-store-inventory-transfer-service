package service

import (
	"testing"

	"github.com/ld/storeinventory/internal/constants"
	"github.com/ld/storeinventory/internal/util"
)

func TestAdminRegisterAllowsNilStore(t *testing.T) {
	svc := newTestUserService()
	u, err := svc.Register("boss", "123456", "老板", constants.RoleAdmin, nil)
	if err != nil {
		t.Fatalf("admin register should allow nil store, got err: %v", err)
	}
	if u.StoreID != nil {
		t.Fatalf("admin StoreID = %v, want nil", *u.StoreID)
	}
}

func TestUpdateProfileKeepsStoreOnNil(t *testing.T) {
	svc := newTestUserService()
	storeID := uint(7)
	u, err := svc.Register("mgr", "123456", "店长", constants.RoleStoreManager, &storeID)
	if err != nil {
		t.Fatalf("register failed: %v", err)
	}
	updated, err := svc.UpdateProfile(u.ID, "新名字", nil)
	if err != nil {
		t.Fatalf("UpdateProfile failed: %v", err)
	}
	if updated.StoreID == nil || *updated.StoreID != 7 {
		t.Fatalf("StoreID changed to %v, want 7", updated.StoreID)
	}
}

func TestGenerateTokenAllowsNilStore(t *testing.T) {
	token, err := util.GenerateToken("secret", 1, 1, "boss", constants.RoleAdmin, nil)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}
	if token == "" {
		t.Fatal("token is empty")
	}
}
