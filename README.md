# ws.go

This code will perform for you the request you ask it to do.

This was developed in order to be able to perform HTTP requests from JavaScript by requesting it to your local request forwarder (this).

You won't be able to perform an HTTP request with JavaScript to a website with CORS set to its own scope, but by using this, you will be able to request an external application to perform the requests for you and return the content.

I made it just to build another project that I wanted to develop fully in JavaScript when found the CORS limitation on my browser.

The idea is to run the server, and send all your HTTP requests from JavaScript to your local "forwarder" and retrieve the results.

The request has to be sent to the server in the following JSON format:

{"URL":"https://github.com/BBerastegui/","Method":"POST","Header":{"HeaderA":["ValueA"],"HeaderB":"ValueB"},"Body":"id=1&page=2"}

The server will perform the request for you and return the following JSON structure with the response:

{"Status":"200","Header":{"HeaderA":["ValueA"],"HeaderB":"ValueB"},"Body":"<html>HTMLCONTENT</html>"}
