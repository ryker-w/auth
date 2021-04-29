package repo

import (
	"github.com/astaxie/beego/orm"
	"github.com/lishimeng/auth/internal/db/model"
	persistence "github.com/lishimeng/go-orm"
)

func GetUserAsLogin(ctx persistence.OrmContext, param string) (u model.AuthUser, err error) {
	cond:= orm.NewCondition()
	cond1 := cond.Or("UserNo", param).Or("Email", param).Or("Phone", param)
	cond2 := cond.And("Status", 20)
	err = ctx.Context.QueryTable(new(model.AuthUser)).SetCond(cond.AndCond(cond1).AndCond(cond2)).One(&u)
	return
}

func GetAuthUserById(ctx persistence.OrmContext, uid int) (u model.AuthUser, err error) {
	u.Id = uid
	err = ctx.Context.Read(&u)
	return
}

func AddUser(ctx persistence.OrmContext, u *model.AuthUser) (err error) {
	_, err = ctx.Context.Insert(u)
	return
}

func UpdateUserPassword(ctx persistence.OrmContext, u *model.AuthUser) (err error) {
	_, err = ctx.Context.Update(u, "Password")
	return
}

func GetAuthUserOrg(ctx persistence.OrmContext, u model.AuthUser) (auo model.AuthUserOrganization, err error) {
	err = ctx.Context.QueryTable(new(model.AuthUserOrganization)).Filter("UserId", u.Id).One(&auo)
	return
}