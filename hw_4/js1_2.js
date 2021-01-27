"use strict"

// ES5
function Post_ES5(author, text, date) {
    this.author = author;
    this.text = text
    this.date = date
}

Post_ES5.prototype.edit = function(text) {
    this.text = text
};

function AttachedPost_ES5(author, text, date) {
    Post_ES5.call(this, author, text, date);
    this.highlighted = false
}

AttachedPost_ES5.prototype.makeTextHighlighted = function () {
    this.highlighted = true
}


//ES6
class Post_ES6 extends  AttachedPost_ES5 {

    constructor(author, text, date) {
        super(author, text, date);
        this.highlighted = false
    }

    edit = function(text) {
        this.text = text
    }
}



let bObj_ES6 = new Post_ES6('Джон', 'Js code', "27.01.2020");
console.log(bObj_ES6.text)
bObj_ES6.edit("Java script");
console.log(bObj_ES6.text)
