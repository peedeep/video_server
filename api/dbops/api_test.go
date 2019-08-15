package dbops

import (
	"testing"
	"strconv"
	"time"
	"fmt"
)

var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate videos")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TextMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	clearTables()
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUserCredentail)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Add user error: %v", err)
	}
}

func testGetUserCredentail(t *testing.T) {
	pwd, err := GetUserCredentail("avenssi")
	if pwd != "123" || err != nil {
		t.Errorf("Get user error: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Delete user error: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredentail("avenssi")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Error of RegetUser: %v", err)
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUserCredential)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testGetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	videoInfo, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	if videoInfo == nil {
		t.Errorf("Error of AddVideoInfo: %v, %v", videoInfo, err)
	}
	tempvid = videoInfo.Id
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
	videoInfo, err := GetVideoInfo(tempvid)
	if err != nil || videoInfo != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUserCredential)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "123456"
	aid := 1
	comment := "I like this video"
	err := AddNewComments(vid, aid, comment)
	if err != nil {
		t.Errorf("Add Comments error: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "123456"
	from := 151476800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	comments, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("List Comments error: %v", err)
	}
	for i, ele := range comments {
		fmt.Printf("comment: %d, %v\n", i, ele)
	}
}
