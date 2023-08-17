const http = require("http");
const { DynamoDBClient, GetItemCommand } = require("@aws-sdk/client-dynamodb"); // CommonJS import

const hostname = "localhost";
const port = 3000;

const server = http.createServer(async (req, res) => {
  try {
    const client = new DynamoDBClient();
    const input = {
      TableName: "spike_low_latency",
      Key: {
        name: {
          S: "test",
        },
      },
      AttributesToGet: ["name", "value"],
    };
    const command = new GetItemCommand(input);
    const response = await client.send(command);

    res.statusCode = 200;
    res.setHeader("Content-Type", "text/plain");
    res.end(JSON.stringify(response));
  } catch (error) {
    console.log(error);
    res.statusCode = 500;
    res.setHeader("Content-Type", "text/plain");
    res.end(JSON.stringify(error));
  }
});

server.listen(port, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
