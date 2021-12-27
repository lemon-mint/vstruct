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

        const go = new Go();
        if (!WebAssembly.instantiateStreaming) {
            WebAssembly.instantiateStreaming = async (resp, importObject) => {
                const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }
        const result = await WebAssembly.instantiateStreaming(fetch("/dist/app.wasm"), go.importObject);
        go.run(result.instance);

        const pkgname = document.getElementById("pkgname");
        const compileBtn = document.getElementById("compile");
        const errordiv = document.getElementById("errordiv");
        const outputdiv = document.getElementById("output");
        const input = document.getElementById("code");
        const error = document.getElementById("error");

        const langgo = document.getElementById("go");
        const langdart = document.getElementById("dart");
        const langrust = document.getElementById("rust");

        compileBtn.addEventListener("click", () => {
            const code = input.value;
            const lang = langgo.checked ? "go" : langdart.checked ? "dart" : langrust.checked ? "rust" : "invalid";
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
