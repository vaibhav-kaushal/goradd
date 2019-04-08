package model

// Code generated by goradd. DO NOT EDIT.

import (
	"context"
	"fmt"
	"github.com/goradd/goradd/pkg/orm/db"
	. "github.com/goradd/goradd/pkg/orm/op"
	"github.com/goradd/goradd/pkg/orm/query"
	"github.com/goradd/goradd/web/examples/model/node"

	//"./node"
	"bytes"
	"encoding/gob"
)

// tmpBase is a base structure to be embedded in a "subclass" and provides the ORM access to the database.
// Do not directly access the internal variables, but rather use the accessor functions, since this class maintains internal state
// related to the variables.

type tmpBase struct {
	d        string
	dIsValid bool
	dIsDirty bool

	i        int
	iIsValid bool
	iIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]interface{}

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update
	_restored bool
}

const (
	TmpDDefault = ""
	TmpIDefault = 0
)

const (
	TmpD = `D`
	TmpI = `I`
)

// Initialize or re-initialize a Tmp database object to default values.
func (o *tmpBase) Initialize() {

	o.d = ""
	o.dIsValid = false
	o.dIsDirty = false

	o.i = 0
	o.iIsValid = false
	o.iIsDirty = false

	o._restored = false
}

func (o *tmpBase) PrimaryKey() string {
	return o.d
}

func (o *tmpBase) D() string {
	if o._restored && !o.dIsValid {
		panic("d was not selected in the last query and so is not valid")
	}
	return o.d
}

// DIsValid returns true if the value was loaded from the database or has been set.
func (o *tmpBase) DIsValid() bool {
	return o.dIsValid
}

// SetD sets the value of D in the object, to be saved later using the Save() function.
func (o *tmpBase) SetD(v string) {
	if o.d != v || !o._restored {
		o.d = v
		o.dIsDirty = true
	}

}

func (o *tmpBase) I() int {
	if o._restored && !o.iIsValid {
		panic("i was not selected in the last query and so is not valid")
	}
	return o.i
}

// IIsValid returns true if the value was loaded from the database or has been set.
func (o *tmpBase) IIsValid() bool {
	return o.iIsValid
}

// SetI sets the value of I in the object, to be saved later using the Save() function.
func (o *tmpBase) SetI(v int) {
	if o.i != v || !o._restored {
		o.i = v
		o.iIsDirty = true
	}

}

// GetAlias returns the alias for the given key.
func (o *tmpBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
		return query.NewAliasValue([]byte{})
	}
}

// LoadTmp queries for a single Tmp object by primary key.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
// If you need a more elaborate query, use QueryTmps() to start a query builder.
func LoadTmp(ctx context.Context, pk string, joinOrSelectNodes ...query.NodeI) *Tmp {
	return QueryTmps().Where(Equal(node.Tmp().D(), pk)).joinOrSelect(joinOrSelectNodes...).Get(ctx)
}

// LoadTmpByD queries for a single Tmp object by the given unique index values.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
// If you need a more elaborate query, use QueryTmps() to start a query builder.
func LoadTmpByD(ctx context.Context, d string, joinOrSelectNodes ...query.NodeI) *Tmp {
	return QueryTmps().
		Where(Equal(node.Tmp().D(), d)).
		joinOrSelect(joinOrSelectNodes...).
		Get(ctx)
}

func QueryTmps() *tmpBuilder {
	return newTmpBuilder()
}

// The tmpBuilder is a private object using the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type tmpBuilder struct {
	base                query.QueryBuilderI
	hasConditionalJoins bool
}

func newTmpBuilder() *tmpBuilder {
	b := &tmpBuilder{
		base: db.GetDatabase("goradd").
			NewBuilder(),
	}
	return b.Join(node.Tmp())
}

// Load terminates the query builder, performs the query, and returns a slice of Tmp objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *tmpBuilder) Load(ctx context.Context) (tmpSlice []*Tmp) {
	results := b.base.Load(ctx)
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Tmp)
		o.load(item, !b.hasConditionalJoins, o, nil, "")
		tmpSlice = append(tmpSlice, o)
	}
	return tmpSlice
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *tmpBuilder) LoadI(ctx context.Context) (tmpSlice []interface{}) {
	results := b.base.Load(ctx)
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Tmp)
		o.load(item, !b.hasConditionalJoins, o, nil, "")
		tmpSlice = append(tmpSlice, o)
	}
	return tmpSlice
}

// Get is a convenience method to return only the first item found in a query. It is equivalent to adding
// Limit(1,0) to the query, and then getting the first item from the returned slice.
// Limits with joins do not currently work, so don't try it if you have a join
// TODO: Change this to Load1 to be more descriptive and avoid confusion with other Getters
func (b *tmpBuilder) Get(ctx context.Context) *Tmp {
	results := b.Limit(1, 0).Load(ctx)
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *tmpBuilder) Expand(n query.NodeI) *tmpBuilder {
	b.base.Expand(n)
	return b
}

// Join adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *tmpBuilder) Join(n query.NodeI, conditions ...query.NodeI) *tmpBuilder {
	var condition query.NodeI
	if len(conditions) > 1 {
		condition = And(conditions)
	} else if len(conditions) == 1 {
		condition = conditions[0]
	}
	b.base.Join(n, condition)
	if condition != nil {
		b.hasConditionalJoins = true
	}
	return b
}

// Where adds a condition to filter what gets selected.
func (b *tmpBuilder) Where(c query.NodeI) *tmpBuilder {
	b.base.Condition(c)
	return b
}

