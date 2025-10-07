use std::env;
use nats;
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Serialize, Debug)]
struct TaskRequest { input: String }

fn heavy_compute(input: &str) -> String {
    // просте псевдо-важке обчислення: реверс + повтор
    let reversed: String = input.chars().rev().collect();
    format!("{}:{}", reversed, input.len())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let nats_url = env::var("NATS_URL").unwrap_or_else(|_| "nats://127.0.0.1:4222".into());
    let nc = nats::connect(nats_url)?;

    let sub = nc.subscribe("tasks")?;
    for msg in sub.messages() {
        if let Ok(task) = serde_json::from_slice::<TaskRequest>(&msg.data) {
            let result = heavy_compute(&task.input);
            let payload = serde_json::to_vec(&serde_json::json!({"result": result}))?;
            let _ = nc.publish("results", payload);
        }
    }

    Ok(())
}

