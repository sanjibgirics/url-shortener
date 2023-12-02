package pkg

// model file will be used to handle the storage part. We are using in-memory
// storage here, so we will use map to store things related to url shortening.
// Can use db here also in the future.

// We will use two maps for faster retrieval.

// shortToActual map will contain generated short code which is used in the short url,
// as key and original url as value.
var shortCodeToOriginalURL = make(map[string]string)

// originalURLToShortCode map will contain original url as key and generated short code as value.
var originalURLToShortCode = make(map[string]string)

// domainsUsageMap map will contain number of times individual domains have used the shortener
// service
var domainsUsageMap = make(map[string]int)
