package response

import (
	"github.com/sewakantor/sw-be/businesses/facility"
)

type ReviewSpecific struct {
	Name   string    `json:"name"`
}

func FromDomainReviewSpecific(domain *facility.Domain) *ReviewSpecific {
	return &ReviewSpecific{
		Name:        domain.Name,
	}
}

func FromDomainReviewsSpecific(data []facility.Domain) []ReviewSpecific {
	var res []ReviewSpecific
	for _, s := range data {
		res = append(res, *FromDomainReviewSpecific(&s))
	}
	return res
}