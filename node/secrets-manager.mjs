import { SecretsManagerClient, GetSecretValueCommand } from "@aws-sdk/client-secrets-manager";

// Node.js版AWS SDKは環境変数AWS_DEFAULT_REGIONを参照しないので明示的に代入する必要がある
const region = process.env.AWS_REGION || process.env.AWS_DEFAULT_REGION || "us-west-1";

const SecretID = "secret/key/test";

try {
    const client = new SecretsManagerClient({ region });

    // Get parameter value
    const { SecretString } = await client.send(
        new GetSecretValueCommand({
            SecretId: SecretID,
        }),
    );
    console.log("value:", SecretString);
} catch (e) {
    console.error(e);
}
