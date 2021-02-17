//пример 1
let a = 1, b = 1, c, d;
c = ++a;
alert(c); // ответ: 2

//a и b присвоено 1. с и d undefined.
// мы взяли "a" со значением 1 и инкрементировали, после чего присвоили его "с".


//пример 2
d = b++;
alert(d); //ответ: 1

//мы взяли b со значением 1, присвоили его "d" и только после этого инкрементировали.

//пример 3
c = 2 + ++a;
alert(c); //ответ: 5

//мы присваеваем "с" 2 плюс инкрементированное "а".
// у "а" значение 2. мы его инкрементируем до 3-х перед тем, как прибавить к 2-м

//пример 4
d = 2 + b++;
alert(d); //ответ: 4

// у "b" значение 2. мы его прибавляем к 2 и присваиваем это "d" и только после этого инкрементируем до 3-х.

alert(a); //3
alert(b); //3