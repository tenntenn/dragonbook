package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"unicode"
)

type Parser struct {
	w         io.Writer
	r         *bufio.Reader
	lookahead rune
}

func NewParser(w io.Writer, r io.Reader) *Parser {
	return &Parser{
		w: w,
		r: bufio.NewReader(r),
	}
}

func (p *Parser) Parse() error {
	if err := p.next(); err != nil {
		return err
	}
	return p.expr()
}

func (p *Parser) next() error {
	c, _, err := p.r.ReadRune()
	if err != nil && err != io.EOF {
		return err
	}
	p.lookahead = c
	return nil
}

func (p *Parser) expr() error {
	if err := p.term(); err != nil {
		return err
	}

	for {
		switch p.lookahead {
		case '+':
			if err := p.match('+'); err != nil {
				return err
			}
			if err := p.term(); err != nil {
				return err
			}
			if _, err := fmt.Fprintf(p.w, "%c", '+'); err != nil {
				return err
			}
		case '-':
			if err := p.match('-'); err != nil {
				return err
			}
			if err := p.term(); err != nil {
				return err
			}
			if _, err := fmt.Fprintf(p.w, "%c", '-'); err != nil {
				return err
			}
		default:
			return nil
		}
	}
}

func (p *Parser) term() error {
	if !unicode.IsDigit(p.lookahead) {
		return errors.New("syntax error")
	}

	if _, err := fmt.Fprintf(p.w, "%c", p.lookahead); err != nil {
		return err
	}

	if err := p.match(p.lookahead); err != nil {
		return err
	}

	return nil
}

func (p *Parser) match(c rune) error {
	if p.lookahead != c {
		return errors.New("syntax error")
	}
	return p.next()
}
