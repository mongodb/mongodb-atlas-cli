use std::fmt;
use std::str::FromStr;

/// Media type suffix of the Atlas Admin API accept header, e.g. `json` in
/// `application/vnd.atlas.preview+json`.
///
/// Only known subtypes are modeled; any other subtype fails to parse.
#[derive(
    Clone,
    Copy,
    Debug,
    Default,
    PartialEq,
    Eq,
    PartialOrd,
    Ord,
    Hash,
    serde::Serialize,
    serde::Deserialize
)]
#[serde(into = "String", try_from = "String")]
pub enum MediaType {
    #[default]
    Json,
    Csv,
    Gzip,
}

impl fmt::Display for MediaType {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Self::Json => write!(f, "json"),
            Self::Csv => write!(f, "csv"),
            Self::Gzip => write!(f, "gzip"),
        }
    }
}

impl From<MediaType> for String {
    fn from(mt: MediaType) -> Self {
        mt.to_string()
    }
}

impl TryFrom<String> for MediaType {
    type Error = MediaTypeParseError;

    fn try_from(s: String) -> Result<Self, Self::Error> {
        s.parse()
    }
}

impl FromStr for MediaType {
    type Err = MediaTypeParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s.trim().to_ascii_lowercase().as_str() {
            "json" => Ok(Self::Json),
            "csv" => Ok(Self::Csv),
            "gzip" => Ok(Self::Gzip),
            other => Err(MediaTypeParseError::Unsupported(other.to_owned())),
        }
    }
}

#[derive(Debug, thiserror::Error)]
#[error("unsupported media type `{0}`")]
pub enum MediaTypeParseError {
    Unsupported(String),
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn display_and_parse_roundtrip() {
        for mt in [MediaType::Json, MediaType::Csv, MediaType::Gzip] {
            assert_eq!(mt.to_string().parse::<MediaType>().unwrap(), mt);
        }
        assert_eq!(MediaType::Json.to_string(), "json");
        assert_eq!(MediaType::Csv.to_string(), "csv");
        assert_eq!(MediaType::Gzip.to_string(), "gzip");
    }

    #[test]
    fn parse_is_case_insensitive() {
        assert_eq!("JSON".parse::<MediaType>().unwrap(), MediaType::Json);
        assert_eq!("Csv".parse::<MediaType>().unwrap(), MediaType::Csv);
    }

    #[test]
    fn rejects_unknown_subtype() {
        assert!("xml".parse::<MediaType>().is_err());
        assert!("".parse::<MediaType>().is_err());
    }

    #[test]
    fn defaults_to_json() {
        assert_eq!(MediaType::default(), MediaType::Json);
    }

    #[test]
    fn serde_roundtrip() {
        let json = serde_json::to_string(&MediaType::Gzip).unwrap();
        assert_eq!(json, "\"gzip\"");
        assert_eq!(serde_json::from_str::<MediaType>(&json).unwrap(), MediaType::Gzip);
    }
}
