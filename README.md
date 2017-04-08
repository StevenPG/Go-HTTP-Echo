# Go-HTTP-Echo
This project was designed to complete the following tasks:
- Allow a service to send JSON requests to a test location
- Display those JSON requests in a logged, or "latest" manner
- Be easily deployable via Go executable, or Docker container

The application came about when testing alerting applications. These alerting applications have capabilities to send JSON/XML to HTTP API endpoints.

However, services to retrieve and use the JSON/XML have not begun development. To begin testing the alerting applications, they must have a valid API endpoint to send data to.

This project seeks to create an easy lightweight location where alerting applications (and any others) can send JSON requests for easy review by testers.

### EndPoints
- host:8000/reset - Reset files on server

- host:8000/info/get - Show all GET requests since last reset

- host:8000/info/delete - Show all DELETE requests since last reset

- host:8001 - Show all POST requests since last reset

- host:8002 - Show all PUT requests since last reset

### Application
By using Golang's asynchronous features and its built in web-server, this application is structured in the following manner:
1. Request comes in
2. Request is written to a file and catalogued on a page of recent requests
3. Request is also presented on the /latest endpoint for quick testing
    - 3a. Intend to push new content to /latest for easy monitoring and testing
4. Retention policy dictates when requests are removed from the application
5. Monitoring page is updated, which contains:
    - 5a. Specs retrieved from web server
    - 5b. Current space on drive
    - 5c. Current CPU/Mem usage (if possible)
    
This is, my first fully featured Go application, feel free to contribute in any manner you please.

This project is licensed under the MIT license, feel free to fork this repository.
