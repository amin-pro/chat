// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"server/ent/message"
	"server/ent/predicate"
	"server/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetType sets the "type" field.
func (mu *MessageUpdate) SetType(s string) *MessageUpdate {
	mu.mutation.SetType(s)
	return mu
}

// SetContent sets the "content" field.
func (mu *MessageUpdate) SetContent(s string) *MessageUpdate {
	mu.mutation.SetContent(s)
	return mu
}

// SetTime sets the "time" field.
func (mu *MessageUpdate) SetTime(t time.Time) *MessageUpdate {
	mu.mutation.SetTime(t)
	return mu
}

// SetLocalID sets the "local_id" field.
func (mu *MessageUpdate) SetLocalID(i int64) *MessageUpdate {
	mu.mutation.ResetLocalID()
	mu.mutation.SetLocalID(i)
	return mu
}

// AddLocalID adds i to the "local_id" field.
func (mu *MessageUpdate) AddLocalID(i int64) *MessageUpdate {
	mu.mutation.AddLocalID(i)
	return mu
}

// SetCanBeDeleted sets the "can_be_deleted" field.
func (mu *MessageUpdate) SetCanBeDeleted(b bool) *MessageUpdate {
	mu.mutation.SetCanBeDeleted(b)
	return mu
}

// AddReceiverIDs adds the "receivers" edge to the User entity by IDs.
func (mu *MessageUpdate) AddReceiverIDs(ids ...int) *MessageUpdate {
	mu.mutation.AddReceiverIDs(ids...)
	return mu
}

// AddReceivers adds the "receivers" edges to the User entity.
func (mu *MessageUpdate) AddReceivers(u ...*User) *MessageUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mu.AddReceiverIDs(ids...)
}

// SetSenderID sets the "sender" edge to the User entity by ID.
func (mu *MessageUpdate) SetSenderID(id int) *MessageUpdate {
	mu.mutation.SetSenderID(id)
	return mu
}

// SetNillableSenderID sets the "sender" edge to the User entity by ID if the given value is not nil.
func (mu *MessageUpdate) SetNillableSenderID(id *int) *MessageUpdate {
	if id != nil {
		mu = mu.SetSenderID(*id)
	}
	return mu
}

