package infrastructure

import (
	"strings"

	"github.com/nikunjratanpara/internal/internal/contract"
)

type Query struct {
	sql    string
	params map[string]any
}

type IMetadataProvider[T contract.IEntity] interface {
	GetSelectColumns() []string
	GetInsertColumns() []string
	GetInsertQuery(entity T, index int) Query
	GetUpdatePart(entity T, index int) Query
	GetUpdateWhereClauses(entity T) Query
	GetDeleteConditions(entity T) Query
	GetParameters(entity T, index int32) []map[string]any
}

type MetadataProvider[T contract.IEntity] struct {
	Table            string
	PrimaryKeyFields []string
	Fields           []string
}

func (meatadataProvider MetadataProvider[T]) GetSelectColumns() []string {
	return meatadataProvider.Fields
}

func (metadataProvider MetadataProvider[T]) GetInsertColumns() []string {
	return metadataProvider.Fields
}

func (metadataProvider MetadataProvider[T]) GetInsertQuery(entity T, index int) Query {
	insertColumnNames := metadataProvider.GetInsertColumns()
	insertColumns := strings.Join(insertColumnNames, ",")

	return Query{sql: "Insert into " + metadataProvider.Table + "(" + insertColumns + ") Values "}
}
