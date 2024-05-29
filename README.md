## Crypto Price Microservice

This project is a practice implementation of a microservice that fetches cryptocurrency prices.
It is designed to demonstrate my understanding and skills in building microservices, utilizing
design patterns such as the Decorator Pattern, and separating services from business logic.
The microservice was built in Go and includes functionalities for logging and metrics, aJSON API
for fetching prices. and a client.

### Learning outcomes
- Building Microservices in Go
- Using the Decorator Pattern to prevent tight coupling and extend functionality
- Utilize Go Contexts to pass data
- Design and Build JSON API

1. Build the image
` docker build -t pricefetcher-image .`

2. Run the image
`docker run -d -p 3000:3000 --name pricefetcher-container pricefetcher-image`

3. Test with Mock Tickers
`http://localhost:3000?ticker=ETH`
`http://localhost:3000?ticker=BTC`
`http://localhost:3000?ticker=RAND`
