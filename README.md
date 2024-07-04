URL Shortner 
Basic url shortner algorithm 

Overview 
Implementing an in-memory URL Shortner in Go using MD5 hasing and base62 encoding. It manages URLs thorugh a `url` struct and an `urlDB` map. The server listens on port 3000 handling root requests. `/shorten` for URL shortening, `/redirect/ID` for redirection back to the original URL. A lightweight yet efficient solution for shortening and managing URLs. 

Installation

To use the URL shortener, you need to have Go installed on your system. You can download and install Go from the official Go website.

Clone the repository to your local machine:

git clone git@github.com:rajjmalla10/URL-Shortener--GO-Lang.git

Running the Server
Navigate to the project directory and run the following command to start the server:

go run main.go

You can change the port in the main.go file if needed.

Shortening a URL

To shorten a URL, send a POST request to the /shorten endpoint with a JSON  containing the original URL:

curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://example.com"}'

Redirecting to the Original URL

To redirect to the original URL, visit the shortened URL in your browser or send a GET request to the /redirect/{id} endpoint, where {id} is the shortened URL ID:

curl http://localhost:8080/redirect/5NxPCOQ5RY7
This will redirect you to the original URL associated with the shortened URL.

Test the Functionality:

You've already provided a test file (main_test.go) with your project. To test the functionality of your URL shortener, simply run at terminal at your root director:

go test


Conclusion
By following these steps, you can clone my  repository, run the server, and use the provided test file to ensure the URL shortener functions correctly.  This approach simplifies testing and usage for anyone interested in my Go-based URL shortener project.