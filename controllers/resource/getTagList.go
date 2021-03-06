package resourceController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/tag"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-gonic/gin"
)

/*
GetTagList : get tag names list
require:
return: tag names list
*/
func GetTagList(c *gin.Context) {

	// get all tags
	tags, err := tagModel.GetAllTags()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch tag list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(tags) > 0 {
		tagsJson, err := util.StructToJsonStr(tags)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "json convert error",
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"tags": tagsJson,
		})
		log.InfoLog.Println(len(tags), "success")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "tag list is empty",
		})
		log.ErrorLog.Println("tag list is empty")
		c.Error(errors.New("tag list is empty"))
	}
}
