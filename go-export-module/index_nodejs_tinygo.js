import fs from 'fs'
import './wasm_exec_tinygo.js'

const mathwasm = fs.readFileSync('./export_js_tinygo.wasm')

console.log(mathwasm)

const go = new Go();
WebAssembly.instantiate(new Uint8Array(mathwasm), go.importObject).then(result => {
    go.run(result.instance)
    var addResult = result.instance.exports.add(42, 24)
    console.log('add result: ', addResult)
    var subResult = result.instance.exports.sub(42, 24)
    console.log('sub result: ', subResult)
    var mulResult = result.instance.exports.mul(42, 24)
    console.log('mul result: ', mulResult)
})
