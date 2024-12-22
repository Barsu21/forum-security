
# Forum Security Project

## Project Description
The project is a secure forum with user registration and authentication functionality, the ability to create posts and comments, and additional security measures such as password encryption and rate limiting.

## Project Structure

```
forum-security/
│
├── cmd/                     # Main code to run the application
│   ├── main.go              # Main executable file of the project
│   └── routes.go            # HTTP routes definition
│
├── database/                # Database logic
│   ├── category.go          # Model and functions for category management
│   ├── comment.go           # Model and functions for comment management
│   ├── database.go          # Core database interaction functions
│   ├── init.go              # Database initialization
│   ├── post.go              # Model and functions for post management
│   ├── reaction.go          # Model and functions for reactions (likes/dislikes)
│   └── user.go              # Model and functions for user management
│
├── handlers/                # HTTP request handlers
│   ├── rate_limit.go        # Rate limiting implementation
│   ├── handlers.go          # Main file with handler registration
│   ├── auth/                # Handlers for authentication and registration
│   ├── comment/             # Handlers for comments
│   └── post/                # Handlers for posts
│
├── internal/                # Utility functions and helpers
│   ├── config.go            # Configuration handling functions
│   ├── email.go             # Email handling (not used for confirmation)
│   ├── generateSessionID.go # Session ID generation
│   ├── hash.go              # Data hashing (e.g., passwords)
│   └── validate.go          # Data validation (e.g., email and password validation)
│
├── models/                  # Data models
│   └── models.go            # Core models (user, post, comment)
│
├── tls/                     # SSL certificates for HTTPS
│   ├── cert.pem             # SSL certificate
│   └── key.pem              # Private key
│
├── uploads/                 # Folder for uploaded images
│   └── index.png            # Example of uploaded image
│
├── web/                     # Static resources and HTML templates
│   ├── css/                 # CSS styles for pages
│   ├── html/                # HTML templates for pages
│   └── images/              # Images used on the pages
│
├── Dockerfile               # Dockerfile for application containerization
├── README.md                # This documentation file
├── config.json              # Project configuration file
├── forum.db                 # Database file
└── go.mod                   # Go module dependencies and versioning
```

## Security Requirements Fulfillment

### HTTPS with SSL Certificates
The project implements HTTPS using SSL certificates located in the `tls` folder. The `cert.pem` and `key.pem` certificates are configured in the `cmd/main.go` file, and the server is started using the `ListenAndServeTLS` method.

### Rate Limiting Mechanism
The rate limiting mechanism is implemented in the `handlers/rate_limit.go` file using the `golang.org/x/time/rate` library. The limit is set to 1 request per second with a burst capacity of up to 5 requests. When the limit is exceeded, a 429 (Too Many Requests) status is returned.

### Password Encryption
User passwords are stored in encrypted form using the `bcrypt` algorithm. The `internal/hash.go` file defines functions for creating a password hash and verifying password hash correspondence.

### Unique Session Identifiers
User sessions are secured with unique identifiers generated using UUID in the `internal/generateSessionID.go` file. Each session receives a unique identifier, protecting against session spoofing.

### Error Handling and HTTP Status Codes
The project includes proper error handling and HTTP status code usage in the `handlers/handlers.go` file. When errors occur, appropriate status codes such as 500 (Internal Server Error) and 429 (Too Many Requests) are returned.

## Launch Instructions

### Dependency Installation
To install all dependencies, run:

```bash
go mod download
```

### Compilation and Launch
Compile and launch the server:

```bash
go run cmd/main.go
```

### Running in Docker
To run the project in Docker, execute the following commands:

```bash
docker build -t forum-security .
docker run -p 8080:8080 forum-security
```

