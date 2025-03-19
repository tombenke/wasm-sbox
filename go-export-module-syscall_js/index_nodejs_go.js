import fs from 'fs'
import './wasm_exec_go.js'

const mathwasm = fs.readFileSync('./export_js_go.wasm')

console.log(mathwasm)

const go = new Go();
WebAssembly.instantiate(new Uint8Array(mathwasm), go.importObject).then(result => {
    go.run(result.instance)
    var addResult = add(42, 24)
    console.log('add result: ', addResult)
    var addResult = add(42, 24)
    console.log('add result: ', addResult)
    var subResult = sub(42, 24)
    console.log('sub result: ', subResult)
    var mulResult = mul(42, 24)
    console.log('mul result: ', mulResult)
})
