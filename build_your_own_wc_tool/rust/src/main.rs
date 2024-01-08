use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let filename = "example.txt"; // 修改为你要统计的文件名

    let file = File::open(filename).expect("Unable to open file");
    let reader = BufReader::new(file);

    let mut line_count = 0;
    let mut word_count = 0;
    let mut char_count = 0;

    for line in reader.lines() {
        let line = line.expect("Unable to read line");
        line_count += 1;
        char_count += line.len();
        word_count += line.split_whitespace().count();
    }

    println!("Lines: {}", line_count);
    println!("Words: {}", word_count);
    println!("Characters: {}", char_count);
}