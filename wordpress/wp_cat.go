package wordpress

import (
	"time"
)

type Category struct {
	CatID int
	PostCategory
}

type PostCategory struct {
	PostType      string `xmlrpc:"post_type"`
	PostStatus    string `xmlrpc:"post_status"`
	PostTitle     string `xmlrpc:"post_title"`
	PostAuthor    int    `xmlrpc:"post_author"`
	PostExcerpt   string `xmlrpc:"post_excerpt"`
	PostContent   string `xmlrpc:"post_content"`
	PostDate      string `xmlrpc:"post_date"`
	PostFormat    string `xmlrpc:"post_format"`
	PostName      string `xmlrpc:"post_name"`
	PostPassword  string `xmlrpc:"post_password"`
	CommentStatus string `xmlrpc:"comment_status"`
	PingStatus    string `xmlrpc:"ping_status"`
	Sticky        int    `xmlrpc:"sticky"`
	PostThumbnail int    `xmlrpc:"post_thumbnail"`
	PostParent    int    `xmlrpc:"post_parent"`
	// Terms         Terms      `xmlrpc:"terms"`
	TermsNames TermsNames `xmlrpc:"terms_names"`
	Enclosure  Enclosure  `xmlrpc:"enclosure"`
}

func (c Category) GetMethod() string {
	return `wp.newTerm`
}

func (c Category) GetArgs(user string, pwd string) interface{} {
	args := make([]interface{}, 4)
	args = append(args, c.CatID, user, pwd, c.PostCategory)
	return args
}

func GetCatById(content string, title string, tags []string, cate []string) (c Category) {
	c.PostCategory = PostCategory{
		PostType:    `post`,
		PostStatus:  `publish`,
		PostTitle:   title,
		PostContent: content,
		PostDate:    time.Now().Format(`2006-01-02 15:04:05`),
		TermsNames: TermsNames{
			PostCategory: cate,
			TagsInput:    tags,
		},
	}
	return c
}

// NewSpecificPost Customize various values by yourself
func NewSpecificCategory(content PostCategory) (c Category) {
	c.PostCategory = content
	return c
}
