//! Force a rebuild when the files the `atlas_cli` derive reads change.
//!
//! The derive reads `spec` + `hierarchy` at expansion time, but cargo only
//! knows about files a build script reports. Without these directives the
//! generated CLI goes stale on config edits.

fn main() {
    println!("cargo:rerun-if-changed=spec.yaml");
    println!("cargo:rerun-if-changed=hierarchy.yaml");
}
