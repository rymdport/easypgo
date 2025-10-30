//go:build !profile

// Package easypgo provides tooling to automatically generate a CPU profile for PGO builds.
package easypgo

// Generate starts a CPU profile and returns a function to stop it.
// The profile generation is only enabled when the "profile" build tag is set.
func Generate() func() {
	return func() {}
}
