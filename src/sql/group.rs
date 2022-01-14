use crate::sql::comment::shouldbespace;
use crate::sql::common::commas;
use crate::sql::idiom::{idiom, Idiom};
use nom::bytes::complete::tag_no_case;
use nom::combinator::opt;
use nom::multi::separated_list1;
use nom::sequence::tuple;
use nom::IResult;
use serde::{Deserialize, Serialize};
use std::fmt;

#[derive(Clone, Debug, Default, Eq, PartialEq, PartialOrd, Serialize, Deserialize)]
pub struct Groups(pub Vec<Group>);

impl fmt::Display for Groups {
	fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
		write!(
			f,
			"GROUP BY {}",
			self.0.iter().map(|ref v| format!("{}", v)).collect::<Vec<_>>().join(", ")
		)
	}
}

#[derive(Clone, Debug, Default, Eq, PartialEq, PartialOrd, Serialize, Deserialize)]
pub struct Group {
	pub group: Idiom,
}

impl fmt::Display for Group {
	fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
		write!(f, "{}", self.group)
	}
}

pub fn group(i: &str) -> IResult<&str, Groups> {
	let (i, _) = tag_no_case("GROUP")(i)?;
	let (i, _) = opt(tuple((shouldbespace, tag_no_case("BY"))))(i)?;
	let (i, _) = shouldbespace(i)?;
	let (i, v) = separated_list1(commas, group_raw)(i)?;
	Ok((i, Groups(v)))
}

fn group_raw(i: &str) -> IResult<&str, Group> {
	let (i, v) = idiom(i)?;
	Ok((
		i,
		Group {
			group: v,
		},
	))
}

#[cfg(test)]
mod tests {

	use super::*;
	use crate::sql::test::Parse;

	#[test]
	fn group_statement() {
		let sql = "GROUP field";
		let res = group(sql);
		assert!(res.is_ok());
		let out = res.unwrap().1;
		assert_eq!(
			out,
			Groups(vec![Group {
				group: Idiom::parse("field")
			}])
		);
		assert_eq!("GROUP BY field", format!("{}", out));
	}

	#[test]
	fn group_statement_by() {
		let sql = "GROUP BY field";
		let res = group(sql);
		assert!(res.is_ok());
		let out = res.unwrap().1;
		assert_eq!(
			out,
			Groups(vec![Group {
				group: Idiom::parse("field")
			}])
		);
		assert_eq!("GROUP BY field", format!("{}", out));
	}

	#[test]
	fn group_statement_multiple() {
		let sql = "GROUP field, other.field";
		let res = group(sql);
		assert!(res.is_ok());
		let out = res.unwrap().1;
		assert_eq!(
			out,
			Groups(vec![
				Group {
					group: Idiom::parse("field")
				},
				Group {
					group: Idiom::parse("other.field")
				},
			])
		);
		assert_eq!("GROUP BY field, other.field", format!("{}", out));
	}
}