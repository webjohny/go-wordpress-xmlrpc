package wordpress

type ModelGetTerms struct {
	ReqGetTerms
}

type ReqGetTerms struct {
	Taxonomy    string `xmlrpc:"taxonomy"`
}

func (m ModelGetTerms) GetMethod() string {
	return `wp.getTerms`
}

func (m ModelGetTerms) GetArgs(user string, pwd string) interface{} {
	args := make([]interface{}, 4)
	args = append(args, user, pwd, m.ReqGetTerms)
	return args
}

func GetTerms(taxonomy string) (m ModelGetTerms) {
	m.ReqGetTerms = ReqGetTerms{
		Taxonomy: taxonomy,
	}
	return m
}
