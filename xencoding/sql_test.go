package xencoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLPretty(t *testing.T) {
	type args struct {
		sql string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				sql: `select foo, bar from baz where foo > 5 and bar < 2`,
			},
			want: `SELECT foo, bar FROM baz WHERE foo > 5 AND bar < 2`,
		},
		{
			name: "2",
			args: args{
				sql: `select a.id, a.name, count( b.id ) from a left join b on a.id = b.id where a.id in ( 1, 2, 3) group by id order by id asc;`,
			},
			want: `SELECT
	a.id, a.name, count(b.id)
FROM
	a LEFT JOIN b ON a.id = b.id
WHERE
	a.id IN (1, 2, 3)
GROUP BY
	id
ORDER BY
	id ASC`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SQLPretty(tt.args.sql), "SQLPretty(%v)", tt.args.sql)
		})
	}
}

func TestSQLMinify(t *testing.T) {
	type args struct {
		sql string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				sql: `select   foo, bar   from baz where foo > 5 and bar < 2`,
			},
			want: `select foo, bar from baz where foo > 5 and bar < 2`,
		},
		{
			name: "2",
			args: args{
				sql: `SELECT
	a.id, a.name, count(b.id)
FROM
	a LEFT JOIN b ON a.id = b.id
WHERE
	a.id IN (1, 2, 3)
GROUP BY
	id
ORDER BY
	id ASC`,
			},
			want: `SELECT a.id, a.name, count(b.id) FROM a LEFT JOIN b ON a.id = b.id WHERE a.id IN (1, 2, 3) GROUP BY id ORDER BY id ASC`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SQLMinify(tt.args.sql), "SQLMinify(%v)", tt.args.sql)
		})
	}
}
