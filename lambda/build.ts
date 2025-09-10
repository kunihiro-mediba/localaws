import { createWriteStream } from "node:fs";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";

import archiver from "archiver";
import { analyzeMetafile, build, stop } from "esbuild";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const outdir = resolve(__dirname, "dist");
const lambdaDir = resolve(dirname(__dirname), "localstack/lambda");
const targets = ["hello"];

const entryPoints = targets.map((name) => `src/${name}/index.ts`);

const { metafile } = await build({
    entryPoints: entryPoints,
    outdir: outdir,
    outbase: "src",
    absWorkingDir: __dirname,

    bundle: true,
    packages: "bundle",
    minify: false,
    keepNames: true,
    metafile: true,

    platform: "node",
    target: ["node22"],
    format: "cjs",
});

await stop();

if (metafile) {
    const result = await analyzeMetafile(metafile, { color: true });
    console.log(result);
}

await Promise.all(targets.map(async (name: string) => {
    const zipFilename = resolve(lambdaDir, `${name}.zip`);
    const writeStream = createWriteStream(zipFilename);
    const archive = archiver.create("zip", {});
    archive.pipe(writeStream);
    archive.file(
        resolve(outdir, `${name}/index.js`),
        { name: "index.js" },
    );
    await archive.finalize();
    console.log(`Created ${zipFilename}`);
}));

console.log("build success");
