# opt
 
This module is supposed to ease up access to golang struct fields that are pointers to primitives - with the premise that nil pointers are equivalent to the zero value of said primitive.

The convenience functions offered by this module essentially abstract away the obligatory nil checks when accessing such fields, thereby fostering easier to read code.
Because the go compiler cannot produce pointers to primitive literals, the same helper functions exist for creating pointers to initialized primitive values using a one liner.

```go
package main

import "github.com/dtcookie/opt"

type Record struct {
    Value *int
}

func main() {
    // record.Value uninitialized
    record := new(Record)

    // manual nil check
    var a int
    if record.Value != nil {
        a = *record.Value
    }

    // nil check performed by opt module
    b := opt.Int(record.Value)

    // assignment using a helper variable
    // the compiler won't accept `record.Value = &3`
    c := 3
    record.Value = &c

    // helper variable abstracted away by opt module
    record.Value = opt.NewInt(3)

}
```