# easypgo
A simple package to generate PGO profiles for Go projects.
Just include this snippet in your Go project's main function and build with `-tags profile` to generate the profile (regular builds will not include any dependencies for the profile generation):

```go
import "github.com/rymdport/easypgo"

func main() {
	stop := easypgo.Generate()
	defer stop()
}
```
