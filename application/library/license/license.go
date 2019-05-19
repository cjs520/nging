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

package license

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/admpub/license_gen/lib"
	"github.com/admpub/log"
	"github.com/admpub/restful/rest"
	"github.com/shirou/gopsutil/cpu"
	"github.com/webx-top/com"
	"github.com/webx-top/echo"
)

var (
	trackerURL         = `//www.webx.top/product/script/nging/tracker.js`
	productURL         = `http://www.webx.top/product/detail/nging`
	licenseURL         = `http://www.webx.top/product/license/nging`
	versionURL         = `http://www.webx.top/product/version/nging`
	licenseFileName    = `license.key`
	licenseFile        = filepath.Join(echo.Wd(), licenseFileName)
	licenseExists      bool
	licenseError       = lib.UnlicensedVersion
	licenseData        *lib.LicenseData
	licenseVersion     string
	machineID          string
	domain             string
	emptyLicense       = lib.LicenseData{}
	ErrLicenseNotFound = errors.New(`License does not exist`)
	SkipLicenseCheck   = true
)

type ServerURL struct {
	Tracker         string //用于统计分析的js地址
	Product         string //该产品的详情介绍页面网址
	License         string //许可证验证和许可证下载API网址
	Version         string //该产品最新版本信息API网址
	LicenseFileName string //许可证文件名称
}

func (s *ServerURL) Apply() {
	trackerURL = s.Tracker
	productURL = s.Product
	licenseURL = s.License
	versionURL = s.Version
	if len(s.LicenseFileName) > 0 {
		licenseFileName = s.LicenseFileName
		licenseFile = filepath.Join(echo.Wd(), licenseFileName)
	}
}

func SetServerURL(s *ServerURL) {
	if s != nil {
		s.Apply()
	}
}

func SetVersion(version string) {
	licenseVersion = version
}

func Version() string {
	return licenseVersion
}

func ProductURL() string {
	return productURL
}

func TrackerURL() string {
	return trackerURL
}

func FilePath() string {
	return licenseFile
}

func FileName() string {
	return licenseFileName
}

func Exists() bool {
	return licenseExists
}

func Error() error {
	return licenseError
}

func License() lib.LicenseData {
	if licenseData == nil {
		return emptyLicense
	}
	return *licenseData
}

// Check 检查权限
func Check(machineID, domain string) error {
	if SkipLicenseCheck {
		return nil
	}
	licenseError = validateFromOfficial(machineID, domain)
	if licenseError != ErrConnectionNotFound {
		return licenseError
	}
	//当官方服务器不可用时才验证本地许可证
	licenseError = Validate()
	return licenseError
}

func Ok(domain string) bool {
	if SkipLicenseCheck {
		return true
	}
	switch licenseError {
	case nil:
		if licenseData == nil {
			licenseError = lib.UnlicensedVersion
			return false
		}
		if !licenseData.Info.Expiration.IsZero() && time.Now().After(licenseData.Info.Expiration) {
			licenseError = lib.ExpiredLicense
			return false
		}
		return true
	default:
		machineID, _ := MachineID()
		err := Check(machineID, domain)
		if err == nil {
			licenseError = nil
			return true
		}
		log.Warn(err)
	}
	return false
}

// Validation 定义验证器
type Validation struct {
	NowVersions []string
}

// Validate 参数验证器
func (v *Validation) Validate(data *lib.LicenseData) error {
	if err := data.CheckExpiration(); err != nil {
		return err
	}
	if err := data.CheckVersion(v.NowVersions...); err != nil {
		return err
	}
	mid, err := MachineID()
	if err != nil {
		return err
	}
	if data.Info.MachineID != mid {
		return lib.InvalidMachineID
	}
	return nil
}

// Validate 验证授权
func Validate() error {
	licenseExists = com.FileExists(FilePath())
	if !licenseExists {
		licenseError = ErrLicenseNotFound
		return licenseError
	}
	b, err := ioutil.ReadFile(FilePath())
	if err != nil {
		return err
	}
	validator := &Validation{
		NowVersions: []string{licenseVersion},
	}
	licenseData, err = lib.CheckLicenseStringAndReturning(string(b), PublicKey(), validator)
	return err
}

