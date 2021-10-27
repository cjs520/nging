/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package server

import (
	"strings"

	"github.com/admpub/log"
	"github.com/admpub/nging/v3/application/library/config"
	"github.com/webx-top/echo"
)

func Service(ctx echo.Context) error {
	logCategories := echo.KVList{}
	if strings.Contains(config.DefaultConfig.Log.LogFile(), `{category}`) {
		ctx.Set(`logWithCategory`, true)
		categories := log.Categories()
		for _, k := range categories {
			v := k
			switch k {
			case `app`:
				v = ctx.T(`Nging日志`)
			case `db`:
				v = ctx.T(`SQL日志`)
			case `echo`:
				v = ctx.T(`Web框架日志`)
			default:
				v = ctx.T(`%s日志`, strings.Title(k))
			}
			logCategories.Add(k, v)
		}
	} else {
		ctx.Set(`logWithCategory`, false)
	}
	ctx.Set(`logCategories`, false)
	return ctx.Render(`server/service`, nil)
}
