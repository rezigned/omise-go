/*
!!! DO NOT MODIFY !!!

autogenerated
 src: gen_list_job.tmpl
 job: &main.GenListJob{Name:"Token"}
  on: Thu Nov 05 09:53:07 +0700 2015
  by: chakrit
*/

package omise

type TokenList struct {
	List
  Data []*Token `json:"data"`
}

func (list *TokenList) Find(id string) *Token {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

