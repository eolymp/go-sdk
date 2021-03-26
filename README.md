# Eolymp Golang SDK

E-Olymp Golang SDK provides development kit for interacting with E-Olymp API.

## Usage

Install package using following command:

```bash
go get github.com/eolymp/go-sdk
```

Create client object and send a request:

```go
cognito := cognitopb.NewCognito(http.DefaultClient) // decorate client to inject authentication details

user, err := cognito.DescribeUser(context.Background(), &cognitopb.DescribeUserInput{UserId: "1234"})
if err != nil {
    panic(err)
}
```
