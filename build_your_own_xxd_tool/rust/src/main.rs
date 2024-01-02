use std::fs::File;
use std::io::{self, Read};

fn main() -> io::Result<()> {
    let mut args = std::env::args().skip(1);
    let file_name = match args.next() {
        Some(name) => name,
        None => {
            eprintln!("Usage: xxd <file>");
            std::process::exit(1);
        }
    };

    let file = File::open(file_name)?;
    let mut reader = io::BufReader::new(file);
    let mut buffer = [0; 16];
    let mut offset = 0;

    loop {
        let bytes_read = reader.read(&mut buffer)?;

        if bytes_read == 0 {
            break;
        }

        print!("{:08x}: ", offset);

        for (i, &byte) in buffer.iter().enumerate() {
            if i < bytes_read {
                print!("{:02x} ", byte);
            } else {
                print!("   ");
            }

            if i == 7 {
                print!(" ");
            }
        }

        print!(" | ");

        for &byte in buffer.iter().take(bytes_read) {
            if byte >= 32 && byte <= 126 {
                print!("{}", byte as char);
            } else {
                print!(".");
            }
        }

        println!();
        offset += bytes_read;
    }

    Ok(())
}