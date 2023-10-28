package main

func CheckNillError(err error) {
	if err != nil {
		panic(err)
	}
}
