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

When running the binary, a profile will be generated for the duration of the program's execution.
Once completed, a `default.pgo` file will be generated in the current working directory.
Do a rebuild without the `profile` for Go to automatically apply the profile to the binary.
