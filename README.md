# Guide to use it
go get github.com/vks16/math


go doc github.com/vks16/math

# in your src code.
```
package main

import (
	"fmt"
	"github.com/vks16/math"
)

func main() {

	num1 := -323

	fmt.Println(math.Abs(num1)

	
	arr := []int{3, 2, 34, 342, 322, -322}

	fmt.Println(math.Max(arr))

	fmt.Println(math.Min(arr)) 
}
```

# Contribution
Add more relevent functions in this math library with detailed documentation.

1. You should always first write a test for a function
2. run the test so that it fails.  ``` go test ```
3. Write the function to pass the test.

