package billing_event

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func getBillingEvent() {
	Articles := []Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	}
}

func postBillingEvent() {

}
