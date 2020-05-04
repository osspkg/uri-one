package httpserv

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"uri-one/providers/db"

	"github.com/deweppro/core/pkg/provider/server/http"
	"github.com/sirupsen/logrus"
)

const runs = "zaqwsxcderfvbgtyhnmjuiklop.0147852369-ZAQWSXCDERFVBGTYHNMJUIKLOP"

var encoder64 = base64.NewEncoding(runs).WithPadding(-1)

func (h *HttpSrv) Get(message *http.Message) {
	url := message.Reader.URL.Path[1:]

	if cache, ok := h.getCache(url); ok {
		message.Writer.Header().Set("Location", cache)
		message.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		message.Writer.WriteHeader(301)
		return
	}

	if rid, err := encoder64.DecodeString(url); err == nil {
		if id, e := strconv.ParseInt(string(rid), 10, 64); e == nil {
			if data, er := h.selectOne(id); er == nil {
				h.setCache(url, data)
				message.Writer.Header().Set("Location", data)
				message.Writer.WriteHeader(301)
				return
			}
		}
	}

	message.Empty(404)
}
func (h *HttpSrv) New(message *http.Message) {
	query := message.Reader.URL.Query()

	message.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	if url, ok := query["go"]; ok {
		message.Encode("text/html; charset=utf-8", func() ([]byte, error) {
			data := strings.Join(url, "")

			id, err := h.insert(data)
			if err != nil {
				return nil, err
			}

			var buf bytes.Buffer
			encoder := base64.NewEncoder(encoder64, &buf)

			if _, err := encoder.Write([]byte(strconv.FormatInt(id, 10))); err != nil {
				return nil, err
			}
			if err := encoder.Close(); err != nil {
				return nil, err
			}

			h.setCache(buf.String(), data)

			return bytes.Join([][]byte{[]byte(h.cfg.Http.Prefix), buf.Bytes()}, []byte("/")), nil
		})
	} else {
		message.Empty(403)
	}

}

func (h *HttpSrv) setCache(key, data string) {
	h.Lock()
	defer h.Unlock()

	h.cache[key] = data
}

func (h *HttpSrv) getCache(key string) (string, bool) {
	h.RLock()
	defer h.RUnlock()

	if data, ok := h.cache[key]; ok {
		return data, true
	}

	return "", false
}

func (h *HttpSrv) insert(name string) (int64, error) {
	result, err := h.db.Exec(db.CSetUrl, name, time.Now().Unix())
	if err != nil {
		return 0, err
	}

	id, er := result.LastInsertId()
	if er != nil {
		return 0, er
	}

	return id, nil
}

func (h *HttpSrv) selectOne(id int64) (string, error) {
	rows, err := h.db.Query(db.CGetUrl, id)
	if err != nil {
		return "", err
	}

	defer func() {
		if e := rows.Close(); e != nil {
			logrus.WithFields(logrus.Fields{
				"err": e.Error(),
			}).Warn("Error rows close")
		}
	}()

	var data string

	for rows.Next() {
		if err := rows.Scan(&id, &data); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err.Error(),
			}).Warn("Error get row")
		}
	}

	if err := rows.Err(); err != nil {
		return "", err
	}

	return data, nil
}
