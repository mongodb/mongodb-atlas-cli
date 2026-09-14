use nutype::nutype;

/// One resource noun of an operation ID — a single camel-case word following the verb.
///
/// Nouns come from path resource identifiers in the IPA operation-ID generator
/// (`operationIdGeneration.js#generateOperationID`). Case is preserved as written so
/// abbreviations round-trip exactly: `X509`, `API`, `Fts` keep their casing.
#[nutype(
    sanitize(trim),
    validate(not_empty),
    derive(Clone, Debug, PartialEq, Eq, PartialOrd, Ord, Hash, Display, AsRef, Deref, FromStr)
)]
pub struct Noun(String);
