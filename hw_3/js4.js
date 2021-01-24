"use strict"

const products = [
    {id: 3, price: 127, photos: ["1.jpg", "2.jpg",]},
    {id: 5, price: 499, photos: []},
    {id: 10, price: 26, photos: ["3.jpg"]},
    {id: 8, price: 78,},
];

let resultHasPhotos = products.filter(product => product.photos != null && product.photos.length > 0);

let resultPrice = products.sort(compare)

function compare(a, b) {
    if (a.price < b.price) {
        return -1;
    }
    if (a.price > b.price) {
        return 1;
    }

    return 0;
}
console.log("Есть фото")
resultHasPhotos.forEach(item => console.log(item))

console.log("Сортировка по цене")
resultPrice.forEach(item => console.log(item))