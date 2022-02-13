const child_process = require('child_process');
const fs = require('fs');
const path = require('path');
const npmbin = child_process.execSync('npm bin').toString().trim();

fs.unlinkSync(path.join(npmbin, 'vstruct'));
