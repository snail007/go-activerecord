package mysql

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func ar() *ActiveRecord {
	ar := new(ActiveRecord)
	ar.Reset()
	return ar
}
func TestFrom(t *testing.T) {
	want := "SELECT * \nFROM `test`"
	got := strings.TrimSpace(ar().From("test").SQL())
	if want != got {
		t.Errorf("TestFrom , except:%s , got:%s", want, got)
	}
}
func TestFromAs(t *testing.T) {
	want := "SELECT * \nFROM `test` AS `asname`"
	got := strings.TrimSpace(ar().FromAs("test", "asname").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}

func TestSelect(t *testing.T) {
	want := "SELECT `a`,`b` \nFROM `test`"
	got := strings.TrimSpace(ar().From("test").Select("a,b").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestJoin(t *testing.T) {
	want := "SELECT `u`.`a`,`test`.`b` \nFROM `test` LEFT JOIN `user` AS `u` ON `u`.`a`=`test`.`a`"
	got := strings.TrimSpace(ar().From("test").Select("u.a,test.b").Join("user", "u", "u.a=test.a", "LEFT").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestWhere(t *testing.T) {
	_ar := ar()
	want := "SELECT * \nFROM `test` \nWHERE `addr` = ? AND `name` = ?"
	want1 := "SELECT * \nFROM `test` \nWHERE `name` = ? AND `addr` = ?"
	got := strings.TrimSpace(_ar.From("test").Where(map[string]interface{}{
		"name": "kitty",
		"addr": "hk",
	}).SQL())
	if want != got && want1 != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestGroupBy(t *testing.T) {
	want := "SELECT * \nFROM `test`  \nGROUP BY `name`,`uid`"
	got := strings.TrimSpace(ar().From("test").GroupBy("name,uid").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestHaving(t *testing.T) {
	want := "SELECT * \nFROM `test`  \nGROUP BY `name`,`uid` \nHAVING count(uid)>3"
	got := strings.TrimSpace(ar().From("test").GroupBy("name,uid").Having("count(uid)>3").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}

func TestOrderBy(t *testing.T) {
	want := "SELECT * \nFROM `test`    \nORDER BY `id` DESC,`name` ASC"
	got := strings.TrimSpace(ar().From("test").OrderBy("id", "desc").OrderBy("name", "asc").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestLimit(t *testing.T) {
	want := "SELECT * \nFROM `test`     \nLIMIT 0,3"
	got := strings.TrimSpace(ar().From("test").Limit(0, 3).SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestInsert(t *testing.T) {
	_ar := ar()
	want := "INSERT INTO  `test` (`name`,`gid`,`addr`,`is_delete`) \nVALUES (?,?,?,?)"
	got := strings.TrimSpace(_ar.Insert("test", map[string]interface{}{
		"name":      "admin",
		"gid":       33,
		"addr":      nil,
		"is_delete": false,
	}).Limit(0, 3).SQL())
	//fmt.Println(_ar.Values())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestReplace(t *testing.T) {
	_ar := ar()
	want := "REPLACE INTO  `test` (`name`,`gid`,`addr`,`is_delete`) \nVALUES (?,?,?,?)"
	got := strings.TrimSpace(_ar.Replace("test", map[string]interface{}{
		"name":      "admin",
		"gid":       33,
		"addr":      nil,
		"is_delete": false,
	}).Limit(0, 3).SQL())
	//fmt.Println(_ar.Values())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}

func TestInsertBatch(t *testing.T) {
	_ar := ar()
	want := "INSERT INTO  `test` (`name`) \nVALUES (?),(?)"
	got := strings.TrimSpace(_ar.InsertBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"name": "admin11",
		},
		map[string]interface{}{
			"name": "admin",
		},
	}).SQL())

	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestReplaceBatch(t *testing.T) {
	_ar := ar()
	want := "REPLACE INTO  `test` (`name`) \nVALUES (?),(?)"
	got := strings.TrimSpace(_ar.ReplaceBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"name": "admin11",
		},
		map[string]interface{}{
			"name": "admin",
		},
	}).SQL())

	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestDelete(t *testing.T) {
	want := "DELETE FROM  `test`"
	got := strings.TrimSpace(ar().Delete("test", nil).SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestUpdate(t *testing.T) {
	_ar := ar()
	want := "UPDATE  `test` \nSET `addr` = NULL"
	got := strings.TrimSpace(_ar.Update("test", map[string]interface{}{
		"addr": nil,
	}, nil).SQL())
	//fmt.Println(_ar.Values())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}

func TestUpdateBatch(t *testing.T) {
	_ar := ar()
	want := "UPDATE  `test` \nSET `name` = CASE \nWHEN `gid` = ? THEN ? \nWHEN `gid` = ? THEN ? \nELSE `name` END \nWHERE gid IN (?,?)"
	got := strings.TrimSpace(_ar.UpdateBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"name": "admin11",
			"gid":  22,
		},
		map[string]interface{}{
			"name": "admin",
			"gid":  33,
		},
	}, []string{"gid"}).SQL())
	//fmt.Println(_ar.Values())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func Test(t *testing.T) {
	group := NewDBGroup("default")
	group.Regist("default", NewDBConfigWith("127.0.0.1", 3306, "test", "root", "admin"))
	group.Regist("blog", NewDBConfigWith("127.0.0.1", 3306, "test", "root", "admin"))
	group.Regist("www", NewDBConfigWith("127.0.0.1", 3306, "test", "root", "admin"))
	db := group.DB("www")
	if db != nil {
		rs, err := db.Query(db.AR().From("test"))
		if err != nil {
			t.Errorf("ERR:%s", err)
		} else {
			fmt.Println(rs.Rows())
		}
	} else {
		fmt.Printf("db group config of name %s not found", "www")
	}
}

type User struct {
	Name       string    `column:"name"`
	ID         int       `column:"id"`
	Weight     uint      `column:"weight"`
	Height     float32   `column:"height"`
	Sex        bool      `column:"sex"`
	CreateTime time.Time `column:"create_time"`
	Foo        string    `column:"foo"`
}

var rawRows = []map[string]interface{}{
	map[string]interface{}{
		"name":        []byte("jack"),
		"id":          []byte("229"),
		"weight":      []byte("60"),
		"height":      []byte("160.3"),
		"sex":         []byte("1"),
		"create_time": []byte("2017-10-10 09:00:09"),
		"pid":         []byte("1"),
	},
	map[string]interface{}{
		"name":        []byte("jack"),
		"id":          []byte("229"),
		"weight":      []byte("60"),
		"height":      []byte("160.3"),
		"sex":         []byte("1"),
		"create_time": []byte("2017-10-10 09:00:09"),
		"pid":         []byte("2"),
	},
}

func TestStruct(t *testing.T) {
	rs := ResultSet{}
	rs.Init(&rawRows)
	s, err := rs.Struct(User{})
	if err != nil {
		t.Errorf("\n==> Except : \nnil\n==> Got : \n%s", err)
	} else {
		if s.(User).Name != "jack" {
			t.Errorf("\n==> Except : \njack\n==> Got : \n%s", s.(User).Name)
		}
		if s.(User).ID != 229 {
			t.Errorf("\n==> Except : \n229\n==> Got : \n%s", s.(User).ID)
		}
		if s.(User).Weight != 60 {
			t.Errorf("\n==> Except : \njack\n==> Got : \n%s", s.(User).Weight)
		}
		if s.(User).Height != 160.3 {
			t.Errorf("\n==> Except : \njack\n==> Got : \n%s", s.(User).Height)
		}
		if s.(User).Sex != true {
			t.Errorf("\n==> Except : \ntrue\n==> Got : \n%s", s.(User).Sex)
		}
		if s.(User).CreateTime.String() != "2017-10-10 09:00:09 +0800 CST" {
			t.Errorf("\n==> Except : \n\"2017-10-10 09:00:09 +0800 CST\"\n==> Got : \n%s", s.(User).CreateTime)
		}
		if s.(User).Sex != true {
			t.Errorf("\n==> Except : \ntrue\n==> Got : \n%s", s.(User).Sex)
		}
	}

}
func TestStructs(t *testing.T) {
	rs := ResultSet{}
	rs.Init(&rawRows)
	sts, err := rs.Structs(User{})
	if err != nil {
		t.Errorf("\n==> Except : \nnil\n==> Got : \n%s", err)
	} else {
		for _, s := range sts {
			if s.(User).Name != "jack" {
				t.Errorf("\n==> Except : \njack\n==> Got : \n%s", s.(User).Name)
			}
			if s.(User).ID != 229 {
				t.Errorf("\n==> Except : \n229\n==> Got : \n%s", s.(User).ID)
			}
			if s.(User).Weight != 60 {
				t.Errorf("\n==> Except : \njack\n==> Got : \n%s", s.(User).Weight)
			}
			if s.(User).Height != 160.3 {
				t.Errorf("\n==> Except : \njack\n==> Got : \n%s", s.(User).Height)
			}
			if s.(User).Sex != true {
				t.Errorf("\n==> Except : \ntrue\n==> Got : \n%s", s.(User).Sex)
			}
			if s.(User).CreateTime.String() != "2017-10-10 09:00:09 +0800 CST" {
				t.Errorf("\n==> Except : \n\"2017-10-10 09:00:09 +0800 CST\"\n==> Got : \n%s", s.(User).CreateTime)
			}
			if s.(User).Sex != true {
				t.Errorf("\n==> Except : \ntrue\n==> Got : \n%s", s.(User).Sex)
			}
		}
	}
}
func TestMapStructs(t *testing.T) {
	rs := ResultSet{}
	rs.Init(&rawRows)
	sts, err := rs.MapStructs("pid", User{})
	if err != nil {
		t.Errorf("\n==> Except : \nnil\n==> Got : \n%s", err)
	} else {
		for _, s := range sts {
			if s.(User).Name != "jack" {
				t.Errorf("\n==> Except : \njack\n==> Got : \n%s", s.(User).Name)
			}
			if s.(User).ID != 229 {
				t.Errorf("\n==> Except : \n229\n==> Got : \n%s", s.(User).ID)
			}
			if s.(User).Weight != 60 {
				t.Errorf("\n==> Except : \njack\n==> Got : \n%s", s.(User).Weight)
			}
			if s.(User).Height != 160.3 {
				t.Errorf("\n==> Except : \njack\n==> Got : \n%s", s.(User).Height)
			}
			if s.(User).Sex != true {
				t.Errorf("\n==> Except : \ntrue\n==> Got : \n%s", s.(User).Sex)
			}
			if s.(User).CreateTime.String() != "2017-10-10 09:00:09 +0800 CST" {
				t.Errorf("\n==> Except : \n\"2017-10-10 09:00:09 +0800 CST\"\n==> Got : \n%s", s.(User).CreateTime)
			}
			if s.(User).Sex != true {
				t.Errorf("\n==> Except : \ntrue\n==> Got : \n%s", s.(User).Sex)
			}
		}
	}
}
func TestUpdateBatch0(t *testing.T) {
	ar := ar().UpdateBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"id":      "id1",
			"gid":     22,
			"name":    "test1",
			"score +": 1,
		}, map[string]interface{}{
			"id":      "id2",
			"gid":     33,
			"name":    "test2",
			"score +": 1,
		},
	}, []string{"id", "gid"})
	fmt.Println(ar.SQL(), ar.Values())
}
