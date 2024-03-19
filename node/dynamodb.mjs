import { DynamoDBClient, ListTablesCommand, PutItemCommand, GetItemCommand } from "@aws-sdk/client-dynamodb";

// Node.js版AWS SDKは環境変数AWS_DEFAULT_REGIONを参照しないので明示的に代入する必要がある
const region = process.env.AWS_REGION || process.env.AWS_DEFAULT_REGION || "us-west-1";

const TableName = "test";

try {
    const client = new DynamoDBClient({ region });

    const listTables = await client.send(new ListTablesCommand());

    console.log(listTables.TableNames);

    // put item
    await client.send(
        new PutItemCommand({
            TableName: TableName,
            Item: {
                id: { S: "1" },
                name: { S: "hoge" },
            },
        }),
    );

    // get item
    const getItem = await client.send(
        new GetItemCommand({
            TableName: TableName,
            Key: {
                id: { S: "1" },
            },
        }),
    );
    console.log(getItem.Item);
} catch (e) {
    console.error(e);
}
