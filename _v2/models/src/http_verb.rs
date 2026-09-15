use nutype::nutype;

#[nutype(
    sanitize(trim, uppercase),
    validate(not_empty),
    derive(
        Clone,
        Debug,
        PartialEq,
        Eq,
        PartialOrd,
        Ord,
        Hash,
        Display,
        AsRef,
        Deref,
        Serialize,
        Deserialize
    )
)]
pub struct Verb(String);
