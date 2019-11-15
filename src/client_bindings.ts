import {readFileSync} from "fs";

async function createWebAssembly() {
    const buf = readFileSync("../compiled_wasm/client.wasm");
    const res = await WebAssembly.instantiate(new Uint8Array(buf).buffer);  
    console.log(res);
}


