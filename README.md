# github.com/pradiptarana/deall-test
This is a repository for deal interview process.

# How To Run
To run the app please run `go mod tidy` and `go run main.go`

# Deployed service
You can access the app at https://deall-test-rho.vercel.app. 

# CURL
Signup
`curl --request POST \
  --url https://deall-test-rana-pradiptas-projects.vercel.app/signup \
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
	"username": "user",
	"password": "pass"
}'`
