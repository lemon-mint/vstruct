import fs from "fs";
import path from "path";

const __dirname = path.resolve();
const npmbin = path.join(__dirname, "bin");

const ext = ".exe";

console.log("removing vstruct from "+path.join(npmbin, "vstruct" + ext));

if (fs.existsSync(path.join(npmbin, "vstruct" + ext))) {
    fs.unlinkSync(path.join(npmbin, 'vstruct' + ext));
}
