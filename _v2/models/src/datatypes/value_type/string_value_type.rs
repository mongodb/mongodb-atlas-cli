use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct ValueTypeString {
    validation: Option<ValueTypeStringValidation>,
}

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub enum ValueTypeStringValidation {
    Length { min: Option<u32>, max: Option<u32> },
    Regex { regex: String },
}

#[cfg(test)]
mod tests {
    use super::*;

    /// `accessKeyID`: `minLength: 16`, `maxLength: 128`.
    #[test]
    fn access_key_id_maps_to_length_validation() {
        let v = ValueTypeString {
            validation: Some(ValueTypeStringValidation::Length {
                min: Some(16),
                max: Some(128),
            }),
        };
        let json = serde_json::to_string(&v).unwrap();
        assert_eq!(json, r#"{"validation":{"Length":{"min":16,"max":128}}}"#);
        assert_eq!(serde_json::from_str::<ValueTypeString>(&json).unwrap(), v);
    }

    /// `customerMasterKeyID`: `minLength: 1`, `maxLength: 2048`.
    #[test]
    fn customer_master_key_id_maps_to_length_validation() {
        let v = ValueTypeString {
            validation: Some(ValueTypeStringValidation::Length {
                min: Some(1),
                max: Some(2048),
            }),
        };
        let json = serde_json::to_string(&v).unwrap();
        assert_eq!(json, r#"{"validation":{"Length":{"min":1,"max":2048}}}"#);
        assert_eq!(serde_json::from_str::<ValueTypeString>(&json).unwrap(), v);
    }

    /// `roleId`: `pattern: ^([a-f0-9]{24})$`.
    #[test]
    fn role_id_maps_to_regex_validation() {
        let v = ValueTypeString {
            validation: Some(ValueTypeStringValidation::Regex {
                regex: "^([a-f0-9]{24})$".to_string(),
            }),
        };
        let json = serde_json::to_string(&v).unwrap();
        assert_eq!(
            json,
            r#"{"validation":{"Regex":{"regex":"^([a-f0-9]{24})$"}}}"#
        );
        assert_eq!(serde_json::from_str::<ValueTypeString>(&json).unwrap(), v);
    }

    /// `secretAccessKey`: only `type: string`, no length or pattern -> no validation.
    #[test]
    fn secret_access_key_has_no_validation() {
        let v = ValueTypeString { validation: None };
        let json = serde_json::to_string(&v).unwrap();
        assert_eq!(json, r#"{"validation":null}"#);
        assert_eq!(serde_json::from_str::<ValueTypeString>(&json).unwrap(), v);
    }

    /// A `minLength`-only constraint leaves `max` empty.
    #[test]
    fn length_validation_optional_bounds() {
        let v: ValueTypeString = serde_json::from_str(
            r#"{"validation":{"Length":{"min":1,"max":null}}}"#,
        )
        .unwrap();
        assert_eq!(
            v.validation,
            Some(ValueTypeStringValidation::Length {
                min: Some(1),
                max: None,
            })
        );
    }
}
