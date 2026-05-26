package timezoneip

import (
	"context"
	"fmt"

	ipinfo "requiems-api/services/networking/ip/info"
	"requiems-api/services/places/timezone"
)

type IPInfoChecker interface {
	CheckInfo(ctx context.Context, ip string) (ipinfo.LookupResponse, error)
}

type TimezoneGetter interface {
	GetTimezoneByCoords(lat, lon float64) (*timezone.Info, error)
	GetTimezoneByCity(city string) (*timezone.Info, error)
}

type Service struct {
	ipInfo   IPInfoChecker
	timezone TimezoneGetter
}

func NewService(i IPInfoChecker, t TimezoneGetter) *Service {
	return &Service{ipInfo: i, timezone: t}
}

type Result struct {
	IP          string  `json:"ip"`
	City        string  `json:"city"`
	CountryCode string  `json:"country_code"`
	Timezone    *string `json:"timezone"`
	UTCOffset   *string `json:"utc_offset"`
	DSTActive   *bool   `json:"dst_active"`
	CurrentTime *string `json:"current_time"`
}

func (s *Service) Resolve(ctx context.Context, ip string) (Result, error) {
	info, err := s.ipInfo.CheckInfo(ctx, ip)
	if err != nil {
		return Result{}, fmt.Errorf("ip_not_found")
	}

	res := Result{
		IP:          ip,
		City:        info.City,
		CountryCode: info.CountryCode,
	}

	tzInfo, err := s.timezone.GetTimezoneByCity(info.City)
	if err != nil {
		// timezone unavailable — return 200 with timezone: null
		return res, nil
	}

	res.Timezone = &tzInfo.Timezone
	res.UTCOffset = &tzInfo.Offset
	res.CurrentTime = &tzInfo.CurrentTime
	dst := tzInfo.IsDST
	res.DSTActive = &dst
	return res, nil
}
