package grenade

import (
	"github.com/icrowley/fake"

	log "github.com/sirupsen/logrus"
)

// LogVolumeSpike causes a spike in the volume of log messages
func LogVolumeSpike() {
	for {
		log.Info(fake.Sentence())
	}
}