// OrderBy  spedifies how the resulting data should be sorted.
func (b *tmpBuilder) OrderBy(nodes ...query.NodeI) *tmpBuilder {
	b.base.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *tmpBuilder) Limit(maxRowCount int, offset int) *tmpBuilder {
	b.base.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *tmpBuilder) Select(nodes ...query.NodeI) *tmpBuilder {
	b.base.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using GetAlias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *tmpBuilder) Alias(name string, n query.NodeI) *tmpBuilder {
	b.base.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *tmpBuilder) Distinct() *tmpBuilder {
	b.base.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *tmpBuilder) GroupBy(nodes ...query.NodeI) *tmpBuilder {
	b.base.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *tmpBuilder) Having(node query.NodeI) *tmpBuilder {
	b.base.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
func (b *tmpBuilder) Count(ctx context.Context, distinct bool, nodes ...query.NodeI) uint {
	return b.base.Count(ctx, distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *tmpBuilder) Delete(ctx context.Context) {
	b.base.Delete(ctx)
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *tmpBuilder) Subquery() *query.SubqueryNode {
	return b.base.Subquery()
}

// joinOrSelect us a private helper function for the Load* functions
func (b *tmpBuilder) joinOrSelect(nodes ...query.NodeI) *tmpBuilder {
	for _, n := range nodes {
		switch n.(type) {
		case query.TableNodeI:
			b.base.Join(n, nil)
		case *query.ColumnNode:
			b.Select(n)
		}
	}
	return b
}

func CountTmpByD(ctx context.Context, d string) uint {
	return QueryTmps().Where(Equal(node.Tmp().D(), d)).Count(ctx, false)
}

func CountTmpByI(ctx context.Context, i int) uint {
	return QueryTmps().Where(Equal(node.Tmp().I(), i)).Count(ctx, false)
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// If linkParent is true we will have child relationships use a pointer back to the parent object. If false, it will create a separate object.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
// Also, if any joins are conditional, that might affect which child objects are included, so in this situation, linkParent should be false
func (o *tmpBase) load(m map[string]interface{}, linkParent bool, objThis *Tmp, objParent interface{}, parentKey string) {
	if v, ok := m["d"]; ok && v != nil {
		if o.d, ok = v.(string); ok {
			o.dIsValid = true
			o.dIsDirty = false
		} else {
			panic("Wrong type found for d.")
		}
	} else {
		o.dIsValid = false
		o.d = ""
	}

	if v, ok := m["i"]; ok && v != nil {
		if o.i, ok = v.(int); ok {
			o.iIsValid = true
			o.iIsDirty = false
		} else {
			panic("Wrong type found for i.")
		}
	} else {
		o.iIsValid = false
		o.i = 0
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}
	o._restored = true
}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *tmpBase) Save(ctx context.Context) {
	if o._restored {
		o.Update(ctx)
	} else {
		o.Insert(ctx)
	}
}

// Update will update the values in the database, saving any changed values.
func (o *tmpBase) Update(ctx context.Context) {
	if !o._restored {
		panic("Cannot update a record that was not originally read from the database.")
	}
	m := o.getModifiedFields()
	if len(m) == 0 {
		return
	}
	d := db.GetDatabase("goradd")
	d.Update(ctx, "tmp", m, "d", fmt.Sprint(o.d))
	o.resetDirtyStatus()
}

// Insert forces the object to be inserted into the database. If the object was loaded from the database originally,
// this will create a duplicate in the database.
func (o *tmpBase) Insert(ctx context.Context) {
	m := o.getModifiedFields()
	if len(m) == 0 {
		return
	}
	d := db.GetDatabase("goradd")
	d.Insert(ctx, "tmp", m)
	o.resetDirtyStatus()
	o._restored = true
}

func (o *tmpBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.dIsDirty {
		fields["d"] = o.d
	}

	if o.iIsDirty {
		fields["i"] = o.i
	}

	return
}

func (o *tmpBase) resetDirtyStatus() {
	o.dIsDirty = false
	o.iIsDirty = false
}

// Delete deletes the associated record from the database.
func (o *tmpBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "tmp", "d", o.d)
}

// DeleteTmp deletes the associated record from the database.
func DeleteTmp(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "tmp", "d", pk)
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *tmpBase) Get(key string) interface{} {

	switch key {
	case "D":
		if !o.dIsValid {
			return nil
		}
		return o.d

	case "I":
		if !o.iIsValid {
			return nil
		}
		return o.i

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database object over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *tmpBase) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err = encoder.Encode(o.d); err != nil {
		return
	}
	if err = encoder.Encode(o.dIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.dIsDirty); err != nil {
		return
	}

	if err = encoder.Encode(o.i); err != nil {
		return
	}
	if err = encoder.Encode(o.iIsValid); err != nil {
		return
	}
	if err = encoder.Encode(o.iIsDirty); err != nil {
		return
	}

	if o._aliases == nil {
		if err = encoder.Encode(false); err != nil {
			return
		}
	} else {
		if err = encoder.Encode(true); err != nil {
			return
		}
		if err = encoder.Encode(o._aliases); err != nil {
			return
		}
	}

	if err = encoder.Encode(o._restored); err != nil {
		return
	}

	return
}

func (o *tmpBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&o.d); err != nil {
		return
	}
	if err = dec.Decode(&o.dIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.dIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.i); err != nil {
		return
	}
	if err = dec.Decode(&o.iIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.iIsDirty); err != nil {
		return
	}

	var hasAliases bool
	if err = dec.Decode(&hasAliases); err != nil {
		return
	}
	if hasAliases {
		if err = dec.Decode(&o._aliases); err != nil {
			return
		}
	}

	if err = dec.Decode(&o._restored); err != nil {
		return
	}

	return err
}