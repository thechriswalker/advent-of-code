#![allow(dead_code)]

pub struct Computer {
  mem: Vec<i32>,
  ip: i32,
}

enum State {
  Input(i32),
  Output(i32),
  Continue,
  Halt,
  Error,
}

const OP_ADD: i32 = 1;
const OP_MULTIPLY: i32 = 2;
const OP_STORE_INPUT: i32 = 3;
const OP_EMIT_OUTPUT: i32 = 4;
const OP_JUMP_TRUE: i32 = 5;
const OP_JUMP_FALSE: i32 = 6;
const OP_LESS_THAN: i32 = 7;
const OP_EQUAL_TO: i32 = 8;
// skip some...
const OP_HALT: i32 = 99;

enum Op {
  Add(i32, i32, i32),
  Multiply(i32, i32, i32),
  Input(i32),
  Output(i32),
  JumpIfTrue(i32, i32),
  JumpIfFalse(i32, i32),
  LessThan(i32, i32, i32),
  EqualTo(i32, i32, i32),
  Halt,
  Error, // unknown code
}

fn no_input() -> i32 {
  panic!("no input handler provided");
}

fn no_output(_: i32) {
  panic!("no output handler provided");
}

impl Computer {
  pub fn new(mem: Vec<i32>) -> Computer {
    Computer { mem: mem, ip: 0 }
  }

  pub fn load(&mut self, mem: Vec<i32>) {
    self.mem = mem;
    self.ip = 0;
  }

  pub fn run_no_io(&mut self) -> bool {
    return self.run(no_input, no_output);
  }

  pub fn run<FI, FO>(&mut self, mut input_handler: FI, mut output_handler: FO) -> bool
  where
    FI: FnMut() -> i32,
    FO: FnMut(i32),
  {
    loop {
      match self.tick() {
        State::Input(a) => self.write(a, input_handler()),
        State::Output(a) => output_handler(a),
        State::Continue => (),        // tick again
        State::Halt => return true,   // clean exit
        State::Error => return false, // fail condition
      }
    }
  }

  pub fn read(&self, register: i32) -> i32 {
    return self.mem[register as usize];
  }

  pub fn write(&mut self, register: i32, value: i32) {
    self.mem[register as usize] = value
  }

  fn tick(&mut self) -> State {
    // println!("tick: ip={}, ins={}", self.ip, self.mem[self.ip as usize]);
    // println!("mem {:?}", self.mem);
    let op = self.read_opcode();
    match op {
      Op::Add(a, b, c) => self.write(c, a + b),
      Op::Multiply(a, b, c) => self.write(c, a * b),
      Op::Input(a) => return State::Input(a),
      Op::Output(a) => return State::Output(a),
      Op::JumpIfFalse(a, b) => {
        if a == 0 {
          self.ip = b
        }
      }
      Op::JumpIfTrue(a, b) => {
        if a != 0 {
          self.ip = b
        }
      }
      Op::LessThan(a, b, c) => match a < b {
        true => self.write(c, 1),
        false => self.write(c, 0),
      },
      Op::EqualTo(a, b, c) => match a == b {
        true => self.write(c, 1),
        false => self.write(c, 0),
      },
      Op::Halt => return State::Halt,
      _ => return State::Error,
    }

    return State::Continue;
  }

  fn read_opcode(&mut self) -> Op {
    let op = self.consume();
    // the actual code is the lower 100 values
    match op % 100 {
      OP_ADD => Op::Add(self.parameter(op, 1), self.parameter(op, 2), self.consume()),
      OP_MULTIPLY => Op::Multiply(self.parameter(op, 1), self.parameter(op, 2), self.consume()),
      OP_STORE_INPUT => Op::Input(self.consume()),
      OP_EMIT_OUTPUT => Op::Output(self.parameter(op, 1)),
      OP_JUMP_FALSE => Op::JumpIfFalse(self.parameter(op, 1), self.parameter(op, 2)),
      OP_JUMP_TRUE => Op::JumpIfTrue(self.parameter(op, 1), self.parameter(op, 2)),
      OP_LESS_THAN => Op::LessThan(self.parameter(op, 1), self.parameter(op, 2), self.consume()),
      OP_EQUAL_TO => Op::EqualTo(self.parameter(op, 1), self.parameter(op, 2), self.consume()),
      OP_HALT => Op::Halt,
      _ => Op::Error,
    }
  }

