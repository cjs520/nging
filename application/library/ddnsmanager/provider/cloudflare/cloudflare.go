package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/admpub/nging/v3/application/library/ddnsmanager/domain/dnsdomain"
	"github.com/admpub/nging/v3/application/library/ddnsmanager/provider"
	"github.com/webx-top/echo"
)

const (
	zonesAPI string = "https://api.cloudflare.com/client/v4/zones"
)

// Cloudflare Cloudflare实现
type Cloudflare struct {
	clientSecret string
	Domains      []*dnsdomain.Domain
	TTL          int
}

// CloudflareZonesResp cloudflare zones返回结果
type CloudflareZonesResp struct {
	CloudflareStatus
	Result []struct {
		ID     string
		Name   string
		Status string
		Paused bool
	}
}

// CloudflareRecordsResp records
type CloudflareRecordsResp struct {
	CloudflareStatus
	Result []CloudflareRecord
}

// CloudflareRecord 记录实体
type CloudflareRecord struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Proxied bool   `json:"proxied"`
	TTL     int    `json:"ttl"`
}

// CloudflareStatus 公共状态
type CloudflareStatus struct {
	Success  bool
	Messages []string
}

const Name = `Cloudflare`

func (*Cloudflare) Name() string {
	return Name
}

func (*Cloudflare) Description() string {
	return ``
}

func (*Cloudflare) SignUpURL() string {
	return ``
}

var configItems = echo.KVList{
	echo.NewKV(`ttl`, `TTL`).SetHKV(`inputType`, `number`),
	echo.NewKV(`clientSecret`, `clientSecret`).SetHKV(`inputType`, `text`),
}

func (*Cloudflare) ConfigItems() echo.KVList {
	return configItems
}

// Init 初始化
func (cf *Cloudflare) Init(settings echo.H, domains []*dnsdomain.Domain) error {
	cf.TTL = settings.Int(`ttl`)
	cf.clientSecret = settings.String(`clientSecret`)
	if cf.TTL < 1 { // 默认600s
		cf.TTL = 1
	}
	return nil
}

func (cf *Cloudflare) Update(recordType string, ipAddr string) error {
	for _, domain := range cf.Domains {
		// get zone
		result, err := cf.getZones(domain)
		if err != nil || len(result.Result) != 1 {
			return err
		}
		zoneID := result.Result[0].ID

		var records CloudflareRecordsResp
		// getDomains 最多更新前50条
		err = cf.request(
			"GET",
			fmt.Sprintf(zonesAPI+"/%s/dns_records?type=%s&name=%s&per_page=50", zoneID, recordType, domain),
			nil,
			&records,
		)

		if err != nil || !records.Success {
			return err
		}

		if len(records.Result) > 0 {
			// 更新
			cf.modify(records, zoneID, domain, recordType, ipAddr)
		} else {
			// 新增
			cf.create(zoneID, domain, recordType, ipAddr)
		}
	}
	return nil
}

// 创建
func (cf *Cloudflare) create(zoneID string, domain *dnsdomain.Domain, recordType string, ipAddr string) {
	ipAddr = domain.IP(ipAddr)
	record := &CloudflareRecord{
		Type:    recordType,
		Name:    domain.String(),
		Content: ipAddr,
		Proxied: false,
		TTL:     cf.TTL,
	}
	var status CloudflareStatus
	err := cf.request(
		"POST",
		fmt.Sprintf(zonesAPI+"/%s/dns_records", zoneID),
		record,
		&status,
	)
	if err == nil && status.Success {
		log.Printf("新增域名解析 %s 成功！IP: %s", domain, ipAddr)
		domain.UpdateStatus = dnsdomain.UpdatedSuccess
	} else {
		log.Printf("新增域名解析 %s 失败！Messages: %s", domain, status.Messages)
		domain.UpdateStatus = dnsdomain.UpdatedFailed
	}
}

// 修改
func (cf *Cloudflare) modify(result CloudflareRecordsResp, zoneID string, domain *dnsdomain.Domain, recordType string, ipAddr string) {
	ipAddr = domain.IP(ipAddr)
	for _, record := range result.Result {
		// 相同不修改
		if record.Content == ipAddr {
			log.Printf("你的IP %s 没有变化, 域名 %s", ipAddr, domain)
			continue
		}
		var status CloudflareStatus
		record.Content = ipAddr
		record.TTL = cf.TTL

		err := cf.request(
			"PUT",
			fmt.Sprintf(zonesAPI+"/%s/dns_records/%s", zoneID, record.ID),
			record,
			&status,
		)

		if err == nil && status.Success {
			log.Printf("更新域名解析 %s 成功！IP: %s", domain, ipAddr)
			domain.UpdateStatus = dnsdomain.UpdatedSuccess
		} else {
			log.Printf("更新域名解析 %s 失败！Messages: %s", domain, status.Messages)
			domain.UpdateStatus = dnsdomain.UpdatedFailed
		}
	}
}

// 获得域名记录列表
func (cf *Cloudflare) getZones(domain *dnsdomain.Domain) (result CloudflareZonesResp, err error) {
	err = cf.request(
		"GET",
		fmt.Sprintf(zonesAPI+"?name=%s&status=%s&per_page=%s", domain.DomainName, "active", "50"),
		nil,
		&result,
	)

	return
}

// request 统一请求接口
func (cf *Cloudflare) request(method string, url string, data interface{}, result interface{}) (err error) {
	jsonStr := make([]byte, 0)
	if data != nil {
		jsonStr, _ = json.Marshal(data)
	}
	var req *http.Request
	req, err = http.NewRequest(
		method,
		url,
		bytes.NewBuffer(jsonStr),
	)
	if err != nil {
		log.Println("http.NewRequest失败. Error: ", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+cf.clientSecret)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 30 * time.Second}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	err = provider.UnmarshalHTTPResponse(resp, url, err, result)

	return
}
