## api-client@1.0.1

This generator creates TypeScript/JavaScript client that utilizes [axios](https://github.com/axios/axios). The generated Node module can be used in the following environments:

Environment
* Node.js
* Webpack
* Browserify

Language level
* ES5 - you must have a Promises/A+ library installed
* ES6

Module system
* CommonJS
* ES6 module system

It can be used in both TypeScript and JavaScript. In TypeScript, the definition will be automatically resolved via `package.json`. ([Reference](https://www.typescriptlang.org/docs/handbook/declaration-files/consumption.html))

### Building

To build and compile the typescript sources to javascript use:
```
npm install
npm run build
```

### Publishing

First build the package then run `npm publish`

### Consuming

navigate to the folder of your consuming project and run one of the following commands.

_published:_

```
npm install api-client@1.0.1 --save
```

_unPublished (not recommended):_

```
npm install PATH_TO_GENERATED_PACKAGE --save
```

### Documentation for API Endpoints

All URIs are relative to *https://requiems.xyz*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdviceApi* | [**v1TextAdviceGet**](docs/AdviceApi.md#v1textadviceget) | **GET** /v1/text/advice | Get Random Advice
*BarcodeApi* | [**v1TechnologyBarcodeBase64Get**](docs/BarcodeApi.md#v1technologybarcodebase64get) | **GET** /v1/technology/barcode/base64 | Generate Barcode (Base64 JSON)
*BarcodeApi* | [**v1TechnologyBarcodeGet**](docs/BarcodeApi.md#v1technologybarcodeget) | **GET** /v1/technology/barcode | Generate Barcode (PNG)
*BinLookupApi* | [**v1FinanceBinBinGet**](docs/BinLookupApi.md#v1financebinbinget) | **GET** /v1/finance/bin/{bin} | BIN Lookup
*ChuckNorrisApi* | [**v1EntertainmentChuckNorrisGet**](docs/ChuckNorrisApi.md#v1entertainmentchucknorrisget) | **GET** /v1/entertainment/chuck-norris | Get Random Chuck Norris Fact
*CitiesApi* | [**v1PlacesCitiesCityGet**](docs/CitiesApi.md#v1placescitiescityget) | **GET** /v1/places/cities/{city} | Get City Info
*ColorConversionApi* | [**v1TechnologyColorGet**](docs/ColorConversionApi.md#v1technologycolorget) | **GET** /v1/technology/color | Convert Color
*CommoditiesApi* | [**v1FinanceCommoditiesCommodityGet**](docs/CommoditiesApi.md#v1financecommoditiescommodityget) | **GET** /v1/finance/commodities/{commodity} | Get Commodity Price
*CounterApi* | [**v1TechnologyCounterNamespaceGet**](docs/CounterApi.md#v1technologycounternamespaceget) | **GET** /v1/technology/counter/{namespace} | Get Counter Value
*CounterApi* | [**v1TechnologyCounterNamespacePost**](docs/CounterApi.md#v1technologycounternamespacepost) | **POST** /v1/technology/counter/{namespace} | Increment Counter
*CryptoApi* | [**v1FinanceCryptoSymbolGet**](docs/CryptoApi.md#v1financecryptosymbolget) | **GET** /v1/finance/crypto/{symbol} | Get Crypto Price
*DadJokesApi* | [**v1EntertainmentJokesDadGet**](docs/DadJokesApi.md#v1entertainmentjokesdadget) | **GET** /v1/entertainment/jokes/dad | Get Random Dad Joke
*DataFormatConversionApi* | [**v1TechnologyFormatPost**](docs/DataFormatConversionApi.md#v1technologyformatpost) | **POST** /v1/technology/format | Convert Format
*DataIntegrityApi* | [**v1SystemsContentModeratePost**](docs/DataIntegrityApi.md#v1systemscontentmoderatepost) | **POST** /v1/systems/content/moderate | Moderate Content
*DataIntegrityApi* | [**v1SystemsTextNormalizePost**](docs/DataIntegrityApi.md#v1systemstextnormalizepost) | **POST** /v1/systems/text/normalize | Normalize Text
*DetectLanguageApi* | [**v1TextDetectLanguagePost**](docs/DetectLanguageApi.md#v1textdetectlanguagepost) | **POST** /v1/text/detect-language | Detect Language
*DeveloperUtilitiesApi* | [**v1TechnologyBase64DecodePost**](docs/DeveloperUtilitiesApi.md#v1technologybase64decodepost) | **POST** /v1/technology/base64/decode | Base64 Decode
*DeveloperUtilitiesApi* | [**v1TechnologyBase64EncodePost**](docs/DeveloperUtilitiesApi.md#v1technologybase64encodepost) | **POST** /v1/technology/base64/encode | Base64 Encode
*DictionaryApi* | [**v1TextDictionaryWordGet**](docs/DictionaryApi.md#v1textdictionarywordget) | **GET** /v1/text/dictionary/{word} | Dictionary Lookup
*DisposableEmailApi* | [**v1NetworkingDisposableBatchPost**](docs/DisposableEmailApi.md#v1networkingdisposablebatchpost) | **POST** /v1/networking/disposable/batch | Check Batch Emails
*DisposableEmailApi* | [**v1NetworkingDisposableCheckPost**](docs/DisposableEmailApi.md#v1networkingdisposablecheckpost) | **POST** /v1/networking/disposable/check | Check Single Email
*DisposableEmailApi* | [**v1NetworkingDisposableDomainDomainGet**](docs/DisposableEmailApi.md#v1networkingdisposabledomaindomainget) | **GET** /v1/networking/disposable/domain/{domain} | Check Domain
*DisposableEmailApi* | [**v1NetworkingDisposableDomainsGet**](docs/DisposableEmailApi.md#v1networkingdisposabledomainsget) | **GET** /v1/networking/disposable/domains | List Domains (Paginated)
*DisposableEmailApi* | [**v1NetworkingDisposableStatsGet**](docs/DisposableEmailApi.md#v1networkingdisposablestatsget) | **GET** /v1/networking/disposable/stats | Get Statistics
*DomainInfoApi* | [**v1NetworkingDomainDomainGet**](docs/DomainInfoApi.md#v1networkingdomaindomainget) | **GET** /v1/networking/domain/{domain} | Get Domain Info
*EmailNormalizeApi* | [**v1TextNormalizeBatchPost**](docs/EmailNormalizeApi.md#v1textnormalizebatchpost) | **POST** /v1/text/normalize/batch | Normalize Email Batch
*EmailNormalizeApi* | [**v1TextNormalizePost**](docs/EmailNormalizeApi.md#v1textnormalizepost) | **POST** /v1/text/normalize | Normalize Email
*EmailValidateApi* | [**v1ValidationEmailBatchPost**](docs/EmailValidateApi.md#v1validationemailbatchpost) | **POST** /v1/validation/email/batch | Validate Emails (Batch)
*EmailValidateApi* | [**v1ValidationEmailPost**](docs/EmailValidateApi.md#v1validationemailpost) | **POST** /v1/validation/email | Validate Email
*EmojiApi* | [**v1EntertainmentEmojiNameGet**](docs/EmojiApi.md#v1entertainmentemojinameget) | **GET** /v1/entertainment/emoji/{name} | Get Emoji by Name
*EmojiApi* | [**v1EntertainmentEmojiRandomGet**](docs/EmojiApi.md#v1entertainmentemojirandomget) | **GET** /v1/entertainment/emoji/random | Get Random Emoji
*EmojiApi* | [**v1EntertainmentEmojiSearchGet**](docs/EmojiApi.md#v1entertainmentemojisearchget) | **GET** /v1/entertainment/emoji/search | Search Emoji
*ExchangeRateApi* | [**v1FinanceConvertGet**](docs/ExchangeRateApi.md#v1financeconvertget) | **GET** /v1/finance/convert | Convert Currency
*ExchangeRateApi* | [**v1FinanceExchangeRateGet**](docs/ExchangeRateApi.md#v1financeexchangerateget) | **GET** /v1/finance/exchange-rate | Get Exchange Rate
*FactsApi* | [**v1EntertainmentFactsGet**](docs/FactsApi.md#v1entertainmentfactsget) | **GET** /v1/entertainment/facts | Get Random Fact
*FitnessExercisesApi* | [**v1HealthBodyPartsGet**](docs/FitnessExercisesApi.md#v1healthbodypartsget) | **GET** /v1/health/body-parts | List Body Parts
*FitnessExercisesApi* | [**v1HealthEquipmentGet**](docs/FitnessExercisesApi.md#v1healthequipmentget) | **GET** /v1/health/equipment | List Equipment
*FitnessExercisesApi* | [**v1HealthExercisesBatchPost**](docs/FitnessExercisesApi.md#v1healthexercisesbatchpost) | **POST** /v1/health/exercises/batch | Batch Get Exercises
*FitnessExercisesApi* | [**v1HealthExercisesGet**](docs/FitnessExercisesApi.md#v1healthexercisesget) | **GET** /v1/health/exercises | List Exercises
*FitnessExercisesApi* | [**v1HealthExercisesIdGet**](docs/FitnessExercisesApi.md#v1healthexercisesidget) | **GET** /v1/health/exercises/{id} | Exercise by ID
*FitnessExercisesApi* | [**v1HealthExercisesRandomGet**](docs/FitnessExercisesApi.md#v1healthexercisesrandomget) | **GET** /v1/health/exercises/random | Random Exercise
*FitnessExercisesApi* | [**v1HealthMusclesGet**](docs/FitnessExercisesApi.md#v1healthmusclesget) | **GET** /v1/health/muscles | List Muscles
*GeocodeApi* | [**v1PlacesGeocodeBatchPost**](docs/GeocodeApi.md#v1placesgeocodebatchpost) | **POST** /v1/places/geocode/batch | Batch Geocode Addresses
*GeocodeApi* | [**v1PlacesGeocodeGet**](docs/GeocodeApi.md#v1placesgeocodeget) | **GET** /v1/places/geocode | Geocode Address
*GeocodeApi* | [**v1PlacesReverseGeocodeBatchPost**](docs/GeocodeApi.md#v1placesreversegeocodebatchpost) | **POST** /v1/places/reverse-geocode/batch | Batch Reverse Geocode
*GeocodeApi* | [**v1PlacesReverseGeocodeGet**](docs/GeocodeApi.md#v1placesreversegeocodeget) | **GET** /v1/places/reverse-geocode | Reverse Geocode
*GlobalDataApi* | [**v1SystemsBusinessCalendarCountryGet**](docs/GlobalDataApi.md#v1systemsbusinesscalendarcountryget) | **GET** /v1/systems/business-calendar/{country} | Business Calendar
*GlobalDataApi* | [**v1SystemsLocationResolvePost**](docs/GlobalDataApi.md#v1systemslocationresolvepost) | **POST** /v1/systems/location/resolve | Resolve Location
*GlobalDataApi* | [**v1SystemsTimezoneFromIpIpGet**](docs/GlobalDataApi.md#v1systemstimezonefromipipget) | **GET** /v1/systems/timezone/from-ip/{ip} | Timezone from IP
*HolidaysApi* | [**v1PlacesHolidaysBatchPost**](docs/HolidaysApi.md#v1placesholidaysbatchpost) | **POST** /v1/places/holidays/batch | Batch Get Holidays
*HolidaysApi* | [**v1PlacesHolidaysGet**](docs/HolidaysApi.md#v1placesholidaysget) | **GET** /v1/places/holidays | Get Holidays
*HoroscopeApi* | [**v1EntertainmentHoroscopeBatchPost**](docs/HoroscopeApi.md#v1entertainmenthoroscopebatchpost) | **POST** /v1/entertainment/horoscope/batch | Batch Daily Horoscopes
*HoroscopeApi* | [**v1EntertainmentHoroscopeSignGet**](docs/HoroscopeApi.md#v1entertainmenthoroscopesignget) | **GET** /v1/entertainment/horoscope/{sign} | Get Daily Horoscope
*IbanApi* | [**v1FinanceIbanBatchPost**](docs/IbanApi.md#v1financeibanbatchpost) | **POST** /v1/finance/iban/batch | Batch Validate IBANs
*IbanApi* | [**v1FinanceIbanIbanGet**](docs/IbanApi.md#v1financeibanibanget) | **GET** /v1/finance/iban/{iban} | Validate IBAN
*IdentityRiskApi* | [**v1SystemsRiskScorePost**](docs/IdentityRiskApi.md#v1systemsriskscorepost) | **POST** /v1/systems/risk/score | Score Risk
*IdentityRiskApi* | [**v1SystemsSignupProtectPost**](docs/IdentityRiskApi.md#v1systemssignupprotectpost) | **POST** /v1/systems/signup/protect | Protect Signup
*IdentityRiskApi* | [**v1SystemsUserVerifyPost**](docs/IdentityRiskApi.md#v1systemsuserverifypost) | **POST** /v1/systems/user/verify | Verify User
*InflationApi* | [**v1FinanceInflationBatchPost**](docs/InflationApi.md#v1financeinflationbatchpost) | **POST** /v1/finance/inflation/batch | Batch Inflation Rates
*InflationApi* | [**v1FinanceInflationGet**](docs/InflationApi.md#v1financeinflationget) | **GET** /v1/finance/inflation | Get Inflation Rate
*IpAsnApi* | [**v1NetworkingIpAsnBatchPost**](docs/IpAsnApi.md#v1networkingipasnbatchpost) | **POST** /v1/networking/ip/asn/batch | Batch ASN Lookup
*IpAsnApi* | [**v1NetworkingIpAsnGet**](docs/IpAsnApi.md#v1networkingipasnget) | **GET** /v1/networking/ip/asn | Lookup ASN (Caller IP)
*IpAsnApi* | [**v1NetworkingIpAsnIpGet**](docs/IpAsnApi.md#v1networkingipasnipget) | **GET** /v1/networking/ip/asn/{ip} | Lookup ASN for IP
*IpInfoApi* | [**v1NetworkingIpGet**](docs/IpInfoApi.md#v1networkingipget) | **GET** /v1/networking/ip | Get IP Info (Caller IP)
*IpInfoApi* | [**v1NetworkingIpInfoBatchPost**](docs/IpInfoApi.md#v1networkingipinfobatchpost) | **POST** /v1/networking/ip/info/batch | Batch IP Geolocation
*IpInfoApi* | [**v1NetworkingIpIpGet**](docs/IpInfoApi.md#v1networkingipipget) | **GET** /v1/networking/ip/{ip} | Get IP Info for IP
*LoremIpsumApi* | [**v1TextLoremBatchPost**](docs/LoremIpsumApi.md#v1textlorembatchpost) | **POST** /v1/text/lorem/batch | Generate Lorem Ipsum (Batch)
*LoremIpsumApi* | [**v1TextLoremGet**](docs/LoremIpsumApi.md#v1textloremget) | **GET** /v1/text/lorem | Generate Lorem Ipsum
*MarkdownApi* | [**v1TechnologyMarkdownBatchPost**](docs/MarkdownApi.md#v1technologymarkdownbatchpost) | **POST** /v1/technology/markdown/batch | Convert Markdown Batch
*MarkdownApi* | [**v1TechnologyMarkdownPost**](docs/MarkdownApi.md#v1technologymarkdownpost) | **POST** /v1/technology/markdown | Convert Markdown to HTML
*MortgageApi* | [**v1FinanceMortgageBatchPost**](docs/MortgageApi.md#v1financemortgagebatchpost) | **POST** /v1/finance/mortgage/batch | Batch Calculate Mortgages
*MortgageApi* | [**v1FinanceMortgageGet**](docs/MortgageApi.md#v1financemortgageget) | **GET** /v1/finance/mortgage | Calculate Mortgage
*MxLookupApi* | [**v1NetworkingMxBatchPost**](docs/MxLookupApi.md#v1networkingmxbatchpost) | **POST** /v1/networking/mx/batch | Batch MX Lookup
*MxLookupApi* | [**v1NetworkingMxDomainGet**](docs/MxLookupApi.md#v1networkingmxdomainget) | **GET** /v1/networking/mx/{domain} | MX Lookup
*NumberBaseConversionApi* | [**v1TechnologyBaseBatchPost**](docs/NumberBaseConversionApi.md#v1technologybasebatchpost) | **POST** /v1/technology/base/batch | Batch Convert Base
*NumberBaseConversionApi* | [**v1TechnologyBaseGet**](docs/NumberBaseConversionApi.md#v1technologybaseget) | **GET** /v1/technology/base | Convert Base
*PasswordGeneratorApi* | [**v1TechnologyPasswordBatchPost**](docs/PasswordGeneratorApi.md#v1technologypasswordbatchpost) | **POST** /v1/technology/password/batch | Batch Generate Passwords
*PasswordGeneratorApi* | [**v1TechnologyPasswordGet**](docs/PasswordGeneratorApi.md#v1technologypasswordget) | **GET** /v1/technology/password | Generate Password
*PaymentsIntelligenceApi* | [**v1SystemsPaymentValidatePost**](docs/PaymentsIntelligenceApi.md#v1systemspaymentvalidatepost) | **POST** /v1/systems/payment/validate | Validate Payment
*PaymentsIntelligenceApi* | [**v1SystemsTransactionRiskPost**](docs/PaymentsIntelligenceApi.md#v1systemstransactionriskpost) | **POST** /v1/systems/transaction/risk | Score Transaction Risk
*PhoneValidationApi* | [**v1ValidationPhoneBatchPost**](docs/PhoneValidationApi.md#v1validationphonebatchpost) | **POST** /v1/validation/phone/batch | Batch Validate Phone Numbers
*PhoneValidationApi* | [**v1ValidationPhoneGet**](docs/PhoneValidationApi.md#v1validationphoneget) | **GET** /v1/validation/phone | Validate Phone Number
*PostalCodeApi* | [**v1PlacesPostalBatchPost**](docs/PostalCodeApi.md#v1placespostalbatchpost) | **POST** /v1/places/postal/batch | Batch Lookup Postal Codes
*PostalCodeApi* | [**v1PlacesPostalCodeGet**](docs/PostalCodeApi.md#v1placespostalcodeget) | **GET** /v1/places/postal/{code} | Lookup Postal Code
*ProfanityApi* | [**v1ValidationProfanityBatchPost**](docs/ProfanityApi.md#v1validationprofanitybatchpost) | **POST** /v1/validation/profanity/batch | Batch Check Profanity
*ProfanityApi* | [**v1ValidationProfanityPost**](docs/ProfanityApi.md#v1validationprofanitypost) | **POST** /v1/validation/profanity | Check Profanity
*QrCodeApi* | [**v1TechnologyQrBase64BatchPost**](docs/QrCodeApi.md#v1technologyqrbase64batchpost) | **POST** /v1/technology/qr/base64/batch | Batch Generate QR Codes (Base64)
*QrCodeApi* | [**v1TechnologyQrBase64Get**](docs/QrCodeApi.md#v1technologyqrbase64get) | **GET** /v1/technology/qr/base64 | Generate QR Code (Base64 JSON)
*QrCodeApi* | [**v1TechnologyQrGet**](docs/QrCodeApi.md#v1technologyqrget) | **GET** /v1/technology/qr | Generate QR Code (PNG)
*QuotesApi* | [**v1EntertainmentQuotesRandomBatchPost**](docs/QuotesApi.md#v1entertainmentquotesrandombatchpost) | **POST** /v1/entertainment/quotes/random/batch | Get Random Quotes (Batch)
*QuotesApi* | [**v1EntertainmentQuotesRandomGet**](docs/QuotesApi.md#v1entertainmentquotesrandomget) | **GET** /v1/entertainment/quotes/random | Get Random Quote
*RandomUserApi* | [**v1TechnologyRandomUserBatchPost**](docs/RandomUserApi.md#v1technologyrandomuserbatchpost) | **POST** /v1/technology/random-user/batch | Batch Generate Users
*RandomUserApi* | [**v1TechnologyRandomUserGet**](docs/RandomUserApi.md#v1technologyrandomuserget) | **GET** /v1/technology/random-user | Get Random User
*RandomWordApi* | [**v1TextWordsBatchPost**](docs/RandomWordApi.md#v1textwordsbatchpost) | **POST** /v1/text/words/batch | Batch Define Words
*RandomWordApi* | [**v1TextWordsRandomGet**](docs/RandomWordApi.md#v1textwordsrandomget) | **GET** /v1/text/words/random | Get Random Word
*SentimentApi* | [**v1TextSentimentBatchPost**](docs/SentimentApi.md#v1textsentimentbatchpost) | **POST** /v1/text/sentiment/batch | Analyze Sentiment (Batch)
*SentimentApi* | [**v1TextSentimentPost**](docs/SentimentApi.md#v1textsentimentpost) | **POST** /v1/text/sentiment | Analyze Sentiment
*SpellCheckApi* | [**v1TextSpellcheckBatchPost**](docs/SpellCheckApi.md#v1textspellcheckbatchpost) | **POST** /v1/text/spellcheck/batch | Check Spelling (Batch)
*SpellCheckApi* | [**v1TextSpellcheckPost**](docs/SpellCheckApi.md#v1textspellcheckpost) | **POST** /v1/text/spellcheck | Check Spelling
*SudokuApi* | [**v1EntertainmentSudokuBatchPost**](docs/SudokuApi.md#v1entertainmentsudokubatchpost) | **POST** /v1/entertainment/sudoku/batch | Batch Generate Sudoku Puzzles
*SudokuApi* | [**v1EntertainmentSudokuGet**](docs/SudokuApi.md#v1entertainmentsudokuget) | **GET** /v1/entertainment/sudoku | Get Sudoku Puzzle
*SwiftCodeApi* | [**v1FinanceSwiftCodeGet**](docs/SwiftCodeApi.md#v1financeswiftcodeget) | **GET** /v1/finance/swift/{code} | Get SWIFT Code
*SwiftCodeApi* | [**v1FinanceSwiftGet**](docs/SwiftCodeApi.md#v1financeswiftget) | **GET** /v1/finance/swift | List SWIFT Codes
*TextSimilarityApi* | [**v1TextSimilarityPost**](docs/TextSimilarityApi.md#v1textsimilaritypost) | **POST** /v1/text/similarity | Compare Text Similarity
*ThesaurusApi* | [**v1TextThesaurusWordGet**](docs/ThesaurusApi.md#v1textthesauruswordget) | **GET** /v1/text/thesaurus/{word} | Thesaurus Lookup
*TimezoneApi* | [**v1PlacesTimezoneBatchPost**](docs/TimezoneApi.md#v1placestimezonebatchpost) | **POST** /v1/places/timezone/batch | Batch Timezone Lookup
*TimezoneApi* | [**v1PlacesTimezoneGet**](docs/TimezoneApi.md#v1placestimezoneget) | **GET** /v1/places/timezone | Get Timezone
*TriviaApi* | [**v1EntertainmentTriviaBatchPost**](docs/TriviaApi.md#v1entertainmenttriviabatchpost) | **POST** /v1/entertainment/trivia/batch | Get Batch Trivia Questions
*TriviaApi* | [**v1EntertainmentTriviaGet**](docs/TriviaApi.md#v1entertainmenttriviaget) | **GET** /v1/entertainment/trivia | Get Trivia Question
*UnitConversionApi* | [**v1TechnologyConvertBatchPost**](docs/UnitConversionApi.md#v1technologyconvertbatchpost) | **POST** /v1/technology/convert/batch | Batch Convert Units
*UnitConversionApi* | [**v1TechnologyConvertGet**](docs/UnitConversionApi.md#v1technologyconvertget) | **GET** /v1/technology/convert | Convert Units
*UnitConversionApi* | [**v1TechnologyConvertUnitsGet**](docs/UnitConversionApi.md#v1technologyconvertunitsget) | **GET** /v1/technology/convert/units | List Available Units
*UseragentApi* | [**v1TechnologyUseragentBatchPost**](docs/UseragentApi.md#v1technologyuseragentbatchpost) | **POST** /v1/technology/useragent/batch | Batch Parse User Agents
*UseragentApi* | [**v1TechnologyUseragentGet**](docs/UseragentApi.md#v1technologyuseragentget) | **GET** /v1/technology/useragent | Parse User Agent
*VpnDetectionApi* | [**v1NetworkingIpVpnIpGet**](docs/VpnDetectionApi.md#v1networkingipvpnipget) | **GET** /v1/networking/ip/vpn/{ip} | Check IP Address
*WhoisApi* | [**v1NetworkingWhoisBatchPost**](docs/WhoisApi.md#v1networkingwhoisbatchpost) | **POST** /v1/networking/whois/batch | Batch WHOIS Lookup
*WhoisApi* | [**v1NetworkingWhoisDomainGet**](docs/WhoisApi.md#v1networkingwhoisdomainget) | **GET** /v1/networking/whois/{domain} | WHOIS Lookup
*WorkingDaysApi* | [**v1PlacesWorkingDaysGet**](docs/WorkingDaysApi.md#v1placesworkingdaysget) | **GET** /v1/places/working-days | Calculate Working Days
*WorldTimeApi* | [**v1PlacesTimeTimezoneGet**](docs/WorldTimeApi.md#v1placestimetimezoneget) | **GET** /v1/places/time/{timezone} | Get Current Time by Timezone


### Documentation For Models

 - [V1EntertainmentChuckNorrisGet200Response](docs/V1EntertainmentChuckNorrisGet200Response.md)
 - [V1EntertainmentChuckNorrisGet200ResponseData](docs/V1EntertainmentChuckNorrisGet200ResponseData.md)
 - [V1EntertainmentEmojiNameGet200Response](docs/V1EntertainmentEmojiNameGet200Response.md)
 - [V1EntertainmentEmojiNameGet200ResponseData](docs/V1EntertainmentEmojiNameGet200ResponseData.md)
 - [V1EntertainmentEmojiRandomGet200Response](docs/V1EntertainmentEmojiRandomGet200Response.md)
 - [V1EntertainmentEmojiRandomGet200ResponseData](docs/V1EntertainmentEmojiRandomGet200ResponseData.md)
 - [V1EntertainmentEmojiSearchGet200Response](docs/V1EntertainmentEmojiSearchGet200Response.md)
 - [V1EntertainmentEmojiSearchGet200ResponseData](docs/V1EntertainmentEmojiSearchGet200ResponseData.md)
 - [V1EntertainmentFactsGet200Response](docs/V1EntertainmentFactsGet200Response.md)
 - [V1EntertainmentFactsGet200ResponseData](docs/V1EntertainmentFactsGet200ResponseData.md)
 - [V1EntertainmentHoroscopeBatchPost200Response](docs/V1EntertainmentHoroscopeBatchPost200Response.md)
 - [V1EntertainmentHoroscopeBatchPost200ResponseData](docs/V1EntertainmentHoroscopeBatchPost200ResponseData.md)
 - [V1EntertainmentHoroscopeBatchPostRequest](docs/V1EntertainmentHoroscopeBatchPostRequest.md)
 - [V1EntertainmentHoroscopeSignGet200Response](docs/V1EntertainmentHoroscopeSignGet200Response.md)
 - [V1EntertainmentHoroscopeSignGet200ResponseData](docs/V1EntertainmentHoroscopeSignGet200ResponseData.md)
 - [V1EntertainmentJokesDadGet200Response](docs/V1EntertainmentJokesDadGet200Response.md)
 - [V1EntertainmentJokesDadGet200ResponseData](docs/V1EntertainmentJokesDadGet200ResponseData.md)
 - [V1EntertainmentQuotesRandomBatchPost200Response](docs/V1EntertainmentQuotesRandomBatchPost200Response.md)
 - [V1EntertainmentQuotesRandomBatchPost200ResponseData](docs/V1EntertainmentQuotesRandomBatchPost200ResponseData.md)
 - [V1EntertainmentQuotesRandomBatchPostRequest](docs/V1EntertainmentQuotesRandomBatchPostRequest.md)
 - [V1EntertainmentQuotesRandomGet200Response](docs/V1EntertainmentQuotesRandomGet200Response.md)
 - [V1EntertainmentQuotesRandomGet200ResponseData](docs/V1EntertainmentQuotesRandomGet200ResponseData.md)
 - [V1EntertainmentSudokuBatchPost200Response](docs/V1EntertainmentSudokuBatchPost200Response.md)
 - [V1EntertainmentSudokuBatchPost200ResponseData](docs/V1EntertainmentSudokuBatchPost200ResponseData.md)
 - [V1EntertainmentSudokuBatchPostRequest](docs/V1EntertainmentSudokuBatchPostRequest.md)
 - [V1EntertainmentSudokuGet200Response](docs/V1EntertainmentSudokuGet200Response.md)
 - [V1EntertainmentSudokuGet200ResponseData](docs/V1EntertainmentSudokuGet200ResponseData.md)
 - [V1EntertainmentTriviaBatchPost200Response](docs/V1EntertainmentTriviaBatchPost200Response.md)
 - [V1EntertainmentTriviaBatchPost200ResponseData](docs/V1EntertainmentTriviaBatchPost200ResponseData.md)
 - [V1EntertainmentTriviaBatchPostRequest](docs/V1EntertainmentTriviaBatchPostRequest.md)
 - [V1EntertainmentTriviaGet200Response](docs/V1EntertainmentTriviaGet200Response.md)
 - [V1EntertainmentTriviaGet200ResponseData](docs/V1EntertainmentTriviaGet200ResponseData.md)
 - [V1FinanceBinBinGet200Response](docs/V1FinanceBinBinGet200Response.md)
 - [V1FinanceBinBinGet200ResponseData](docs/V1FinanceBinBinGet200ResponseData.md)
 - [V1FinanceCommoditiesCommodityGet200Response](docs/V1FinanceCommoditiesCommodityGet200Response.md)
 - [V1FinanceCommoditiesCommodityGet200ResponseData](docs/V1FinanceCommoditiesCommodityGet200ResponseData.md)
 - [V1FinanceConvertGet200Response](docs/V1FinanceConvertGet200Response.md)
 - [V1FinanceConvertGet200ResponseData](docs/V1FinanceConvertGet200ResponseData.md)
 - [V1FinanceCryptoSymbolGet200Response](docs/V1FinanceCryptoSymbolGet200Response.md)
 - [V1FinanceCryptoSymbolGet200ResponseData](docs/V1FinanceCryptoSymbolGet200ResponseData.md)
 - [V1FinanceExchangeRateGet200Response](docs/V1FinanceExchangeRateGet200Response.md)
 - [V1FinanceExchangeRateGet200ResponseData](docs/V1FinanceExchangeRateGet200ResponseData.md)
 - [V1FinanceIbanBatchPost200Response](docs/V1FinanceIbanBatchPost200Response.md)
 - [V1FinanceIbanBatchPost200ResponseData](docs/V1FinanceIbanBatchPost200ResponseData.md)
 - [V1FinanceIbanBatchPostRequest](docs/V1FinanceIbanBatchPostRequest.md)
 - [V1FinanceIbanIbanGet200Response](docs/V1FinanceIbanIbanGet200Response.md)
 - [V1FinanceIbanIbanGet200ResponseData](docs/V1FinanceIbanIbanGet200ResponseData.md)
 - [V1FinanceInflationBatchPost200Response](docs/V1FinanceInflationBatchPost200Response.md)
 - [V1FinanceInflationBatchPost200ResponseData](docs/V1FinanceInflationBatchPost200ResponseData.md)
 - [V1FinanceInflationBatchPostRequest](docs/V1FinanceInflationBatchPostRequest.md)
 - [V1FinanceInflationGet200Response](docs/V1FinanceInflationGet200Response.md)
 - [V1FinanceInflationGet200ResponseData](docs/V1FinanceInflationGet200ResponseData.md)
 - [V1FinanceMortgageBatchPost200Response](docs/V1FinanceMortgageBatchPost200Response.md)
 - [V1FinanceMortgageBatchPost200ResponseData](docs/V1FinanceMortgageBatchPost200ResponseData.md)
 - [V1FinanceMortgageBatchPostRequest](docs/V1FinanceMortgageBatchPostRequest.md)
 - [V1FinanceMortgageGet200Response](docs/V1FinanceMortgageGet200Response.md)
 - [V1FinanceMortgageGet200ResponseData](docs/V1FinanceMortgageGet200ResponseData.md)
 - [V1FinanceSwiftCodeGet200Response](docs/V1FinanceSwiftCodeGet200Response.md)
 - [V1FinanceSwiftCodeGet200ResponseData](docs/V1FinanceSwiftCodeGet200ResponseData.md)
 - [V1HealthBodyPartsGet200Response](docs/V1HealthBodyPartsGet200Response.md)
 - [V1HealthBodyPartsGet200ResponseData](docs/V1HealthBodyPartsGet200ResponseData.md)
 - [V1HealthEquipmentGet200Response](docs/V1HealthEquipmentGet200Response.md)
 - [V1HealthEquipmentGet200ResponseData](docs/V1HealthEquipmentGet200ResponseData.md)
 - [V1HealthExercisesBatchPost200Response](docs/V1HealthExercisesBatchPost200Response.md)
 - [V1HealthExercisesBatchPost200ResponseData](docs/V1HealthExercisesBatchPost200ResponseData.md)
 - [V1HealthExercisesBatchPostRequest](docs/V1HealthExercisesBatchPostRequest.md)
 - [V1HealthExercisesGet200Response](docs/V1HealthExercisesGet200Response.md)
 - [V1HealthExercisesGet200ResponseData](docs/V1HealthExercisesGet200ResponseData.md)
 - [V1HealthExercisesIdGet200Response](docs/V1HealthExercisesIdGet200Response.md)
 - [V1HealthExercisesIdGet200ResponseData](docs/V1HealthExercisesIdGet200ResponseData.md)
 - [V1HealthExercisesRandomGet200Response](docs/V1HealthExercisesRandomGet200Response.md)
 - [V1HealthMusclesGet200Response](docs/V1HealthMusclesGet200Response.md)
 - [V1HealthMusclesGet200ResponseData](docs/V1HealthMusclesGet200ResponseData.md)
 - [V1NetworkingDisposableBatchPost200Response](docs/V1NetworkingDisposableBatchPost200Response.md)
 - [V1NetworkingDisposableBatchPost200ResponseData](docs/V1NetworkingDisposableBatchPost200ResponseData.md)
 - [V1NetworkingDisposableBatchPostRequest](docs/V1NetworkingDisposableBatchPostRequest.md)
 - [V1NetworkingDisposableCheckPost200Response](docs/V1NetworkingDisposableCheckPost200Response.md)
 - [V1NetworkingDisposableCheckPost200ResponseData](docs/V1NetworkingDisposableCheckPost200ResponseData.md)
 - [V1NetworkingDisposableCheckPostRequest](docs/V1NetworkingDisposableCheckPostRequest.md)
 - [V1NetworkingDisposableDomainDomainGet200Response](docs/V1NetworkingDisposableDomainDomainGet200Response.md)
 - [V1NetworkingDisposableDomainDomainGet200ResponseData](docs/V1NetworkingDisposableDomainDomainGet200ResponseData.md)
 - [V1NetworkingDisposableDomainsGet200Response](docs/V1NetworkingDisposableDomainsGet200Response.md)
 - [V1NetworkingDisposableDomainsGet200ResponseData](docs/V1NetworkingDisposableDomainsGet200ResponseData.md)
 - [V1NetworkingDisposableStatsGet200Response](docs/V1NetworkingDisposableStatsGet200Response.md)
 - [V1NetworkingDisposableStatsGet200ResponseData](docs/V1NetworkingDisposableStatsGet200ResponseData.md)
 - [V1NetworkingDomainDomainGet200Response](docs/V1NetworkingDomainDomainGet200Response.md)
 - [V1NetworkingDomainDomainGet200ResponseData](docs/V1NetworkingDomainDomainGet200ResponseData.md)
 - [V1NetworkingIpAsnBatchPost200Response](docs/V1NetworkingIpAsnBatchPost200Response.md)
 - [V1NetworkingIpAsnBatchPost200ResponseData](docs/V1NetworkingIpAsnBatchPost200ResponseData.md)
 - [V1NetworkingIpAsnBatchPostRequest](docs/V1NetworkingIpAsnBatchPostRequest.md)
 - [V1NetworkingIpAsnGet200Response](docs/V1NetworkingIpAsnGet200Response.md)
 - [V1NetworkingIpAsnGet200ResponseData](docs/V1NetworkingIpAsnGet200ResponseData.md)
 - [V1NetworkingIpAsnIpGet200Response](docs/V1NetworkingIpAsnIpGet200Response.md)
 - [V1NetworkingIpAsnIpGet200ResponseData](docs/V1NetworkingIpAsnIpGet200ResponseData.md)
 - [V1NetworkingIpGet200Response](docs/V1NetworkingIpGet200Response.md)
 - [V1NetworkingIpGet200ResponseData](docs/V1NetworkingIpGet200ResponseData.md)
 - [V1NetworkingIpInfoBatchPost200Response](docs/V1NetworkingIpInfoBatchPost200Response.md)
 - [V1NetworkingIpInfoBatchPost200ResponseData](docs/V1NetworkingIpInfoBatchPost200ResponseData.md)
 - [V1NetworkingIpIpGet200Response](docs/V1NetworkingIpIpGet200Response.md)
 - [V1NetworkingIpIpGet200ResponseData](docs/V1NetworkingIpIpGet200ResponseData.md)
 - [V1NetworkingIpVpnIpGet200Response](docs/V1NetworkingIpVpnIpGet200Response.md)
 - [V1NetworkingIpVpnIpGet200ResponseData](docs/V1NetworkingIpVpnIpGet200ResponseData.md)
 - [V1NetworkingMxBatchPostRequest](docs/V1NetworkingMxBatchPostRequest.md)
 - [V1NetworkingMxDomainGet200Response](docs/V1NetworkingMxDomainGet200Response.md)
 - [V1NetworkingMxDomainGet200ResponseData](docs/V1NetworkingMxDomainGet200ResponseData.md)
 - [V1NetworkingWhoisBatchPost200Response](docs/V1NetworkingWhoisBatchPost200Response.md)
 - [V1NetworkingWhoisBatchPost200ResponseData](docs/V1NetworkingWhoisBatchPost200ResponseData.md)
 - [V1NetworkingWhoisBatchPostRequest](docs/V1NetworkingWhoisBatchPostRequest.md)
 - [V1NetworkingWhoisDomainGet200Response](docs/V1NetworkingWhoisDomainGet200Response.md)
 - [V1NetworkingWhoisDomainGet200ResponseData](docs/V1NetworkingWhoisDomainGet200ResponseData.md)
 - [V1PlacesCitiesCityGet200Response](docs/V1PlacesCitiesCityGet200Response.md)
 - [V1PlacesCitiesCityGet200ResponseData](docs/V1PlacesCitiesCityGet200ResponseData.md)
 - [V1PlacesGeocodeBatchPost200Response](docs/V1PlacesGeocodeBatchPost200Response.md)
 - [V1PlacesGeocodeBatchPost200ResponseData](docs/V1PlacesGeocodeBatchPost200ResponseData.md)
 - [V1PlacesGeocodeBatchPostRequest](docs/V1PlacesGeocodeBatchPostRequest.md)
 - [V1PlacesGeocodeGet200Response](docs/V1PlacesGeocodeGet200Response.md)
 - [V1PlacesGeocodeGet200ResponseData](docs/V1PlacesGeocodeGet200ResponseData.md)
 - [V1PlacesHolidaysBatchPost200Response](docs/V1PlacesHolidaysBatchPost200Response.md)
 - [V1PlacesHolidaysBatchPost200ResponseData](docs/V1PlacesHolidaysBatchPost200ResponseData.md)
 - [V1PlacesHolidaysBatchPostRequest](docs/V1PlacesHolidaysBatchPostRequest.md)
 - [V1PlacesHolidaysGet200Response](docs/V1PlacesHolidaysGet200Response.md)
 - [V1PlacesHolidaysGet200ResponseData](docs/V1PlacesHolidaysGet200ResponseData.md)
 - [V1PlacesPostalBatchPost200Response](docs/V1PlacesPostalBatchPost200Response.md)
 - [V1PlacesPostalBatchPost200ResponseData](docs/V1PlacesPostalBatchPost200ResponseData.md)
 - [V1PlacesPostalBatchPostRequest](docs/V1PlacesPostalBatchPostRequest.md)
 - [V1PlacesPostalCodeGet200Response](docs/V1PlacesPostalCodeGet200Response.md)
 - [V1PlacesPostalCodeGet200ResponseData](docs/V1PlacesPostalCodeGet200ResponseData.md)
 - [V1PlacesReverseGeocodeBatchPost200Response](docs/V1PlacesReverseGeocodeBatchPost200Response.md)
 - [V1PlacesReverseGeocodeBatchPost200ResponseData](docs/V1PlacesReverseGeocodeBatchPost200ResponseData.md)
 - [V1PlacesReverseGeocodeBatchPostRequest](docs/V1PlacesReverseGeocodeBatchPostRequest.md)
 - [V1PlacesReverseGeocodeGet200Response](docs/V1PlacesReverseGeocodeGet200Response.md)
 - [V1PlacesReverseGeocodeGet200ResponseData](docs/V1PlacesReverseGeocodeGet200ResponseData.md)
 - [V1PlacesTimeTimezoneGet200Response](docs/V1PlacesTimeTimezoneGet200Response.md)
 - [V1PlacesTimeTimezoneGet200ResponseData](docs/V1PlacesTimeTimezoneGet200ResponseData.md)
 - [V1PlacesTimezoneBatchPost200Response](docs/V1PlacesTimezoneBatchPost200Response.md)
 - [V1PlacesTimezoneBatchPost200ResponseData](docs/V1PlacesTimezoneBatchPost200ResponseData.md)
 - [V1PlacesTimezoneBatchPostRequest](docs/V1PlacesTimezoneBatchPostRequest.md)
 - [V1PlacesTimezoneGet200Response](docs/V1PlacesTimezoneGet200Response.md)
 - [V1PlacesTimezoneGet200ResponseData](docs/V1PlacesTimezoneGet200ResponseData.md)
 - [V1PlacesWorkingDaysGet200Response](docs/V1PlacesWorkingDaysGet200Response.md)
 - [V1PlacesWorkingDaysGet200ResponseData](docs/V1PlacesWorkingDaysGet200ResponseData.md)
 - [V1SystemsBusinessCalendarCountryGet200Response](docs/V1SystemsBusinessCalendarCountryGet200Response.md)
 - [V1SystemsBusinessCalendarCountryGet200ResponseData](docs/V1SystemsBusinessCalendarCountryGet200ResponseData.md)
 - [V1SystemsContentModeratePost200Response](docs/V1SystemsContentModeratePost200Response.md)
 - [V1SystemsContentModeratePost200ResponseData](docs/V1SystemsContentModeratePost200ResponseData.md)
 - [V1SystemsContentModeratePostRequest](docs/V1SystemsContentModeratePostRequest.md)
 - [V1SystemsLocationResolvePost200Response](docs/V1SystemsLocationResolvePost200Response.md)
 - [V1SystemsLocationResolvePost200ResponseData](docs/V1SystemsLocationResolvePost200ResponseData.md)
 - [V1SystemsLocationResolvePostRequest](docs/V1SystemsLocationResolvePostRequest.md)
 - [V1SystemsPaymentValidatePost200Response](docs/V1SystemsPaymentValidatePost200Response.md)
 - [V1SystemsPaymentValidatePost200ResponseData](docs/V1SystemsPaymentValidatePost200ResponseData.md)
 - [V1SystemsPaymentValidatePostRequest](docs/V1SystemsPaymentValidatePostRequest.md)
 - [V1SystemsRiskScorePost200Response](docs/V1SystemsRiskScorePost200Response.md)
 - [V1SystemsRiskScorePost200ResponseData](docs/V1SystemsRiskScorePost200ResponseData.md)
 - [V1SystemsRiskScorePostRequest](docs/V1SystemsRiskScorePostRequest.md)
 - [V1SystemsSignupProtectPost200Response](docs/V1SystemsSignupProtectPost200Response.md)
 - [V1SystemsSignupProtectPost200ResponseData](docs/V1SystemsSignupProtectPost200ResponseData.md)
 - [V1SystemsSignupProtectPostRequest](docs/V1SystemsSignupProtectPostRequest.md)
 - [V1SystemsTextNormalizePost200Response](docs/V1SystemsTextNormalizePost200Response.md)
 - [V1SystemsTextNormalizePost200ResponseData](docs/V1SystemsTextNormalizePost200ResponseData.md)
 - [V1SystemsTextNormalizePostRequest](docs/V1SystemsTextNormalizePostRequest.md)
 - [V1SystemsTimezoneFromIpIpGet200Response](docs/V1SystemsTimezoneFromIpIpGet200Response.md)
 - [V1SystemsTimezoneFromIpIpGet200ResponseData](docs/V1SystemsTimezoneFromIpIpGet200ResponseData.md)
 - [V1SystemsTransactionRiskPost200Response](docs/V1SystemsTransactionRiskPost200Response.md)
 - [V1SystemsTransactionRiskPost200ResponseData](docs/V1SystemsTransactionRiskPost200ResponseData.md)
 - [V1SystemsTransactionRiskPostRequest](docs/V1SystemsTransactionRiskPostRequest.md)
 - [V1SystemsUserVerifyPost200Response](docs/V1SystemsUserVerifyPost200Response.md)
 - [V1SystemsUserVerifyPost200ResponseData](docs/V1SystemsUserVerifyPost200ResponseData.md)
 - [V1SystemsUserVerifyPostRequest](docs/V1SystemsUserVerifyPostRequest.md)
 - [V1TechnologyBarcodeBase64Get200Response](docs/V1TechnologyBarcodeBase64Get200Response.md)
 - [V1TechnologyBarcodeBase64Get200ResponseData](docs/V1TechnologyBarcodeBase64Get200ResponseData.md)
 - [V1TechnologyBase64DecodePost200Response](docs/V1TechnologyBase64DecodePost200Response.md)
 - [V1TechnologyBase64DecodePost200ResponseData](docs/V1TechnologyBase64DecodePost200ResponseData.md)
 - [V1TechnologyBase64DecodePostRequest](docs/V1TechnologyBase64DecodePostRequest.md)
 - [V1TechnologyBase64EncodePost200Response](docs/V1TechnologyBase64EncodePost200Response.md)
 - [V1TechnologyBase64EncodePost200ResponseData](docs/V1TechnologyBase64EncodePost200ResponseData.md)
 - [V1TechnologyBase64EncodePostRequest](docs/V1TechnologyBase64EncodePostRequest.md)
 - [V1TechnologyBaseBatchPost200Response](docs/V1TechnologyBaseBatchPost200Response.md)
 - [V1TechnologyBaseBatchPost200ResponseData](docs/V1TechnologyBaseBatchPost200ResponseData.md)
 - [V1TechnologyBaseBatchPostRequest](docs/V1TechnologyBaseBatchPostRequest.md)
 - [V1TechnologyBaseGet200Response](docs/V1TechnologyBaseGet200Response.md)
 - [V1TechnologyBaseGet200ResponseData](docs/V1TechnologyBaseGet200ResponseData.md)
 - [V1TechnologyColorGet200Response](docs/V1TechnologyColorGet200Response.md)
 - [V1TechnologyColorGet200ResponseData](docs/V1TechnologyColorGet200ResponseData.md)
 - [V1TechnologyConvertBatchPost200Response](docs/V1TechnologyConvertBatchPost200Response.md)
 - [V1TechnologyConvertBatchPost200ResponseData](docs/V1TechnologyConvertBatchPost200ResponseData.md)
 - [V1TechnologyConvertBatchPostRequest](docs/V1TechnologyConvertBatchPostRequest.md)
 - [V1TechnologyConvertGet200Response](docs/V1TechnologyConvertGet200Response.md)
 - [V1TechnologyConvertGet200ResponseData](docs/V1TechnologyConvertGet200ResponseData.md)
 - [V1TechnologyConvertUnitsGet200Response](docs/V1TechnologyConvertUnitsGet200Response.md)
 - [V1TechnologyConvertUnitsGet200ResponseData](docs/V1TechnologyConvertUnitsGet200ResponseData.md)
 - [V1TechnologyCounterNamespaceGet200Response](docs/V1TechnologyCounterNamespaceGet200Response.md)
 - [V1TechnologyCounterNamespaceGet200ResponseData](docs/V1TechnologyCounterNamespaceGet200ResponseData.md)
 - [V1TechnologyCounterNamespacePost200Response](docs/V1TechnologyCounterNamespacePost200Response.md)
 - [V1TechnologyCounterNamespacePost200ResponseData](docs/V1TechnologyCounterNamespacePost200ResponseData.md)
 - [V1TechnologyFormatPost200Response](docs/V1TechnologyFormatPost200Response.md)
 - [V1TechnologyFormatPost200ResponseData](docs/V1TechnologyFormatPost200ResponseData.md)
 - [V1TechnologyFormatPostRequest](docs/V1TechnologyFormatPostRequest.md)
 - [V1TechnologyMarkdownBatchPost200Response](docs/V1TechnologyMarkdownBatchPost200Response.md)
 - [V1TechnologyMarkdownBatchPost200ResponseData](docs/V1TechnologyMarkdownBatchPost200ResponseData.md)
 - [V1TechnologyMarkdownBatchPostRequest](docs/V1TechnologyMarkdownBatchPostRequest.md)
 - [V1TechnologyMarkdownPost200Response](docs/V1TechnologyMarkdownPost200Response.md)
 - [V1TechnologyMarkdownPost200ResponseData](docs/V1TechnologyMarkdownPost200ResponseData.md)
 - [V1TechnologyMarkdownPostRequest](docs/V1TechnologyMarkdownPostRequest.md)
 - [V1TechnologyPasswordBatchPost200Response](docs/V1TechnologyPasswordBatchPost200Response.md)
 - [V1TechnologyPasswordBatchPost200ResponseData](docs/V1TechnologyPasswordBatchPost200ResponseData.md)
 - [V1TechnologyPasswordBatchPostRequest](docs/V1TechnologyPasswordBatchPostRequest.md)
 - [V1TechnologyPasswordGet200Response](docs/V1TechnologyPasswordGet200Response.md)
 - [V1TechnologyPasswordGet200ResponseData](docs/V1TechnologyPasswordGet200ResponseData.md)
 - [V1TechnologyQrBase64BatchPost200Response](docs/V1TechnologyQrBase64BatchPost200Response.md)
 - [V1TechnologyQrBase64BatchPost200ResponseData](docs/V1TechnologyQrBase64BatchPost200ResponseData.md)
 - [V1TechnologyQrBase64BatchPostRequest](docs/V1TechnologyQrBase64BatchPostRequest.md)
 - [V1TechnologyQrBase64Get200Response](docs/V1TechnologyQrBase64Get200Response.md)
 - [V1TechnologyQrBase64Get200ResponseData](docs/V1TechnologyQrBase64Get200ResponseData.md)
 - [V1TechnologyRandomUserBatchPost200Response](docs/V1TechnologyRandomUserBatchPost200Response.md)
 - [V1TechnologyRandomUserBatchPost200ResponseData](docs/V1TechnologyRandomUserBatchPost200ResponseData.md)
 - [V1TechnologyRandomUserBatchPostRequest](docs/V1TechnologyRandomUserBatchPostRequest.md)
 - [V1TechnologyRandomUserGet200Response](docs/V1TechnologyRandomUserGet200Response.md)
 - [V1TechnologyRandomUserGet200ResponseData](docs/V1TechnologyRandomUserGet200ResponseData.md)
 - [V1TechnologyUseragentBatchPost200Response](docs/V1TechnologyUseragentBatchPost200Response.md)
 - [V1TechnologyUseragentBatchPost200ResponseData](docs/V1TechnologyUseragentBatchPost200ResponseData.md)
 - [V1TechnologyUseragentBatchPostRequest](docs/V1TechnologyUseragentBatchPostRequest.md)
 - [V1TechnologyUseragentGet200Response](docs/V1TechnologyUseragentGet200Response.md)
 - [V1TechnologyUseragentGet200ResponseData](docs/V1TechnologyUseragentGet200ResponseData.md)
 - [V1TextAdviceGet200Response](docs/V1TextAdviceGet200Response.md)
 - [V1TextAdviceGet200ResponseData](docs/V1TextAdviceGet200ResponseData.md)
 - [V1TextAdviceGet200ResponseMetadata](docs/V1TextAdviceGet200ResponseMetadata.md)
 - [V1TextDetectLanguagePost200Response](docs/V1TextDetectLanguagePost200Response.md)
 - [V1TextDetectLanguagePost200ResponseData](docs/V1TextDetectLanguagePost200ResponseData.md)
 - [V1TextDetectLanguagePostRequest](docs/V1TextDetectLanguagePostRequest.md)
 - [V1TextDictionaryWordGet200Response](docs/V1TextDictionaryWordGet200Response.md)
 - [V1TextDictionaryWordGet200ResponseData](docs/V1TextDictionaryWordGet200ResponseData.md)
 - [V1TextLoremBatchPost200Response](docs/V1TextLoremBatchPost200Response.md)
 - [V1TextLoremBatchPost200ResponseData](docs/V1TextLoremBatchPost200ResponseData.md)
 - [V1TextLoremBatchPostRequest](docs/V1TextLoremBatchPostRequest.md)
 - [V1TextLoremGet200Response](docs/V1TextLoremGet200Response.md)
 - [V1TextLoremGet200ResponseData](docs/V1TextLoremGet200ResponseData.md)
 - [V1TextNormalizeBatchPost200Response](docs/V1TextNormalizeBatchPost200Response.md)
 - [V1TextNormalizeBatchPost200ResponseData](docs/V1TextNormalizeBatchPost200ResponseData.md)
 - [V1TextNormalizeBatchPostRequest](docs/V1TextNormalizeBatchPostRequest.md)
 - [V1TextNormalizePost200Response](docs/V1TextNormalizePost200Response.md)
 - [V1TextNormalizePost200ResponseData](docs/V1TextNormalizePost200ResponseData.md)
 - [V1TextNormalizePostRequest](docs/V1TextNormalizePostRequest.md)
 - [V1TextSentimentBatchPost200Response](docs/V1TextSentimentBatchPost200Response.md)
 - [V1TextSentimentBatchPost200ResponseData](docs/V1TextSentimentBatchPost200ResponseData.md)
 - [V1TextSentimentBatchPostRequest](docs/V1TextSentimentBatchPostRequest.md)
 - [V1TextSentimentPost200Response](docs/V1TextSentimentPost200Response.md)
 - [V1TextSentimentPost200ResponseData](docs/V1TextSentimentPost200ResponseData.md)
 - [V1TextSentimentPostRequest](docs/V1TextSentimentPostRequest.md)
 - [V1TextSimilarityPost200Response](docs/V1TextSimilarityPost200Response.md)
 - [V1TextSimilarityPost200ResponseData](docs/V1TextSimilarityPost200ResponseData.md)
 - [V1TextSimilarityPostRequest](docs/V1TextSimilarityPostRequest.md)
 - [V1TextSpellcheckBatchPost200Response](docs/V1TextSpellcheckBatchPost200Response.md)
 - [V1TextSpellcheckBatchPost200ResponseData](docs/V1TextSpellcheckBatchPost200ResponseData.md)
 - [V1TextSpellcheckBatchPostRequest](docs/V1TextSpellcheckBatchPostRequest.md)
 - [V1TextSpellcheckPost200Response](docs/V1TextSpellcheckPost200Response.md)
 - [V1TextSpellcheckPost200ResponseData](docs/V1TextSpellcheckPost200ResponseData.md)
 - [V1TextSpellcheckPostRequest](docs/V1TextSpellcheckPostRequest.md)
 - [V1TextThesaurusWordGet200Response](docs/V1TextThesaurusWordGet200Response.md)
 - [V1TextThesaurusWordGet200ResponseData](docs/V1TextThesaurusWordGet200ResponseData.md)
 - [V1TextWordsBatchPost200Response](docs/V1TextWordsBatchPost200Response.md)
 - [V1TextWordsBatchPost200ResponseData](docs/V1TextWordsBatchPost200ResponseData.md)
 - [V1TextWordsBatchPostRequest](docs/V1TextWordsBatchPostRequest.md)
 - [V1TextWordsRandomGet200Response](docs/V1TextWordsRandomGet200Response.md)
 - [V1TextWordsRandomGet200ResponseData](docs/V1TextWordsRandomGet200ResponseData.md)
 - [V1ValidationEmailBatchPost200Response](docs/V1ValidationEmailBatchPost200Response.md)
 - [V1ValidationEmailBatchPost200ResponseData](docs/V1ValidationEmailBatchPost200ResponseData.md)
 - [V1ValidationEmailBatchPostRequest](docs/V1ValidationEmailBatchPostRequest.md)
 - [V1ValidationEmailPost200Response](docs/V1ValidationEmailPost200Response.md)
 - [V1ValidationEmailPost200ResponseData](docs/V1ValidationEmailPost200ResponseData.md)
 - [V1ValidationEmailPostRequest](docs/V1ValidationEmailPostRequest.md)
 - [V1ValidationPhoneBatchPostRequest](docs/V1ValidationPhoneBatchPostRequest.md)
 - [V1ValidationPhoneGet200Response](docs/V1ValidationPhoneGet200Response.md)
 - [V1ValidationPhoneGet200ResponseData](docs/V1ValidationPhoneGet200ResponseData.md)
 - [V1ValidationProfanityBatchPost200Response](docs/V1ValidationProfanityBatchPost200Response.md)
 - [V1ValidationProfanityBatchPost200ResponseData](docs/V1ValidationProfanityBatchPost200ResponseData.md)
 - [V1ValidationProfanityBatchPostRequest](docs/V1ValidationProfanityBatchPostRequest.md)
 - [V1ValidationProfanityPost200Response](docs/V1ValidationProfanityPost200Response.md)
 - [V1ValidationProfanityPost200ResponseData](docs/V1ValidationProfanityPost200ResponseData.md)
 - [V1ValidationProfanityPostRequest](docs/V1ValidationProfanityPostRequest.md)


<a id="documentation-for-authorization"></a>
## Documentation For Authorization


Authentication schemes defined for the API:
<a id="requiems-api-key"></a>
### requiems-api-key

- **Type**: API key
- **API key parameter name**: requiems-api-key
- **Location**: HTTP header

