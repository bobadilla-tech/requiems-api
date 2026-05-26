package paymentvalidate

import (
	"context"
	"strings"
	"sync"

	"requiems-api/services/finance/bin"
	"requiems-api/services/finance/iban"
	"requiems-api/services/finance/swift"
)

type BINLooker interface {
	Lookup(ctx context.Context, raw string) (bin.LookupResponse, error)
}

type IBANParser interface {
	Parse(ctx context.Context, raw string) (iban.ParseResponse, error)
}

type SWIFTLooker interface {
	Lookup(ctx context.Context, raw string) (swift.LookupResponse, error)
}

type Service struct {
	bin   BINLooker
	iban  IBANParser
	swift SWIFTLooker
}

func NewService(b BINLooker, i IBANParser, s SWIFTLooker) *Service {
	return &Service{bin: b, iban: i, swift: s}
}

type BINResult struct {
	Valid       bool   `json:"valid"`
	Scheme      string `json:"scheme"`
	CardType    string `json:"card_type"`
	CardLevel   string `json:"card_level"`
	CountryCode string `json:"country_code"`
	Issuer      string `json:"issuer"`
	Prepaid     bool   `json:"prepaid"`
	Luhn        bool   `json:"luhn"`
}

type IBANResult struct {
	Valid       bool   `json:"valid"`
	CountryCode string `json:"country_code"`
	BankCode    string `json:"bank_code"`
	Account     string `json:"account_number"`
}

type SWIFTResult struct {
	Valid       bool   `json:"valid"`
	Institution string `json:"institution"`
	Country     string `json:"country"`
	Branch      string `json:"branch"`
}

type Consistency struct {
	OK    bool     `json:"ok"`
	Flags []string `json:"flags"`
}

type Result struct {
	BIN         *BINResult   `json:"bin"`
	IBAN        *IBANResult  `json:"iban"`
	SWIFT       *SWIFTResult `json:"swift"`
	Consistency Consistency  `json:"consistency"`
}

func (s *Service) Validate(ctx context.Context, req Request) (Result, error) {
	type binOut struct {
		r   bin.LookupResponse
		err error
	}
	type ibanOut struct {
		r   iban.ParseResponse
		err error
	}
	type swiftOut struct {
		r   swift.LookupResponse
		err error
	}

	binCh := make(chan binOut, 1)
	ibanCh := make(chan ibanOut, 1)
	swiftCh := make(chan swiftOut, 1)

	var wg sync.WaitGroup

	if req.BIN != "" {
		wg.Go(func() {
			r, err := s.bin.Lookup(ctx, req.BIN)
			binCh <- binOut{r, err}
		})
	} else {
		binCh <- binOut{}
	}

	if req.IBAN != "" {
		wg.Go(func() {
			r, err := s.iban.Parse(ctx, req.IBAN)
			ibanCh <- ibanOut{r, err}
		})
	} else {
		ibanCh <- ibanOut{}
	}

	if req.SWIFT != "" {
		wg.Go(func() {
			r, err := s.swift.Lookup(ctx, req.SWIFT)
			swiftCh <- swiftOut{r, err}
		})
	} else {
		swiftCh <- swiftOut{}
	}

	wg.Wait()

	binResult := <-binCh
	ibanResult := <-ibanCh
	swiftResult := <-swiftCh

	var out Result

	var binCountry, ibanCountry, swiftCountry string
	var binBankCode4, ibanBankCode string

	if req.BIN != "" {
		if binResult.err == nil {
			out.BIN = &BINResult{
				Valid:       true,
				Scheme:      binResult.r.Scheme,
				CardType:    binResult.r.CardType,
				CardLevel:   binResult.r.CardLevel,
				CountryCode: binResult.r.CountryCode,
				Issuer:      binResult.r.IssuerName,
				Prepaid:     binResult.r.Prepaid,
				Luhn:        binResult.r.Luhn,
			}
			binCountry = strings.ToUpper(binResult.r.CountryCode)
		} else {
			out.BIN = &BINResult{Valid: false}
		}
	}

	if req.IBAN != "" {
		if ibanResult.err == nil {
			out.IBAN = &IBANResult{
				Valid:       ibanResult.r.Valid,
				CountryCode: ibanResult.r.Country[:2],
				BankCode:    ibanResult.r.BankCode,
				Account:     ibanResult.r.Account,
			}
			ibanCountry = strings.ToUpper(ibanResult.r.IBAN[:2])
			ibanBankCode = strings.ToUpper(ibanResult.r.BankCode)
		} else {
			out.IBAN = &IBANResult{Valid: false}
		}
	}

	if req.SWIFT != "" {
		if swiftResult.err == nil {
			out.SWIFT = &SWIFTResult{
				Valid:       true,
				Institution: swiftResult.r.BankName,
				Country:     swiftResult.r.CountryCode,
				Branch:      swiftResult.r.City,
			}
			// BIC country is chars 5–6 of the SWIFT code (ISO alpha-2)
			if len(swiftResult.r.SwiftCode) >= 6 {
				swiftCountry = strings.ToUpper(swiftResult.r.SwiftCode[4:6])
			}
			// BIC bank prefix = first 4 chars
			if len(swiftResult.r.SwiftCode) >= 4 {
				binBankCode4 = strings.ToUpper(swiftResult.r.SwiftCode[:4])
			}
		} else {
			out.SWIFT = &SWIFTResult{Valid: false}
		}
	}

	out.Consistency = checkConsistency(binCountry, ibanCountry, swiftCountry, ibanBankCode, binBankCode4)
	return out, nil
}

func checkConsistency(binCC, ibanCC, swiftCC, ibanBankCode, swiftBICPrefix string) Consistency {
	flags := make([]string, 0, 4)

	if binCC != "" && ibanCC != "" && binCC != ibanCC {
		flags = append(flags, "country_mismatch_bin_iban")
	}
	if binCC != "" && swiftCC != "" && binCC != swiftCC {
		flags = append(flags, "country_mismatch_bin_swift")
	}
	if ibanCC != "" && swiftCC != "" && ibanCC != swiftCC {
		flags = append(flags, "country_mismatch_iban_swift")
	}
	if ibanBankCode != "" && swiftBICPrefix != "" {
		// IBAN bank code must match the first 4 chars of the SWIFT BIC.
		if !strings.HasPrefix(strings.ToUpper(ibanBankCode), swiftBICPrefix) &&
			!strings.HasPrefix(swiftBICPrefix, strings.ToUpper(ibanBankCode)) {
			flags = append(flags, "bank_mismatch_iban_swift")
		}
	}

	return Consistency{OK: len(flags) == 0, Flags: flags}
}
