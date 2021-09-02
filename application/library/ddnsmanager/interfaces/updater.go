package interfaces

import (
	"github.com/admpub/nging/v3/application/library/ddnsmanager/domain/dnsdomain"
	"github.com/webx-top/echo"
)

type Updater interface {
	Name() string
	Description() string
	SignUpURL() string
	Init(providerSettings echo.H, domains []*dnsdomain.Domain) error
	Update(recordType string, ip string) error
	ConfigItems() echo.KVList
}
