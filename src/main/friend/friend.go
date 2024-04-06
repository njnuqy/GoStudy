package friend

import (
	"fmt"
	"sync"
)

var (
	cache = &sync.Pool{
		New: func() interface{} {
			return &option{sex: 0}
		},
	}
)

type Option func(*option)

type option struct {
	sex    int
	age    int
	height int
	weight int
	hobby  string
}

func (o *option) reset() {
	o.sex = 0
	o.age = 0
	o.height = 0
	o.weight = 0
	o.hobby = ""
}

func getOption() *option {
	return cache.Get().(*option)
}

func releaseOption(opt *option) {
	opt.reset()
	cache.Put(opt)
}

// WithSex setup sex, 1=female 2=male
func WithSex(sex int) Option {
	return func(opt *option) {
		opt.sex = sex
	}
}

// WithAge setup age
func WithAge(age int) Option {
	return func(opt *option) {
		opt.age = age
	}
}

// WithHeight set up height
func WithHeight(height int) Option {
	return func(opt *option) {
		opt.height = height
	}
}

// WithWeight set up weight
func WithWeight(weight int) Option {
	return func(opt *option) {
		opt.weight = weight
	}
}

// WithHobby set up Hobby
func WithHobby(hobby string) Option {
	return func(opt *option) {
		opt.hobby = hobby
	}
}

func Find(where string, options ...Option) (string, error) {
	friend := fmt.Sprintf("从 %s 找朋友\n", where)

	opt := getOption()
	defer func() {
		releaseOption(opt)
	}()

	for _, f := range options {
		f(opt)
	}

	if opt.sex == 1 {
		sex := "性别：女性"
		friend += fmt.Sprintf("%s\n", sex)
	}
	if opt.sex == 2 {
		sex := "性别：男性"
		friend += fmt.Sprintf("%s\n", sex)
	}

	if opt.age != 0 {
		age := fmt.Sprintf("年龄：%d岁", opt.age)
		friend += fmt.Sprintf("%s\n", age)
	}

	if opt.height != 0 {
		height := fmt.Sprintf("身高：%dcm", opt.height)
		friend += fmt.Sprintf("%s\n", height)
	}

	if opt.weight != 0 {
		weight := fmt.Sprintf("体重：%dkg", opt.weight)
		friend += fmt.Sprintf("%s\n", weight)
	}

	if opt.hobby != "" {
		hobby := fmt.Sprintf("爱好：%s", opt.hobby)
		friend += fmt.Sprintf("%s\n", hobby)
	}

	return friend, nil
}
