const child_process = require('child_process');
const fs = require('fs');
const path = require('path');
const https = require('https');
const packageJSON = JSON.parse(fs.readFileSync(path.join(__dirname, '../package.json'), 'utf8'));

const npmbin = child_process.execSync('npm bin').toString().trim();

//console.log('npm bin:', npmbin);

fs.mkdirSync(npmbin, { recursive: true });

const filepath = path.join(npmbin, 'vstruct');

const ws = fs.createWriteStream(filepath);

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

const url = `https://github.com/lemon-mint/vstruct/releases/download/v${version}/vstruct_${version}_${platform}_${arch}.tar.gz`
//console.log(url);
https.get(
    url,
    function (response) {
        response.pipe(ws);
    }
);
