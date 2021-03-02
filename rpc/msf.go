package rpc

import (
	"bytes"
	"fmt"
	"net/http"

	"gopkg.in/vmihailenco/msgpack.v2"
)

type sessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type loginReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

type loginRes struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type logoutReq struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

type logoutRes struct {
	Result string `msgpack:"result"`
}

type SessionListResponse struct {
	ID          uint32 `msgpack:",omitempty"`
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_pack"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"description"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort int    `msgpack:"session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid"`
}

type Metaspoilt struct {
	host     string
	user     string
	password string
	token    string
}

func New(host, user, password string) (*Metaspoilt, error) {
	metaspoilt := &Metaspoilt{
		host:     host,
		password: password,
		user:     user,
	}

	if err := metaspoilt.Login(); err != nil {
		return nil, err
	}
	return metaspoilt, nil
}

func (meta *Metaspoilt) send(req interface{}, res interface{}) error {
	buf := new(bytes.Buffer)
	msgpack.NewEncoder(buf).Encode(req)
	dest := fmt.Sprintf("http://%s/api", meta.host)
	r, err := http.Post(dest, "binary/message-pack", buf)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func (meta *Metaspoilt) Login() error {
	ctx := &loginReq{
		Method:   "auth.login",
		Username: meta.user,
		Password: meta.password,
	}
	var res loginRes
	if err := meta.send(ctx, &res); err != nil {
		return err
	}
	fmt.Printf("%+v\n", res)
	meta.token = res.Token
	return nil
}

func (meta *Metaspoilt) Logout() error {
	ctx := &logoutReq{
		Method:      "auth.logout",
		Token:       meta.token,
		LogoutToken: meta.token,
	}
	var res logoutRes
	if err := meta.send(ctx, &res); err != nil {
		return err
	}
	meta.token = ""
	return nil
}
