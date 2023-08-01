package utils

func CheckError(err error, message string) {
	if err != nil {
		print(message)
		panic(err)
	}
}
