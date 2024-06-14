## Golang JWT Authentication System

This project provides a Golang implementation for user authentication using JSON Web Tokens (JWT), GORM for database interaction, PostgreSQL as the database, and Gin for routing and request handling.

### Prerequisites

- Golang (version 1.18 or later recommended): [https://go.dev/doc/install](https://go.dev/doc/install)
- PostgreSQL (version 12 or later recommended): [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
- A code editor or IDE (e.g., Visual Studio Code)

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/naman1402/jwt-auth-go.git
   ```

2. Navigate to the project directory:

   ```bash
   cd jwt-auth-go
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

### Configuration

1. Rename the file  `.env.example` to `.env` in the project root directory.
2. Add the following environment variables to `.env`:


### Usage

**1. User Registration:**

- Send a POST request to `/signin` with a JSON body containing `email` and `password` fields:

   ```json
   {
     "email": "user@example.com",
     "password": "your_password"
   }
   ```

- The server will hash the password securely and store the user in the database. A successful response will include a status code of 201 (Created).

**2. User Login:**

- Send a POST request to `/login` with a JSON body containing `email` and `password` fields:

   ```json
   {
     "email": "user@example.com",
     "password": "your_password"
   }
   ```

- The server will verify the credentials and generate a JWT token if successful. The response will include the token and a status code of 200 (OK).

### Security Considerations

- Use a strong, random string for the `JWT_SECRET` environment variable.
- Consider implementing additional security measures like rate limiting and input validation.
- Stay updated with the latest security best practices for JWT and authentication systems.

### Further Development

- Implement role-based authorization (RBAC) to control user access to specific resources.
- Add refresh tokens for longer-lasting sessions.
- Consider using a framework like OAuth 2.0 for more complex authentication flows.

### License

This project is licensed under the MIT License (see LICENSE file).
