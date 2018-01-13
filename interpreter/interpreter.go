package interpreter

import (
	"strings"
)

type node interface {
	Parse(context *context)
	ToString() string
}

type ProgramNode struct {
	commandListNode node
}

func (pn *ProgramNode) Parse(context *context) {
	context.skipToken("program")
	pn.commandListNode = &commandListNode{}
	pn.commandListNode.Parse(context)
}

func (pn *ProgramNode) ToString() string {
	return "program: " + pn.commandListNode.ToString()
}

type commandListNode struct {
	list []node
}

func (cln *commandListNode) Parse(context *context) {
	for {
		if context.currentToken == "end" {
			break
		} else {
			commandNode := &commandNode{}
			commandNode.Parse(context)
			cln.list = append(cln.list, commandNode)
		}
	}
}

func (cln *commandListNode) ToString() string {
	var str string
	for _, l := range cln.list {
		str += l.ToString()
	}
	return str
}

type commandNode struct {
	node node
}

func (cn *commandNode) Parse(context *context) {
	cn.node = &primitiveCommandNode{}
	cn.node.Parse(context)
}

func (cn *commandNode) ToString() string {
	return cn.node.ToString()
}

type primitiveCommandNode struct {
	name string
}

func (pcn *primitiveCommandNode) Parse(context *context) {
	pcn.name = context.currentToken
	context.skipToken(pcn.name)
}

func (pcn *primitiveCommandNode) ToString() string {
	return pcn.name + " "
}

type context struct {
	text         string
	tokens       []string
	currentToken string
}

func NewContext(text string) *context {
	context := &context{
		text:   text,
		tokens: strings.Fields(text),
	}
	context.nextToken()
	return context
}

func (c *context) nextToken() string {
	if len(c.tokens) == 0 {
		c.currentToken = ""
	} else {
		c.currentToken = c.tokens[0]
		c.tokens = c.tokens[1:]
	}
	return c.currentToken
}

func (c *context) skipToken(token string) {
	c.nextToken()
}