// Save 保存授权文件
func Save(b []byte) error {
	return ioutil.WriteFile(licenseFile, b, os.ModePerm)
}

// Generate 生成演示版证书
func Generate(privBytes []byte, pemSaveDirs ...string) error {
	var err error
	if privBytes == nil {
		var pubBytes []byte
		pubBytes, privBytes, err = lib.GenerateCertificateData(2048)
		if err != nil {
			return err
		}
		publicKey = string(pubBytes)
		var pemSaveDir string
		if len(pemSaveDirs) > 0 {
			pemSaveDir = pemSaveDirs[0]
		} else {
			pemSaveDir = filepath.Join(echo.Wd(), `data`)
		}
		if len(pemSaveDir) > 0 {
			err = ioutil.WriteFile(filepath.Join(pemSaveDir, `nging.pem.pub`), pubBytes, os.ModePerm)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(filepath.Join(pemSaveDir, `nging.pem`), privBytes, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	info := &lib.LicenseInfo{
		Name:       `demo`,
		LicenseID:  `0`,
		Version:    licenseVersion,
		Expiration: time.Now().Add(30 * 24 * time.Hour),
	}
	info.MachineID, err = MachineID()
	if err != nil {
		return err
	}
	licBytes, err := lib.GenerateLicense(info, string(privBytes))
	if err != nil {
		return err
	}
	return Save(licBytes)
}

// MachineID 生成当前机器的机器码
func MachineID() (string, error) {
	if len(machineID) > 0 {
		return machineID, nil
	}
	addrs, err := lib.MACAddresses(false)
	if err != nil {
		return ``, err
	}
	if len(addrs) < 1 {
		return ``, lib.ErrorMachineID
	}
	cpuInfo, err := cpu.Info()
	if err != nil {
		return ``, err
	}
	var cpuID string
	if len(cpuInfo) > 0 {
		cpuID = cpuInfo[0].PhysicalID
		if len(cpuID) == 0 {
			cpuID = com.Md5(com.Dump(cpuInfo, false))
		}
	}
	machineID = com.MakePassword(lib.Hash(addrs[0])+`#`+cpuID, `coscms`, 3, 8, 19)
	return machineID, err
}

// FullLicenseURL 包含完整参数的授权网址
func FullLicenseURL(machineID, domain string) string {
	return licenseURL + `?os=` + runtime.GOOS + `&arch=` + runtime.GOARCH + `&version=` + licenseVersion + `&machineID=` + machineID + `&domain=` + domain + `&time=` + time.Now().Format(`20060102-150405`)
}

// Download 从官方服务器重新下载许可证
func Download(machineID, domain string) error {
	operation := `获取授权证书失败：%v`
	response := rest.Get(FullLicenseURL(machineID, domain) + `&pipe=download`)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(operation, response.Status)
	}
	officialResp := &OfficialResp{}
	err := response.FillUp(officialResp)
	if err != nil {
		return fmt.Errorf(operation, err)
	}
	if officialResp.Code != 1 {
		return fmt.Errorf(`%v`, officialResp.Info)
	}
	if officialResp.Data == nil {
		return errors.New(`下载证书失败：官方数据返回异常`)
	}
	if com.FileExists(licenseFile) {
		err = os.Rename(licenseFile, licenseFile+`.`+time.Now().Format(`20060102150405`))
		if err != nil {
			return err
		}
	}
	licenseData = &officialResp.Data.LicenseData
	b, err := com.JSONEncode(licenseData, `  `)
	if err != nil {
		b = []byte(err.Error())
	}
	err = ioutil.WriteFile(licenseFile, b, os.ModePerm)
	if err != nil {
		return fmt.Errorf(`保存授权证书失败：%v`, err)
	}
	return err
}
