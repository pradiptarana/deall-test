# github.com/pradiptarana/deall-test
This is a repository for deal interview process. The login process generate a JWT to be use for further access to another endpoint. I use pattern that divide the code into repository, usecase and transport. The repository part is for accesing the storage/database/3rd party. The usecase part is for business logic. The transport part is layer for user to access the service. The pattern useful for decoupling for the low level. Example when we want to change the database, we only need to change the repository part only. Same with the transport, if we want to build RPC, we only need to add implementation at transport only.

# How To Run
To run the app please run `go mod tidy` and `go run main.go`

# Deployed service
You can access the app at https://deall-test-rho.vercel.app. 

# CURL
Signup
`curl --request POST \
  --url https://deall-test-rho.vercel.app/api/signup \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/8.4.5' \
  --data '{
	"id": 1,
	"username": "pradiptarana",
	"password": "pass",
	"fullname": "rana pradipta",
	"email": "ranapradipta4@gmail.com",
	"age": 25,
	"location": "jakarta"
}'`


Login
`curl --request POST \
  --url https://deall-test-rho.vercel.app/api/login \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/8.4.5' \
  --data '{
	"username": "pradiptarana",
	"password": "pass"
}'`
