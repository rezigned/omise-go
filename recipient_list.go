/*
!!! DO NOT MODIFY !!!

autogenerated
 src: gen_list_job.tmpl
 job: &main.GenListJob{Name:"Recipient"}
  on: Thu Nov 05 09:53:07 +0700 2015
  by: chakrit
*/

package omise

type RecipientList struct {
	List
  Data []*Recipient `json:"data"`
}

func (list *RecipientList) Find(id string) *Recipient {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

