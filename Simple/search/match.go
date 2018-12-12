package search

import (
	"log"
	"fmt"
)

//Result保存搜索的结果
type Result struct {
	Field   string
	Content string
}

//Matcher定义了要实现的新搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//Match函数，为每个数据源单独启动goroutine来执行这个函数，并发地执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//对特定的匹配器执行搜索
	searchResult, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}
	for _, result := range searchResult {
		results <- result
	}
}

//Display函数从每个单独的goroutine接收到结果后在终端窗口输出
func Display(results chan *Result) {
	//通道会一直被阻塞，直到有结果写入
	//一旦通道被关闭，for循环就会终止
	for result := range results {
		fmt.Printf("%s:\n%s\n", result.Field, result.Content)
	}
}