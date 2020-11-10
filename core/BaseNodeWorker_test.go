package core_test

import (
	"testing"

	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/loader"
	"github.com/stretchr/testify/require"
)

// 自定义action节点
type actionTest struct {
	Action
}

func (a *actionTest) Initialize(setting *config.BTNodeCfg) {
	a.Action.Initialize(setting)
}

func (a *actionTest) OnTick(tick *Tick) b3.Status {
	tick.Blackboard.SetMem("OnTick", true)
	return b3.SUCCESS
}

func (a *actionTest) OnEnter(tick *Tick) {
	tick.Blackboard.SetMem("OnEnter", true)
}

func (a *actionTest) OnOpen(tick *Tick) {
	tick.Blackboard.SetMem("OnOpen", true)
}

func (a *actionTest) OnClose(tick *Tick) {
	tick.Blackboard.SetMem("OnClose", true)
}

func (a *actionTest) OnExit(tick *Tick) {
	tick.Blackboard.SetMem("OnExit", true)
}

func TestBaseNodeWorker(t *testing.T) {
	assert := require.New(t)

	maps := loader.NewRegisterStructMaps()
	maps.Register("actionTest", new(actionTest))

	tree := NewBehaviorTree()
	node, nodeErr := maps.New("actionTest")
	assert.NoError(nodeErr)
	assert.NotNil(node)
	baseNode := node.(IBaseNode)
	assert.NotNil(baseNode)

	spec := &config.BTNodeCfg{
		Id:          "id",
		Name:        "name",
		Category:    b3.ACTION,
		Title:       "title",
		Description: "desc",
	}

	baseNode.Ctor()
	baseNode.Initialize(spec)
	baseNode.SetBaseNodeWorker(baseNode.(IBaseWorker))

	tree.SetRoot(baseNode)

	board := NewBlackboard()
	tree.Tick(tree, board)

	assert.True(board.GetMem("OnEnter").(bool))
	assert.True(board.GetMem("OnOpen").(bool))
	assert.True(board.GetMem("OnTick").(bool))
	assert.True(board.GetMem("OnClose").(bool))
	assert.True(board.GetMem("OnExit").(bool))
}
