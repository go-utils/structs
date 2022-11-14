# structs
structure utils

# What is this?
Provides convenience functions on structures.  

# Usage

### GetNilFields

```go
package main

import (
	"log"

	"github.com/go-utils/structs"
)

type (
	Repository1 struct{}
	Repository2 struct{}
	Repository3 struct{}
)

type ServiceRepositories struct {
	Repository1 *Repository1
	Repository2 *Repository2
	Repository3 *Repository3
}

type Service struct {
	repos ServiceRepositories
}

func NewService(repos ServiceRepositories) *Service {
	if nilFields := structs.GetNilFields(repos); len(nilFields) > 0 {
		log.Fatalf("%+v", nilFields)
		// Output: ["ServiceRepositories.Repository2", "ServiceRepositories.Repository3"]
	}
	return &Service{repos: repos}
}

func main() {
	repos := ServiceRepositories{
		Repository1: new(Repository1),
	}
	service := NewService(repos)
	...
}
```

---

### GetStructName

```go
package main

import (
	"fmt"
	"log"

	"github.com/go-utils/structs"
)

type Hoge struct{}

func main() {
	hoge := new(Hoge)
	hogeName := structs.GetStructName(hoge)
	fmt.Printf("struct name: %#v\n", hogeName)
	// Output: "Hoge"
}
```

## License
- Under the [MIT License](./LICENSE)
- Copyright (C) 2022 go-utils
