//go:build profile

package easypgo

import (
	"io"
	"log"
	"os"
	"runtime/pprof"
)

// Generate starts a CPU profile and returns a function to stop it.
// The profile generation is only enabled when the "profile" build tag is set.
func Generate() func() {
	src, err := os.CreateTemp("", "profile-*.pgo")
	if err != nil {
		log.Fatalln("Could not create CPU profile: ", err)
	}

	if err := pprof.StartCPUProfile(src); err != nil {
		log.Fatalln("Could not start CPU profile: ", err)
	}

	return func() {
		pprof.StopCPUProfile()
		defer src.Close()

		dst, err := os.Create("default.pgo")
		if err != nil {
			log.Fatalln("Could create default.pgo: ", err)
		}

		defer dst.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			log.Fatalln("Could not copy CPU profile to default.pgo: ", err)
		}

		err = os.Remove(src.Name())
		if err != nil {
			log.Fatalln("Could not remove temporary CPU profile: ", err)
		}
	}
}
