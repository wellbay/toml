package toml

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var tomlText = `
# test comments preserving

version = "v1.0.1" # 这里是软件的版本
release = {release_date = "2023-03-27", stable = true, author = "zhangyi", hobbies = [ # comment for hobbies
    "basketball", # comment for hobbies.basketball
    "football", # comment for hobbies.football
    "play game", # comment for hobbies.play_game
]} # comment for release key
release_history = [
    {data = "2020-01-01", version = "v0.0.1"}, {date = "2021-01-01", version = "v0.1"}, # version 0.1
    {date = "2022-01-01", version = "v0.2"}, # version 0.2
    {date = "2023-01-01", version = "v0.3"}  # version 0.3
]



[object]
name = "zhangsan"  # 姓名
age = +35  # 年龄
gender = "male" # 性别
married = true # 是否已婚
# 这里是所有运动的枚举
arr = [ #comment for arr
    "basketball",   # 篮球
    "football", # 足球
    "ping pong", # 乒乓球
    # this is comment for sub array
    [
        "foo", # sub array foo
        "bar", # sub array bar
        "foobar" # sub array foobar
    ], # comment for sub array
    {user = "foo", age = 32, is_student = false, friends = ["Tom", "Jerry", "Wangwangwang"] }, # 行内表注释
]    # also comment for arr


[[dataway]]
urls = "https://www.baidu.com" # comment for dataway[0].urls
timeout = "5s"
max_retry_count = 10

[profile]
endpoint = "/v1/write/profile"
tags = ["foo=bar", "foo1=bar1"]

[[dataway]]
urls = "https://www.google.com"
timeout = "5s"
max_retry_count = 10

[[dataway.admin]]
name = "admin1"
level = 10
[[dataway.admin]]
name = "admin2"
level = 5

#comment for object.country[0]
[[object.country]]
name = "American"
short_name = "USA"

#comment for object.country[1]
[[object.country]]
name = "China"
short_name = "CN"
`

func TestEncodeComment(t *testing.T) {
	var x interface{}
	meta, err := Decode(tomlText, &x)
	if err != nil {
		log.Fatal(err)
	}

	enc := NewEncoder(os.Stdout)
	if err := enc.EncodeWithComments(x, meta); err != nil {
		log.Fatal(err)
	}
}

func TestDecodeComment(t *testing.T) {
	var x interface{}
	meta, err := Decode(tomlText, &x)
	if err != nil {
		t.Fatal(err)
	}
	dumpKeys(meta.comments, "")
}

func dumpKeys(arr map[segment]*KeySegments, indent string) {
	for _, kc := range arr {
		if kc.documentComment != nil {
			fmt.Println(indent + kc.documentComment.String())
		}
		fmt.Print(indent + kc.String())
		if kc.lineTailComment != nil {
			fmt.Printf(" %s", kc.lineTailComment.String())
		}

		fmt.Println("\n--------------------------------------")

		if len(kc.children) > 0 {
			dumpKeys(kc.children, indent+"\t\t")
		}
	}
}
