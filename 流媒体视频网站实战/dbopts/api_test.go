package dbopts

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
流程 init(loginName,truncate table)-->test()-->truncate table
*/
var tempvid string

func TestMain(m *testing.M) {
	clearTable()
	m.Run()
	clearTable()
}

func clearTable() {
	Init()
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("add", testAddUser)
	t.Run("get", testGetUser)
	t.Run("del", testDeleteUser)
	t.Run("reget", testReGetUser)

}

func testAddUser(t *testing.T) {
	err := AddUserCredential("chory", "123")
	if err != nil {
		t.Errorf("add user error: %v/n", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("chory")
	if pwd != "123" || err != nil {
		t.Fatal("GetUser error ", err)
	}
}
func testDeleteUser(t *testing.T) {
	err := DeleteUser("chory", "123")
	if err != nil {
		t.Errorf("delete user error: %v/n", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("chory")
	if err != nil {
		t.Errorf("reGetUser error:%v/n", err)
	}
	if pwd != "" {
		t.Errorf("deleted user error: %v/n", err)
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTable()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTable()
	t.Run("AddUser", testAddUser)
	t.Run("AddCommnets", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}
