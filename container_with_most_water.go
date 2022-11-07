package main

func maxArea(height []int) int {
    max:=0
    var i, j=0, len(height)-1
    for i<len(height)&&j>=0{
        minheight:=height[i]
        width:=0
        if j-i>=0{
            width=j-i
        }else{
            width=i-j
        }
        if height[j] < height[i]{
            minheight=height[j]
            j--
        }else{
            i++
        }
        area:=width*minheight
        if area > max{
            max=area
        }
    }  
    return max
}