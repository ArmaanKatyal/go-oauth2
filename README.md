# Go-OAuth2

Go-OAuth is a Go-based application designed for integrating OAuth2.0 authentication, specifically with Google's OAuth service. It leverages Go's robustness and efficiency to securely manage user authentication and session management using Redis. This solution offers a REST API interface, enabling users to easily authenticate via Google OAuth and manage sessions in a Redis datastore.

## Prerequisites

Before setting up Go-OAuth, ensure you have the following:

-   Latest stable version of Go.
-   Access to a Google Cloud Platform account for OAuth 2.0 credentials.

## Usage

To start using the OAuth2.0 authentication flow with your application, send a GET request to the authentication endpoint to redirect users to Google's OAuth consent page:

### Authentication Endpoint:

```plaintext
GET http://localhost:8080/google_login
```

### Callback Endpoint:

```plaintext
GET http://localhost:8080/google_callback
```

This endpoint exchanges the authentication code for tokens and establishes a session.

### Profile Endpoint:

```plaintext
GET http://localhost:8080/profile
```

This endpoint returns the logged-in user's profile information.

## Configuration

Configure the application through a .env file or environment variables, including:

-   Google OAuth 2.0 credentials (CLIENT_ID, CLIENT_SECRET).

Example of .env configuration:

```plaintext
CLIENT_ID=your_google_client_id
CLIENT_SECRET=your_google_client_secret
JWT_SECRET=your_jwt_secret
```

## Running the Application

1. Clone the repository and navigate to the project directory.
2. Ensure all prerequisites are met.
3. Configure the application as described in the Configuration section.
4. Run the application:

```bash
docker-compose up -d
```

This starts the server, making the authentication endpoints available for OAuth flows.
