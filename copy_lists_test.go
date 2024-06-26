package fieldmask_utils_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	fieldmask_utils "github.com/mennanov/fieldmask-utils"
)

func TestStructToStruct_ListUpdate(t *testing.T) {
	listHandlers := make(map[string]fieldmask_utils.ListHandler)
	idVariableName := "Id"
	listHandlers["Field1"] = fieldmask_utils.ListHandler{
		OverrideFullList:    false,
		ListKeyVariableName: &idVariableName,
	}
	opts := fieldmask_utils.WithListHandlers(listHandlers)
	type AListItem struct {
		Id    string
		Name  string
		Count int
	}
	type A struct {
		Field1 []AListItem
		Field2 int
		Field3 []AListItem
	}
	src := &A{
		Field1: []AListItem{
			{
				Id:    "2",
				Name:  "newtest2",
				Count: 2,
			},
			{
				Id:    "1",
				Name:  "newtest1",
				Count: 2,
			},
		},
		Field2: 1,
	}
	dst := &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
			{
				Id:    "2",
				Name:  "test2",
				Count: 2,
			},
			{
				Id:    "3",
				Name:  "test3",
				Count: 3,
			},
		},
		Field2: 0,
	}
	mask := fieldmask_utils.MaskFromString("Field1{1{Name}}")
	err := fieldmask_utils.StructToStruct(mask, src, dst, opts)
	require.NoError(t, err)
	assert.Equal(t, &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "newtest1",
				Count: 1,
			},
			{
				Id:    "2",
				Name:  "test2",
				Count: 2,
			},
			{
				Id:    "3",
				Name:  "test3",
				Count: 3,
			},
		},
		Field2: 0,
	}, dst)
}

func TestStructToStruct_ListRemove(t *testing.T) {
	listHandlers := make(map[string]fieldmask_utils.ListHandler)
	idVariableName := "Id"
	listHandlers["Field1"] = fieldmask_utils.ListHandler{
		OverrideFullList:    false,
		ListKeyVariableName: &idVariableName,
	}
	opts := fieldmask_utils.WithListHandlers(listHandlers)
	type AListItem struct {
		Id    string
		Name  string
		Count int
	}
	type A struct {
		Field1 []AListItem
		Field2 int
		Field3 []AListItem
	}
	src := &A{
		Field1: []AListItem{},
		Field2: 1,
	}
	dst := &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
			{
				Id:    "2",
				Name:  "test2",
				Count: 2,
			},
			{
				Id:    "3",
				Name:  "test3",
				Count: 3,
			},
		},
		Field2: 0,
	}
	mask := fieldmask_utils.MaskFromString("Field1{1}")
	err := fieldmask_utils.StructToStruct(mask, src, dst, opts)
	require.NoError(t, err)
	assert.Equal(t, &A{
		Field1: []AListItem{
			{
				Id:    "2",
				Name:  "test2",
				Count: 2,
			},
			{
				Id:    "3",
				Name:  "test3",
				Count: 3,
			},
		},
		Field2: 0,
	}, dst)
}

func TestStructToStruct_ListAdd(t *testing.T) {
	listHandlers := make(map[string]fieldmask_utils.ListHandler)
	listHandlers["Field3"] = fieldmask_utils.ListHandler{
		OverrideFullList: false,
	}
	opts := fieldmask_utils.WithListHandlers(listHandlers)
	type AListItem struct {
		Id    string
		Name  string
		Count int
	}
	type A struct {
		Field1 []AListItem
		Field2 int
		Field3 []AListItem
	}
	src := &A{
		Field2: 1,
		Field3: []AListItem{
			{
				Id:    "2",
				Name:  "newtest2",
				Count: 2,
			},
		},
	}
	dst := &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
		Field2: 0,
		Field3: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
	}
	mask := fieldmask_utils.MaskFromString("Field3")
	err := fieldmask_utils.StructToStruct(mask, src, dst, opts)
	require.NoError(t, err)
	assert.Equal(t, &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
		Field2: 0,
		Field3: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
			{
				Id:    "2",
				Name:  "newtest2",
				Count: 2,
			},
		},
	}, dst)
}

func TestStructToStruct_ListAddWithId(t *testing.T) {
	listHandlers := make(map[string]fieldmask_utils.ListHandler)
	idVariableName := "Id"
	listHandlers["Field3"] = fieldmask_utils.ListHandler{
		OverrideFullList:    false,
		ListKeyVariableName: &idVariableName,
	}
	opts := fieldmask_utils.WithListHandlers(listHandlers)
	type AListItem struct {
		Id    string
		Name  string
		Count int
	}
	type A struct {
		Field1 []AListItem
		Field2 int
		Field3 []AListItem
	}
	src := &A{
		Field2: 1,
		Field3: []AListItem{
			{
				Id:    "2",
				Name:  "newtest2",
				Count: 2,
			},
		},
	}
	dst := &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
		Field2: 0,
		Field3: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
	}
	mask := fieldmask_utils.MaskFromString("Field3")
	err := fieldmask_utils.StructToStruct(mask, src, dst, opts)
	require.NoError(t, err)
	assert.Equal(t, &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
		Field2: 0,
		Field3: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
			{
				Id:    "2",
				Name:  "newtest2",
				Count: 2,
			},
		},
	}, dst)
}

