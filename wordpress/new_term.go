package wordpress

type ModelNewTerm struct {
	TermID int
	ReqNewTerm
}

type ReqNewTerm struct {
	Name        string `xmlrpc:"name"`
	Taxonomy    string `xmlrpc:"taxonomy"`
	Slug    	string `xmlrpc:"slug"`
	Description string `xmlrpc:"description"`
	Parent    	int    `xmlrpc:"parent"`
}

func (m ModelNewTerm) GetMethod() string {
	return `wp.newTerm`
}

func (m ModelNewTerm) GetArgs(user string, pwd string) interface{} {
	args := make([]interface{}, 4)
	args = append(args, m.TermID, user, pwd, m.ReqNewTerm)
	return args
}

func NewTerm(name string, taxonomy string, slug string, description string, parentId int) (m ModelNewTerm) {
	m.ReqNewTerm = ReqNewTerm{
		Name: name,
		Taxonomy: taxonomy,
		Slug: slug,
		Description: description,
		Parent: parentId,
	}
	return m
}
