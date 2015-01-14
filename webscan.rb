# Read arguments sent in the request

require "net/http"
require "uri"
require "json"
require "cgi"
#$response_hash = {:headers => {}, :body => ""}
#response_json = JSON.generate(response_hash)
#puts response_json

# Initialize a TCPServer object that will listen
server = TCPServer.new('localhost', 8080)

def dorequest(url)
	uri = URI.parse(url)
	http = Net::HTTP.new(uri.host, uri.port)
	res = http.request(Net::HTTP::Get.new(uri.request_uri))
	return handleResponse(res)
end

def handleResponse(res)
	@response_hash = {:headers => {}, :body => ""}
	res.each do |h|
		@response_hash[:headers][h] = res[h]
	end
	@response_hash[:body] = res.body
	return JSON.generate(@response_hash)
end

loop do
	socket = server.accept
	request = socket.gets

	puts "REQUEST MADE: "+request;

	# This outputs the first parameter
	#puts parsedrequest.first#.scan(".*GET.*")
	parsedrequest.each do |args|
		puts args
	end
	# Set here an URL if yo want to test	
	response = dorequest("http://localhost");
	
	socket.print response
	socket.close
end

