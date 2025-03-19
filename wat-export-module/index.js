import fs from 'fs'

const mathwasm = fs.readFileSync('./math.wasm')

console.log(mathwasm)

const math = await WebAssembly.instantiate(new Uint8Array(mathwasm)).then(res => res.instance.exports)

console.log("square(11)", math.square(11))
console.log("subtract(7, 3)", math.subtract(7, 3))
console.log("add(1, 1)", math.add(1, 1))
