package main

func maxOperations(nums []int, k int) int {
    var m= make(map[int]int,0)
    var count int
    for _, a:= range nums{
        m[a]++
        if m[k-a]>0{
            if a!=k-a || m[a]>=2{
                count++
                m[a]--
                m[k-a]--
            }
        }
    }
    return count
}