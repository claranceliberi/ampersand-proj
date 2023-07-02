package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ToInt(value string) uint {
	i, _ := strconv.ParseUint(value, 10, 64)
	return uint(i)
}


func ToString(number float64) string {
	return strconv.FormatUint(uint64(number), 10)
}

func RecordExistsByID(db *gorm.DB, model interface{}, id float64) (bool, error) {
	result := db.First(model, "id = ?", uint(id))
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func HandleDatabaseCheck(c *gin.Context, exists bool, err error, notExistMessage string) {
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": notExistMessage})
		return
	}
}
