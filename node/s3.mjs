import { S3Client, ListBucketsCommand, PutObjectCommand, GetObjectCommand } from "@aws-sdk/client-s3";

// Node.js版AWS SDKは環境変数AWS_DEFAULT_REGIONを参照しないので明示的に代入する必要がある
const region = process.env.AWS_REGION || process.env.AWS_DEFAULT_REGION || "us-west-1";

const BucketName = "my-bucket";
const ObjectKey = "filename.txt";

try {
    const client = new S3Client({
        region,
        forcePathStyle: true,
    });

    // List buckets
    const { Buckets } = await client.send(new ListBucketsCommand({}));
    for (const bucket of Buckets) {
        console.log(bucket.Name);
    }

    // Put object
    await client.send(
        new PutObjectCommand({
            Bucket: BucketName,
            Key: ObjectKey,
            Body: "Hello,World!\n",
        }),
    );

    // Get object
    const { Body } = await client.send(
        new GetObjectCommand({
            Bucket: BucketName,
            Key: ObjectKey,
        }),
    );
    const body = await Body?.transformToString("utf8");
    console.log(body);
} catch (e) {
    console.error("error:", e);
}
