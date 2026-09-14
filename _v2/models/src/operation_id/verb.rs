use nutype::nutype;

/// Action verb of an operation ID — the first camel-case word, lower-cased.
///
/// Mirrors the `methodName` the IPA operation-ID generator derives the leading verb from:
/// a standard method (`get`, `list`, `create`, `update`, `delete`, `reset`) or a custom
/// method action verb (`revoke`, `add`, `remove`, ...). It must start with a letter, matching
/// the camel-case method names the IPA rules accept (`IPA-10x operation-id` rulesets).
#[nutype(
    sanitize(trim, lowercase),
    validate(not_empty),
    derive(Clone, Debug, PartialEq, Eq, PartialOrd, Ord, Hash, Display, AsRef, Deref, FromStr)
)]
pub struct Verb(String);
