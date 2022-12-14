package controller

import (
	"GoBlog/common"
	"GoBlog/internal/middleware/jwt"
	"GoBlog/internal/model"
	"GoBlog/internal/service"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"log"
	"net/http"
	"strconv"
)

// GetComments 获取当前文章的所有评论
//todo paginated list
func GetComments(w http.ResponseWriter, r *http.Request) {
	//获取ArticleID
	data := r.URL.Query().Get("aid")
	aid, _ := strconv.Atoi(data)
	commentList, err := service.GetComments(aid)
	if err != nil {
		common.Error(w, err)
	}
	common.Success(w, commentList)
}

// PostComment 发布评论，返回生成的评论的cid
func PostComment(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")
	//todo 测试中间件传递token部分开始
	_, claim, err := jwt.ParseToken(tokenStr)
	if err != nil {
		common.Error(w, errors.New("请先登录！"))
		return
	}
	log.Println("PostComment: success pass token")
	log.Println("PostComment: 执行操作的用户Uid:", claim.Uid)
	//以上为测试部分
	data := common.GetRequestJsonParams(r)
	uid, _ := strconv.Atoi(data["uid"].(string))
	aid, _ := strconv.Atoi(data["aid"].(string))

	comment := &model.Comment{
		Cid:     int(idgen.NextId()),
		Content: data["content"].(string),
		Uid:     uid,
		Aid:     aid,
	}
	cid := strconv.Itoa(comment.Cid)
	err = service.PostComment(comment)
	if err != nil {
		common.Error(w, err)
	}
	common.Success(w, cid)
}

// DeleteComment 根据评论的cid删除
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	data := common.GetRequestJsonParams(r)
	cid, _ := strconv.Atoi(data["cid"].(string))
	err := service.DeleteComment(cid)
	if err != nil {
		common.Error(w, err)
	}
	common.Success(w, nil)
}

// Reply2Comment 回复评论ReplyType字段为0，回复的是回复则为1
func Reply2Comment(w http.ResponseWriter, r *http.Request) {
	//data := common.GetRequestJsonParams(r)

}
