/*
!!! DO NOT MODIFY !!!

autogenerated
 src: gen_list_job.tmpl
 job: &main.GenListJob{Name:"Customer"}
  on: Thu Nov 05 09:53:07 +0700 2015
  by: chakrit
*/

package omise

type CustomerList struct {
	List
  Data []*Customer `json:"data"`
}

func (list *CustomerList) Find(id string) *Customer {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

