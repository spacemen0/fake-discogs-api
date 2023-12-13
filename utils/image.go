package utils

import (
	"fmt"
	"os"
)

func DeleteImageFile(imageUrl string) error {
	err := os.Remove(fmt.Sprintf("images/%s.jpg", imageUrl))
	if err != nil {
		return err
	}
	return nil
}
