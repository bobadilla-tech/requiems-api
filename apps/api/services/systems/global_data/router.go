package globaldata

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"

	"requiems-api/platform/config"
	ipinfo "requiems-api/services/networking/ip/info"
	"requiems-api/services/places/geocode"
	"requiems-api/services/places/holidays"
	"requiems-api/services/places/timezone"
	workingdays "requiems-api/services/places/working-days"
	businesscalendar "requiems-api/services/systems/global_data/business_calendar"
	locationresolve "requiems-api/services/systems/global_data/location_resolve"
	timezoneip "requiems-api/services/systems/global_data/timezone_ip"

	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
)

// Deps holds the external clients needed by the Global Data system.
type Deps struct {
	IPIClient *ipi.Client
	Cfg       config.Config
	RDB       *redis.Client
}

// RegisterRoutes mounts all Global Data system endpoints on r.
func RegisterRoutes(r chi.Router, deps Deps) {
	timezoneSvc, err := timezone.NewService()
	if err != nil {
		log.Printf("global_data: failed to initialize timezone service: %v", err)
	}

	holidaysSvc := holidays.NewService()
	workingDaysSvc := workingdays.NewService()

	calSvc := businesscalendar.NewService(holidaysSvc, workingDaysSvc)
	businesscalendar.RegisterRoutes(r, calSvc)

	infoSvc := ipinfo.NewService(deps.IPIClient)
	tzIPSvc := timezoneip.NewService(infoSvc, timezoneSvc)
	timezoneip.RegisterRoutes(r, tzIPSvc)

	geocodeSvc := geocode.NewService(deps.Cfg.NominatimURL, &http.Client{Timeout: 10 * time.Second}, deps.RDB)
	locSvc := locationresolve.NewService(geocodeSvc, timezoneSvc, holidaysSvc, workingDaysSvc)
	locationresolve.RegisterRoutes(r, locSvc)
}
