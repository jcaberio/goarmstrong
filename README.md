goarmstrong
===========

A package that checks if a number is an Armstrong number.

An [Armstrong number](https://en.wikipedia.org/wiki/Narcissistic_number) is an n-digit number that is equal to the sum of the nth powers of its digits.

For example, 371 is an Armstrong number since 3<sup>3</sup> + 7<sup>3</sup> + 1<sup>3</sup> = 371

Here's a [list of Armstrong numbers](https://oeis.org/A005188).

Installation
------------

```
go get github.com/jcaberio/goarmstrong
```

Usage
-----

```go
package main

import (
  "fmt"
	"github.com/jcaberio/goarmstrong"
)

func main() {
	v := goarmstrong.IsArmstrong(912985153)
	fmt.Println(v)
}

```
