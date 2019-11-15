import {readFileSync} from "fs";

async function createWebAssembly() {
    const buf = readFileSync("../client.wasm");
    const res = await WebAssembly.instantiate(new Uint8Array(buf).buffer);  
    console.log(res);
}

