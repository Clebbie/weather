## Quickstart

You need the set the environment variable "WEATHERSTACK_KEY"

You will also need to install buf (see <https://buf.build/docs/installation>)

- buf gives some quality of life features like proto file linting and proto file generation
- run `buf generate --clean` to generate the proto structs. This can be added to any CI/CD pipeline to ensure the service is always using the latest build.

If you want the binary you can `go build` or if you just want to run the program `go run main.go`

Now with the service running we can use either gRPC or REST.

Here are the endpoints:

### gRPC

`localhost:8080/WeatherService/CheckWeather`

### REST

POST `localhost:8080/weather.v1.WeatherService/CheckWeather`

## Technology Used

In this project, I am using Buf's ConnectRPC. This allows this project to serve both HTTP requests and gRPC requests. I think This serves the project best by allowing me to expose a browser friendly api, as well a speed and effecient gRPC connections for other apps or tools to use.

I chose to implement the challenge as a micro service to keep the program modular and allow it to drop into an existing system or build an
application around it.

I also implemented the cache as a hashmap, but it is just implementing an interface so in the future I could swap it out with a db or redis.

## Whats next?

If I were to continue to develop this project, I would add a docker file for this first. With that, it would really be ready to deploy. I would also like to add more flags in the requests such as using fahrenheit instead of only celsius.
