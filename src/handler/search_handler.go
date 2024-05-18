package handler

import (
	"github.com/gin-gonic/gin"
)

func searchHandler(c *gin.Context) {
	// 从 Gin 的 Context 中获取 Elasticsearch 客户端
	//esClient, _ := c.Value("esClient").(*elasticsearch.Client)

	// 获取搜索关键词
	//keyword := c.Query("q")
	//
	//// 构建 Elasticsearch 查询
	//searchService := esClient.Search().
	//	Index("your_index_name"). // 指定索引名称
	//	Query(elasticsearch.NewQueryStringQuery(keyword))
	//
	//// 执行搜索查询
	//searchResult, err := searchService.Do(c)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//
	//// 处理搜索结果
	//var hits []interface{}
	//for _, hit := range searchResult.Hits.Hits {
	//	hits = append(hits, hit.Source)
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"hits": hits,
	//})
}
