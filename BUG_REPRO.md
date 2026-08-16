# BUG 复现说明（store-inventory-transfer__003）

## Bug 是什么
SKU 批量导入时复用了同一个 `model.SKU` 结构体指针，导致所有导入的 SKU 都指向最后一个商品的编码，批量导入后数据全部串成同一条。

## 如何触发
```bash
go test ./internal/service -run 'TestBatchImportKeepsDistinctCodes' -count=1
```

## 错误信息
```
--- FAIL: TestBatchImportKeepsDistinctCodes
    sku_batch_test.go:56: created[0].Code = SKU-C, want SKU-A
```
