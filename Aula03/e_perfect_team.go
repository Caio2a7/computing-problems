package main
import "fmt"
func main(){
	var cases int
	fmt.Scan(&cases)

	for range cases{
		perfect := 0
		coders, maths, any := 0,0,0
		fmt.Scan(&coders)
		fmt.Scan(&maths)
		fmt.Scan(&any)
		for coders > 0 && maths > 0 {
			if any <= 0 {
				if coders > 1 && coders >= maths{
					any++
					coders--
				} else if maths > 1 && maths > coders{
					any++
					maths--
				} else {
					break
				}
			}
			coders--
			maths--
			any--
			perfect++
		}
		fmt.Printf("%d\n", perfect)
	}
}
