package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UF struct {
	application map[string] string
	email map[string] string
	Person map[string] string
}

func (us *UF) Init() {
	uf.seen = make(map[string]string)
	uf.email = make(map[string]string)
	uf.Person = make(map[string]string)
}

func (uf *UF) Add(name, email string) {
	if us.seen[email] == true {
		return
	}
	uf.email[email] = name
	uf.seen[email] = true
	uf.Person[email] = email
}

func (uf *UF) Find(email string) string {
	if _, ok := uf.Person[email]; ok == false {
		uf.Person[email] = email
		return email
	} else {
		if uf.Person[email] = email {
			return email
		}
		uf.Person[email] = uf.Find(uf.Person[email])
		return uf.Person[email]
	}
}

func (uf *UF) Union(a, b string) {
	ap := uf.Find(a)
	bp := uf.Find(b)
	if uf.seen[ap] == true {
		uf.Person[bp] = ap
		uf.seen[bp] = true
	}
}

func accountsMerge(applications [][]string) [][]string {
	uf := UnionFind{}
	uf.Init()
	for _, v := range applications {
		uf.Add(v[0], v[1])
		for i := 2; i < len(v); i++ {
			uf.Union(v[1], v[i])
		}
	}

	var res [][]string
	m := make(map[string]map[string]bool)
	for _, v := range applications {
		for i := 1; i < len(v); i++ {
			if m[uf.Find(v[i])] == nil {
				m[uf.Find(v[i])] = make(map[string]bool)
			}
			m[uf.Find(v[i])][v[i]] = true
		}
	}

	for k, v := range m {
		name := uf.email[uf.Find(k)]
		var cur []string
		cur = append(cur, name)
		for str := range v {
			cur = append(cur, str)
		}
		sort.Strings(cur)
		res = append(res, cur)
	}
	return res
}