# struct

> Golang 结构体工具

> Golang struct toolkit

---

## Mysql 转 Gorm Struct + 字段枚举

> 生成枚举需要注释满足格式 "注释: 0, 枚举说明; 1, 枚举说明;"

```shell
# 需要去除 sql 中的 `, shell 会默认执行 `` 中的文本
struct gorm "CREATE TABLE user_info (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(20) NOT NULL DEFAULT '' COMMENT '姓名',
  age tinyint(4) NOT NULL DEFAULT '0' COMMENT '年龄',
  age_unit tinyint(1) NOT NULL DEFAULT '0' COMMENT '年龄单位: 0, 岁; 1, 月; 2, 天;',
  gender tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别: 0, 未知; 1, 男; 2, 女;',
  create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (id)
) ENGINE=InnoDB COMMENT='用户信息表';"
```

```go
// UserInfo ...
type UserInfo struct {
	Id	int	`json:"id" gorm:"column:id" comment:""`
	Name	string	`json:"name" gorm:"column:name" comment:"姓名"`
	Age	int8	`json:"age" gorm:"column:age" comment:"年龄"`
	AgeUnit	AgeUnit	`json:"ageUnit" gorm:"column:age_unit" comment:"年龄单位: 0, 岁; 1, 月; 2, 天;"`
	Gender	Gender	`json:"gender" gorm:"column:gender" comment:"性别: 0, 未知; 1, 男; 2, 女;"`
	CreateTime	time.Time	`json:"createTime" gorm:"column:create_time" comment:"创建时间"`
	UpdateTime	time.Time	`json:"updateTime" gorm:"column:update_time" comment:"更新时间"`
}

// AgeUnit 年龄单位
type AgeUnit int8

const (
	AgeUnit0 AgeUnit = iota
	AgeUnit1
	AgeUnit2
)

func (t AgeUnit) String() string {
	switch t {
	case AgeUnit0:
		return "岁"
	case AgeUnit1:
		return "月"
	case AgeUnit2:
		return "天"
	default:
		return ""
	}
}

// Gender 性别
type Gender int8

const (
	Gender0 Gender = iota
	Gender1
	Gender2
)

func (t Gender) String() string {
	switch t {
	case Gender0:
		return "未知"
	case Gender1:
		return "男"
	case Gender2:
		return "女"
	default:
		return ""
	}
}
```

---

## Mysql 转 Common Struct

```shell
struct common "CREATE TABLE user_info (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(20) NOT NULL DEFAULT '' COMMENT '姓名',
  age tinyint(4) NOT NULL DEFAULT '0' COMMENT '年龄',
  age_unit tinyint(1) NOT NULL DEFAULT '0' COMMENT '年龄单位: 0, 岁; 1, 月; 2, 天;',
  gender tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别: 0, 未知; 1, 男; 2, 女;',
  create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (id)
) ENGINE=InnoDB COMMENT='用户信息表';"
```

```go
// UserInfo ...
type UserInfo struct {
	Id	int	`json:"id"`
	Name	string	`json:"name"`
	Age	int8	`json:"age"`
	AgeUnit	int8	`json:"ageUnit"`
	Gender	int8	`json:"gender"`
	CreateTime	time.Time	`json:"createTime"`
	UpdateTime	time.Time	`json:"updateTime"`
}
```

---

## 生成枚举

> 满足格式: "类型名称:类型:注释:0, 枚举说明;1, 枚举说明;"

```shell
struct enum "AgeUnit:int8:年龄单位: 0, 岁; 1, 月; 2, 天;"
```

```go
// AgeUnit 年龄单位
type AgeUnit int8

const (
	AgeUnit0 AgeUnit = iota
	AgeUnit1
	AgeUnit2
)

func (t AgeUnit) String() string {
	switch t {
	case AgeUnit0:
		return "岁"
	case AgeUnit1:
		return "月"
	case AgeUnit2:
		return "天"
	default:
		return ""
	}
}
```
