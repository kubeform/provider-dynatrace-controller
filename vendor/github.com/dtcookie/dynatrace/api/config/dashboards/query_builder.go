package dashboards

import "fmt"

type querybuilder struct {
	query     string
	separator string
}

func (qb *querybuilder) Append(name string, value string) {
	qb.query = qb.query + qb.separator + name + "=" + value
	qb.separator = "&"
}

func (qb *querybuilder) Build(prefix string) string {
	if len(prefix) == 0 {
		return qb.query
	}
	if len(qb.query) == 0 {
		return prefix
	}
	return fmt.Sprintf("%s?%s", prefix, qb.query)
}
