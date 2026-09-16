mod operation;
#[cfg(test)]
mod quirks;
mod spec;

pub use operation::{Operation, OperationParseError};
pub use spec::{Spec, SpecParseError};
