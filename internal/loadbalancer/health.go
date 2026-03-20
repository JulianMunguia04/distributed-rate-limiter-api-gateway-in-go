package loadbalancer

import (
	"log"
	"net/http"
	"time"
)

func HealthCheck(backends []*Backend) {
	for {
		for _, b := range backends {

			resp, err := http.Get(b.URL.String() + "/health")

			if err != nil || resp.StatusCode != http.StatusOK {
				if b.IsAlive() {
					log.Println("Backend DOWN:", b.URL)
				}
				b.SetAlive(false)
			} else {
				if !b.IsAlive() {
					log.Println("Backend UP:", b.URL)
				}
				b.SetAlive(true)
			}
		}

		time.Sleep(5 * time.Second)
	}
}
