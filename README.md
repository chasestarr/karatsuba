## Karatsuba ([Karatsuba](https://en.wikipedia.org/wiki/Karatsuba_algorithm) multiplication algorithm)

[![GoDoc](https://godoc.org/github.com/chasestarr/karatsuba?status.svg)](https://godoc.org/github.com/chasestarr/karatsuba)

```go
import (
  "fmt"

  "github.com/chasestarr/karatsuba"
)

func main() {
  product := karatsuba.Multiply("1234" "5678") // "7006652"
  fmt.Println(product)
}
```