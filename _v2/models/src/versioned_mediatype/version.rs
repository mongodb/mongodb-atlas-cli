use std::fmt;
use std::str::FromStr;

use super::date::VersionDateParseError;
use super::VersionDate;

/// Version part of an Atlas Admin API accept header.
///
/// Variants listed in the MMS versioned API lifecycle order: preview → upcoming → stable.
#[derive(
    Clone, Copy, Debug, PartialEq, Eq, Hash, serde::Serialize, serde::Deserialize
)]
#[serde(into = "String", try_from = "String")]
pub enum Version {
    Preview,
    Upcoming(VersionDate),
    Stable(VersionDate),
}

impl fmt::Display for Version {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Self::Stable(date) => write!(f, "{date}"),
            Self::Upcoming(date) => write!(f, "{date}.upcoming"),
            Self::Preview => write!(f, "preview"),
        }
    }
}

impl From<Version> for String {
    fn from(v: Version) -> Self {
        v.to_string()
    }
}

impl TryFrom<String> for Version {
    type Error = VersionParseError;

    fn try_from(s: String) -> Result<Self, Self::Error> {
        s.parse()
    }
}

impl FromStr for Version {
    type Err = VersionParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let s = s.trim();
        if s.eq_ignore_ascii_case("preview") {
            return Ok(Self::Preview);
        }
        if let Some(prefix) = s.strip_suffix(".upcoming") {
            return Ok(Self::Upcoming(VersionDate::from_str(prefix)?));
        }
        Ok(Self::Stable(VersionDate::from_str(s)?))
    }
}

#[derive(Debug, thiserror::Error)]
pub enum VersionParseError {
    #[error("invalid version: {0}")]
    InvalidDate(#[from] VersionDateParseError),
}

#[cfg(test)]
mod tests {
    use super::*;

    fn date(y: &str) -> VersionDate {
        VersionDate::from_str(y).unwrap()
    }

    #[test]
    fn display_matches_extension_types() {
        assert_eq!(Version::Stable(date("2024-10-23")).to_string(), "2024-10-23");
        assert_eq!(
            Version::Upcoming(date("2025-09-22")).to_string(),
            "2025-09-22.upcoming"
        );
        assert_eq!(Version::Preview.to_string(), "preview");
    }

    #[test]
    fn parse_roundtrip() {
        for v in [
            Version::Stable(date("2024-10-23")),
            Version::Upcoming(date("2025-09-22")),
            Version::Preview,
        ] {
            assert_eq!(v.to_string().parse::<Version>().unwrap(), v);
        }
    }

    #[test]
    fn parse_is_case_insensitive_for_preview() {
        assert_eq!("PREVIEW".parse::<Version>().unwrap(), Version::Preview);
    }

    #[test]
    fn rejects_invalid_input() {
        for bad in ["garbage", "preview.upcoming", "2024-13-01", "", "2024-10-23foo"] {
            assert!(bad.parse::<Version>().is_err(), "should reject `{bad}`");
        }
    }

    #[test]
    fn serde_roundtrip() {
        let v = Version::Upcoming(date("2025-09-22"));
        let json = serde_json::to_string(&v).unwrap();
        assert_eq!(json, "\"2025-09-22.upcoming\"");
        let back: Version = serde_json::from_str(&json).unwrap();
        assert_eq!(back, v);
    }
}
