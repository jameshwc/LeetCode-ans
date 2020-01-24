func reverseBits(num uint32) uint32 {
    var ans, cnt uint32
    for cnt != 32{
        b := (num>>cnt) & 1
        ans = ans | (b<<(31-cnt))
        cnt++
    }
    return ans
}