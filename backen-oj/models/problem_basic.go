package models

/**
问题表
*/
import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"`
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problem_categories"` // 关联问题分类表
	Content           string             `gorm:"column:content;type:text;" json:"content"`
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`
	CategoryId        uint               `gorm:"column:category_id;type:int(11);" json:"category_id"`
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"` // 最大运行时长
	MaxMem            int                `gorm:"column:max_mem;type:int(11);" json:"max_mem"`         // 最大运行内存
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {

	tx := DB.Model(new(ProblemBasic)).Where("title like ? OR content like ? ", "%"+keyword+"%", "%"+keyword+"%").
		Preload("ProblemCategories.CategoryBasic")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_baisc.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )", categoryIdentity)
	}

	return tx

}
