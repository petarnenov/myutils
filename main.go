package myutils

func CheckNillError(err error) {
	if err != nil {
		panic(err)
	}
}
