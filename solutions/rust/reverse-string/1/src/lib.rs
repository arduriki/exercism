pub fn reverse(input: &str) -> String {
    // Create a new String to store the reversed string
    let mut reversed_string = String::new();
    
    // Iterate over the characters of the input string
    // in reverse order.
    for c in input.chars().rev() {
        // Append each character to the reversed string
        reversed_string.push(c);
    }
    
    reversed_string
}
