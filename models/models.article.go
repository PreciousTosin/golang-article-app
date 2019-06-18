package models

import "errors"

// import "errors"

type article struct {
	ID      int
	Title   string
	Content string
}

// We're storing the article list in memory
var ArticleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func GetAllArticles() []article {
	return ArticleList
}

func GetArticleByID(id int) (*article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

/********************* testing block start *************************************/
var tmpArticleList []article

// This function is used to store the main lists into the temporary one
// for testing
func saveLists() {
	tmpArticleList = ArticleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	ArticleList = tmpArticleList
}

/********************* testing block ends *************************************/
