import { SSMClient, GetParameterCommand } from "@aws-sdk/client-ssm";

// Node.js版AWS SDKは環境変数AWS_DEFAULT_REGIONを参照しないので明示的に代入する必要がある
const region = process.env.AWS_REGION || process.env.AWS_DEFAULT_REGION || "us-west-1";

const ParameterName = "parameter-name-test";

try {
    const client = new SSMClient({ region });

    // Get parameter value
    const { Parameter } = await client.send(
        new GetParameterCommand({
            Name: ParameterName,
            WithDecryption: true,
        }),
    );
    console.log("value:", Parameter?.Value);
} catch (e) {
    console.error(e);
}
