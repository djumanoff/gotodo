package utils

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

const (
	defaultP   = 1
	defaultIpp = 10
)

type ListParams struct {
	Page         int
	ItemsPerPage int
	OrderField   string
	OrderAsc     bool
}

func (params ListParams) SQLOrderAndPaging() string {
	q := ""
	if params.OrderField != "" {
		q += "ORDER BY " + params.OrderField
		if params.OrderAsc {
			q += " ASC"
		} else {
			q += " DESC"
		}
	}
	if params.ItemsPerPage > 0 {
		q += " LIMIT " + strconv.Itoa(params.ItemsPerPage)
		q += " OFFSET " + strconv.Itoa(params.ItemsPerPage*(params.Page-1))
	}
	return q
}

func ParseFromRequest(r *http.Request, query interface{}) *ListParams {
	var ipp, p int
	var err error
	lp := &ListParams{}
	q := r.URL.Query().Get("q")
	pStr := r.URL.Query().Get("p")
	ippStr := r.URL.Query().Get("ipp")
	_ = bson.UnmarshalJSON([]byte(q), query)
	if pStr == "" {
		p = defaultP
	} else {
		p, err = strconv.Atoi(pStr)
		if err != nil {
			p = defaultP
		}
	}

	if ippStr == "" {
		ipp = defaultIpp
	} else {
		ipp, err = strconv.Atoi(ippStr)
		if err != nil {
			ipp = defaultIpp
		}
	}
	lp.Page = p
	lp.ItemsPerPage = ipp
	lp.OrderAsc = true

	orderField := r.URL.Query().Get("order")
	if orderField != "" {
		if orderField[0] == '-' {
			lp.OrderAsc = false
			lp.OrderField = orderField[1:]
		} else {
			lp.OrderField = orderField
		}
	}

	return lp
}
