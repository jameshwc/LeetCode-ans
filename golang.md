# Golang 研究心得

> 這裡記錄了我解題遇到的屬於golang本身的特別情況

## return pointer to local variable
```go
// 0002. Add Two Numbers
func test() *ListNode {
    return &ListNode{Val: 10, Next: nil}
}
```
上面這段code在c++ 是不合法的，原因是回傳值是在function裡面宣告的，而c++在離開function後，該function的local variable就會被destruct，於是該指標指向的位址就非預期的值。但golang在這裡很神奇的會分析該指標有沒有escape local stack，如果有的話就會將該變數的位址分配在heap，故上面這段寫法是可行的，在golang圈裡也非常常見。
[參考資料](https://stackoverflow.com/questions/13715237/return-pointer-to-local-struct)