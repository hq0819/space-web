package utils

import (
	"github.com/jlaffaye/ftp"
	"io"
	"space-web/setting"
	"time"
)

func LoginFtp(conf *setting.FtpConfig) (*ftp.ServerConn, error) {
	dial, _ := ftp.Dial("114.115.213.117:21", ftp.DialWithTimeout(5*time.Second))
	err := dial.Login("self", "805345702")
	return dial, err
}

func Upload(fileName string, reader io.Reader) error {
	conn, err := LoginFtp(setting.FtpConf)
	if err != nil {
		return err
	}
	return conn.Stor(fileName, reader)
}
