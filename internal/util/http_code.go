package util

func GetHttpCode(code string) int {
	switch {
	case code == "200":
		return 200
	case code == "400":
		return 400
	case code == "401":
		return 400
	default:
		return 500
	}
}
