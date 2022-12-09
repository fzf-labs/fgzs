package tree

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestFindParentNode(t *testing.T) {
	type args struct {
		node     *Tree
		allNodes []Tree
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestFindParentNode",
			args: args{node: &Tree{Id: 3, ParentId: 1, Title: "二级子集1"}, allNodes: []Tree{
				{Id: 1, ParentId: 0, Title: "总览1"},
				{Id: 2, ParentId: 0, Title: "总览2"},
				{Id: 3, ParentId: 1, Title: "二级子集1"},
				{Id: 4, ParentId: 1, Title: "二级子集2"},
				{Id: 5, ParentId: 2, Title: "二级子集3"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(&tt.args.node)
			FindParentNode(tt.args.node, tt.args.allNodes)
			fmt.Println(&tt.args.node)
			marshal, err := json.Marshal(tt.args.node)
			if err != nil {
				return
			}
			fmt.Println(string(marshal))
		})
	}
}

func TestFindSubNode(t *testing.T) {
	type args struct {
		node     *Tree
		allNodes []Tree
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestFindSubNode",
			args: args{node: &Tree{Id: 2, ParentId: 0, Title: "总览2"}, allNodes: []Tree{
				{Id: 1, ParentId: 0, Title: "总览1"},
				{Id: 2, ParentId: 0, Title: "总览2"},
				{Id: 3, ParentId: 1, Title: "二级子集1"},
				{Id: 4, ParentId: 1, Title: "二级子集2"},
				{Id: 5, ParentId: 2, Title: "二级子集3"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FindSubNode(tt.args.node, tt.args.allNodes)
			marshal, err := json.Marshal(tt.args.node)
			if err != nil {
				return
			}
			fmt.Println(string(marshal))
		})
	}
}

func TestGenerateTree(t *testing.T) {
	type args struct {
		list []Tree
	}
	tests := []struct {
		name string
		args args
		want []Tree
	}{
		{
			name: "tree",
			args: args{list: []Tree{
				{Id: 1, ParentId: 0, Title: "总览1"},
				{Id: 2, ParentId: 0, Title: "总览2"},
				{Id: 3, ParentId: 1, Title: "二级子集1"},
				{Id: 4, ParentId: 1, Title: "二级子集2"},
				{Id: 5, ParentId: 2, Title: "二级子集3"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateTree(tt.args.list)
			marshal, err := json.Marshal(got)
			if err != nil {
				return
			}
			fmt.Println(string(marshal))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrees_Len(t *testing.T) {
	tests := []struct {
		name string
		t    Trees
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrees_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		t    Trees
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrees_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		t    Trees
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_recursiveTree(t *testing.T) {
	type args struct {
		tree     *Tree
		allNodes []Tree
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recursiveTree(tt.args.tree, tt.args.allNodes)
		})
	}
}
