use super::VersionDate;

// TODO: implement from_str/to_string/display + serialize/deserialize + useful derives (compare, etc)
pub enum Version {
    Stable(VersionDate),
    Preview,
    Upcoming(VersionDate),
}
