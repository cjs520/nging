// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingFtpUser []*NgingFtpUser

func (s Slice_NgingFtpUser) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFtpUser) RangeRaw(fn func(m *NgingFtpUser) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFtpUser) GroupBy(keyField string) map[string][]*NgingFtpUser {
	r := map[string][]*NgingFtpUser{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFtpUser{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFtpUser) KeyBy(keyField string) map[string]*NgingFtpUser {
	r := map[string]*NgingFtpUser{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFtpUser) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFtpUser) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFtpUser) FromList(data interface{}) Slice_NgingFtpUser {
	values, ok := data.([]*NgingFtpUser)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFtpUser{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingFtpUser(ctx echo.Context) *NgingFtpUser {
	m := &NgingFtpUser{}
	m.SetContext(ctx)
	return m
}

// NgingFtpUser FTP用户
type NgingFtpUser struct {
	base    factory.Base
	objects []*NgingFtpUser

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Username    string `db:"username" bson:"username" comment:"用户名" json:"username" xml:"username"`
	Password    string `db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	Banned      string `db:"banned" bson:"banned" comment:"是否禁止连接" json:"banned" xml:"banned"`
	Directory   string `db:"directory" bson:"directory" comment:"授权目录(一行一个) " json:"directory" xml:"directory"`
	IpWhitelist string `db:"ip_whitelist" bson:"ip_whitelist" comment:"IP白名单(一行一个) " json:"ip_whitelist" xml:"ip_whitelist"`
	IpBlacklist string `db:"ip_blacklist" bson:"ip_blacklist" comment:"IP黑名单(一行一个) " json:"ip_blacklist" xml:"ip_blacklist"`
	Created     uint   `db:"created" bson:"created" comment:"创建时间 " json:"created" xml:"created"`
	Updated     uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	GroupId     uint   `db:"group_id" bson:"group_id" comment:"用户组" json:"group_id" xml:"group_id"`
}

// - base function

func (a *NgingFtpUser) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingFtpUser) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFtpUser) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFtpUser) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFtpUser) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFtpUser) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFtpUser) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFtpUser) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFtpUser) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingFtpUser) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFtpUser) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingFtpUser) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingFtpUser) Objects() []*NgingFtpUser {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFtpUser) XObjects() Slice_NgingFtpUser {
	return Slice_NgingFtpUser(a.Objects())
}

func (a *NgingFtpUser) NewObjects() factory.Ranger {
	return &Slice_NgingFtpUser{}
}

func (a *NgingFtpUser) InitObjects() *[]*NgingFtpUser {
	a.objects = []*NgingFtpUser{}
	return &a.objects
}

func (a *NgingFtpUser) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFtpUser) Short_() string {
	return "nging_ftp_user"
}

func (a *NgingFtpUser) Struct_() string {
	return "NgingFtpUser"
}

func (a *NgingFtpUser) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingFtpUser) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFtpUser) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingFtpUser) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingFtpUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpUser(*v))
		case []*NgingFtpUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpUser(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFtpUser) GroupBy(keyField string, inputRows ...[]*NgingFtpUser) map[string][]*NgingFtpUser {
	var rows Slice_NgingFtpUser
	if len(inputRows) > 0 {
		rows = Slice_NgingFtpUser(inputRows[0])
	} else {
		rows = Slice_NgingFtpUser(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFtpUser) KeyBy(keyField string, inputRows ...[]*NgingFtpUser) map[string]*NgingFtpUser {
	var rows Slice_NgingFtpUser
	if len(inputRows) > 0 {
		rows = Slice_NgingFtpUser(inputRows[0])
	} else {
		rows = Slice_NgingFtpUser(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFtpUser) AsKV(keyField string, valueField string, inputRows ...[]*NgingFtpUser) param.Store {
	var rows Slice_NgingFtpUser
	if len(inputRows) > 0 {
		rows = Slice_NgingFtpUser(inputRows[0])
	} else {
		rows = Slice_NgingFtpUser(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFtpUser) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingFtpUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpUser(*v))
		case []*NgingFtpUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpUser(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFtpUser) Insert() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingFtpUser) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingFtpUser) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Updatex()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(a).Updatex(); err != nil {
		return
	}
	err = DBI.Fire("updated", a, mw, args...)
	return
}

func (a *NgingFtpUser) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdateByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).UpdateByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingFtpUser) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdatexByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).UpdatexByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingFtpUser) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFtpUser) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["banned"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["banned"] = "N"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingFtpUser) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(keysValues).Update()
	}
	m := *a
	m.FromRow(keysValues.Map())
	if err = DBI.FireUpdate("updating", &m, keysValues.Keys(), mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(keysValues).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, keysValues.Keys(), mw, args...)
}

func (a *NgingFtpUser) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Banned) == 0 {
			a.Banned = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Banned) == 0 {
			a.Banned = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingFtpUser) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingFtpUser) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Deletex()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).Deletex(); err != nil {
		return
	}
	err = DBI.Fire("deleted", a, mw, args...)
	return
}

func (a *NgingFtpUser) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFtpUser) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingFtpUser) Reset() *NgingFtpUser {
	a.Id = 0
	a.Username = ``
	a.Password = ``
	a.Banned = ``
	a.Directory = ``
	a.IpWhitelist = ``
	a.IpBlacklist = ``
	a.Created = 0
	a.Updated = 0
	a.GroupId = 0
	return a
}

func (a *NgingFtpUser) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Username"] = a.Username
		r["Password"] = a.Password
		r["Banned"] = a.Banned
		r["Directory"] = a.Directory
		r["IpWhitelist"] = a.IpWhitelist
		r["IpBlacklist"] = a.IpBlacklist
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		r["GroupId"] = a.GroupId
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Username":
			r["Username"] = a.Username
		case "Password":
			r["Password"] = a.Password
		case "Banned":
			r["Banned"] = a.Banned
		case "Directory":
			r["Directory"] = a.Directory
		case "IpWhitelist":
			r["IpWhitelist"] = a.IpWhitelist
		case "IpBlacklist":
			r["IpBlacklist"] = a.IpBlacklist
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		case "GroupId":
			r["GroupId"] = a.GroupId
		}
	}
	return r
}

func (a *NgingFtpUser) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "username":
			a.Username = param.AsString(value)
		case "password":
			a.Password = param.AsString(value)
		case "banned":
			a.Banned = param.AsString(value)
		case "directory":
			a.Directory = param.AsString(value)
		case "ip_whitelist":
			a.IpWhitelist = param.AsString(value)
		case "ip_blacklist":
			a.IpBlacklist = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "group_id":
			a.GroupId = param.AsUint(value)
		}
	}
}

func (a *NgingFtpUser) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "Username":
			a.Username = param.AsString(vv)
		case "Password":
			a.Password = param.AsString(vv)
		case "Banned":
			a.Banned = param.AsString(vv)
		case "Directory":
			a.Directory = param.AsString(vv)
		case "IpWhitelist":
			a.IpWhitelist = param.AsString(vv)
		case "IpBlacklist":
			a.IpBlacklist = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "GroupId":
			a.GroupId = param.AsUint(vv)
		}
	}
}

func (a *NgingFtpUser) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["username"] = a.Username
		r["password"] = a.Password
		r["banned"] = a.Banned
		r["directory"] = a.Directory
		r["ip_whitelist"] = a.IpWhitelist
		r["ip_blacklist"] = a.IpBlacklist
		r["created"] = a.Created
		r["updated"] = a.Updated
		r["group_id"] = a.GroupId
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "username":
			r["username"] = a.Username
		case "password":
			r["password"] = a.Password
		case "banned":
			r["banned"] = a.Banned
		case "directory":
			r["directory"] = a.Directory
		case "ip_whitelist":
			r["ip_whitelist"] = a.IpWhitelist
		case "ip_blacklist":
			r["ip_blacklist"] = a.IpBlacklist
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		case "group_id":
			r["group_id"] = a.GroupId
		}
	}
	return r
}

func (a *NgingFtpUser) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFtpUser) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
