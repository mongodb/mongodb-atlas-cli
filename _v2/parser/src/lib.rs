use nutype::nutype;

pub struct Operation {
    pub http_verb: HTTPVerb,
    pub version_str: Version,
    pub url_template: String,
    pub url_param_names: Vec<String>,
    pub is_paginated: bool,
    pub response_type: String,
    /// "json" (default) or "gzip"
    pub response_format: String,
}

macro_rules! simple_str_nutype {
    ($type_name:ident) => {
        #[nutype(
            sanitize(trim),
            validate(not_empty),
            derive(Clone, Debug, PartialEq, Eq, PartialOrd, Ord, Hash, AsRef, Deref)
        )]
        pub struct $type_name(String);
    };
}

simple_str_nutype!(HTTPVerb);
simple_str_nutype!(Version);
