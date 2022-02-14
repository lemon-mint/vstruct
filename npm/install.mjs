import fs from "fs";
import path from "path";
import { Octokit } from "octokit";
const octokit = new Octokit();
import got from 'got';
import tarfs from "tar-fs";
import zlib from "zlib";

const __dirname = path.resolve();
const packageJSON = JSON.parse(fs.readFileSync(path.join(__dirname, './package.json'), 'utf8'));
const npmbin = path.join(__dirname, "bin");

//console.log('npm bin:', npmbin);

fs.mkdirSync(npmbin, { recursive: true });

const filepath = path.join(npmbin, 'vstruct');

//const ws = fs.createWriteStream(filepath);

const version = packageJSON.version;

const ARCHS = {
    "x64": "amd64",
    "x32": "386",
    "ia32": "386",
    "arm": "arm_6",
    "arm64": "arm64",
    "ppc64": "ppc64",
}

if (!ARCHS[process.arch]) {
    console.error(`Unsupported architecture: ${process.arch}`);
    process.exit(1);
}
const arch = ARCHS[process.arch];

const PLATFORMS = {
    "darwin": "darwin",
    "win32": "windows",
    "linux": "linux",
    "freebsd": "freebsd",
    "openbsd": "openbsd",
}
if (!PLATFORMS[process.platform]) {
    console.error(`Unsupported platform: ${process.platform}`);
    process.exit(1);
}
const platform = PLATFORMS[process.platform];

// const url = `https://github.com/lemon-mint/vstruct/releases/download/v${version}/vstruct_${version}_${platform}_${arch}.tar.gz`
// //console.log(url);
// https.get(
//     url,
//     function (response) {
//         response.pipe(ws);
//     }
// );

const owner = 'lemon-mint';
const repo = 'vstruct';
const tag = `v${version}`;

(async function () {
    const release = await octokit.rest.repos.getReleaseByTag({
        owner,
        repo,
        tag,
    });
    
    if (release.status !== 200) {
        console.error('Failed to get release');
        process.exit(1);
    }

    const assets = release.data.assets;
    const asset = assets.find(a => a.name === `vstruct_${version}_${platform}_${arch}.tar.gz`);
    if (!asset) {
        console.error('Failed to find asset');
        process.exit(1);
    }

    console.log(`Downloading ${asset.name}`);
    const f = await got.get(asset.browser_download_url, {
        headers: {
            'Accept': 'application/octet-stream',
        },
        responseType: 'buffer',
        followRedirect: true,
    })
    fs.writeFileSync(filepath+".tar.gz", f.body);

    console.log(`Extracting ${asset.name}`);
    fs.createReadStream(filepath+".tar.gz").pipe(zlib.createGunzip()).pipe(tarfs.extract(filepath+".dir")).on('finish', () => {
        if (fs.existsSync(filepath+".tar.gz")) {
            fs.unlinkSync(filepath+".tar.gz");
        }
        let ext = "";
        if (platform === "windows") {
            ext = ".exe";
        }
        console.log("Installing vstruct to "+path.join(npmbin, "vstruct" + ext));
        fs.copyFileSync(path.join(filepath+".dir", "vstruct"+ext), filepath+ext);
        if (fs.existsSync(filepath+".dir")) {
            fs.rmdirSync(filepath+".dir", { recursive: true, force: true });
        }
    });
})();
