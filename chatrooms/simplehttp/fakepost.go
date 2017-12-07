package main

type FakePost struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (p *FakePost) Retrieve(id int) (err error) {
	p.Id = id
	return
}

func (p *FakePost) Create() (err error) {
	return
}

func (p *FakePost) Update() (err error) {
	return
}

func (p *FakePost) Delete() (err error) {
	return
}