func TestStructToStruct_ListUpdateAndAdd(t *testing.T) {
	listHandlers := make(map[string]fieldmask_utils.ListHandler)
	idVariableName := "Id"
	listHandlers["Field1"] = fieldmask_utils.ListHandler{
		OverrideFullList:    false,
		ListKeyVariableName: &idVariableName,
	}
	listHandlers["Field3"] = fieldmask_utils.ListHandler{
		OverrideFullList: false,
	}
	opts := fieldmask_utils.WithListHandlers(listHandlers)
	type AListItem struct {
		Id    string
		Name  string
		Count int
	}
	type A struct {
		Field1 []AListItem
		Field2 int
		Field3 []AListItem
	}
	src := &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "newtest1",
				Count: 2,
			},
		},
		Field2: 1,
		Field3: []AListItem{
			{
				Id:    "2",
				Name:  "newtest2",
				Count: 2,
			},
		},
	}
	dst := &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
			{
				Id:    "2",
				Name:  "test2",
				Count: 2,
			},
			{
				Id:    "3",
				Name:  "test3",
				Count: 3,
			},
		},
		Field2: 0,
		Field3: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
	}
	mask := fieldmask_utils.MaskFromString("Field1{1{Name}}, Field3")
	err := fieldmask_utils.StructToStruct(mask, src, dst, opts)
	require.NoError(t, err)
	assert.Equal(t, &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "newtest1",
				Count: 1,
			},
			{
				Id:    "2",
				Name:  "test2",
				Count: 2,
			},
			{
				Id:    "3",
				Name:  "test3",
				Count: 3,
			},
		},
		Field2: 0,
		Field3: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
			{
				Id:    "2",
				Name:  "newtest2",
				Count: 2,
			},
		},
	}, dst)
}

func TestStructToStruct_ListOverride(t *testing.T) {
	listHandlers := make(map[string]fieldmask_utils.ListHandler)
	listHandlers["Field3"] = fieldmask_utils.ListHandler{
		OverrideFullList: true,
	}
	opts := fieldmask_utils.WithListHandlers(listHandlers)
	type AListItem struct {
		Id    string
		Name  string
		Count int
	}
	type A struct {
		Field1 []AListItem
		Field2 int
		Field3 []AListItem
	}
	src := &A{
		Field2: 1,
		Field3: []AListItem{
			{
				Id:    "2",
				Name:  "newtest2",
				Count: 2,
			},
		},
	}
	dst := &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
		Field2: 0,
		Field3: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
	}
	mask := fieldmask_utils.MaskFromString("Field3")
	err := fieldmask_utils.StructToStruct(mask, src, dst, opts)
	require.NoError(t, err)
	assert.Equal(t, &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
		Field2: 0,
		Field3: []AListItem{
			{
				Id:    "2",
				Name:  "newtest2",
				Count: 2,
			},
		},
	}, dst)
}

func TestStructToStruct_NestedCombinations(t *testing.T) {
	listHandlers := make(map[string]fieldmask_utils.ListHandler)
	idVariableName := "Id"
	listHandlers["Field1"] = fieldmask_utils.ListHandler{
		OverrideFullList:    false,
		ListKeyVariableName: &idVariableName,
	}
	listHandlers["Field1.Attributes"] = fieldmask_utils.ListHandler{
		OverrideFullList:    false,
		ListKeyVariableName: &idVariableName,
	}
	listHandlers["Field1.Attributes.Attributes"] = fieldmask_utils.ListHandler{
		OverrideFullList: true,
	}
	opts := fieldmask_utils.WithListHandlers(listHandlers)
	type AListItem struct {
		Id         string
		Name       string
		Count      int
		Attributes []AListItem
	}
	type A struct {
		Field1 []AListItem
		Field2 int
		Field3 []AListItem
	}
	src := &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "newtestlevel1",
				Count: 1,
				Attributes: []AListItem{
					{
						Id:    "1",
						Name:  "newtestlevel2",
						Count: 2,
						Attributes: []AListItem{
							{
								Id:    "2",
								Name:  "newtestlevel3",
								Count: 2,
							},
						},
					},
				},
			},
		},
	}
	dst := &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "testlevel1",
				Count: 1,
				Attributes: []AListItem{
					{
						Id:    "1",
						Name:  "testlevel2-1",
						Count: 1,
						Attributes: []AListItem{
							{
								Id:    "1",
								Name:  "testlevel3-1",
								Count: 1,
							},
							{
								Id:    "2",
								Name:  "testlevel3-2",
								Count: 2,
							},
						},
					},
					{
						Id:    "2",
						Name:  "testlevel2-2",
						Count: 2,
					},
				},
			},
		},
		Field2: 0,
		Field3: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
	}
	mask := fieldmask_utils.MaskFromString("Field1{1{Attributes{1{Name, Attributes}}}}")
	err := fieldmask_utils.StructToStruct(mask, src, dst, opts)
	require.NoError(t, err)
	assert.Equal(t, &A{
		Field1: []AListItem{
			{
				Id:    "1",
				Name:  "testlevel1",
				Count: 1,
				Attributes: []AListItem{
					{
						Id:    "1",
						Name:  "newtestlevel2",
						Count: 1,
						Attributes: []AListItem{
							{
								Id:    "2",
								Name:  "newtestlevel3",
								Count: 2,
							},
						},
					},
					{
						Id:    "2",
						Name:  "testlevel2-2",
						Count: 2,
					},
				},
			},
		},
		Field2: 0,
		Field3: []AListItem{
			{
				Id:    "1",
				Name:  "test1",
				Count: 1,
			},
		},
	}, dst)
}
