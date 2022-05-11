use anyhow::Result;
use spin_sdk::{
    http::{Request, Response},
    http_component,
};
use std::time::{SystemTime, UNIX_EPOCH};

/// A simple Spin HTTP component.
#[http_component]
fn hello_world(_req: Request) -> Result<Response> {
    println!("'hello, world!' from the WAsm module!");
    let nanos = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .subsec_nanos();
    println!("nanos: {}", nanos);
    Ok(http::Response::builder()
        .status(200)
        .body(Some("Writing a very simple Spin component in Rust\n".into()))?)
}
