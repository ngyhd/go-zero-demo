package logic

import (
	"go-zero-demo/post/model"
	"go-zero-demo/post/post"
)

// ConvertPostToData 将 model.Post 转换为 post.PostData
func ConvertPostToData(p *model.Post) *post.PostData {
	return &post.PostData{
		Id:       p.Id,
		UserId:   p.UserId,
		Title:    p.Title,
		Content:  p.Content,
		Views:    p.Views,
		Likes:    p.Likes,
		Comments: p.Comments,
		Shares:   p.Shares,
		Collects: p.Collects,
	}
}

// ConvertPostsToDataList 将 model.Post 列表转换为 post.PostData 列表
func ConvertPostsToDataList(posts []model.Post) []*post.PostData {
	infos := make([]*post.PostData, 0, len(posts))
	for _, p := range posts {
		infos = append(infos, ConvertPostToData(&p))
	}
	return infos
}
