# Go Firebase API Backend Template

A starter template for building RESTful APIs in Go using Gin framework, and Firebase as Database.

### Features

-   Implements the Clean Architecture pattern for a scalable and maintainable codebase
-   Uses the Gin framework for efficient and fast handling of HTTP requests

### Build with
- Go
- Gin
- Firebase

### Getting Started

#### Prerequisites

-   Go version 1.21.5 or higher

#### Clone the Repo
Follow these simple steps:

```
$ git clone https://github.com/MaathavanJkr/go-firebase-api-template.git
$ cd go-firebase-api-template
$ cp .config-sample.yaml .config.yaml # create a copy of the example configuration file
```

#### Firebase

- Create a new firebase project.
- Go to Project Settings.
- Switch to Service accounts Tab.
- Click on Generate a new private key button.
- Save the file (Service account key) in a path you remember. (Default: Save it as "credentials.json" in the app root folder)

#### Configuration

##### App

-   `port`: The port on which the server will listen (defaults to 8080)
-   `environment`: The environment the application is running in (defaults to "development")

##### Firebase

-   `service_account_key_path`: Path to the service account key file (defaults to "credentials.json")


#### Run the App
To use it in development:
```
$ go run ./cmd/app
```

#### Build and Run the App
To build and run:
```
$ go build ./cmd/app
$ ./app
```
### Folder Structure

```
/go-firebase-api-template
|-- /cmd
|   |-- /app
|       |-- main.go
|
|-- /pkg
|   |-- /config
|   |   |-- config.go
|   |
|   |-- /controllers
|   |   |-- user_controller.go
|   |
|   |-- /database
|   |   |-- firebase.go
|   |
|   |-- /models
|   |   |-- user_model.go
|   |
|   |-- /routes
|       |-- routes.go
|
|-- .gitignore
|-- config-sample.yaml
|-- config.yaml
|-- credentials.json
|-- go.mod
|-- go.sum
|-- LICENSE
|-- README.md
```

### Contributing

This project is open for contributions and suggestions. If you have an idea for a new feature or a bug fix, don't hesitate to open a pull request

### License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/snykk/go-rest-boilerplate/blob/master/LICENSE) file for details.