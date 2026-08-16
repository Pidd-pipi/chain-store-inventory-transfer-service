# BUG 复现说明（store-inventory-transfer__001）

## Bug 是什么
门店库存建档与库存不足的错误哨兵传播链断裂：
- 首次给某个门店+SKU 建档库存时，仓库层把“不存在”错误丢失哨兵，服务层又把“不存在”误判成“资源状态冲突”，导致新库存无法建档。
- 库存不足时服务层返回“参数不合法”，而不是“库存不足”哨兵。
- `util.AsAppError` 对库存不足错误映射为 500 内部错误，而不是 400 业务错误。

## 如何触发
```bash
go test ./internal/service -run 'TestEnsureCreatesInventoryWhenAbsent|TestCheckSufficientReturnsStockNotEnough|TestAsAppErrorMapsStockNotEnough' -count=1
```

## 错误信息
```
--- FAIL: TestEnsureCreatesInventoryWhenAbsent
    error_propagation_test.go:20: Ensure should create inventory, got err: ensure inventory store[1] sku[10]: resource conflict
--- FAIL: TestCheckSufficientReturnsStockNotEnough
    error_propagation_test.go:30: seed ensure failed: ensure inventory store[1] sku[10]: resource conflict
--- FAIL: TestAsAppErrorMapsStockNotEnough
    error_propagation_test.go:47: AsAppError HTTPStatus = 500, want 400
```
