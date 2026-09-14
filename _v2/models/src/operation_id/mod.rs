mod noun;
mod verb;

use std::fmt;
use std::str::FromStr;

pub use noun::Noun;
pub use verb::Verb;

use noun::NounError;
use verb::VerbError;

/// Components of an Atlas Admin API operation ID, deconstructed from the IPA operation-ID
/// generator in `mongodb/openapi` (`tools/spectral/ipa/rulesets/functions/utils/operationIdGeneration.js`).
///
/// An operationId is `{verb}{Noun}{Noun}...`, e.g. `getProject` is the verb `get` and the
/// noun `Project`; `createX509AuthenticationDatabaseUser` is the verb `create` with the
/// nouns `X509`, `Authentication`, `Database`, `User`.
///
/// Words are split with the same boundary as the `CAMEL_CASE_WITH_ABBREVIATIONS` regex the
/// generator uses. Noun casing is preserved (`X509`, `API`), so `FromStr`/`Display` round-trip
/// any published operationId exactly.
#[derive(
    Clone, Debug, PartialEq, Eq, PartialOrd, Ord, Hash, serde::Serialize, serde::Deserialize
)]
#[serde(into = "String", try_from = "String")]
pub struct OperationId {
    pub verb: Verb,
    pub nouns: Vec<Noun>,
}

impl fmt::Display for OperationId {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", self.verb)?;
        for noun in &self.nouns {
            write!(f, "{noun}")?;
        }
        Ok(())
    }
}

impl From<OperationId> for String {
    fn from(id: OperationId) -> Self {
        id.to_string()
    }
}

impl TryFrom<String> for OperationId {
    type Error = OperationIdParseError;

    fn try_from(s: String) -> Result<Self, Self::Error> {
        s.parse()
    }
}

impl FromStr for OperationId {
    type Err = OperationIdParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let s = s.trim();
        let Some(first) = s.chars().next() else {
            return Err(OperationIdParseError::MissingVerb);
        };
        if !first.is_ascii_alphabetic()
            || !s
                .chars()
                .all(|c| c.is_ascii_alphanumeric())
        {
            return Err(OperationIdParseError::Invalid(s.to_owned()));
        }

        let mut words = camel_words(s).into_iter();
        let Some(verb_word) = words.next() else {
            return Err(OperationIdParseError::MissingVerb);
        };
        if verb_word.is_empty() {
            return Err(OperationIdParseError::MissingVerb);
        }

        let verb = Verb::from_str(&verb_word).map_err(OperationIdParseError::InvalidVerb)?;
        let mut nouns = Vec::new();
        for (index, word) in words.enumerate() {
            let noun = Noun::from_str(&word)
                .map_err(|source| OperationIdParseError::InvalidNoun { index, source })?;
            nouns.push(noun);
        }

        Ok(Self { verb, nouns })
    }
}

/// Splits a camelCase string into words at the same boundaries as the
/// `CAMEL_CASE_WITH_ABBREVIATIONS` regex from `operationIdGeneration.js`:
/// a lowercase/digit-to-uppercase boundary, or the end of an acronym run (`X509`, `API`).
fn camel_words(s: &str) -> Vec<String> {
    let chars: Vec<char> = s.chars().collect();
    let mut words = Vec::new();
    let mut start = 0;

    for i in 0..chars.len() {
        if !chars[i].is_ascii_uppercase() {
            continue;
        }
        if i == start {
            continue;
        }
        let prev = chars[i - 1];
        let boundary = prev.is_ascii_lowercase()
            || prev.is_ascii_digit()
            || (prev.is_ascii_uppercase()
                && chars.get(i + 1).is_some_and(|c| c.is_ascii_lowercase()));
        if boundary {
            words.push(chars[start..i].iter().collect());
            start = i;
        }
    }

    words.push(chars[start..].iter().collect());
    words
}

#[derive(Debug, thiserror::Error)]
pub enum OperationIdParseError {
    /// Contains characters outside `[A-Za-z0-9]`.
    #[error("invalid operation ID `{0}`, expected only ASCII alphanumeric characters")]
    Invalid(String),
    /// No verb word, e.g. the empty string.
    #[error("operation ID is missing its action verb")]
    MissingVerb,
    #[error(transparent)]
    InvalidVerb(#[from] VerbError),
    #[error("invalid operation ID noun at index {index}: {source}")]
    InvalidNoun {
        index: usize,
        source: NounError,
    },
}

#[cfg(test)]
mod tests {
    use super::*;

    fn noun(s: &str) -> Noun {
        s.parse().unwrap()
    }

    #[test]
    fn parses_into_verb_and_nouns() {
        let id: OperationId = "getProject".parse().unwrap();
        assert_eq!(id.verb.as_ref(), "get");
        assert_eq!(id.nouns, vec![noun("Project")]);

        let id: OperationId = "createX509AuthenticationDatabaseUser".parse().unwrap();
        assert_eq!(id.verb.as_ref(), "create");
        assert_eq!(
            id.nouns,
            vec![
                noun("X509"),
                noun("Authentication"),
                noun("Database"),
                noun("User"),
            ]
        );
    }

    #[test]
    fn preserves_abbreviations() {
        let id: OperationId = "getOpenAPI".parse().unwrap();
        assert_eq!(id.verb.as_ref(), "get");
        assert_eq!(id.nouns, vec![noun("Open"), noun("API")]);
    }

    #[test]
    fn verb_only_operation_ids_parse() {
        let id: OperationId = "get".parse().unwrap();
        assert_eq!(id.verb.as_ref(), "get");
        assert!(id.nouns.is_empty());
    }

    #[test]
    fn display_and_parse_roundtrip() {
        for id in [
            "getProject",
            "createX509AuthenticationDatabaseUser",
            "listFederationSettingConnectedOrgConfigRoleMappings",
            "updateGroupAiModelApiCloudGeographyModelGroupNameRateLimits",
            "getOpenAPI",
            "resetGroupAiModelApiCloudGeographyModelGroupNameRateLimits",
            "get",
        ] {
            let parsed: OperationId = id.parse().unwrap();
            assert_eq!(parsed.to_string(), id, "roundtrip failed for `{id}`");
        }
    }

    #[test]
    fn rejects_invalid_input() {
        for bad in [
            "",            // empty
            "   ",         // whitespace-only
            "1234",        // verb must start with a letter
            "get_project", // separator, not alphanumeric
            "get project",
            "get/project",
        ] {
            assert!(
                bad.parse::<OperationId>().is_err(),
                "should reject `{bad}`"
            );
        }
    }

    #[test]
    fn serde_roundtrip() {
        let id: OperationId = "createX509AuthenticationDatabaseUser".parse().unwrap();
        let json = serde_json::to_string(&id).unwrap();
        assert_eq!(json, "\"createX509AuthenticationDatabaseUser\"");
        let back: OperationId = serde_json::from_str(&json).unwrap();
        assert_eq!(back, id);
    }
}
