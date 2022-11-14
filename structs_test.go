package structs

import (
	"reflect"
	"testing"
)

type (
	Case1 struct{}
	Case2 struct{}
	Case3 struct{}
	Case4 struct{}
	Case5 interface{}
)

type Example1 struct {
	Case1 *Case1
	Case2 *Case2
	Case3 *Case3
	Case4 *Case4
}

type Example2 struct {
	Case1 *Case1
	Case2 *Case2
	Case3 *Case3
	Case4 *Case4
	*Example1
}

type Example3 struct {
	Interface Case5
	Example2
}

func TestGetNilFields(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		args args
		want []string
	}{
		{
			args: args{
				value: new(Example1),
			},
			want: []string{
				"Example1.Case1",
				"Example1.Case2",
				"Example1.Case3",
				"Example1.Case4",
			},
		},
		{
			args: args{
				value: Example1{
					Case2: new(Case2),
				},
			},
			want: []string{
				"Example1.Case1",
				"Example1.Case3",
				"Example1.Case4",
			},
		},
		{
			args: args{
				value: Example2{
					Case1: new(Case1),
					Case2: new(Case2),
					Case3: new(Case3),
					Case4: new(Case4),
				},
			},
			want: []string{
				"Example2.Example1",
			},
		},
		{
			args: args{
				value: &Example2{
					Case1:    new(Case1),
					Case2:    new(Case2),
					Case3:    new(Case3),
					Case4:    new(Case4),
					Example1: new(Example1),
				},
			},
			want: []string{
				"Example2.Example1.Case1",
				"Example2.Example1.Case2",
				"Example2.Example1.Case3",
				"Example2.Example1.Case4",
			},
		},
		{
			args: args{
				value: Example2{
					Case1: new(Case1),
					Case2: new(Case2),
					Case3: new(Case3),
					Case4: new(Case4),
					Example1: &Example1{
						Case2: new(Case2),
					},
				},
			},
			want: []string{
				"Example2.Example1.Case1",
				"Example2.Example1.Case3",
				"Example2.Example1.Case4",
			},
		},
		{
			args: args{
				value: Example3{
					Example2: Example2{
						Case2: new(Case2),
					},
				},
			},
			want: []string{
				"Example3.Interface",
				"Example3.Example2.Case1",
				"Example3.Example2.Case3",
				"Example3.Example2.Case4",
				"Example3.Example2.Example1",
			},
		},
		{
			args: args{
				value: new(Example3),
			},
			want: []string{
				"Example3.Interface",
				"Example3.Example2.Case1",
				"Example3.Example2.Case2",
				"Example3.Example2.Case3",
				"Example3.Example2.Case4",
				"Example3.Example2.Example1",
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GetNilFields(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNilFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStructName(t *testing.T) {
	var testStruct struct{}
	type args struct {
		obj any
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				obj: Example1{},
			},
			want: "Example1",
		},
		{
			args: args{
				obj: Case1{},
			},
			want: "Case1",
		},
		{
			args: args{
				obj: testStruct,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GetStructName(tt.args.obj); got != tt.want {
				t.Errorf("GetStructName() = %v, want %v", got, tt.want)
			}
		})
	}
}
