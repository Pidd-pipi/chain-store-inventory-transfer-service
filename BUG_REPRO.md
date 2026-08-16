# BUG 复现说明（store-inventory-transfer__004）

## Bug 是什么
用户门店绑定的 nil 处理缺失：
- 总部管理员/总部用户未绑定门店时注册被拒绝。
- 店长更新资料时，如果请求未携带 `store_id`，会把原有门店绑定清空。
- 给未绑定门店的用户签发 JWT 时发生 nil 指针解引用 panic。

## 如何触发
```bash
go test ./internal/service -run 'TestAdminRegisterAllowsNilStore|TestUpdateProfileKeepsStoreOnNil|TestGenerateTokenAllowsNilStore'
```

## 错误信息
```
--- FAIL: TestAdminRegisterAllowsNilStore
    user_nil_test.go:14: admin register should allow nil store, got err: register: validation failed
--- FAIL: TestUpdateProfileKeepsStoreOnNil
    user_nil_test.go:33: StoreID changed to <nil>, want 7
panic: runtime error: invalid memory address or nil pointer dereference
```
