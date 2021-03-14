С какими интерфейсами встречались:

func Print(a ...interface{}) (n int, err error) {
    return Fprint(os.Stdout, a...)
}

func Scanln(a ...interface{}) (n int, err error) {
    return Fscanln(os.Stdin, a...)
}

