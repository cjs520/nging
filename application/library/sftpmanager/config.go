package sftpmanager

import (
	"io"
	"log"
	"os"

	"github.com/admpub/nging/v5/application/library/config"
	webTerminalConfig "github.com/admpub/web-terminal/config"
	webTerminalSSH "github.com/admpub/web-terminal/library/ssh"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Config 配置
type Config struct {
	Host       string `db:"host" bson:"host" comment:"主机名" json:"host" xml:"host"`
	Port       int    `db:"port" bson:"port" comment:"端口" json:"port" xml:"port"`
	Username   string `db:"username" bson:"username" comment:"用户名" json:"username" xml:"username"`
	Password   string `db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	PrivateKey string `db:"private_key" bson:"private_key" comment:"私钥内容" json:"private_key" xml:"private_key"`
	Passphrase string `db:"passphrase" bson:"passphrase" comment:"私钥口令" json:"passphrase" xml:"passphrase"`
	Charset    string `db:"charset" bson:"charset" comment:"字符集" json:"charset" xml:"charset"`
}

func (c *Config) MakeClientConfig(input io.Reader, output io.Writer) (*ssh.ClientConfig, error) {
	account := &webTerminalConfig.AccountConfig{
		User:     c.Username,
		Password: config.FromFile().Decode(c.Password),
		Charset:  c.Charset,
	}
	if len(c.PrivateKey) > 0 {
		account.PrivateKey = []byte(c.PrivateKey)
	}
	if len(c.Passphrase) > 0 {
		account.Passphrase = []byte(config.FromFile().Decode(c.Passphrase))
	}
	return webTerminalConfig.NewSSHStandard(os.Stdin, log.Writer(), account)
}

func (c *Config) MakeClient() (*webTerminalSSH.SSH, error) {
	clientCfg, err := c.MakeClientConfig(os.Stdin, log.Writer())
	if err != nil {
		return nil, err
	}
	hostCfg := webTerminalConfig.NewHostConfig(clientCfg, c.Host, c.Port)
	sshCfg := webTerminalConfig.NewSSHConfig(hostCfg)
	sshClient := webTerminalSSH.New(sshCfg)
	err = sshClient.Connect()
	return sshClient, err
}

func (c *Config) Connect() (*sftp.Client, error) {
	sshClient, err := c.MakeClient()
	err = sshClient.Connect()
	if err != nil {
		return nil, err
	}
	return sftp.NewClient(sshClient.Client)
}

func DefaultConnector(c *Config) (*sftp.Client, error) {
	return c.Connect()
}
