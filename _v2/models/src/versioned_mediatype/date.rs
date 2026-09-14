use std::fmt;
use std::str::FromStr;

use nutype::nutype;

/// Year of a version date (e.g. `2024` in `2024-10-23`), validated to a 4-digit year as
/// required by the Atlas Admin API accept headers (`RequestVersion.ACCEPT_HEADER_REGEX`).
#[nutype(
    validate(greater_or_equal = 1000, less_or_equal = 9999),
    derive(Clone, Copy, Debug, PartialEq, Eq, PartialOrd, Ord, Hash, Display, FromStr, Into)
)]
pub struct Year(u32);

/// Month of a version date, validated to the `1..=12` range.
#[nutype(
    validate(greater_or_equal = 1, less_or_equal = 12),
    derive(Clone, Copy, Debug, PartialEq, Eq, PartialOrd, Ord, Hash, Display, FromStr, Into)
)]
pub struct Month(u32);

/// Day of a version date, validated to the `1..=31` range.
#[nutype(
    validate(greater_or_equal = 1, less_or_equal = 31),
    derive(Clone, Copy, Debug, PartialEq, Eq, PartialOrd, Ord, Hash, Display, FromStr, Into)
)]
pub struct Day(u32);

/// A calendar date identifying an Atlas Admin API version, formatted `YYYY-MM-DD`
/// (zero-padded like Java's `java.time.LocalDate#toString`).
#[derive(
    Clone, Copy, Debug, PartialEq, Eq, PartialOrd, Ord, Hash, serde::Serialize, serde::Deserialize
)]
#[serde(into = "String", try_from = "String")]
pub struct VersionDate {
    pub year: Year,
    pub month: Month,
    pub day: Day,
}

impl fmt::Display for VersionDate {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "{:04}-{:02}-{:02}",
            u32::from(self.year),
            u32::from(self.month),
            u32::from(self.day)
        )
    }
}

impl From<VersionDate> for String {
    fn from(v: VersionDate) -> Self {
        v.to_string()
    }
}

impl TryFrom<String> for VersionDate {
    type Error = VersionDateParseError;

    fn try_from(s: String) -> Result<Self, Self::Error> {
        s.parse()
    }
}

impl FromStr for VersionDate {
    type Err = VersionDateParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut parts = s.trim().split('-');
        let (year, month, day) = match (parts.next(), parts.next(), parts.next(), parts.next()) {
            (Some(year), Some(month), Some(day), None) => (year, month, day),
            _ => return Err(VersionDateParseError::InvalidFormat(s.to_owned())),
        };

        let year = Year::from_str(year).map_err(VersionDateParseError::InvalidYear)?;
        let month = Month::from_str(month).map_err(VersionDateParseError::InvalidMonth)?;
        let day = Day::from_str(day).map_err(VersionDateParseError::InvalidDay)?;

        Ok(Self { year, month, day })
    }
}

#[derive(Debug, thiserror::Error)]
pub enum VersionDateParseError {
    /// Input not of the form `YYYY-MM-DD`.
    #[error("invalid version date `{0}`, expected `YYYY-MM-DD`")]
    InvalidFormat(String),
    #[error("invalid year: {0}")]
    InvalidYear(YearParseError),
    #[error("invalid month: {0}")]
    InvalidMonth(MonthParseError),
    #[error("invalid day: {0}")]
    InvalidDay(DayParseError),
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn parses_and_displays() {
        let date = VersionDate::from_str("2024-10-23").unwrap();
        assert_eq!(date.to_string(), "2024-10-23");
        assert_eq!(
            date,
            VersionDate {
                year: Year::from_str("2024").unwrap(),
                month: Month::from_str("10").unwrap(),
                day: Day::from_str("23").unwrap(),
            }
        );
    }

    #[test]
    fn display_zero_pads() {
        let date = VersionDate::from_str("2024-1-2").unwrap();
        assert_eq!(date.to_string(), "2024-01-02");
    }

    #[test]
    fn rejects_invalid_input() {
        for bad in [
            "2024-13-01",    // month 13
            "2024-01-32",    // day 32
            "999-01-01",     // year not 4 digits
            "10000-01-01",   // year > 9999
            "2024-01",       // missing day
            "2024-01-01-01", // too many fields
            "foo",
            "",
        ] {
            assert!(VersionDate::from_str(bad).is_err(), "should reject `{bad}`");
        }
    }

    #[test]
    fn orders_chronologically() {
        let early = VersionDate::from_str("2024-10-23").unwrap();
        let later = VersionDate::from_str("2025-01-01").unwrap();
        assert!(early < later);
        assert!(later > early);
        assert_eq!(early, VersionDate::from_str("2024-10-23").unwrap());
    }

    #[test]
    fn serde_roundtrip() {
        let date = VersionDate::from_str("2024-10-23").unwrap();
        let json = serde_json::to_string(&date).unwrap();
        assert_eq!(json, "\"2024-10-23\"");
        let back: VersionDate = serde_json::from_str(&json).unwrap();
        assert_eq!(back, date);
    }
}
