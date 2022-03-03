# struct

> Golang 结构体工具, Golang Struct Toolkit

---

## 生成 Gorm Struct

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
	Id         uint64          `json:"id" gorm:"column:id"`
	Name       string          `json:"name" gorm:"column:name"`
	Age        int8            `json:"age" gorm:"column:age"`
	AgeUnit    UserInfoAgeUnit `json:"age_unit" gorm:"column:age_unit"`
	Gender     UserInfoGender  `json:"gender" gorm:"column:gender"`
	CreateTime time.Time       `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time       `json:"update_time" gorm:"column:update_time"`
}

// TableName ...
func (t *UserInfo) TableName() string {
	return "user_info"
}

// UserInfoAgeUnit 年龄单位
type UserInfoAgeUnit int8

const (
	UserInfoAgeUnit0 UserInfoAgeUnit = iota
	UserInfoAgeUnit1
	UserInfoAgeUnit2
)

func (t UserInfoAgeUnit) String() string {
	switch t {
	case UserInfoAgeUnit0:
		return "岁"
	case UserInfoAgeUnit1:
		return "月"
	case UserInfoAgeUnit2:
		return "天"
	default:
		return ""
	}
}

// UserInfoGender 性别
type UserInfoGender int8

const (
	UserInfoGender0 UserInfoGender = iota
	UserInfoGender1
	UserInfoGender2
)

func (t UserInfoGender) String() string {
	switch t {
	case UserInfoGender0:
		return "未知"
	case UserInfoGender1:
		return "男"
	case UserInfoGender2:
		return "女"
	default:
		return ""
	}
}
```

---

## 生成普通 Struct

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
	Id         uint64    `json:"id"`
	Name       string    `json:"name"`
	Age        int8      `json:"age"`
	AgeUnit    int8      `json:"age_unit"`
	Gender     int8      `json:"gender"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
```

---

## 生成枚举代码

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

---

## 解析 sql 文件批量生成 Gorm Struct 文件

> sql 建议从数据库相关工具直接导出结构来使用

```shell
struct sql ~/Downloads/demo.sql

# Success: 写入文件成功: ~/Downloads/demo.sql-output/a.go
# Success: 写入文件成功: ~/Downloads/demo.sql-output/b.go
# Success: 写入文件成功: ~/Downloads/demo.sql-output/c.go
# Success: 写入文件成功: ~/Downloads/demo.sql-output/d.go
# 🍺🍺🍺 共生成 4 个文件
```
