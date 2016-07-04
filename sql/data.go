// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

func (p *Parser) parseData() (exp []Expr, err error) {

	if tok, _, exi := p.mightBe(SET, MERGE, CONTENT); exi {

		if is(tok, SET) {
			if exp, err = p.parseSet(); err != nil {
				return nil, err
			}
		}

		if is(tok, MERGE) {
			if exp, err = p.parseMerge(); err != nil {
				return nil, err
			}
		}

		if is(tok, CONTENT) {
			if exp, err = p.parseContent(); err != nil {
				return nil, err
			}
		}

	}

	return

}

func (p *Parser) parseSet() (mul []Expr, err error) {

	var tok Token
	var lit string

	for {

		one := &BinaryExpression{}

		tok, lit, err = p.shouldBe(IDENT)
		if err != nil {
			return nil, &ParseError{Found: lit, Expected: []string{"field name"}}
		}

		one.LHS, err = declare(tok, lit)
		if err != nil {
			return nil, err
		}

		tok, lit, err = p.shouldBe(EQ, INC, DEC)
		if err != nil {
			return nil, err
		}
		one.Op = lit

		tok, lit, err = p.shouldBe(IDENT, NULL, VOID, NOW, DATE, TIME, TRUE, FALSE, STRING, REGION, NUMBER, DOUBLE, JSON, ARRAY)
		if err != nil {
			return nil, &ParseError{Found: lit, Expected: []string{"field value"}}
		}

		one.RHS, err = declare(tok, lit)
		if err != nil {
			return nil, err
		}

		mul = append(mul, one)

		if _, _, exi := p.mightBe(COMMA); !exi {
			p.unscan()
			break
		}

	}

	return mul, nil

}

func (p *Parser) parseDiff() (exp []Expr, err error) {

	one := &DiffExpression{}

	tok, lit, err := p.shouldBe(JSON)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	val, err := declare(tok, lit)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	one.JSON = val

	exp = append(exp, one)

	return

}

func (p *Parser) parseMerge() (exp []Expr, err error) {

	one := &MergeExpression{}

	tok, lit, err := p.shouldBe(JSON)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	val, err := declare(tok, lit)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	one.JSON = val

	exp = append(exp, one)

	return

}

func (p *Parser) parseContent() (exp []Expr, err error) {

	one := &ContentExpression{}

	tok, lit, err := p.shouldBe(JSON)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	val, err := declare(tok, lit)
	if err != nil {
		return nil, &ParseError{Found: lit, Expected: []string{"json"}}
	}

	one.JSON = val

	exp = append(exp, one)

	return

}