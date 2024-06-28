# Experimental polymorphic pseudonym setup.

- Based on my original proof-of-concept written in Rust, which uses the libpep crate. 
- All components (transcryptor, access manager, pools, keyserver etc) are inside the same program. In a normal setup these components are separated and talk with GRPC/HTTP or similar.
- Minimalistic implementation of ScalarNonZero and GroupElement, but enough for this proof-of-concept.


## Usage

    go run ./...

There are 2 pieces of information send from a "watch" into the storage facility: 2 AES keys. These keys should be read back by another party (The Doctor).

```
GoPP
TRAC[0000] [KS] GenerateFactor: SF
TRAC[0000] [KS] GenerateFactor: DOC
This is work that is done inside our smartwatch that knows our BSN
**************************************************************
AES key1 generated:  04110660fcc605ea850149be4e58c384874b78e1a4009d806f9e7a9a78786e7f
**************************************************************
TRAC[0000] [GP] GeneratePseudonym: 950000024 f83820febc18b518bdc9d85a6337fcc15bc88d8d92d4fb4742db2b669868aa78
DEBU[0000] ! epid(A)@SF: a6e6b58cf8a463cc39e1264e5117d1a0a4b23b0f75b84f8728fe99e8487d5f63aa7877f02d9a89f6bd75200367bc1b0ad194bbf86b576daf7d4403df8522eb2b6c4927b0679e4a125a5962a98a0d744c7cf6c4466a77f34a4641f1b0d3abdd6a
TRAC[0000] [SF] Store: {[166 230 181 140 248 164 99 204 57 225 38 78 81 23 209 160 164 178 59 15 117 184 79 135 40 254 153 232 72 125 95 99] [170 120 119 240 45 154 137 246 189 117 32 3 103 188 27 10 209 148 187 248 107 87 109 175 125 68 3 223 133 34 235 43] [108 73 39 176 103 158 74 18 90 89 98 169 138 13 116 76 124 246 196 70 106 119 243 74 70 65 241 176 211 171 221 106]} {[174 11 122 207 227 6 117 23 82 115 161 21 76 37 212 70 212 105 201 232 133 201 62 172 240 175 25 238 108 159 69 54] [118 219 86 176 96 209 10 223 47 226 66 235 1 225 27 202 163 28 38 83 210 231 90 32 193 52 140 90 246 103 52 122] [182 74 165 225 48 18 248 217 203 185 77 4 5 37 92 118 34 219 31 192 213 205 62 54 130 214 151 176 60 237 71 116]}
TRAC[0000] [SF] Decrypted key: 20a5411eb6a117659e8b63827d1b29d628cb83b9835b9e046c648d47657bdb6f
**************************************************************
AES key2 generated:  90301cf1cbea971c7006f722350cec87d88adfccd025e189d0e2748861539e69
**************************************************************
DEBU[0000] ! epid(A)@SF: 2ad3375d6b74abf810eabe7b7f1efa1c3dca25e6562197cdd0b949c6ba1a7a673c243c9704af078d0b85a7a6565389f277c5c06c37a6dab7ff7c62634965f8596c4927b0679e4a125a5962a98a0d744c7cf6c4466a77f34a4641f1b0d3abdd6a
TRAC[0000] [SF] Store: {[42 211 55 93 107 116 171 248 16 234 190 123 127 30 250 28 61 202 37 230 86 33 151 205 208 185 73 198 186 26 122 103] [60 36 60 151 4 175 7 141 11 133 167 166 86 83 137 242 119 197 192 108 55 166 218 183 255 124 98 99 73 101 248 89] [108 73 39 176 103 158 74 18 90 89 98 169 138 13 116 76 124 246 196 70 106 119 243 74 70 65 241 176 211 171 221 106]} {[26 56 233 0 43 66 108 127 122 42 59 141 217 92 206 104 204 213 253 174 120 88 144 58 83 17 130 163 179 86 150 13] [150 222 31 66 152 111 136 169 176 50 202 126 97 174 233 123 124 35 249 30 126 241 214 14 50 31 81 150 217 28 141 78] [182 74 165 225 48 18 248 217 203 185 77 4 5 37 92 118 34 219 31 192 213 205 62 54 130 214 151 176 60 237 71 116]}
TRAC[0000] [SF] Decrypted key: 20a5411eb6a117659e8b63827d1b29d628cb83b9835b9e046c648d47657bdb6f
All done. The smartwatch has send 2 pieces of data to the storage facility for this BSN.


At the doctors. The doctor wants to retrieve the data for a specific BSN
TRAC[0000] [GP] GeneratePseudonym: 950000024 f83820febc18b518bdc9d85a6337fcc15bc88d8d92d4fb4742db2b669868aa78
DEBU[0000] ! epid(A)@SF: bef68ac0ccb933ead39bb967f4196a2099e7f9aee35e4f826538c08c4e28604522dfe74877d4857aaf50feff6162cc83d61a4ac57ce216f8c0ad96f0a89bda3b6c4927b0679e4a125a5962a98a0d744c7cf6c4466a77f34a4641f1b0d3abdd6a
TRAC[0000] [SF] Retrieve: {[190 246 138 192 204 185 51 234 211 155 185 103 244 25 106 32 153 231 249 174 227 94 79 130 101 56 192 140 78 40 96 69] [34 223 231 72 119 212 133 122 175 80 254 255 97 98 204 131 214 26 74 197 124 226 22 248 192 173 150 240 168 155 218 59] [108 73 39 176 103 158 74 18 90 89 98 169 138 13 116 76 124 246 196 70 106 119 243 74 70 65 241 176 211 171 221 106]}
TRAC[0000] [SF] Decrypted key: 20a5411eb6a117659e8b63827d1b29d628cb83b9835b9e046c648d47657bdb6f
Data retrieved from the storage for bsn:  2
INFO[0000] Decrypted AES key: 04110660fcc605ea850149be4e58c384874b78e1a4009d806f9e7a9a78786e7f
INFO[0000] Decrypted AES key: 90301cf1cbea971c7006f722350cec87d88adfccd025e189d0e2748861539e69
```

Here the last two decrypted keys are the ones the watch has send to the storage facility. Note that the storage facility is not able to decrypt this data, nor is the access manager or the transcryptor.


## TODO:
- Read through the whitepaper and see if we got things correctly.
