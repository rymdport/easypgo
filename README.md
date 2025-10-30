# easypgo
A simple package to generate PGO profiles for Go projects.

## Usage
Just include this snippet in your Go project's main function and build with `-tags profile` to generate the profile:

```go
import "github.com/rymdport/easypgo"

func main() {
	stop := easypgo.Generate()
	defer stop()
}
```

Binaries built without the tag will not generate any profiles (or import any of the dependencies for it).
When running the binary built with the `profile` tag, a profile will be generated for the duration of the program's execution.
Once completed, a `default.pgo` file will be stored in the current working directory.
Go automatically picks up the profile and uses it to optimize the binary further; just do a rebuild without the `profile` tag.
