package tsql

import "github.com/tbshill/goparsec"

var _ Node = &BlockCommentNode{}

type BlockCommentNode struct {
	Text string
}

func (b *BlockCommentNode) Parse(in string) (string, error) {
	var rem string
	var err error

	if _, rem, err = ExpectMultiLineCommentStart(in); err != nil {
		return in, err
	}

	if b.Text, rem, err = goparsec.ExpectUntil(ExpectMultiLineCommentEnd)(rem); err != nil {
		return in, err
	}

	if _, rem, err = ExpectMultiLineCommentEnd(rem); err != nil {
		return in, err
	}

	return rem, nil
}

func (b *BlockCommentNode) String() string {
	return BlockCommentStart + b.Text + BlockCommentEnd
}