// SetSender sets the "sender" edge to the User entity.
func (mu *MessageUpdate) SetSender(u *User) *MessageUpdate {
	return mu.SetSenderID(u.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearReceivers clears all "receivers" edges to the User entity.
func (mu *MessageUpdate) ClearReceivers() *MessageUpdate {
	mu.mutation.ClearReceivers()
	return mu
}

// RemoveReceiverIDs removes the "receivers" edge to User entities by IDs.
func (mu *MessageUpdate) RemoveReceiverIDs(ids ...int) *MessageUpdate {
	mu.mutation.RemoveReceiverIDs(ids...)
	return mu
}

// RemoveReceivers removes "receivers" edges to User entities.
func (mu *MessageUpdate) RemoveReceivers(u ...*User) *MessageUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mu.RemoveReceiverIDs(ids...)
}

// ClearSender clears the "sender" edge to the User entity.
func (mu *MessageUpdate) ClearSender() *MessageUpdate {
	mu.mutation.ClearSender()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldType,
		})
	}
	if value, ok := mu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldContent,
		})
	}
	if value, ok := mu.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldTime,
		})
	}
	if value, ok := mu.mutation.LocalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: message.FieldLocalID,
		})
	}
	if value, ok := mu.mutation.AddedLocalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: message.FieldLocalID,
		})
	}
	if value, ok := mu.mutation.CanBeDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldCanBeDeleted,
		})
	}
	if mu.mutation.ReceiversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.ReceiversTable,
			Columns: message.ReceiversPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedReceiversIDs(); len(nodes) > 0 && !mu.mutation.ReceiversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.ReceiversTable,
			Columns: message.ReceiversPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ReceiversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.ReceiversTable,
			Columns: message.ReceiversPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.SenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.SenderTable,
			Columns: []string{message.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.SenderTable,
			Columns: []string{message.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetType sets the "type" field.
func (muo *MessageUpdateOne) SetType(s string) *MessageUpdateOne {
	muo.mutation.SetType(s)
	return muo
}

// SetContent sets the "content" field.
func (muo *MessageUpdateOne) SetContent(s string) *MessageUpdateOne {
	muo.mutation.SetContent(s)
	return muo
}

// SetTime sets the "time" field.
func (muo *MessageUpdateOne) SetTime(t time.Time) *MessageUpdateOne {
	muo.mutation.SetTime(t)
	return muo
}

// SetLocalID sets the "local_id" field.
func (muo *MessageUpdateOne) SetLocalID(i int64) *MessageUpdateOne {
	muo.mutation.ResetLocalID()
	muo.mutation.SetLocalID(i)
	return muo
}

// AddLocalID adds i to the "local_id" field.
func (muo *MessageUpdateOne) AddLocalID(i int64) *MessageUpdateOne {
	muo.mutation.AddLocalID(i)
	return muo
}

// SetCanBeDeleted sets the "can_be_deleted" field.
func (muo *MessageUpdateOne) SetCanBeDeleted(b bool) *MessageUpdateOne {
	muo.mutation.SetCanBeDeleted(b)
	return muo
}

// AddReceiverIDs adds the "receivers" edge to the User entity by IDs.
func (muo *MessageUpdateOne) AddReceiverIDs(ids ...int) *MessageUpdateOne {
	muo.mutation.AddReceiverIDs(ids...)
	return muo
}

// AddReceivers adds the "receivers" edges to the User entity.
func (muo *MessageUpdateOne) AddReceivers(u ...*User) *MessageUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return muo.AddReceiverIDs(ids...)
}

// SetSenderID sets the "sender" edge to the User entity by ID.
func (muo *MessageUpdateOne) SetSenderID(id int) *MessageUpdateOne {
	muo.mutation.SetSenderID(id)
	return muo
}

// SetNillableSenderID sets the "sender" edge to the User entity by ID if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableSenderID(id *int) *MessageUpdateOne {
	if id != nil {
		muo = muo.SetSenderID(*id)
	}
	return muo
}

// SetSender sets the "sender" edge to the User entity.
func (muo *MessageUpdateOne) SetSender(u *User) *MessageUpdateOne {
	return muo.SetSenderID(u.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearReceivers clears all "receivers" edges to the User entity.
func (muo *MessageUpdateOne) ClearReceivers() *MessageUpdateOne {
	muo.mutation.ClearReceivers()
	return muo
}

// RemoveReceiverIDs removes the "receivers" edge to User entities by IDs.
func (muo *MessageUpdateOne) RemoveReceiverIDs(ids ...int) *MessageUpdateOne {
	muo.mutation.RemoveReceiverIDs(ids...)
	return muo
}

// RemoveReceivers removes "receivers" edges to User entities.
func (muo *MessageUpdateOne) RemoveReceivers(u ...*User) *MessageUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return muo.RemoveReceiverIDs(ids...)
}

// ClearSender clears the "sender" edge to the User entity.
func (muo *MessageUpdateOne) ClearSender() *MessageUpdateOne {
	muo.mutation.ClearSender()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Message.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldType,
		})
	}
	if value, ok := muo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldContent,
		})
	}
	if value, ok := muo.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldTime,
		})
	}
	if value, ok := muo.mutation.LocalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: message.FieldLocalID,
		})
	}
	if value, ok := muo.mutation.AddedLocalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: message.FieldLocalID,
		})
	}
	if value, ok := muo.mutation.CanBeDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldCanBeDeleted,
		})
	}
	if muo.mutation.ReceiversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.ReceiversTable,
			Columns: message.ReceiversPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedReceiversIDs(); len(nodes) > 0 && !muo.mutation.ReceiversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.ReceiversTable,
			Columns: message.ReceiversPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ReceiversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.ReceiversTable,
			Columns: message.ReceiversPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.SenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.SenderTable,
			Columns: []string{message.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.SenderTable,
			Columns: []string{message.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
