//go:build profile

package easypgo

import (
	"log"
	"os"
	"runtime/pprof"
)

// Generate starts a CPU profile and returns a function to stop it.
// The profile generation is only enabled when the "profile" build tag is set.
func Generate() func() {
	f, err := os.CreateTemp("", "profile-*.pgo")
	if err != nil {
		log.Fatalln("Could not create CPU profile: ", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatalln("Could not start CPU profile: ", err)
	}

	return func() {
		pprof.StopCPUProfile()
		err := f.Close()
		if err != nil {
			log.Fatalln("Could not close CPU profile: ", err)
		}
		err = os.Rename(f.Name(), "default.pgo")
		if err != nil {
			log.Fatalln("Could not move CPU profile into default.pgo: ", err)
		}
	}
}
