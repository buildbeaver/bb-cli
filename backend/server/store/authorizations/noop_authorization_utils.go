package authorizations

import (
	"github.com/buildbeaver/buildbeaver/common/models"
	"github.com/doug-martin/goqu/v9"
)

// WithIsAuthorizedListFilter filters the supplied dataset to resources that the specified identity is
// authorized to perform operation on. Set resourceIDColumnName to the name of the id column of the
// resource table being searched.
//
// NOTE: It's important to add this filter to your query immediately after declaring the select/from. This is because
// this filter derives off of the supplied dataset, which will copy all WHERE and JOIN clauses that have already been
// set. But why derive if it creates this ordering problem, you ask? It's because we *want* to copy the dialect
// that's set on the dataset (this filter does some things (e.g. the union) that have different syntax depending
// on the underlying database).
//
// Note that the no-op version of this filter does not filter out any resources, i.e. any identity is allowed to
// see anything.
func WithIsAuthorizedListFilter(
	dataset *goqu.SelectDataset,
	identityID models.IdentityID,
	operation models.Operation,
	resourceIDColumnName string,
) *goqu.SelectDataset {
	return dataset
}
