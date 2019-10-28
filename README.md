# GoMovie Rest 🎥
For those who wants to learn the Go language simple rest

1. git clone project
2. go run main.go


```
get: http://localhost:8080/api/movies
get: http://localhost:8080/api/movies/1

post: http://localhost:8080/api/movies
```

To add movie to json send post request to above url with following body
```
{
	"name": "Joker",
	"release": " 4 October 2019",
	"director": {
		"firstname": "Todd",
		"lastname": "Phillips"
	}
}
```
