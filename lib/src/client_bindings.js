"use strict";
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs = __importStar(require("fs"));
async function createWebAssembly() {
    const buf = fs.readFileSync("../client.wasm");
    const res = await WebAssembly.instantiate(new Uint8Array(buf).buffer);
    console.log(res);
}
