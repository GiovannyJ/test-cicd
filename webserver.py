import http.server
import socketserver

# Define the port to run the server on
PORT = 4040

# Create a handler to serve the HTML files
handler = http.server.SimpleHTTPRequestHandler

# Change the default index.html to hello.html
handler.extensions_map[".html"] = "text/html"

# Create a socket server
with socketserver.TCPServer(("", PORT), handler) as httpd:
    print(f"Serving on port {PORT}")
    httpd.serve_forever()
