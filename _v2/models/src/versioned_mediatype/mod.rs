mod date;
mod mediatype;
mod version;

use std::fmt;
use std::str::FromStr;

pub use date::VersionDate;
pub use mediatype::MediaType;
pub use version::Version;

use mediatype::MediaTypeParseError;
use version::VersionParseError;

const ACCEPT_HEADER_PREFIX: &str = "application/vnd.atlas.";

/// A full Atlas Admin API versioned accept header, e.g. `application/vnd.atlas.2024-10-23+json`.
///
/// Mirrors `RequestVersion` from MMS: a version (preview, upcoming date or stable date) plus a
/// media type subtype. Format follows the `VersioningConstants` accept-header templates, e.g.
/// `application/vnd.atlas.{version}+{subtype}`.
#[derive(
    Clone, Copy, Debug, PartialEq, Eq, Hash, serde::Serialize, serde::Deserialize
)]
#[serde(into = "String", try_from = "String")]
pub struct VersionedAcceptHeader {
    pub version: Version,
    pub mediatype: MediaType,
}

impl fmt::Display for VersionedAcceptHeader {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "{ACCEPT_HEADER_PREFIX}{}+{}",
            self.version, self.mediatype
        )
    }
}

impl From<VersionedAcceptHeader> for String {
    fn from(h: VersionedAcceptHeader) -> Self {
        h.to_string()
    }
}

impl TryFrom<String> for VersionedAcceptHeader {
    type Error = VersionedAcceptHeaderParseError;

    fn try_from(s: String) -> Result<Self, Self::Error> {
        s.parse()
    }
}

impl FromStr for VersionedAcceptHeader {
    type Err = VersionedAcceptHeaderParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let s = s.trim();
        // Accept headers are case-insensitive (Pattern.CASE_INSENSITIVE on the Java side).
        let s = s.to_ascii_lowercase();
        let rest = s.strip_prefix(ACCEPT_HEADER_PREFIX).ok_or_else(|| {
            VersionedAcceptHeaderParseError::InvalidPrefix(s.to_owned())
        })?;
        let (version, mediatype) = rest
            .split_once('+')
            .ok_or(VersionedAcceptHeaderParseError::MissingMediaType)?;

        let version = Version::from_str(version)?;
        let mediatype = MediaType::from_str(mediatype)?;

        Ok(Self { version, mediatype })
    }
}

#[derive(Debug, thiserror::Error)]
pub enum VersionedAcceptHeaderParseError {
    /// Does not start with `application/vnd.atlas.`.
    #[error("not a versioned accept header `{0}`, expected prefix `application/vnd.atlas.`")]
    InvalidPrefix(String),
    /// Missing the `+<subtype>` suffix.
    #[error("versioned accept header is missing its `+<subtype>` suffix")]
    MissingMediaType,
    #[error(transparent)]
    InvalidVersion(#[from] VersionParseError),
    #[error(transparent)]
    UnsupportedMediaType(#[from] MediaTypeParseError),
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn display_and_parse_roundtrip() {
        for header in [
            "application/vnd.atlas.preview+json",
            "application/vnd.atlas.2024-10-23+json",
            "application/vnd.atlas.2025-09-22.upcoming+csv",
            "application/vnd.atlas.2024-08-05+json",
        ] {
            let parsed: VersionedAcceptHeader = header.parse().unwrap();
            assert_eq!(parsed.to_string(), header);
        }
    }

    #[test]
    fn parse_is_case_insensitive() {
        assert!(matches!(
            "Application/VND.ATLAS.PREVIEW+JSON".parse::<VersionedAcceptHeader>(),
            Ok(VersionedAcceptHeader {
                version: Version::Preview,
                mediatype: MediaType::Json,
            })
        ));
    }

    #[test]
    fn rejects_invalid_headers() {
        for bad in [
            "application/vnd.other.2024-10-23+json", // wrong prefix
            "application/vnd.atlas.2024-10-23",      // missing media type suffix
            "application/vnd.atlas.2024-13-01+json", // invalid version date
            "application/vnd.atlas.preview+xml",     // unsupported media type
            "application/vnd.atlas.+json",           // empty version
            "",
        ] {
            assert!(
                bad.parse::<VersionedAcceptHeader>().is_err(),
                "should reject `{bad}`"
            );
        }
    }

    #[test]
    fn serde_roundtrip() {
        let header = "application/vnd.atlas.2024-10-23+json"
            .parse::<VersionedAcceptHeader>()
            .unwrap();
        let json = serde_json::to_string(&header).unwrap();
        assert_eq!(json, "\"application/vnd.atlas.2024-10-23+json\"");
        let back: VersionedAcceptHeader = serde_json::from_str(&json).unwrap();
        assert_eq!(back, header);
    }
}
