mod date;
mod mediatype;
mod version;

pub use date::VersionDate;
pub use mediatype::MediaType;
pub use version::Version;

// TODO: implement from_str/to_string/display + serialize/deserialize + useful derives (compare, etc)
pub struct VersionedAcceptHeader {
    pub version: Version,
    pub mediatype: MediaType,
}
