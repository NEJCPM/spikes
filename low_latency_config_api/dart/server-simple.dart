import 'dart:io';

void main() {
  HttpServer.bind("0.0.0.0", 3002).then((HttpServer server) async {
    server.listen((HttpRequest request) async {
      request.response.write('Hello!');
      request.response.close();
    });
  });
}
