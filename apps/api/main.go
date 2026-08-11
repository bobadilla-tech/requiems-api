package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	sentry "github.com/getsentry/sentry-go"

	"requiems-api/app"
	"requiems-api/platform/config"
)

// shutdownTimeout bounds how long the server waits for in-flight requests to
// drain after receiving a shutdown signal before forcing an exit.
const shutdownTimeout = 15 * time.Second

// @title						Requiems API
// @version					1.0.0
// @description				Unified access to enterprise-grade APIs — email validation, text utilities, and more. Authenticate with the `requiems-api-key` header.
// @host						api.requiems.xyz
// @schemes					https
//
// @tag.name					advice
// @tag.description			Get random pieces of advice and wisdom for inspiration
// @tag.name					barcode
// @tag.description			Generate barcodes in multiple formats (Code 128, Code 93, Code 39, EAN-8, EAN-13), returned as a PNG image or base64-encoded JSON
// @tag.name					base64
// @tag.description			Encode strings to Base64 and decode Base64 back to plain text. Supports standard and URL-safe (base64url) variants.
// @tag.name					bin-lookup
// @tag.description			Pass the first 6–8 digits of any payment card and get back the issuing bank, card network, type, and country
// @tag.name					chuck-norris
// @tag.description			Get a random Chuck Norris fact from a curated built-in database. Every call returns a different fact with a stable unique ID.
// @tag.name					cities
// @tag.description			Look up city metadata including population, timezone, and coordinates
// @tag.name					color-conversion
// @tag.description			Convert color values between HEX, RGB, HSL, and CMYK. Every response includes all four representations.
// @tag.name					commodities
// @tag.description			Historical and current annual average prices for 16 major commodities — precious metals, energy, and agricultural goods. Sourced from FRED.
// @tag.name					counter
// @tag.description			Atomic, namespace-isolated hit counter
// @tag.name					crypto
// @tag.description			Get live cryptocurrency prices, 24h change, market cap, and trading volume for 20+ major coins including BTC, ETH, SOL, and more.
// @tag.name					dad-jokes
// @tag.description			Get a random dad joke. Classic groan-worthy puns and wholesome humor, served one at a time.
// @tag.name					data-format-conversion
// @tag.description			Convert structured data between JSON, YAML, CSV, XML, and TOML in a single API call.
// @tag.name					data-integrity
// @tag.description			Validate, moderate, and normalize user input at the boundary of your product.
// @tag.name					detect-language
// @tag.description			Detect the language of any text with confidence scoring
// @tag.name					dictionary
// @tag.description			Get word definitions, phonetics, usage examples, and synonyms
// @tag.name					disposable-email
// @tag.description			Check whether an email domain belongs to a known disposable or temporary email provider.
// @tag.name					domain-info
// @tag.description			Look up DNS records and check domain availability. Returns A, AAAA, MX, NS, TXT, and CNAME records alongside a registration availability flag.
// @tag.name					email-normalize
// @tag.description			Normalize email addresses to their canonical form. Lowercased, trimmed, and canonicalized with provider-specific rules including alias-domain resolution.
// @tag.name					email-validate
// @tag.description			Full email validation in one call. Syntax check, MX record lookup, disposable domain detection, normalization, and typo suggestions.
// @tag.name					emoji
// @tag.description			Look up emoji by name, search by keyword, or get a random emoji with full Unicode metadata.
// @tag.name					exchange-rate
// @tag.description			Get live currency exchange rates and convert amounts between currencies. Rates are sourced from the ECB and cached for up to one hour.
// @tag.name					facts
// @tag.description			Get random interesting facts from a curated database. Filter by category — science, history, technology, nature, space, or food.
// @tag.name					fitness-exercises
// @tag.description			Browse 1,500+ exercises with step-by-step instructions, target muscles, secondary muscles, equipment requirements, and body part filters.
// @tag.name					geocode
// @tag.description			Convert addresses to coordinates and coordinates back to addresses
// @tag.name					global-data
// @tag.description			Resolve addresses, coordinates, and IPs into complete location profiles.
// @tag.name					holidays
// @tag.description			Get a list of holidays for a specific country and year
// @tag.name					horoscope
// @tag.description			Get daily horoscope readings for all 12 zodiac signs
// @tag.name					iban
// @tag.description			Validate IBAN numbers and extract the bank code and account number. Supports all countries in the official SWIFT IBAN Registry.
// @tag.name					identity-risk
// @tag.description			Compose email, phone, IP, and behavioral signals into a single risk decision.
// @tag.name					inflation
// @tag.description			Historical and current CPI inflation rates for 241 countries, sourced from the World Bank. Includes up to 30 years of annual data.
// @tag.name					ip-asn
// @tag.description			Look up Autonomous System Number (ASN), organization, ISP, and network route information for any IP address.
// @tag.name					ip-info
// @tag.description			Get geolocation data for any IP address including country, city, ISP, and VPN detection.
// @tag.name					lorem-ipsum
// @tag.description			Generate placeholder text for design mockups and prototypes
// @tag.name					markdown
// @tag.description			Convert Markdown to HTML in a single API call. Optionally sanitize the output to strip unsafe tags and prevent XSS.
// @tag.name					mortgage
// @tag.description			Calculate monthly mortgage payments and full amortization schedules. Pass principal, annual interest rate, and loan term in years.
// @tag.name					mx-lookup
// @tag.description			Look up MX (Mail Exchange) records for any domain, sorted by priority.
// @tag.name					number-base-conversion
// @tag.description			Convert integers between binary, octal, decimal, and hexadecimal. Accepts optional 0x, 0b, and 0o prefixes.
// @tag.name					password-generator
// @tag.description			Generate cryptographically secure random passwords with customizable length and character sets
// @tag.name					payments-intelligence
// @tag.description			Validate financial identifiers and score transaction risk in a single call.
// @tag.name					phone-validation
// @tag.description			Validate phone numbers globally. Detect carrier, country, type, and VOIP or virtual risk using only phone metadata.
// @tag.name					postal-code
// @tag.description			Look up city, state, and coordinates for any postal code worldwide
// @tag.name					profanity
// @tag.description			Detect and censor profanity in text for content moderation
// @tag.name					qr-code
// @tag.description			Generate QR codes from any text or URL, returned as a PNG image or base64-encoded JSON
// @tag.name					quotes
// @tag.description			Access a database of inspirational and famous quotes
// @tag.name					random-user
// @tag.description			Generate random fake user profiles for testing and prototyping — names, emails, phone numbers, addresses, and avatars
// @tag.name					sentiment
// @tag.description			Analyze the sentiment of any text and get a positive, negative, or neutral classification with a confidence score and full class breakdown.
// @tag.name					spell-check
// @tag.description			Check spelling and get correction suggestions for misspelled words
// @tag.name					sudoku
// @tag.description			Generate Sudoku puzzles with solutions across multiple difficulty levels
// @tag.name					swift-code
// @tag.description			Validate and look up bank information by SWIFT/BIC code, including institution, location, and branch metadata.
// @tag.name					text-similarity
// @tag.description			Compare two texts and get a cosine similarity score between 0 and 1
// @tag.name					thesaurus
// @tag.description			Find synonyms and antonyms for any word to enhance vocabulary and writing
// @tag.name					timezone
// @tag.description			Get timezone information for any location by coordinates or city name
// @tag.name					trivia
// @tag.description			Get random trivia questions with multiple-choice answers. Filter by category and difficulty.
// @tag.name					unit-conversion
// @tag.description			Convert between units of measurement — length, weight, volume, temperature, area, and speed
// @tag.name					useragent
// @tag.description			Parse user agent strings to extract browser, OS, device type, and bot detection
// @tag.name					vpn-detection
// @tag.description			Detect if an IP address belongs to a VPN, proxy, Tor exit node, or hosting provider. Returns threat scores and fraud indicators.
// @tag.name					whois
// @tag.description			Get domain registration details including registrar, name servers, status, creation and expiry dates, and DNSSEC information.
// @tag.name					words
// @tag.description			Generate random words for testing, games, and creative projects
// @tag.name					working-days
// @tag.description			Calculate the number of working days between two dates with optional country-specific holidays
// @tag.name					world-time
// @tag.description			Get the current time for any IANA timezone by name
//
// @security					requiems-api-key
// @securityDefinitions.apikey	requiems-api-key
// @in							header
// @name						requiems-api-key
func main() {
	os.Exit(run())
}

func run() int {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	if cfg.Environment != "development" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:              cfg.SentryDSN,
			Environment:      cfg.Environment,
			TracesSampleRate: 0.01,
		}); err != nil {
			logger.Error("sentry.Init failed", "error", err)
		}
	}

	appInstance, err := app.New(ctx, cfg)

	if err != nil {
		logger.Error("failed to initialise app", "error", err)
		sentry.Flush(2 * time.Second)
		return 1
	}

	addr := fmt.Sprintf(":%s", cfg.Port)

	server := &http.Server{
		Addr:              addr,
		Handler:           appInstance.Handler(),
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	serveErr := make(chan error, 1)

	go func() {
		logger.Info("API server listening", "addr", addr)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
			return
		}

		serveErr <- nil
	}()

	select {
	case err := <-serveErr:
		if err != nil {
			logger.Error("server error", "error", err)
			sentry.Flush(2 * time.Second)
			appInstance.Close()
			return 1
		}
	case <-ctx.Done():
		logger.Info("shutdown signal received, draining in-flight requests")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("graceful shutdown failed", "error", err)
		}
	}

	appInstance.Close()
	sentry.Flush(2 * time.Second)
	logger.Info("shutdown complete")
	return 0
}
