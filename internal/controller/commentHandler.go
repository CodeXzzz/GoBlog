package controller

import (
	"SummerProject/internal/model"
	"SummerProject/internal/service"
	"SummerProject/utils"
	"net/http"
	"strconv"
)

// GetComments 获取当前文章的所有评论
func GetComments(w http.ResponseWriter, r *http.Request) {
	//获取ArticleID
	data := r.URL.Query().Get("aid")
	aid, _ := strconv.Atoi(data)
	commentList, err := service.GetComments(aid)
	if err != nil {
		utils.Error(w, err)
	}
	utils.Success(w, commentList)
}

// PostComment 发布评论，返回生成的评论的cid
func PostComment(w http.ResponseWriter, r *http.Request) {
	data := utils.GetRequestJsonParams(r)
	uid, _ := strconv.Atoi(data["uid"].(string))
	aid, _ := strconv.Atoi(data["aid"].(string))

	comment := &model.Comment{
		Cid:     utils.EncodeID(),
		Content: data["content"].(string),
		Uid:     uid,
		Aid:     aid,
	}
	cid := strconv.Itoa(comment.Cid)
	err := service.PostComment(comment)
	if err != nil {
		utils.Error(w, err)
	}
	utils.Success(w, cid)
}

// DeleteComment 根据评论的cid删除
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	data := utils.GetRequestJsonParams(r)
	cid, _ := strconv.Atoi(data["cid"].(string))
	err := service.DeleteComment(cid)
	if err != nil {
		utils.Error(w, err)
	}
	utils.Success(w, nil)
}

// Reply2Comment 回复评论ReplyType字段为0，回复的是回复则为1
func Reply2Comment(w http.ResponseWriter, r *http.Request) {
	//data := utils.GetRequestJsonParams(r)

}