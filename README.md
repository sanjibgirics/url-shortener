# URL Shortener

Shorten the length of given URL. Users can use the shortened URL to reach the original URL.
From the shortened URL they can also get back the original URL.

This also has a feature of redirection API i.e, redirecting to the original URL once
short URL is clicked.

It also has metric API which on call will return the top 3 domains that have used the URL shortener the most

## To build
From `cmd` directory execute `go build -o urlshortener`

## To run
`./urlshortener`, if built already following above.
To build and run directly from `cmd` directory execute `go run main.go`

## To call the APIs

Can use curl or postman or browser.

### To make short URL from long URL using curl
```
curl -X POST -H "Content-Type: application/json" -d "{\"url\": \"http://www.longurl.com/l/wefewfewfwefewfwefwefwefewfwe\"}" http://localhost:8083/shorturl
output:
{"originalURL":"http://www.longurl.com/l/wefewfewfwefewfwefwefwefewfwe","shortURL":"http://localhost:8083/s/X886sZ"}
```
### To get back original URL from the short URL using curl
```
curl -X POST -H "Content-Type: application/json" -d "{\"url\": \"http://localhost:8083/s/X886sZ\"}" http://localhost:8083/originalurl
output:
{"originalURL":"http://www.longurl.com/l/wefewfewfwefewfwefwefwefewfwe","shortURL":"http://localhost:8083/s/X886sZ"}
```

### For redirection
Directly enter the short URL link in browser

### For getting top domain metrics using curl
```
curl -X GET http://localhost:8083/topdomainsmetric | jq '.'
output:
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100   121  100   121    0     0   311k      0 --:--:-- --:--:-- --:--:--  118k
[
    {
    "Domain": "www.googgle677975.com",
    "Usage": 4
    },
    {
    "Domain": "www.xyz.com",
    "Usage": 2
    },
    {
    "Domain": "www.longurl.com",
    "Usage": 1
    }
]
```
jq has been just used to format the output

## To run tests
While the application is running, from the root directory of the project, run `go test ./test -v`

## Containerization using docker

### To build docker image with our url shortener service
`sudo docker build -t urlshortner .`

### To run url shortener docker container
`sudo docker run -d -p 8083:8083 urlshortener`

Also, the image is present in my dockerhub:
https://hub.docker.com/r/sanjibgiri/urlshortener
