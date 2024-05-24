# Emulation Notes

## Tips
 - wiki.nesdev.com
 - look at data sheets for the 6502
 - 

## Intro

### Emulation Overview
 - CPU: 2A03 (6502 with in-built audio integration)
 - CPU connects to components via a bus
 - CPU can address a 16-bit address space
 - Components include: CPU, PPU, APU, Cartridge, RAM, Controller Ports
 - memory map: 0x0000 - 0x7FFF
 - CPU can address 64KB of memory (most of which is not used)
 - APU is part of the CPU and is mapped to 0x4000 - 0x4017 (24 bytes)
 - The cartridge is mapped to 0x4020 - 0xFFFF, the end of the address space
 - PPU (2C02) is the graphics chip and is mapped to 0x2000 - 0x2007
 - Every clock cycle the PPU reads a byte from the cartridge
 - The CPU and PPU run at different clock speeds
 - The PPU has its own bus to communicate with the memory mapped to graphics, VRAM, and palettes
 - The graphics (8KB) exists in the cartridge and is stored at 0x0000 - 0x1FFF
 - VRAM is 2KB and is mapped to 0x2000 - 0x27FF
 - Palettes are 32 bytes and are mapped to 0x3F00 - 0x3FFF 
 - The PPU also communicates with the OAM which is not on any bus and is used to store the current sprite on the screen
 - The PPU is clocked at 3 times the speed of the CPU
 - The CPU uses DMA to write to the OAM
 - There are additional circuitry on the cartridges called mappers which are used to extend the capabilities of the NES
 - Mappers are used for bank switching, where the CPU can configure the mapper to give different data for the same address ranges.
 - ive been going through something, 1855 days. Ive been going through something. 

### Bitwise Operatins
 - Can perform logical operations with a mask to extract specific bits
 - AND: Extract bits
 - OR: Set bits
 - XOR: Toggle bits
 - NOT: Invert bits
 - Shifts: Move bits left or right to multiply or divide by 2

### Bitfields
 - Bitfields are used to store multiple boolean values in a single byte
 - Each bit in a byte can be used to store a boolean value
 - Bitfields are used to store multiple boolean values in a single byte
 - Each bit in a byte can be used to store a boolean value
 - go does not have a built-in bitfield type so will need to create a register type 


## CPU

### Notes
 - The CPU is a 6502
 - The CPU outputs a 16-bit address to the bus, and outputs a 8-bit data to the bus
 - It can read and write data. These functions use the same inputs, so there needs to be a r/w flag
 - The CPU needs a clock signal to operate. The clock forces the CPU to read the address bus and execute the instruction.
 - The CPU is connected to the bus via data and address lines
 - The CPU outputs an address to the bus and the device at that address will output data to the bus for the CPU to read
 - The CPU has the following registers:
  - A: Accumulator
  - X: Index Register
  - Y: Index Register
  These store 8-bit values
  - PC: Program Counter
  - stkp: Stack Pointer
  - status: Processor Status Register
  - We need to take into account the size and duration of each of the 56 instructions
  - We can represent the CPU as a struct with the registers as fields
  - We can represent the instructions as a 16x16 matrix where the first byte indexes the row and the second byte indexes the column e.g. 0x41 is LDA $01
  - The sequence of events for each instruction is:
    - Read the byte at the program counter
    - index the instruction matrix with the byte to get the addessing mode and the number of cycles
    - Read any bytes required by the addressing mode
    - Execute the instruction 
    - Wait for the number of cycles
  - Standard interrupts can be ignored but the non-maskable interrupt (NMI) must be handled

### Code
1. Create bus and cpu structs
  - Define what devices are connected to the bus (CPU and RAM - reset to zero)

2. Create a function to read and write to the bus
3. Connect the CPU to the bus, the cpu will read and write to the bus
4. Create an enum for the bits of the status register
5. Create the registers
6. Create a get and set function for the registers
7. Create 12 functions for the addressing modes
9. Create a function for each of the 56 instructions and one for illegal instructions
10. Create a function for the clock cycle, reset, and request interrupt and non-maskable interrupt (the last 3 are asynchronous)
11. Create a fetch function to read the byte at the program counter
12. Store the variable locations and cycle counts for each instruction 
13. create the 16x16 matrix for the instructions
14. create dissassembler function and store the opcodes and instructions in a map


