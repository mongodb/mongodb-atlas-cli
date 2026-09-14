use crate::parameter::Parameter;

// TODO: implement to_string
// Represents a parameterized url, example: /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/{databaseName}/{collectionName}/collStats/measurements:
//
pub struct ParameterizedUrl {
    pub parts: Vec<Part>,
}

pub enum Part {
    Const(String),
    Parameter(Parameter),
}
