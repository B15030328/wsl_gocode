package gokitService

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	mymux "github.com/gorilla/mux"
)

/*
3、transport层处理HTTP 编码解码
*/

func DecodeUserNameRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if uid, ok := mymux.Vars(r)["uid"]; ok {
		uid, _ := strconv.Atoi(uid)
		return UserRequest{UserId: uid}, nil
	}
	return nil, errors.New("uid 参数错误")
}

func EncodeUserNameResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
