package main


func countingValleys(steps int32, path string) int32 {
    var ans int32=0
    var curr int32 =0
    for _,v:=range path{
        if rune(v)=='D'{
            curr-=1
        }else{
            curr+=1
            if (curr==0){
                ans++
            }
        }
    }
    return ans
}