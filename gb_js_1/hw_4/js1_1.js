"use strict"

// ES5
function Product_ES5(name, price) {
    this.name = name;
    this.price = price
}

Product_ES5.prototype.make25PercentDiscount = function() {
    let result = this.price - (this.price / 100 * 25)
    console.log(result)
};

let bObj = new Product_ES5('Шоколад', 500);
bObj.make25PercentDiscount();


//ES6
class Product_ES6 {
    constructor(name, price) {
        this.name = name;
        this.price = price
    }

    make25PercentDiscount = function() {
        let result = this.price - (this.price / 100 * 25)
        console.log(result)
    }
}



let bObj_ES6 = new Product_ES6('Хлеб', 150);
bObj_ES6.make25PercentDiscount();
