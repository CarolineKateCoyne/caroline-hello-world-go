# Kinsta - Hello World - Go
An example of how to set your Go application up to enable deployment on Kinsta App Hosting services.

> Kinsta’s Application Hosting is a service to run your web apps and any databases side by side in a hassle-free environment, tailored for developer needs and ease of use.
App Hosting is currently in an invite-only beta phase, sign up for a test account at [kinsta.com/app-hosting](https://kinsta.com/app-hosting).

## Dependency Management

During the deployment process Kinsta will automatically install dependencies defined in your go.mod file.

## Web Server Setup

### Port
Kinsta automatically sets the `PORT` environment variable. You should **not** define it yourself and you should **not** hard-code it into the application. Use
`os.Getenv("PORT")` in your code when referring to the server port.

### Start Command
When deploying an application Kinsta will automatically create a web process.