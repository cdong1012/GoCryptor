# func bufferHashing(target []byte) uint64 {
# 	var result uint64 = 0

# 	for each := range target {
# 		result = ((result + uint64(each)) % 0xf6) ^ uint64(each)
# 	}
# 	return result
# }

# uint64_t ror(uint64_t v, unsigned int bits) 
# {
#     return (v>>bits) | (v<<(8*sizeof(uint64_t)-bits));
# }

def ror(v, bits):
    return (v>>bits) | (v<<(32-bits))
    

def bufferHashing(target):
    result = 0
    for each in target:
        result = ord(each) + ror(result, 15)
        
    return result

print(hex(bufferHashing("thebat")))
print(hex(bufferHashing("dssds")))
print("   _____        _____                  _             \n  / ____|      / ____|                | |            \n | |  __  ___ | |     _ __ _   _ _ __ | |_ ___  _ __ \n | | |_ |/ _ \| |    | '__| | | | '_ \| __/ _ \| '__|\n | |__| | (_) | |____| |  | |_| | |_) | || (_) | |   \n  \_____|\___/ \_____|_|   \__, | .__/ \__\___/|_|   \n                            __/ | |                  \n                           |___/|_|                  --> Your ID: %s\n--> Your key: %x\n")