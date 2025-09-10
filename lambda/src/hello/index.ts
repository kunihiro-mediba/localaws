import type { APIGatewayEvent, Handler } from "aws-lambda";

export const handler: Handler<APIGatewayEvent> = async (event, context) => {
    console.log({ event, context });
    const { httpMethod, path } = event;
    return {
        statusCode: 200,
        body: JSON.stringify({
            message: "Hello, world!",
            httpMethod,
            path,
        }),
    };
};
