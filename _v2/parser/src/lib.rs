mod operation;
#[cfg(test)]
mod quirks;
mod spec;

pub use operation::{ConversionOptions, Operation, OperationParseError, OperationVersion};
pub use spec::{Spec, SpecParseError};
