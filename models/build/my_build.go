package build

import "github.com/golyu/sql-build"

var (
	myRule = sqlBuild.Rule{
		IntValue:     -1,
		Int8Value:    -1,
		Int16Value:   -1,
		Int32Value:   -1,
		Int64Value:   -1,
		UintValue:    0,
		Uint8Value:   0,
		Uint16Value:  0,
		Uint32Value:  0,
		Uint64Value:  0,
		Float32Value: -1,
		Float64Value: -1,
		StringValue:  "",
	}
)
// insert
type InsertImpl struct {
	insertBuild sqlBuild.InsertBuild
}

func (i *InsertImpl) Value(value interface{}) *InsertImpl {
	i.insertBuild.Value(value, myRule)
	return i
}
func (i *InsertImpl) Values(value interface{}) *InsertImpl {
	i.insertBuild.Values(value, myRule)
	return i
}
func (i *InsertImpl) String() (string,error) {
	return i.insertBuild.String()
}

func NewInsert(table string) *InsertImpl {
	insert := new(InsertImpl)
	insert.insertBuild.Insert(table)
	return insert
}

// update
type UpdateImpl struct {
	updateBuild sqlBuild.UpdateBuild
}

func (u *UpdateImpl) Set(value interface{}, key string) *UpdateImpl {
	u.updateBuild.Set(value, key, myRule)
	return u
}
func (u *UpdateImpl) Set_(value interface{}, key string) *UpdateImpl {
	u.updateBuild.Set_(value, key, myRule)
	return u
}
func (u *UpdateImpl) Where(value interface{}, key string) *UpdateImpl {
	u.updateBuild.Where(value, key, myRule)
	return u
}
func (u *UpdateImpl) Where_(value interface{}, key string) *UpdateImpl {
	u.updateBuild.Where_(value, key, myRule)
	return u
}
func (i *UpdateImpl) String() (string,error) {
	return i.updateBuild.String()
}

func NewUpdate(table string) *UpdateImpl {
	update := new(UpdateImpl)
	update.updateBuild.Update(table)
	return update
}

//select
type SelectImpl struct {
	selectBuild sqlBuild.SelectBuild
}

func (u *SelectImpl) Where(value interface{}, key string) *SelectImpl {
	u.selectBuild.Where(value, key, myRule)
	return u
}
func (u *SelectImpl) Where_(value interface{}, key string) *SelectImpl {
	u.selectBuild.Where_(value, key, myRule)
	return u
}
func (i *SelectImpl) String() (string,error) {
	return i.selectBuild.String()
}
func NewSelect(table string) *SelectImpl {
	update := new(SelectImpl)
	update.selectBuild.Select(table)
	return update
}

//delete
type DeleteImpl struct {
	deleteBuild sqlBuild.DeleteBuild
}

func (u *DeleteImpl) Where(value interface{}, key string) *DeleteImpl {
	u.deleteBuild.Where(value, key, myRule)
	return u
}
func (u *DeleteImpl) Where_(value interface{}, key string) *DeleteImpl {
	u.deleteBuild.Where_(value, key, myRule)
	return u
}

func (i *DeleteImpl) String() (string,error) {
	return i.deleteBuild.String()
}
func NewDelete(table string) *DeleteImpl {
	update := new(DeleteImpl)
	update.deleteBuild.Delete(table)
	return update
}
