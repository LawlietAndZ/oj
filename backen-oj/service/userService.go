package service

import (
	"backen-oj/define"
	"backen-oj/helper"
	"backen-oj/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetUserDetail
// @Tags 公共方法
// @Summary 用户详情
// @Param identity query string false "用户的唯一标识"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user-detail [get]
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户唯一标识不能为空",
		})
		return
	}
	data := new(models.UserBasic)
	err := models.DB.Omit("password").Where("identity = ? ", identity).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取用户失败，identity为:" + identity + " Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// Login
// @Tags 公共方法
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息为空",
		})
		return
	}
	password = helper.GetMd5(password)

	data := new(models.UserBasic)
	err := models.DB.Where("name = ? AND password = ? ", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserBasic Error:" + err.Error(),
		})
		return
	}

	token, err := helper.GenerateToken(data.Identity, data.Name, data.IsAdmin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GenerateToken Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token":    token,
			"is_admin": 1,
		},
	})
}

// SendCode
// @Tags 公共方法
// @Summary 发送验证码
// @Param email formData string true "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	code := helper.GetRand()
	models.RDB.Set(c,email,code,time.Second*300)
	err := helper.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Send Code Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

// Register
// @Tags 公共方法
// @Summary 用户注册
// @Param mail formData string true "mail"
// @Param code formData string true "code"
// @Param name formData string true "name"
// @Param password formData string true "password"
// @Param phone formData string false "phone"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /register [post]]
func Register(c *gin.Context) {
	mail := c.PostForm("mail")
	userCode := c.PostForm("code")
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	if mail == "" || userCode == "" || name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	//验证验证码
	syscode, err := models.RDB.Get(c, mail).Result()
	if err != nil {
		log.Println("获取验证码错误"+err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取验证码失败" ,
		})
		return
	}
	if userCode != syscode {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}


	//判断邮箱是否存在
	var cnt int64

	err=models.DB.Where("mail = ?",mail).Model(new(models.UserBasic)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "查找是否存在邮箱失败" + err.Error(),
		})
		return
	}
	if cnt>0{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该邮箱已被注册" ,
		})
		return
	}

	userIdentity := helper.GetUUID()
	data := &models.UserBasic{
		Identity: userIdentity,
		Name:     name,
		Password: helper.GetMd5(password),
		Phone:    phone,
		Mail:     mail,
	}
	err = models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "创建用户失败" + err.Error(),
		})
		return
	}

	//到这里说明创建用户成功
	//生成token
	token,err:=helper.GenerateToken(userIdentity,name,data.IsAdmin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "生成token失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})

}

// GetRankList
// @Tags 公共方法
// @Summary 用户排行榜
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /rank-list [get]
func GetRankList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size

	var count int64
	list := make([]*models.UserBasic, 0)
	err = models.DB.Model(new(models.UserBasic)).Count(&count).Order("pass_num DESC, submit_num ASC").
		Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Rank List Error:" + err.Error(),
		})
		return
	}
	for _, v := range list {
		mail := strings.Split(v.Mail, "@")
		if len(mail) >= 2 {
			v.Mail = string(mail[0][0]) + "**@" + mail[1]
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}