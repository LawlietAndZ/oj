package service

import (
	"backen-oj/define"
	"backen-oj/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "请输入当前页，默认第一页"
// @Param size query int false "size"
// @Param keyword query string false "模糊搜索keyword"
// @Param category_identity query string false "分类的唯一标识"
// @Success 200 {string} json "{"code":"200","list":"","count":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {

	var count int64

	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("page转换出错", err)
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("size转换出错", err)
	}

	page = (page - 1) * size

	//获取前端数据
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("categoryIdentity")
	tx := models.GetProblemList(keyword, categoryIdentity)

	list := make([]*models.ProblemBasic, 0)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("获取问题列表失败", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})

}

// GetProblemDetail
// @Tags 公共方法
// @Summary 问题详情
// @Param identity query string false "problem identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {

	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不可以为空",
		})
		return
	}


	data := new(models.ProblemBasic)
	err := models.DB.Where("identity = ?", identity).
		Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "问题不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get ProblemDetail Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})

}
