// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingFileThumb []*NgingFileThumb

func (s Slice_NgingFileThumb) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFileThumb) RangeRaw(fn func(m *NgingFileThumb) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFileThumb) GroupBy(keyField string) map[string][]*NgingFileThumb {
	r := map[string][]*NgingFileThumb{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFileThumb{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFileThumb) KeyBy(keyField string) map[string]*NgingFileThumb {
	r := map[string]*NgingFileThumb{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFileThumb) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFileThumb) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFileThumb) FromList(data interface{}) Slice_NgingFileThumb {
	values, ok := data.([]*NgingFileThumb)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFileThumb{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingFileThumb 图片文件缩略图
type NgingFileThumb struct {
	base    factory.Base
	objects []*NgingFileThumb

	Id        uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"主键" json:"id" xml:"id"`
	FileId    uint64 `db:"file_id" bson:"file_id" comment:"文件ID" json:"file_id" xml:"file_id"`
	Size      uint64 `db:"size" bson:"size" comment:"文件大小" json:"size" xml:"size"`
	Width     uint   `db:"width" bson:"width" comment:"宽度(像素)" json:"width" xml:"width"`
	Height    uint   `db:"height" bson:"height" comment:"高度(像素)" json:"height" xml:"height"`
	Dpi       uint   `db:"dpi" bson:"dpi" comment:"分辨率" json:"dpi" xml:"dpi"`
	SaveName  string `db:"save_name" bson:"save_name" comment:"保存名称" json:"save_name" xml:"save_name"`
	SavePath  string `db:"save_path" bson:"save_path" comment:"保存路径" json:"save_path" xml:"save_path"`
	ViewUrl   string `db:"view_url" bson:"view_url" comment:"访问网址" json:"view_url" xml:"view_url"`
	UsedTimes uint   `db:"used_times" bson:"used_times" comment:"被使用的次数" json:"used_times" xml:"used_times"`
	Md5       string `db:"md5" bson:"md5" comment:"缩略图文件MD5值" json:"md5" xml:"md5"`
}

// - base function

func (a *NgingFileThumb) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingFileThumb) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFileThumb) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFileThumb) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFileThumb) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFileThumb) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFileThumb) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFileThumb) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFileThumb) Namer() func(string) string {
	return a.base.Namer()
}

func (a *NgingFileThumb) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFileThumb) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingFileThumb) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingFileThumb) Objects() []*NgingFileThumb {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFileThumb) XObjects() Slice_NgingFileThumb {
	return Slice_NgingFileThumb(a.Objects())
}

func (a *NgingFileThumb) NewObjects() factory.Ranger {
	return &Slice_NgingFileThumb{}
}

func (a *NgingFileThumb) InitObjects() *[]*NgingFileThumb {
	a.objects = []*NgingFileThumb{}
	return &a.objects
}

func (a *NgingFileThumb) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFileThumb) Short_() string {
	return "nging_file_thumb"
}

func (a *NgingFileThumb) Struct_() string {
	return "NgingFileThumb"
}

func (a *NgingFileThumb) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingFileThumb) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFileThumb) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingFileThumb) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFileThumb:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileThumb(*v))
		case []*NgingFileThumb:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileThumb(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFileThumb) GroupBy(keyField string, inputRows ...[]*NgingFileThumb) map[string][]*NgingFileThumb {
	var rows Slice_NgingFileThumb
	if len(inputRows) > 0 {
		rows = Slice_NgingFileThumb(inputRows[0])
	} else {
		rows = Slice_NgingFileThumb(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFileThumb) KeyBy(keyField string, inputRows ...[]*NgingFileThumb) map[string]*NgingFileThumb {
	var rows Slice_NgingFileThumb
	if len(inputRows) > 0 {
		rows = Slice_NgingFileThumb(inputRows[0])
	} else {
		rows = Slice_NgingFileThumb(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFileThumb) AsKV(keyField string, valueField string, inputRows ...[]*NgingFileThumb) param.Store {
	var rows Slice_NgingFileThumb
	if len(inputRows) > 0 {
		rows = Slice_NgingFileThumb(inputRows[0])
	} else {
		rows = Slice_NgingFileThumb(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFileThumb) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFileThumb:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileThumb(*v))
		case []*NgingFileThumb:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFileThumb(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFileThumb) Add() (pk interface{}, err error) {
	a.Id = 0
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingFileThumb) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingFileThumb) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFileThumb) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

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

func (a *NgingFileThumb) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Id = 0
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
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

func (a *NgingFileThumb) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingFileThumb) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFileThumb) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingFileThumb) Reset() *NgingFileThumb {
	a.Id = 0
	a.FileId = 0
	a.Size = 0
	a.Width = 0
	a.Height = 0
	a.Dpi = 0
	a.SaveName = ``
	a.SavePath = ``
	a.ViewUrl = ``
	a.UsedTimes = 0
	a.Md5 = ``
	return a
}

func (a *NgingFileThumb) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["FileId"] = a.FileId
		r["Size"] = a.Size
		r["Width"] = a.Width
		r["Height"] = a.Height
		r["Dpi"] = a.Dpi
		r["SaveName"] = a.SaveName
		r["SavePath"] = a.SavePath
		r["ViewUrl"] = a.ViewUrl
		r["UsedTimes"] = a.UsedTimes
		r["Md5"] = a.Md5
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "FileId":
			r["FileId"] = a.FileId
		case "Size":
			r["Size"] = a.Size
		case "Width":
			r["Width"] = a.Width
		case "Height":
			r["Height"] = a.Height
		case "Dpi":
			r["Dpi"] = a.Dpi
		case "SaveName":
			r["SaveName"] = a.SaveName
		case "SavePath":
			r["SavePath"] = a.SavePath
		case "ViewUrl":
			r["ViewUrl"] = a.ViewUrl
		case "UsedTimes":
			r["UsedTimes"] = a.UsedTimes
		case "Md5":
			r["Md5"] = a.Md5
		}
	}
	return r
}

func (a *NgingFileThumb) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "file_id":
			a.FileId = param.AsUint64(value)
		case "size":
			a.Size = param.AsUint64(value)
		case "width":
			a.Width = param.AsUint(value)
		case "height":
			a.Height = param.AsUint(value)
		case "dpi":
			a.Dpi = param.AsUint(value)
		case "save_name":
			a.SaveName = param.AsString(value)
		case "save_path":
			a.SavePath = param.AsString(value)
		case "view_url":
			a.ViewUrl = param.AsString(value)
		case "used_times":
			a.UsedTimes = param.AsUint(value)
		case "md5":
			a.Md5 = param.AsString(value)
		}
	}
}

func (a *NgingFileThumb) Set(key interface{}, value ...interface{}) {
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
			a.Id = param.AsUint64(vv)
		case "FileId":
			a.FileId = param.AsUint64(vv)
		case "Size":
			a.Size = param.AsUint64(vv)
		case "Width":
			a.Width = param.AsUint(vv)
		case "Height":
			a.Height = param.AsUint(vv)
		case "Dpi":
			a.Dpi = param.AsUint(vv)
		case "SaveName":
			a.SaveName = param.AsString(vv)
		case "SavePath":
			a.SavePath = param.AsString(vv)
		case "ViewUrl":
			a.ViewUrl = param.AsString(vv)
		case "UsedTimes":
			a.UsedTimes = param.AsUint(vv)
		case "Md5":
			a.Md5 = param.AsString(vv)
		}
	}
}

func (a *NgingFileThumb) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["file_id"] = a.FileId
		r["size"] = a.Size
		r["width"] = a.Width
		r["height"] = a.Height
		r["dpi"] = a.Dpi
		r["save_name"] = a.SaveName
		r["save_path"] = a.SavePath
		r["view_url"] = a.ViewUrl
		r["used_times"] = a.UsedTimes
		r["md5"] = a.Md5
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "file_id":
			r["file_id"] = a.FileId
		case "size":
			r["size"] = a.Size
		case "width":
			r["width"] = a.Width
		case "height":
			r["height"] = a.Height
		case "dpi":
			r["dpi"] = a.Dpi
		case "save_name":
			r["save_name"] = a.SaveName
		case "save_path":
			r["save_path"] = a.SavePath
		case "view_url":
			r["view_url"] = a.ViewUrl
		case "used_times":
			r["used_times"] = a.UsedTimes
		case "md5":
			r["md5"] = a.Md5
		}
	}
	return r
}

func (a *NgingFileThumb) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFileThumb) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
