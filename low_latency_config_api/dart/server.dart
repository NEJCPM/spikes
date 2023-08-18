import 'dart:io';
import 'dart:convert';
import 'package:aws_dynamodb_api/dynamodb-2012-08-10.dart';

void main() {
  HttpServer.bind("0.0.0.0", 3002).then((HttpServer server) async {
    server.listen((HttpRequest request) async {
      final service = DynamoDB(region: 'us-east-1');

      Map<String, AttributeValue> dbMap = Map();
      dbMap["name"] = AttributeValue(s: 'test');

      final result =
          await service.getItem(key: dbMap, tableName: 'spike_low_latency');

      request.response.write(jsonEncode(result.item));
      request.response.close();
    });
  });
}
