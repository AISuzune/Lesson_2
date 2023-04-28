package main

import (
	"fmt"
	"time"
)

func PrintStruct(s interface{}) {
	fmt.Printf("Type: %T\nValue: %v\n", s, s)
}

type Question struct {
	Id          int64     //问题的Id
	User        User      //提问者用户
	Topic       string    //主题
	Title       string    //标题
	Content     string    //内容
	Comments    int64     //评论数(Answer)
	Collections int64     //收藏数
	Likes       int64     //点赞数
	Pageviews   int64     //浏览量
	Followers   int64     //订阅数
	Answer      []Answer  //回答
	CreatedAt   time.Time //创建时间
	DeletedAt   time.Time //删除时间
	UpdateAt    time.Time // 跟新时间
}

type User struct {
	Id   int64  //用户Id
	Name string // 用户姓名
}

type Answer struct {
	Id          int64
	User        User      //回答(Answer)的用户
	Pid         int64     //Answer的父Id(Question)
	Content     string    //回答的内容
	Likes       int64     // 喜欢数
	Disagrees   int64     //不同意
	Agrees      int64     //同意
	Collections int64     //收藏
	Followers   int64     //订阅
	Comment     []Comment //评论
	CreatedAt   time.Time
	DeletedAt   time.Time
	UpdateAt    time.Time
}

type Comment struct {
	Id                 int64
	Type               int // 0为一级评论 1为二级评论
	User               User
	ReplyId            int64
	Content            string
	Likes              int64
	ChildCommentNumber int64
	CreateAt           time.Time
	Ip                 string //Ip 地址
	Pid                int64
}

func (q *Question) AddQuestion() {
	q.User = User{
		Id:   1,
		Name: "zty",
	}
	q.Topic = "yuan神"
	q.Content = "为什么你喜欢yuan神呢？"
	q.CreatedAt = time.Now()
	q.UpdateAt = time.Now()
}

func (q *Question) AddAnswer() {
	answer := Answer{
		Id:      1,
		Content: "yuan神长很帅",
		User: User{
			Id:   2,
			Name: "Zty",
		},
		Pid: 1,
	}
	answer.AddComment()
	q.Answer = []Answer{answer}
}

func (a *Answer) AddComment() {
	comment1 := Comment{
		Id:      1,
		Content: "是的 男神",
		User: User{
			Id:   3,
			Name: "zty",
		},
		Pid: 1,
	}
	comment2 := Comment{
		Id:      2,
		Content: "yuan神 YYDS",
		User: User{
			Id:   4,
			Name: "ty",
		},
		ReplyId: 1,
		Pid:     1,
	}
	a.Comment = []Comment{comment1, comment2}
}

func (q *Question) Print() {
	PrintStruct(q)
}

func main() {
	Q := &Question{
		Id: 1,
		User: User{
			Id:   1,
			Name: "Z",
		},
		Topic:     "yuan神",
		Content:   "为什么你喜欢yuan神呢？",
		CreatedAt: time.Now(),
		Answer: []Answer{
			{
				Id:      1,
				Content: "yuan神长很帅",
				User: User{
					Id:   2,
					Name: "Zty",
				},
				Pid: 1,
				Comment: []Comment{
					{
						Id:      1,
						Content: "是的 男神",
						User: User{
							Id:   3,
							Name: "zty",
						},
						Pid: 1,
					},
					{
						Id:      2,
						Content: "yuan神 YYDS",
						User: User{
							Id:   4,
							Name: "ty",
						},
						ReplyId: 1,
						Pid:     1,
					},
				},
			},
		},
		UpdateAt: time.Now(),
	}
	Q.Print()
}
