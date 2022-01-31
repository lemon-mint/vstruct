var v_loaded = false;

(function () {
    const GoVersion = "1.17";

    const RuntimeURL = `https://cdn.jsdelivr.net/gh/golang/go@release-branch.go${GoVersion}/misc/wasm/wasm_exec.min.js`;

    function loadJS(url) {
        return new Promise((resolve, reject) => {
            var script = document.createElement("script");
            script.src = url;
            script.onload = () => {
                resolve();
            };
            script.onerror = () => {
                reject(new Error("Failed to load " + url));
            };
            document.head.appendChild(script);
        });
    }

    async function main() {
        await loadJS(RuntimeURL);

        const wasmURL = "/dist/app.min.wasm";

        const go = new Go();

        try {
            const result = await WebAssembly.instantiateStreaming(fetch(wasmURL), go.importObject);
            go.run(result.instance);
        } catch (e) {
            console.log(e);
            const wasmdata = await (await fetch(wasmURL)).arrayBuffer();
            const result = await WebAssembly.instantiate(wasmdata, go.importObject);
            go.run(result.instance);
        }
        
        const pkgname = document.getElementById("pkgname");
        const compileBtn = document.getElementById("compile");
        const errordiv = document.getElementById("errordiv");
        const outputdiv = document.getElementById("output");
        const input = document.getElementById("code");
        const error = document.getElementById("error");

        const langgo = document.getElementById("go");
        const langdart = document.getElementById("dart");
        const langrust = document.getElementById("rust");
        const langpython = document.getElementById("python");

        compileBtn.addEventListener("click", () => {
            const code = input.value;
            const lang = langgo.checked ? "go" : langdart.checked ? "dart" : langrust.checked ? "rust" : langpython.checked ? "python" : "invalid";
            const result = compile(lang, pkgname.value, code);
            if (result) {
                errordiv.hidden = true;
                error.innerHTML = "";

                if (result.err == "") {
                    outputdiv.value = result.code;
                } else {
                    errordiv.hidden = false;
                    error.innerText = result.err;
                }
            } else {
                errordiv.hidden = false;
                error.innerHTML = "Compilation failed";
            }
        });
    }

    main().then(() => {
        console.log("main done");
    }).catch(e => {
        console.error(e);
    });
})();

function compile(lang, pkgname, code) {
    if (!v_loaded) {
        return false;
    }
    return JSON.parse(v_compile(lang, pkgname, code));
}
