use nutype::nutype;

#[nutype(
    sanitize(trim),
    validate(not_empty),
    derive(Debug, PartialEq, Eq, PartialOrd, Ord, Hash, AsRef, Display, Clone)
)]
pub struct OperationId(String);
