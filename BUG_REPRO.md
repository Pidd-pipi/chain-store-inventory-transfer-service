# BUG 复现说明（store-inventory-transfer__002）

## Bug 是什么
低库存预警与补货建议的边界计算错误：
- 库存数量刚好等于安全库存时，被错误标成“低库存”。
- 缺货商品的建议补货数量明显偏少（本应补至安全库存的 1.5 倍）。
- 滞销商品判断方向反了，库存充足但销量极少的商品被误判为快消。

## 如何触发
```bash
go test ./internal/service -run 'TestInventoryStatusBoundary|TestCalculateSuggestQtyBoundary|TestIsSlowMovingBoundary|TestReplenishBoundary'
```

## 错误信息
```
--- FAIL: TestInventoryStatusBoundary
    replenish_boundary_test.go:72: InventoryStatusText(50,50) = 低库存, want 正常
--- FAIL: TestCalculateSuggestQtyBoundary
    replenish_boundary_test.go:78: CalculateSuggestQty(20,80) = 60, want 100
--- FAIL: TestIsSlowMovingBoundary
    replenish_boundary_test.go:84: IsSlowMoving(500,10) = false, want true
--- FAIL: TestReplenishBoundary
    replenish_boundary_test.go:98: SuggestQty = 0, want 100
```
