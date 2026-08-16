# BUG 复现说明（store-inventory-transfer__005）

## Bug 是什么
门店与 SKU 新建时的唯一性冲突误判：
- 重复编码的门店被当作创建成功返回，而不是返回“资源状态冲突”。
- 重复编码的 SKU 被错误返回“请求参数不合法”，而不是“资源状态冲突”。

## 如何触发
```bash
go test ./internal/service -run 'TestStoreCreateDuplicateReturnsConflict|TestSKUCreateDuplicateReturnsConflict' -count=1
```

## 错误信息
```
--- FAIL: TestStoreCreateDuplicateReturnsConflict
    create_conflict_test.go:44: StoreService.Create duplicate should return error
--- FAIL: TestSKUCreateDuplicateReturnsConflict
    create_conflict_test.go:58: err = create sku[code=SKU-1]: validation failed, want errors.Is util.ErrConflict
```
