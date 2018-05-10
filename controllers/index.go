package controllers

import (
	"msw/services"
	"msw/utils"
)

//首页

type IndexController struct {
	BaseController
}

func (this *IndexController)Index(){
	//首页轮播图展示
	page := 1
	slideShow := services.GetDishList(page,utils.PAGESIZE3,"ORDER BY release_date DESC",nil)
	this.Data["dish"] = slideShow.Object

	//首页菜单列表展示
	menuShow := services.GetMenuList(page,utils.PAGESIZE6,"ORDER BY popular_count DESC",nil)
	this.Data["menu"] = menuShow.Object
	//首页菜谱列表展示

	//TODO
	//首页最新消息展示

	/*fmt.Print(resp.Object)*/
	this.TplName = "site/index.html"
}

