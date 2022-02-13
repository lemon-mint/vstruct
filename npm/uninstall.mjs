import child_process from "child_process";
import fs from "fs";
import path from "path";
const npmbin = child_process.execSync('npm bin').toString().trim();

let ext = "";
if (process.platform === "win32") {
    ext = ".exe";
}

if (fs.existsSync(path.join(npmbin, "vstruct" + ext))) {
    fs.unlinkSync(path.join(npmbin, 'vstruct' + ext));
}
