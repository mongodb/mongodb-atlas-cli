use nutype::nutype;

#[nutype(
    validate(predicate = |values| !values.is_empty() && values.iter().all(|v| !v.is_empty())),
    derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)
)]
pub struct ValueTypeEnum(Vec<String>);

#[cfg(test)]
mod tests {
    use super::ValueTypeEnum;

    #[test]
    fn validates() {
        assert!(ValueTypeEnum::try_new(vec!["a".into(), "b".into()]).is_ok());
        assert!(ValueTypeEnum::try_new(Vec::<String>::new()).is_err());
        assert!(ValueTypeEnum::try_new(vec!["a".into(), String::new()]).is_err());
    }
}
