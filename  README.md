# Go Web Server with Cat Breeds and Word Count API

This is a Go web server that provides two endpoints: one for checking word count in a string and another for retrieving cat breeds data from an external API.

## Requirements

- Go 1.20 or later
- Gin (Gin-Gonic) framework

## Installation

1. Make sure you have Go installed on your system. You can download it from the official Go website: https://golang.org/

2. Install the Gin framework by running the following command:


3. Clone this repository or download the source code files.

4. Navigate to the project directory.

5. Build and run the application:
    
    go run .



The server will start running on port 8080.

## Endpoints

### 1. Check Word Count

- **URL**: POST /
- **Request Payload**: JSON object with a "str" property containing the string to be checked.
- **Response**: JSON response with an appropriate message and status code based on the word count.





## Example:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"str": "This is a sample string with more than eight words."}' http://localhost:8080/

Response:
{
"message": "200 OK"
}



2.Get Cat Breeds
URL: GET /cat-breeds
Response: JSON response containing cat breeds data fetched from the external API.

Example:

curl http://localhost:8080/cat-breeds

Response:
{
  "data": {
    "Country1": [
      {
        "breed": "Breed1",
        "origin": "Origin1",
        "coat": "Coat1",
        "pattern": "Pattern1"
      },
      ...
    ],
    "Country2": [
      {
        "breed": "Breed2",
        "origin": "Origin2",
        "coat": "Coat2",
        "pattern": "Pattern2"
      },
      ...
    ],
    ...
  },
  "last_page": 5
}
```