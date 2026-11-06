## api-client@1.0.0

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
npm install api-client@1.0.0 --save
```

_unPublished (not recommended):_

```
npm install PATH_TO_GENERATED_PACKAGE --save
```

### Documentation for API Endpoints

All URIs are relative to *https://api.requiems.xyz*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdviceApi* | [**v1TextAdviceGet**](docs/AdviceApi.md#v1textadviceget) | **GET** /v1/text/advice | Get Random Advice
*BarcodeApi* | [**v1TechnologyBarcodeBase64Get**](docs/BarcodeApi.md#v1technologybarcodebase64get) | **GET** /v1/technology/barcode/base64 | Generate Barcode (Base64 JSON)
*BarcodeApi* | [**v1TechnologyBarcodeGet**](docs/BarcodeApi.md#v1technologybarcodeget) | **GET** /v1/technology/barcode | Generate Barcode (PNG)
*Base64Api* | [**v1TechnologyBase64DecodePost**](docs/Base64Api.md#v1technologybase64decodepost) | **POST** /v1/technology/base64/decode | Decode
*Base64Api* | [**v1TechnologyBase64EncodePost**](docs/Base64Api.md#v1technologybase64encodepost) | **POST** /v1/technology/base64/encode | Encode
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
*DetectLanguageApi* | [**v1TextDetectLanguagePost**](docs/DetectLanguageApi.md#v1textdetectlanguagepost) | **POST** /v1/text/detect-language | Detect Language
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
*FitnessExercisesApi* | [**v1HealthExercisesGet**](docs/FitnessExercisesApi.md#v1healthexercisesget) | **GET** /v1/health/exercises | List Exercises
*FitnessExercisesApi* | [**v1HealthExercisesIdGet**](docs/FitnessExercisesApi.md#v1healthexercisesidget) | **GET** /v1/health/exercises/{id} | Get Exercise by ID
*FitnessExercisesApi* | [**v1HealthExercisesRandomGet**](docs/FitnessExercisesApi.md#v1healthexercisesrandomget) | **GET** /v1/health/exercises/random | Random Exercise
*FitnessExercisesApi* | [**v1HealthMusclesGet**](docs/FitnessExercisesApi.md#v1healthmusclesget) | **GET** /v1/health/muscles | List Muscles
*GeocodeApi* | [**v1PlacesGeocodeGet**](docs/GeocodeApi.md#v1placesgeocodeget) | **GET** /v1/places/geocode | Geocode Address
*GeocodeApi* | [**v1PlacesReverseGeocodeGet**](docs/GeocodeApi.md#v1placesreversegeocodeget) | **GET** /v1/places/reverse-geocode | Reverse Geocode
*HolidaysApi* | [**v1PlacesHolidaysBatchPost**](docs/HolidaysApi.md#v1placesholidaysbatchpost) | **POST** /v1/places/holidays/batch | Batch Get Holidays
*HolidaysApi* | [**v1PlacesHolidaysGet**](docs/HolidaysApi.md#v1placesholidaysget) | **GET** /v1/places/holidays | Get Holidays
*HoroscopeApi* | [**v1EntertainmentHoroscopeSignGet**](docs/HoroscopeApi.md#v1entertainmenthoroscopesignget) | **GET** /v1/entertainment/horoscope/{sign} | Get Daily Horoscope
*IbanApi* | [**v1FinanceIbanBatchPost**](docs/IbanApi.md#v1financeibanbatchpost) | **POST** /v1/finance/iban/batch | Batch Validate IBANs
*IbanApi* | [**v1FinanceIbanIbanGet**](docs/IbanApi.md#v1financeibanibanget) | **GET** /v1/finance/iban/{iban} | Validate IBAN
*InflationApi* | [**v1FinanceInflationBatchPost**](docs/InflationApi.md#v1financeinflationbatchpost) | **POST** /v1/finance/inflation/batch | Batch Inflation Rates
*InflationApi* | [**v1FinanceInflationGet**](docs/InflationApi.md#v1financeinflationget) | **GET** /v1/finance/inflation | Get Inflation Rate
*IpAsnApi* | [**v1NetworkingIpAsnGet**](docs/IpAsnApi.md#v1networkingipasnget) | **GET** /v1/networking/ip/asn | Lookup ASN (Caller IP)
*IpAsnApi* | [**v1NetworkingIpAsnIpGet**](docs/IpAsnApi.md#v1networkingipasnipget) | **GET** /v1/networking/ip/asn/{ip} | Lookup ASN for IP
*IpInfoApi* | [**v1NetworkingIpGet**](docs/IpInfoApi.md#v1networkingipget) | **GET** /v1/networking/ip | Get IP Info (Caller IP)
*IpInfoApi* | [**v1NetworkingIpIpGet**](docs/IpInfoApi.md#v1networkingipipget) | **GET** /v1/networking/ip/{ip} | Get IP Info for IP
*LoremIpsumApi* | [**v1TextLoremGet**](docs/LoremIpsumApi.md#v1textloremget) | **GET** /v1/text/lorem | Generate Lorem Ipsum
*MarkdownApi* | [**v1TechnologyMarkdownPost**](docs/MarkdownApi.md#v1technologymarkdownpost) | **POST** /v1/technology/markdown | Convert Markdown to HTML
*MortgageApi* | [**v1FinanceMortgageGet**](docs/MortgageApi.md#v1financemortgageget) | **GET** /v1/finance/mortgage | Calculate Mortgage
*MxLookupApi* | [**v1NetworkingMxDomainGet**](docs/MxLookupApi.md#v1networkingmxdomainget) | **GET** /v1/networking/mx/{domain} | MX Lookup
*NumberBaseConversionApi* | [**v1TechnologyBaseGet**](docs/NumberBaseConversionApi.md#v1technologybaseget) | **GET** /v1/technology/base | Convert Base
*PasswordGeneratorApi* | [**v1TechnologyPasswordGet**](docs/PasswordGeneratorApi.md#v1technologypasswordget) | **GET** /v1/technology/password | Generate Password
*PhoneValidationApi* | [**v1ValidationPhoneBatchPost**](docs/PhoneValidationApi.md#v1validationphonebatchpost) | **POST** /v1/validation/phone/batch | Batch Validate Phone Numbers
*PhoneValidationApi* | [**v1ValidationPhoneGet**](docs/PhoneValidationApi.md#v1validationphoneget) | **GET** /v1/validation/phone | Validate Phone Number
*PostalCodeApi* | [**v1PlacesPostalCodeGet**](docs/PostalCodeApi.md#v1placespostalcodeget) | **GET** /v1/places/postal/{code} | Lookup Postal Code
*ProfanityApi* | [**v1TextProfanityPost**](docs/ProfanityApi.md#v1textprofanitypost) | **POST** /v1/text/profanity | Check Profanity
*QrCodeApi* | [**v1TechnologyQrBase64Get**](docs/QrCodeApi.md#v1technologyqrbase64get) | **GET** /v1/technology/qr/base64 | Generate QR Code (Base64 JSON)
*QrCodeApi* | [**v1TechnologyQrGet**](docs/QrCodeApi.md#v1technologyqrget) | **GET** /v1/technology/qr | Generate QR Code (PNG)
*QuotesApi* | [**v1TextQuotesRandomGet**](docs/QuotesApi.md#v1textquotesrandomget) | **GET** /v1/text/quotes/random | Get Random Quote
*RandomUserApi* | [**v1TechnologyRandomUserGet**](docs/RandomUserApi.md#v1technologyrandomuserget) | **GET** /v1/technology/random-user | Get Random User
*RandomWordApi* | [**v1TextWordsRandomGet**](docs/RandomWordApi.md#v1textwordsrandomget) | **GET** /v1/text/words/random | Get Random Word
*SentimentApi* | [**v1TextSentimentPost**](docs/SentimentApi.md#v1textsentimentpost) | **POST** /v1/text/sentiment | Analyze Sentiment
*SpellCheckApi* | [**v1TextSpellcheckPost**](docs/SpellCheckApi.md#v1textspellcheckpost) | **POST** /v1/text/spellcheck | Check Spelling
*SudokuApi* | [**v1EntertainmentSudokuBatchPost**](docs/SudokuApi.md#v1entertainmentsudokubatchpost) | **POST** /v1/entertainment/sudoku/batch | Batch Generate Sudoku Puzzles
*SudokuApi* | [**v1EntertainmentSudokuGet**](docs/SudokuApi.md#v1entertainmentsudokuget) | **GET** /v1/entertainment/sudoku | Get Sudoku Puzzle
*SwiftCodeApi* | [**v1FinanceSwiftCodeGet**](docs/SwiftCodeApi.md#v1financeswiftcodeget) | **GET** /v1/finance/swift/{code} | Get SWIFT Code
*SwiftCodeApi* | [**v1FinanceSwiftGet**](docs/SwiftCodeApi.md#v1financeswiftget) | **GET** /v1/finance/swift | List SWIFT Codes
*TextSimilarityApi* | [**v1TextSimilarityPost**](docs/TextSimilarityApi.md#v1textsimilaritypost) | **POST** /v1/text/similarity | Compare Text Similarity
*ThesaurusApi* | [**v1TextThesaurusWordGet**](docs/ThesaurusApi.md#v1textthesauruswordget) | **GET** /v1/text/thesaurus/{word} | Thesaurus Lookup
*TimezoneApi* | [**v1PlacesTimezoneGet**](docs/TimezoneApi.md#v1placestimezoneget) | **GET** /v1/places/timezone | Get Timezone
*TriviaApi* | [**v1EntertainmentTriviaGet**](docs/TriviaApi.md#v1entertainmenttriviaget) | **GET** /v1/entertainment/trivia | Get Trivia Question
*UnitConversionApi* | [**v1TechnologyConvertGet**](docs/UnitConversionApi.md#v1technologyconvertget) | **GET** /v1/technology/convert | Convert Units
*UnitConversionApi* | [**v1TechnologyConvertUnitsGet**](docs/UnitConversionApi.md#v1technologyconvertunitsget) | **GET** /v1/technology/convert/units | List Available Units
*UseragentApi* | [**v1TechnologyUseragentGet**](docs/UseragentApi.md#v1technologyuseragentget) | **GET** /v1/technology/useragent | Parse User Agent
*VpnDetectionApi* | [**v1NetworkingIpVpnIpGet**](docs/VpnDetectionApi.md#v1networkingipvpnipget) | **GET** /v1/networking/ip/vpn/{ip} | Check IP Address
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
 - [V1EntertainmentHoroscopeSignGet200Response](docs/V1EntertainmentHoroscopeSignGet200Response.md)
 - [V1EntertainmentHoroscopeSignGet200ResponseData](docs/V1EntertainmentHoroscopeSignGet200ResponseData.md)
 - [V1EntertainmentJokesDadGet200Response](docs/V1EntertainmentJokesDadGet200Response.md)
 - [V1EntertainmentJokesDadGet200ResponseData](docs/V1EntertainmentJokesDadGet200ResponseData.md)
 - [V1EntertainmentSudokuBatchPost200Response](docs/V1EntertainmentSudokuBatchPost200Response.md)
 - [V1EntertainmentSudokuBatchPost200ResponseData](docs/V1EntertainmentSudokuBatchPost200ResponseData.md)
 - [V1EntertainmentSudokuBatchPostRequest](docs/V1EntertainmentSudokuBatchPostRequest.md)
 - [V1EntertainmentSudokuGet200Response](docs/V1EntertainmentSudokuGet200Response.md)
 - [V1EntertainmentSudokuGet200ResponseData](docs/V1EntertainmentSudokuGet200ResponseData.md)
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
 - [V1FinanceMortgageGet200Response](docs/V1FinanceMortgageGet200Response.md)
 - [V1FinanceMortgageGet200ResponseData](docs/V1FinanceMortgageGet200ResponseData.md)
 - [V1FinanceSwiftCodeGet200Response](docs/V1FinanceSwiftCodeGet200Response.md)
 - [V1FinanceSwiftCodeGet200ResponseData](docs/V1FinanceSwiftCodeGet200ResponseData.md)
 - [V1HealthBodyPartsGet200Response](docs/V1HealthBodyPartsGet200Response.md)
 - [V1HealthBodyPartsGet200ResponseData](docs/V1HealthBodyPartsGet200ResponseData.md)
 - [V1HealthEquipmentGet200Response](docs/V1HealthEquipmentGet200Response.md)
 - [V1HealthEquipmentGet200ResponseData](docs/V1HealthEquipmentGet200ResponseData.md)
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
 - [V1NetworkingIpAsnGet200Response](docs/V1NetworkingIpAsnGet200Response.md)
 - [V1NetworkingIpAsnGet200ResponseData](docs/V1NetworkingIpAsnGet200ResponseData.md)
 - [V1NetworkingIpAsnIpGet200Response](docs/V1NetworkingIpAsnIpGet200Response.md)
 - [V1NetworkingIpAsnIpGet200ResponseData](docs/V1NetworkingIpAsnIpGet200ResponseData.md)
 - [V1NetworkingIpGet200Response](docs/V1NetworkingIpGet200Response.md)
 - [V1NetworkingIpGet200ResponseData](docs/V1NetworkingIpGet200ResponseData.md)
 - [V1NetworkingIpIpGet200Response](docs/V1NetworkingIpIpGet200Response.md)
 - [V1NetworkingIpIpGet200ResponseData](docs/V1NetworkingIpIpGet200ResponseData.md)
 - [V1NetworkingIpVpnIpGet200Response](docs/V1NetworkingIpVpnIpGet200Response.md)
 - [V1NetworkingIpVpnIpGet200ResponseData](docs/V1NetworkingIpVpnIpGet200ResponseData.md)
 - [V1NetworkingMxDomainGet200Response](docs/V1NetworkingMxDomainGet200Response.md)
 - [V1NetworkingMxDomainGet200ResponseData](docs/V1NetworkingMxDomainGet200ResponseData.md)
 - [V1NetworkingWhoisDomainGet200Response](docs/V1NetworkingWhoisDomainGet200Response.md)
 - [V1NetworkingWhoisDomainGet200ResponseData](docs/V1NetworkingWhoisDomainGet200ResponseData.md)
 - [V1PlacesCitiesCityGet200Response](docs/V1PlacesCitiesCityGet200Response.md)
 - [V1PlacesCitiesCityGet200ResponseData](docs/V1PlacesCitiesCityGet200ResponseData.md)
 - [V1PlacesGeocodeGet200Response](docs/V1PlacesGeocodeGet200Response.md)
 - [V1PlacesGeocodeGet200ResponseData](docs/V1PlacesGeocodeGet200ResponseData.md)
 - [V1PlacesHolidaysBatchPost200Response](docs/V1PlacesHolidaysBatchPost200Response.md)
 - [V1PlacesHolidaysBatchPost200ResponseData](docs/V1PlacesHolidaysBatchPost200ResponseData.md)
 - [V1PlacesHolidaysBatchPostRequest](docs/V1PlacesHolidaysBatchPostRequest.md)
 - [V1PlacesHolidaysGet200Response](docs/V1PlacesHolidaysGet200Response.md)
 - [V1PlacesHolidaysGet200ResponseData](docs/V1PlacesHolidaysGet200ResponseData.md)
 - [V1PlacesPostalCodeGet200Response](docs/V1PlacesPostalCodeGet200Response.md)
 - [V1PlacesPostalCodeGet200ResponseData](docs/V1PlacesPostalCodeGet200ResponseData.md)
 - [V1PlacesReverseGeocodeGet200Response](docs/V1PlacesReverseGeocodeGet200Response.md)
 - [V1PlacesReverseGeocodeGet200ResponseData](docs/V1PlacesReverseGeocodeGet200ResponseData.md)
 - [V1PlacesTimeTimezoneGet200Response](docs/V1PlacesTimeTimezoneGet200Response.md)
 - [V1PlacesTimeTimezoneGet200ResponseData](docs/V1PlacesTimeTimezoneGet200ResponseData.md)
 - [V1PlacesTimezoneGet200Response](docs/V1PlacesTimezoneGet200Response.md)
 - [V1PlacesTimezoneGet200ResponseData](docs/V1PlacesTimezoneGet200ResponseData.md)
 - [V1PlacesWorkingDaysGet200Response](docs/V1PlacesWorkingDaysGet200Response.md)
 - [V1PlacesWorkingDaysGet200ResponseData](docs/V1PlacesWorkingDaysGet200ResponseData.md)
 - [V1TechnologyBarcodeBase64Get200Response](docs/V1TechnologyBarcodeBase64Get200Response.md)
 - [V1TechnologyBarcodeBase64Get200ResponseData](docs/V1TechnologyBarcodeBase64Get200ResponseData.md)
 - [V1TechnologyBase64DecodePost200Response](docs/V1TechnologyBase64DecodePost200Response.md)
 - [V1TechnologyBase64DecodePost200ResponseData](docs/V1TechnologyBase64DecodePost200ResponseData.md)
 - [V1TechnologyBase64DecodePostRequest](docs/V1TechnologyBase64DecodePostRequest.md)
 - [V1TechnologyBase64EncodePost200Response](docs/V1TechnologyBase64EncodePost200Response.md)
 - [V1TechnologyBase64EncodePost200ResponseData](docs/V1TechnologyBase64EncodePost200ResponseData.md)
 - [V1TechnologyBase64EncodePostRequest](docs/V1TechnologyBase64EncodePostRequest.md)
 - [V1TechnologyBaseGet200Response](docs/V1TechnologyBaseGet200Response.md)
 - [V1TechnologyBaseGet200ResponseData](docs/V1TechnologyBaseGet200ResponseData.md)
 - [V1TechnologyColorGet200Response](docs/V1TechnologyColorGet200Response.md)
 - [V1TechnologyColorGet200ResponseData](docs/V1TechnologyColorGet200ResponseData.md)
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
 - [V1TechnologyMarkdownPost200Response](docs/V1TechnologyMarkdownPost200Response.md)
 - [V1TechnologyMarkdownPost200ResponseData](docs/V1TechnologyMarkdownPost200ResponseData.md)
 - [V1TechnologyMarkdownPostRequest](docs/V1TechnologyMarkdownPostRequest.md)
 - [V1TechnologyPasswordGet200Response](docs/V1TechnologyPasswordGet200Response.md)
 - [V1TechnologyPasswordGet200ResponseData](docs/V1TechnologyPasswordGet200ResponseData.md)
 - [V1TechnologyQrBase64Get200Response](docs/V1TechnologyQrBase64Get200Response.md)
 - [V1TechnologyQrBase64Get200ResponseData](docs/V1TechnologyQrBase64Get200ResponseData.md)
 - [V1TechnologyRandomUserGet200Response](docs/V1TechnologyRandomUserGet200Response.md)
 - [V1TechnologyRandomUserGet200ResponseData](docs/V1TechnologyRandomUserGet200ResponseData.md)
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
 - [V1TextLoremGet200Response](docs/V1TextLoremGet200Response.md)
 - [V1TextLoremGet200ResponseData](docs/V1TextLoremGet200ResponseData.md)
 - [V1TextNormalizeBatchPost200Response](docs/V1TextNormalizeBatchPost200Response.md)
 - [V1TextNormalizeBatchPost200ResponseData](docs/V1TextNormalizeBatchPost200ResponseData.md)
 - [V1TextNormalizeBatchPostRequest](docs/V1TextNormalizeBatchPostRequest.md)
 - [V1TextNormalizePost200Response](docs/V1TextNormalizePost200Response.md)
 - [V1TextNormalizePost200ResponseData](docs/V1TextNormalizePost200ResponseData.md)
 - [V1TextNormalizePostRequest](docs/V1TextNormalizePostRequest.md)
 - [V1TextProfanityPost200Response](docs/V1TextProfanityPost200Response.md)
 - [V1TextProfanityPost200ResponseData](docs/V1TextProfanityPost200ResponseData.md)
 - [V1TextProfanityPostRequest](docs/V1TextProfanityPostRequest.md)
 - [V1TextQuotesRandomGet200Response](docs/V1TextQuotesRandomGet200Response.md)
 - [V1TextQuotesRandomGet200ResponseData](docs/V1TextQuotesRandomGet200ResponseData.md)
 - [V1TextSentimentPost200Response](docs/V1TextSentimentPost200Response.md)
 - [V1TextSentimentPost200ResponseData](docs/V1TextSentimentPost200ResponseData.md)
 - [V1TextSentimentPostRequest](docs/V1TextSentimentPostRequest.md)
 - [V1TextSimilarityPost200Response](docs/V1TextSimilarityPost200Response.md)
 - [V1TextSimilarityPost200ResponseData](docs/V1TextSimilarityPost200ResponseData.md)
 - [V1TextSimilarityPostRequest](docs/V1TextSimilarityPostRequest.md)
 - [V1TextSpellcheckPost200Response](docs/V1TextSpellcheckPost200Response.md)
 - [V1TextSpellcheckPost200ResponseData](docs/V1TextSpellcheckPost200ResponseData.md)
 - [V1TextSpellcheckPostRequest](docs/V1TextSpellcheckPostRequest.md)
 - [V1TextThesaurusWordGet200Response](docs/V1TextThesaurusWordGet200Response.md)
 - [V1TextThesaurusWordGet200ResponseData](docs/V1TextThesaurusWordGet200ResponseData.md)
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


<a id="documentation-for-authorization"></a>
## Documentation For Authorization


Authentication schemes defined for the API:
<a id="requiems-api-key"></a>
### requiems-api-key

- **Type**: API key
- **API key parameter name**: requiems-api-key
- **Location**: HTTP header

