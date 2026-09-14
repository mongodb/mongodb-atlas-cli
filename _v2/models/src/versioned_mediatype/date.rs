// TODO: implement from_str/to_string/display + serialize/deserialize + useful derives (compare, etc)
pub struct VersionDate {
    pub year: Year,
    pub month: Month,
    pub day: Day,
}

// TODO: do validation using nutype crate (derives) + add max possible useful derives
pub struct Year(u32);

// TODO: do validation using nutype crate (derives) + add max possible useful derives
pub struct Month(u32);

// TODO: do validation using nutype crate (derives) + add max possible useful derives
pub struct Day(u32);
