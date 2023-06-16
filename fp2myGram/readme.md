# Final Project - MyGram API

### Scalable Web Services with Go - Digitalent ✕ Hacktiv8

MyGram is a free photo sharing app written in Go. People can share, view, and comment photos by everyone. Anyone can create an account by registering an email address and creating a username.

## Getting Started

To start running this project locally,

```bash
git clone https://github.com/musshal/mygram-api.git
```

Open mygram-api folder and install all required dependencies

```bash
cd mygram-api && go mod tidy
```

Copy the example env file and adjust the env file

```
cp .env.example .env
```

Start the server


```bash
go run main.go
```

Check the MyGram API documentation

```html
http://localhost:8080/swagger/index.html
```