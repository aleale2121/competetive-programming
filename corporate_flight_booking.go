package main

func corpFlightBookings(bookings [][]int, n int) []int {
    ans:=make([]int,n,n)
    for i:=0;i<len(bookings);i++{
        for k:=bookings[i][0]-1;k<bookings[i][1];k++{
            ans[k]+=bookings[i][2]
        }
    }
    return ans
    
}