function division(a, b) {
    return a / b;
}

function multiplication(a, b) {
    return a * b;
}

function plus(a, b) {
    return a + b;
}

function minus(a, b) {
    return a - b;
}

//задание 5

function operation(a, b, operation) {
    switch (operation) {
        case division:
            division(a, b);
            break;

        case multiplication:
            multiplication(a, b);
            break;

        case plus:
            plus(a, b);
            break;

        case minus:
            minus(a, b)
            break;

    }
}