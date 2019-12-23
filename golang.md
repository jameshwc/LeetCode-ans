# Golang 研究心得

> 這裡記錄了我解題遇到的屬於golang本身的特別情況

## return pointer to local variable
```go
// 0002. Add Two Numbers
func test() *ListNode{
    return &ListNode{Val: 10, Next: nil}
}
```