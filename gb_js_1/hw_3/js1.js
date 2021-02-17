"use strict";
let count = 0
for (; count <= 10; count++) {
    let result = count % 2

    if(count === 0) {
        console.log(count + " это ноль")

    } else if(result === 0) {
        console.log(count + " это четное число")

    } else if(result === 1) {
        console.log(count + " это нечетное число")

    }

}