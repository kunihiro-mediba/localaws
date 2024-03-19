import { KMSClient, EncryptCommand, DecryptCommand } from "@aws-sdk/client-kms";

// Node.js版AWS SDKは環境変数AWS_DEFAULT_REGIONを参照しないので明示的に代入する必要がある
const region = process.env.AWS_REGION || process.env.AWS_DEFAULT_REGION || "us-west-1";

const KeyId = "alias/kmskey/test";
const PlainText = "Hello,KMS!";

try {
    const client = new KMSClient({ region });

    const encrypt = await client.send(
        new EncryptCommand({
            KeyId,
            Plaintext: Buffer.from(PlainText),
        }),
    );
    const ciphertextBase64 = Buffer.from(encrypt.CiphertextBlob).toString("base64");
    console.log("ciphertext: " + ciphertextBase64);

    const decrypt = await client.send(
        new DecryptCommand({
            CiphertextBlob: Buffer.from(ciphertextBase64, "base64"),
        }),
    );
    console.log("plaintext: " + Buffer.from(decrypt.Plaintext).toString("utf-8"));
} catch (e) {
    console.error(e);
}
