package rwutil

import "fmt"

func GetFileFormat(name string, contentType string) string {
	switch contentType {
	case "image/jpg":
		return fmt.Sprintf("%s.%s", name, "jpg")
	case "image/jpeg":
		return fmt.Sprintf("%s.%s", name, "jpeg")
	case "image/png":
		return fmt.Sprintf("%s.%s", name, "png")
	default:
		return name
	}
}
