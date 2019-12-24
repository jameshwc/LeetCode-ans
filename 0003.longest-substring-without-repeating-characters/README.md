# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

用*idx*記錄字元有沒有出現過，如果有，出現過最右邊的位置在哪。
*left*和*i*夾成了一個窗口，確認現在這個字元過去出現的位置有沒有在窗口內，如果有將窗口長度縮短。

> 隨便寫寫居然過了sample...submit再根據一些測資寫特判，多丟了兩次就過了...
> 當然，第一次AC的code沒這麼漂亮