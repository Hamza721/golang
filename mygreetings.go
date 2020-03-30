package GoPackage
import (
	"fmt"
)
var greetingString ="Hello World!"

func Program() int  {
		var age int
	var check bool
    fmt.Println(" 1 – Print Covid19 cases in Pakistan","\n",
"2 – Print Covid19 cases in SouthKorea","\n",
"3 – Print Covid19 cases in France","\n",
"4 – Print my personalized message to Coronavirus","\n",
"0 – Exit")
check=true
	for ; check; {

		fmt.Println("Please check an option!")
		fmt.Scan(&age)

		if age==1 {
			fmt.Println("number of cases in Pakistan is: 1,420")
		} else if age==2 {
			fmt.Println("number of cases in South Korea is: 9,478")
		} else if age==3 {
			fmt.Println("number of cases in Farance is: 32,964")
		} else if age==4 {
			fmt.Println("StayHome_StaySave")
		} else if age==0 {
			check=false
		} else if age > 4 {
			fmt.Println(" invalid input")
		}
	}
		// return greetingString + "-" + name
		return 0

}
