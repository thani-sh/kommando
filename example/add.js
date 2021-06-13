#!/usr/bin/env node

const [ x, y ] = process.argv
    .slice(2)
    .map(n => parseInt(n))

console.log(x + y)