  fn consume(&mut self) -> i32 {
    // get value in memory
    // increment instruction pointer,
    let v = self.mem[self.ip as usize];
    self.ip += 1;
    return v;
  }

  fn parameter(&mut self, op: i32, n: u32) -> i32 {
    // this one uses the op and the parameter position to work
    // out if this parameter should be fetched in positional or immediate mode
    let mode = (op / 10i32.pow(n + 1)) % 2;
    match mode {
      // the digit in the relevant 100s, 1000s, 10000s, column is 0 => positional mode
      // so the value is a register
      0 => {
        let v = self.consume();
        let o = self.read(v);
        // println!("pos mode: reg={}, value={}", v, o);
        return o;
      }
      // the digit is 1, so immediate mode
      1 => {
        let v = self.consume();
        // println!("imm mode: value={}", v);
        return v;
      }
      _ => panic!("What sort of crazy number was that? {}", mode),
    }
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_day_2_part_1() {
    let mut pc = Computer::new(vec![
      1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 6, 19, 1, 19, 6, 23, 2, 23, 6, 27, 2,
      6, 27, 31, 2, 13, 31, 35, 1, 9, 35, 39, 2, 10, 39, 43, 1, 6, 43, 47, 1, 13, 47, 51, 2, 6, 51,
      55, 2, 55, 6, 59, 1, 59, 5, 63, 2, 9, 63, 67, 1, 5, 67, 71, 2, 10, 71, 75, 1, 6, 75, 79, 1,
      79, 5, 83, 2, 83, 10, 87, 1, 9, 87, 91, 1, 5, 91, 95, 1, 95, 6, 99, 2, 10, 99, 103, 1, 5,
      103, 107, 1, 107, 6, 111, 1, 5, 111, 115, 2, 115, 6, 119, 1, 119, 6, 123, 1, 123, 10, 127, 1,
      127, 13, 131, 1, 131, 2, 135, 1, 135, 5, 0, 99, 2, 14, 0, 0,
    ]);
    pc.write(1, 12);
    pc.write(2, 2);
    pc.run_no_io();
    assert_eq!(pc.read(0), 3224742);
  }

  #[test]

  fn test_day_2_part_2() {
    let code = vec![
      1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 6, 19, 1, 19, 6, 23, 2, 23, 6, 27, 2,
      6, 27, 31, 2, 13, 31, 35, 1, 9, 35, 39, 2, 10, 39, 43, 1, 6, 43, 47, 1, 13, 47, 51, 2, 6, 51,
      55, 2, 55, 6, 59, 1, 59, 5, 63, 2, 9, 63, 67, 1, 5, 67, 71, 2, 10, 71, 75, 1, 6, 75, 79, 1,
      79, 5, 83, 2, 83, 10, 87, 1, 9, 87, 91, 1, 5, 91, 95, 1, 95, 6, 99, 2, 10, 99, 103, 1, 5,
      103, 107, 1, 107, 6, 111, 1, 5, 111, 115, 2, 115, 6, 119, 1, 119, 6, 123, 1, 123, 10, 127, 1,
      127, 13, 131, 1, 131, 2, 135, 1, 135, 5, 0, 99, 2, 14, 0, 0,
    ];
    let mut pc = Computer::new(vec![]);
    let target = 19690720;
    for n in 1..100 {
      for v in 1..100 {
        // clone and load
        pc.load(code.to_vec());
        // set the 2 register
        pc.write(1, n);
        pc.write(2, v);
        pc.run_no_io();
        if pc.read(0) == target {
          // we are done.
          assert_eq!(n * 100 + v, 7960);
          return;
        }
      }
    }
    panic!("should have found a solution")
  }

  #[test]
  fn test_day_5_part_2() {
    // aka is-eight
    // the example program uses an input instruction to ask for a single number.
    // The program will then output 999 if the input value is below 8,
    // output 1000 if the input value is equal to 8,
    // or output 1001 if the input value is greater than 8.
    let code = vec![
      3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0,
      1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105,
      1, 46, 98, 99,
    ];
    let mut pc = Computer::new(code.to_vec());
    let mut last_output: i32 = 0;
    pc.run(|| 8, |x| last_output = x);
    assert_eq!(last_output, 1000);

    pc.load(code.to_vec());
    pc.run(|| 9, |x| last_output = x);
    assert_eq!(last_output, 1001);

    pc.load(code.to_vec());
    pc.run(|| 0, |x| last_output = x);
    assert_eq!(last_output, 999);
  }
}
